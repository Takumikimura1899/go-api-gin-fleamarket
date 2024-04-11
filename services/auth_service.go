package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

// passwordをハッシュ化してユーザーを作成
func (s *AuthService) SignUp(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

// TODO: tokenの生成処理を実装する 仮でemailを返す
func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &foundUser.Email, nil
}
