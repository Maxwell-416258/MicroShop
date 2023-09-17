package handler

//暴露服务
import (
	"context"
	"user/domain/model"
	"user/domain/service"
	user "user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {
	//将请求的用户信息与用户模型结构体关联
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

// 登录
func (u *User) Login(ctx context.Context, userLogin *user.UserLoginRequest, loginResponse *user.UserLoginResponse) error {
	isOK, err := u.UserDataService.CheckPwd(userLogin.UserName, userLogin.Pwd)
	if err != nil {
		return err
	}
	loginResponse.IsSuccess = isOK
	return nil
}

// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userinfoRequest *user.UserInfoRequest, userinfoResponse *user.UserInfoResponse) error {
	getUser, err := u.UserDataService.FindUserByName(userinfoRequest.UserName)
	if err != nil {
		return err
	}
	userinfoResponse = UserForResponse(getUser)
	return nil
}

// 类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
