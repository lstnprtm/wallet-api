package usecase

import (
	"errors"
	"github.com/lstnprtm/wallet-api/internal/domain"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	repo domain.AuthRepository
}

func NewAuthUsecase(r domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{repo: r}
}

func (u *authUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func (u *authUsecase) Register(username, password string) error {
	_, err := u.repo.GetByUsername(username)
	if err == nil {
		return errors.New("username already taken")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.repo.CreateUser(username, string(hashed))
}
