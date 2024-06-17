package usecase

import (
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/user/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserUsecase struct {
	userRepo *repository.UserRepo
}

func NewUserUsecase(repo *repository.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (u *UserUsecase) RegisterUser(user *model.RegisterUser) error {
	dbUser := &entity.User{
		Username: user.Username,
		Password: user.Password,
		Type: user.Type,
	}

	err := u.userRepo.CreateUser(dbUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetAllUsers() ([]model.UserDetail, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUsecase) DeleteAll() error {
	err := u.userRepo.DeleteAll()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Login(user *model.LoginUser) (string, error) {
	userDetail, err := u.userRepo.Login(user.Username, user.Password)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"userID": userDetail.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"userType": userDetail.Type,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}