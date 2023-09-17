package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(username string, pwd string) (isOk bool, err error)
}

func NewUserDataService(userReponsitory repository.IUserRepository) IUserDataService {
	return &UserDataService{userReponsitory}
}

type UserDataService struct {
	UserResponsitory repository.IUserRepository
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

// 插入用户
func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte) //写入数据库时把密码加密再写入
	return u.UserResponsitory.CreateUser(user)
}

// 删除用户
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserResponsitory.DeleteUserByID(userID)
}

// 更新用户
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	//判断是否更新了密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	//log
	return u.UserResponsitory.UpdateUser(user)
}

// 查找用户
func (u *UserDataService) FindUserByName(userName string) (user *model.User, err error) {
	return u.UserResponsitory.FindUserByName(userName)
}

// 比对用户名和密码
func (u *UserDataService) CheckPwd(UserName string, pwd string) (isOk bool, err error) {
	user, err := u.UserResponsitory.FindUserByName(UserName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
