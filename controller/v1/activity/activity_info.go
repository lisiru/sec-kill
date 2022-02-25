package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"sec-kill/pkg/code"
	"sec-kill/pkg/response"
)

type ActivityInfoParam struct {
	ActivityId uint `form:"activity_id" binding:"required"`
}


// 获取活动的信息
func (a *ActivityController) ActivityInfo(c *gin.Context)  {
	param:=ActivityInfoParam{}
	if err:=c.ShouldBindQuery(&param);err!=nil{
		response.WriteResponse(c,errors.WithCode(code.ErrParamNotValid,err.Error()),nil)
		return
	}
	where:=map[string]interface{}{"id":param.ActivityId}

	info,err:=a.service.Activity().GetActivityInfo(c,where)
	if err != nil {
		response.WriteResponse(c,err,nil)
		return
	}
	res:=convertActivityInfoDTO(info)
	response.WriteResponse(c,nil,res)
	return





}
