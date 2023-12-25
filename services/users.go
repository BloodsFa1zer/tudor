package services

import (
	"fmt"

	dbmappers "study_marketplace/domen/mappers/db_mappers"
	entities "study_marketplace/domen/models/entities_models"
	config "study_marketplace/pkg/infrastructure/config"
	"study_marketplace/pkg/repositories"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/gomail.v2"
)

type UserService interface {
	UserLogin(ctx *gin.Context, user *entities.User) (string, error)
	UserRegister(ctx *gin.Context, user *entities.User) (string, *entities.User, error)
	UserInfo(ctx *gin.Context, userId int64) (*entities.User, error)
	ProviderAuth(ctx *gin.Context, userInfo *entities.User) (string, error)
	UserPatch(ctx *gin.Context, patch *entities.User) (*entities.User, error)
	PasswordReset(ctx *gin.Context, email string) (bool, error)
	PasswordCreate(ctx *gin.Context, userID int64, newPassword string) error
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

func (s *userService) UserLogin(ctx *gin.Context, inputuser *entities.User) (string, error) {
	user, err := s.db.GetUserByEmail(ctx, inputuser.Email)
	if err != nil {
		return "", err
	}
	if err = s.comparePass(user.Password, inputuser.Password); err != nil {
		err := fmt.Errorf("invalid email or password")
		return "", err
	}

	token, err := s.genToken(inputuser.ID, inputuser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) UserRegister(ctx *gin.Context, user *entities.User) (string, *entities.User, error) {
	user, err := s.db.CreateUser(ctx, dbmappers.UserToCreateUserParams(user))
	if err != nil {
		return "", nil, err
	}
	token, err := s.genToken(user.ID, user.Email)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (s *userService) UserInfo(ctx *gin.Context, userId int64) (*entities.User, error) {
	user, err := s.db.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ProviderAuth(ctx *gin.Context, userInfo *entities.User) (string, error) {
	user, err := s.db.CreateorUpdateUser(ctx, userInfo)
	if err != nil {
		return "", err
	}
	token, err := s.genToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) UserPatch(ctx *gin.Context, patch *entities.User) (*entities.User, error) {
	patch, err := s.db.UpdateUser(ctx, patch)
	if err != nil {
		return nil, err
	}
	return patch, nil
}

func (s *userService) PasswordReset(ctx *gin.Context, email string) (bool, error) {
	var validEmail bool = true
	user, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return !validEmail, fmt.Errorf("failed request to DB")
	}
	_, err = s.emailSend(email, *user)
	if err != nil {
		return false, err
	}
	return validEmail, nil
}

func (s *userService) PasswordCreate(ctx *gin.Context, userID int64, newPassword string) error {
	pass := s.hashPass(newPassword)
	_, err := s.db.UpdateUser(ctx, &entities.User{ID: userID, Password: pass})
	if err != nil {
		return fmt.Errorf("failed request to DB")
	}
	return nil
}

func (s *userService) emailSend(userEmail string, user entities.User) (bool, error) {
	token, err := s.genToken(user.ID, "user.Name")
	if err != nil {
		return false, fmt.Errorf("failed to generate token")
	}

	from := s.conf.GoogleEmailAddress
	response := s.newEmail("user.Name", token).message
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", userEmail)
	msg.SetHeader("Subject", "Password reset")
	msg.SetBody("text/html", response)

	postman := gomail.NewDialer("smtp.gmail.com", 587, from, s.conf.GoogleEmailSecret)

	if err := postman.DialAndSend(msg); err != nil {
		return false, fmt.Errorf("failed to send email")
	}

	return true, nil
}
