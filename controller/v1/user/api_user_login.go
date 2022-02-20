package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"sec-kill/pkg/code"
	"sec-kill/pkg/response"
)

type LoginByPhoneCodeParam struct {
	Phone string `json:"phone" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func (us *UserController) LoginByPhoneCode(c *gin.Context)   {
	param:=&LoginByPhoneCodeParam{}
	if err:=c.ShouldBindJSON(param);err!=nil{
		response.WriteResponse(c, errors.WithCode(code.ErrParamNotValid, err.Error()), nil)
		return
	}
	res,err:=us.service.Users().LoginByPhoneCode(c,param.Phone,param.Code)
	if err != nil {
		response.WriteResponse(c,err,nil)
		return
	}
	responseDTO:=ConvertLoginDTO(res)
	response.WriteResponse(c,nil,responseDTO)
	return

}
