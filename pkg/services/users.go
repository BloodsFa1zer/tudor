package services

import (
	"context"
	"fmt"

	dbmappers "study_marketplace/pkg/domen/mappers/db_mappers"
	entities "study_marketplace/pkg/domen/models/entities"
	config "study_marketplace/pkg/infrastructure/config"
	"study_marketplace/pkg/repositories"
)

type UserService interface {
	UserLogin(ctx context.Context, inputuser *entities.User) (string, *entities.User, error)
	UserRegister(ctx context.Context, user *entities.User) (string, *entities.User, error)
	UserInfo(ctx context.Context, userId int64) (*entities.User, error)
	UserPatch(ctx context.Context, patch *entities.User) (string, *entities.User, error)
	PasswordReset(ctx context.Context, email string) error
	PasswordCreate(ctx context.Context, userID int64, newPassword string) error
}

type userService struct {
	db          repositories.UsersRepository
	conf        *config.Config
	genToken    func(userid int64, userName string) (string, error)
	hashPass    func(password string) string
	comparePass func(hashedPassword string, password string) error
	sendEmail   func(token string, to string) error
}

func NewUserService(
	conf *config.Config,
	gTF func(userid int64, userName string) (string, error),
	hPass func(password string) string,
	cPass func(hashedPassword string, password string) error,
	sendEmail func(token string, to string) error,
	db repositories.UsersRepository) UserService {
	return &userService{db, conf, gTF, hPass, cPass, sendEmail}
}

func (s *userService) UserLogin(ctx context.Context, inputuser *entities.User) (string, *entities.User, error) {
	user, err := s.db.GetUserByEmail(ctx, inputuser.Email)
	if err != nil {
		return "", nil, err
	}
	if err = s.comparePass(user.Password, inputuser.Password); err != nil {
		err := fmt.Errorf("invalid email or password")
		return "", nil, err
	}
	token, err := s.genToken(user.ID, user.Email)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (s *userService) UserRegister(ctx context.Context, user *entities.User) (string, *entities.User, error) {
	user.Password = s.hashPass(user.Password)
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

func (s *userService) UserInfo(ctx context.Context, userId int64) (*entities.User, error) {
	user, err := s.db.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ProviderAuth(ctx context.Context, userInfo *entities.User) (string, error) {
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

func (s *userService) UserPatch(ctx context.Context, patch *entities.User) (string, *entities.User, error) {
	patch, err := s.db.UpdateUser(ctx, patch)
	if err != nil {
		return "", nil, err
	}
	token, err := s.genToken(patch.ID, patch.Email)
	if err != nil {
		return "", nil, err
	}
	return token, patch, nil
}

func (s *userService) PasswordReset(ctx context.Context, email string) error {
	user, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("failed request to DB")
	}

	token, err := s.genToken(user.ID, user.Email)
	if err != nil {
		return fmt.Errorf("failed to generate token")
	}

	if err := s.sendEmail(token, user.Email); err != nil {
		return err
	}
	return nil
}

func (s *userService) PasswordCreate(ctx context.Context, userID int64, newPassword string) error {
	pass := s.hashPass(newPassword)
	_, err := s.db.UpdateUser(ctx, &entities.User{ID: userID, Password: pass})
	if err != nil {
		return fmt.Errorf("failed request to DB")
	}
	return nil
}

// func (s *userService) emailSend(userEmail string, user entities.User) (bool, error) {

// 	response := s.newEmail(user.Email, token).message
// 	msg := gomail.NewMessage()

// 	msg.SetHeader("From", fmt.Sprintf("%s <%s>", s.conf.GoogleEmailSenderName, s.conf.GoogleEmailAddress))
// 	msg.SetHeader("To", userEmail)
// 	msg.SetHeader("Subject", "Password reset")
// 	msg.SetBody("text/html", response)
// 	postman := gomail.NewDialer("smtp.gmail.com", 587, s.conf.GoogleEmailAddress, s.conf.GoogleEmailSecret)

// 	if err := postman.DialAndSend(msg); err != nil {
// 		return false, fmt.Errorf("failed to send email")
// 	}

// 	return true, nil
// }
