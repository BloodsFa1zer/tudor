package services

import (
	"fmt"
	"reflect"
	"time"

	config "study_marketplace/config"
	"study_marketplace/internal/database/queries"
	"study_marketplace/internal/database/repositories"
	"study_marketplace/models"

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
	db   repositories.UsersRepository
	conf *config.Config
}

func NewUserService(conf *config.Config, db repositories.UsersRepository) UserService {
	return &userService{db, conf}
}

func (t *userService) UserLogin(ctx *gin.Context, inputModel models.InLogin) (string, error) {

	user, err := t.db.GetUserByEmail(ctx, inputModel.Email)
	if err != nil {
		return "error", err
	}

	cmpPassword := ComparePassword(user.Password, inputModel.Password)
	if cmpPassword != nil {
		err := fmt.Errorf("invalid email or password")
		return "error", err
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "error", err
	}

	return token, nil
}

func (t *userService) UserRegister(ctx *gin.Context, inputModel queries.User) (queries.User, error) {
	isEmailExist, err := t.db.IsUserEmailExist(ctx, inputModel.Email)
	if err != nil {
		err = fmt.Errorf("db search error")
		return queries.User{}, err
	}
	if isEmailExist {
		err = fmt.Errorf("user with such email already registred")
		return queries.User{}, err
	}

	hashedPassword := hashPassword(inputModel.Password)

	args := &queries.CreateUserParams{
		Name:      inputModel.Name,
		Email:     inputModel.Email,
		Password:  hashedPassword,
		Photo:     "default.jpeg",
		Verified:  false,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := t.db.CreateUser(ctx, *args)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (t *userService) UserInfo(ctx *gin.Context, userId int64) (queries.User, error) {
	user, err := t.db.GetUserById(ctx, userId)
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

func (t *userService) UserPatch(ctx *gin.Context, patch queries.User) (queries.User, error) {

	user, err := t.db.GetUserById(ctx, patch.ID)

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
		patch.Password = hashPassword(patch.Password)
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

	patchedUser, err := t.db.UpdateUser(ctx, *userTmp)
	if err != nil {
		fmt.Println("Faield to update user")
	}

	return patchedUser, nil
}

func (t *userService) PasswordReset(ctx *gin.Context, email models.EmailRequest) (bool, error) {
	validEmail, _ := t.db.IsUserEmailExist(ctx, email.Email)

	if !validEmail {
		fmt.Println("Email not found")
		return validEmail, fmt.Errorf("Email not found")
	}

	user, err := t.db.GetUserByEmail(ctx, email.Email)
	if err != nil {
		return validEmail, fmt.Errorf("Failed request to DB.")
	}

	_, err = t.EmailSend(email.Email, user)
	if err != nil {
		return false, err
	}

	return validEmail, nil
}

func (t *userService) PasswordCreate(ctx *gin.Context, userID int64, newPassword models.UserPassword) error {
	if newPassword.Password == "" {
		return fmt.Errorf("New password not valid.")
	}
	patchPassword := hashPassword(newPassword.Password)

	user, err := t.db.GetUserById(ctx, userID)
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

	_, err = t.db.UpdateUser(ctx, updateUser)

	if err != nil {
		return err
	}

	return nil
}

func (t *userService) EmailSend(userEmail string, user queries.User) (bool, error) {
	token, err := GenerateToken(user)
	if err != nil {
		return false, fmt.Errorf("Failed to generate token.")
	}

	from := configs.GOOGLE_EMAIL_ADDRESS
	response := NewEmail(user.Name, token).message
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", userEmail)
	msg.SetHeader("Subject", "Password reset")
	msg.SetBody("text/html", response)

	postman := gomail.NewDialer("smtp.gmail.com", 587, from, configs.GOOGLE_EMAIL_SECRET)

	if err := postman.DialAndSend(msg); err != nil {
		return false, fmt.Errorf("Failed to send email.")
	}

	return true, nil
}
