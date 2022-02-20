package user

type LoginByPhoneCodeDTO struct {
	UserId uint64 `json:"user_id"`
	Username string `json:"username"`
	LoginNum uint `json:"login_num"`
}