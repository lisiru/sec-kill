package user

import "sec-kill/model"

func ConvertLoginDTO(user *model.User) *LoginByPhoneCodeDTO  {
	return &LoginByPhoneCodeDTO{
		Username: user.Username,
		UserId: user.Uid,
		LoginNum: user.LoginNum,

	}
}