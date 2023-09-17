package repository

import (
	"github.com/jinzhu/gorm"
	"user/domain/model"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	//创建用户
	CreateUser(user *model.User) (int64, error)
	//根据用户ID删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(user *model.User) error
	//查找所有用户
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表

func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

// 根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

// 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

//

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userALl []model.User, srr error) {
	return userALl, u.mysqlDb.Find(&userALl).Error
}
