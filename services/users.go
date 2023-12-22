package services

import (
	"fmt"
	"reflect"
	"time"

	"study_marketplace/database/queries"
	"study_marketplace/database/repositories"
	"study_marketplace/domen/models"
	config "study_marketplace/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/gomail.v2"
)

type UserService interface {
	UserLogin(ctx *gin.Context, inputModel models.InLogin) (string, error)
	UserRegister(ctx *gin.Context, inputModel queries.User) (queries.User, error)
	UserInfo(ctx *gin.Context, userId int64) (queries.User, error)
	GetOrCreateUser(ctx *gin.Context, userInfo models.GoogleResponse) (queries.User, error)
	UserPatch(ctx *gin.Context, patch queries.User) (queries.User, error)
	PasswordReset(ctx *gin.Context, email models.EmailRequest) (bool, error)
	PasswordCreate(ctx *gin.Context, userID int64, newPassword models.UserPassword) error
	EmailSend(userEmail string, user queries.User) (bool, error)
}

type userService struct {
	db          repositories.UsersRepository
	conf        *config.Config
	genToken    func(userid int64, userName string) (string, error)
	hashPass    func(password string) string
	comparePass func(hashedPassword string, password string) error
}

func NewUserService(
	conf *config.Config,
	gTF func(userid int64, userName string) (string, error),
	hPass func(password string) string,
	cPass func(hashedPassword string, password string) error,
	db repositories.UsersRepository) UserService {
	return &userService{db, conf, gTF, hPass, cPass}
}

func (s *userService) UserLogin(ctx *gin.Context, inputModel models.InLogin) (string, error) {

	user, err := s.db.GetUserByEmail(ctx, inputModel.Email)
	if err != nil {
		return "error", err
	}

	cmpPassword := s.comparePass(user.Password, inputModel.Password)
	if cmpPassword != nil {
		err := fmt.Errorf("invalid email or password")
		return "error", err
	}

	token, err := s.genToken(user.ID, user.Name)
	if err != nil {
		return "error", err
	}

	return token, nil
}

func (s *userService) UserRegister(ctx *gin.Context, inputModel queries.User) (queries.User, error) {
	isEmailExist, err := s.db.IsUserEmailExist(ctx, inputModel.Email)
	if err != nil {
		err = fmt.Errorf("db search error")
		return queries.User{}, err
	}
	if isEmailExist {
		err = fmt.Errorf("user with such email already registred")
		return queries.User{}, err
	}

	args := &queries.CreateUserParams{
		Name:      inputModel.Name,
		Email:     inputModel.Email,
		Password:  s.hashPass(inputModel.Password),
		Photo:     "default.jpeg",
		Verified:  false,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := s.db.CreateUser(ctx, *args)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (s *userService) UserInfo(ctx *gin.Context, userId int64) (queries.User, error) {
	user, err := s.db.GetUserById(ctx, userId)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (t *userService) GetOrCreateUser(ctx *gin.Context, userInfo models.GoogleResponse) (queries.User, error) {
	isEmailExist, err := t.db.IsUserEmailExist(ctx, userInfo.Email)
	if err != nil {
		fmt.Println("email search query failed")
	}

	var user queries.User

	if isEmailExist {
		user, err = t.db.GetUserByEmail(ctx, userInfo.Email)
		if err != nil {
			fmt.Println("failed to find user")
		}
	} else {
		args := &queries.CreateUserParams{
			Name:      userInfo.Name,
			Email:     userInfo.Email,
			Password:  "",
			Photo:     "default.jpeg",
			Verified:  false,
			Role:      "user",
			UpdatedAt: time.Now(),
		}

		user, err = t.db.CreateUser(ctx, *args)
		if err != nil {
			fmt.Println("Faield to create user")
		}
	}
	return user, nil
}

func (s *userService) UserPatch(ctx *gin.Context, patch queries.User) (queries.User, error) {
	user, err := s.db.GetUserById(ctx, patch.ID)
	userTmp := &queries.UpdateUserParams{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Photo:     user.Photo,
		Verified:  user.Verified,
		Password:  user.Password,
		Role:      user.Role,
		UpdatedAt: user.UpdatedAt,
	}

	if patch.Password != "" {
		patch.Password = s.hashPass(patch.Password)
	}

	userValue := reflect.ValueOf(userTmp).Elem()
	patchValue := reflect.ValueOf(patch)
	for i := 0; i < userValue.NumField(); i++ {
		field := userValue.Field(i)
		updateField := patchValue.Field(i)

		if updateField.IsValid() && !updateField.IsZero() {
			field.Set(updateField)
		}
	}

	userTmp.UpdatedAt = time.Now()

	patchedUser, err := s.db.UpdateUser(ctx, *userTmp)
	if err != nil {
		fmt.Println("Faield to update user")
	}

	return patchedUser, nil
}

func (s *userService) PasswordReset(ctx *gin.Context, email models.EmailRequest) (bool, error) {
	validEmail, _ := s.db.IsUserEmailExist(ctx, email.Email)

	if !validEmail {
		fmt.Println("Email not found")
		return validEmail, fmt.Errorf("Email not found")
	}

	user, err := s.db.GetUserByEmail(ctx, email.Email)
	if err != nil {
		return validEmail, fmt.Errorf("Failed request to DB.")
	}

	_, err = s.EmailSend(email.Email, user)
	if err != nil {
		return false, err
	}

	return validEmail, nil
}

func (s *userService) PasswordCreate(ctx *gin.Context, userID int64, newPassword models.UserPassword) error {
	if newPassword.Password == "" {
		return fmt.Errorf("New password not valid.")
	}
	patchPassword := s.hashPass(newPassword.Password)

	user, err := s.db.GetUserById(ctx, userID)
	if err != nil {
		return fmt.Errorf("Failed to find user.")
	}

	updateUser := queries.UpdateUserParams{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Photo:     user.Photo,
		Verified:  user.Verified,
		Password:  patchPassword,
		Role:      user.Role,
		UpdatedAt: time.Now(),
	}

	_, err = s.db.UpdateUser(ctx, updateUser)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) EmailSend(userEmail string, user queries.User) (bool, error) {
	token, err := s.genToken(user.ID, user.Name)
	if err != nil {
		return false, fmt.Errorf("Failed to generate token.")
	}

	from := s.conf.GoogleEmailAddress
	response := s.newEmail(user.Name, token).message
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", userEmail)
	msg.SetHeader("Subject", "Password reset")
	msg.SetBody("text/html", response)

	postman := gomail.NewDialer("smtp.gmail.com", 587, from, s.conf.GoogleEmailSecret)

	if err := postman.DialAndSend(msg); err != nil {
		return false, fmt.Errorf("Failed to send email.")
	}

	return true, nil
}
