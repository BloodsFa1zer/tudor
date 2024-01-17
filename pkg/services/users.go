package services

import (
	"context"
	"fmt"

	entities "study_marketplace/pkg/domain/models/entities"
	"study_marketplace/pkg/repositories"
)

type UserService interface {
	UserLogin(ctx context.Context, inputuser *entities.User) (string, *entities.User, error)
	UserRegister(ctx context.Context, user *entities.User) (string, *entities.User, error)
	UserInfo(ctx context.Context, userId int64) (*entities.User, error)
	UserPatch(ctx context.Context, patch *entities.User) (string, *entities.User, error)
	PasswordReset(ctx context.Context, email string) error
	PasswordChange(ctx context.Context, userID int64, currentPassword string, newPassword string) error
	PasswordCreate(ctx context.Context, userID int64, newPassword string) error
	EmailChange(ctx context.Context, userID int64, currentPassword string, newEmail string) error
}

type userService struct {
	db          repositories.UsersRepository
	genToken    func(userid int64, userName string) (string, error)
	hashPass    func(password string) string
	comparePass func(hashedPassword string, password string) error
	sendEmail   func(token string, to string) error
}

func NewUserService(
	gTF func(userid int64, userName string) (string, error),
	hPass func(password string) string,
	cPass func(hashedPassword string, password string) error,
	sendEmail func(token string, to string) error,
	db repositories.UsersRepository) UserService {
	return &userService{db, gTF, hPass, cPass, sendEmail}
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
	user, err := s.db.CreateUser(ctx, user)
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

func (s *userService) PasswordChange(ctx context.Context, userID int64, currentPassword string, newPassword string) error {
	user, err := s.db.GetUserById(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed request to DB")
	}

	if err := s.comparePass(user.Password, currentPassword); err != nil {
		return fmt.Errorf("current password is wrong")
	}

	userWithNewPassword := &entities.User{ID: user.ID, Password: s.hashPass(newPassword)}

	_, err = s.db.UpdateUser(ctx, userWithNewPassword)
	if err != nil {
		return fmt.Errorf("failed to update password in the database")
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

func (s *userService) EmailChange(ctx context.Context, userID int64, currentPassword string, newEmail string) error {
	user, err := s.db.GetUserById(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed request to DB")
	}

	if err := s.comparePass(user.Password, currentPassword); err != nil {
		return fmt.Errorf("current password is wrong")
	}

	if user.Email == newEmail {
		return fmt.Errorf("current email and new email are equal")
	}

	userWithNewEmail := &entities.User{ID: user.ID, Email: newEmail}

	_, err = s.db.UpdateUser(ctx, userWithNewEmail)
	if err != nil {
		return fmt.Errorf("failed to update email in the database")
	}

	return nil
}
