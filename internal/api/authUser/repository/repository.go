package repository

import (
	"foodOrder/domain/entity"
	"foodOrder/domain/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	userDB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{userDB: db}
}

func (u *UserRepo) GetUserByUsername(username string) (*model.UserDetail, error) {
	var user model.UserDetail
	if err := u.userDB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) CreateUser(user *entity.User) (error) {
	dbTx := u.userDB.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Create(user).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}

func (u *UserRepo) GetAllUsers() ([]model.UserDetail, error) {
	var users []entity.User
	if err := u.userDB.Find(&users).Error; err != nil {
		return nil, err
	}

	var userDetail []model.UserDetail
	for _, user := range users {
		userDetail = append(userDetail, model.UserDetail{
			ID:        user.ID.String(),
			Username:  user.Username,
			Type:      user.Type,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return userDetail, nil
}

func (u *UserRepo) DeleteAll() error {
	dbTx := u.userDB.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Where("1 = 1").Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}

func (u *UserRepo) Login(username, password string) (*model.UserDetail, error) {
	var user entity.User
	if err := u.userDB.Where("username = ? AND password = ?",
		username, password).First(&user).Error; err != nil {
		return nil, err
	}

	userDetail := model.UserDetail{
		ID:        user.ID.String(),
		Username:  user.Username,
		Type:      user.Type,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &userDetail, nil

}