package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"sec-kill/pkg/code"
	"sec-kill/pkg/response"
)

type SendPhoneCodeParam struct {
	Phone string `json:"phone" binding:"required"`
}

func (us *UserController) SendPhoneCode(c *gin.Context) {
	param := &SendPhoneCodeParam{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.WriteResponse(c, errors.WithCode(code.ErrParamNotValid, err.Error()), nil)
		return
	}
	us.service.Users().SendPhoneCode(c, param.Phone)
	response.WriteResponse(c, nil, nil)
	return

}
