package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"sec-kill/model"
	"sec-kill/pkg/code"
	"sec-kill/pkg/response"
	"sec-kill/pkg/util"
)

type CreateActivityParam struct {
	ActivityName string `json:"activity_name" binding:"required"`
	ActivityGoodId uint `json:"activity_good_id" binding:"required"`
	ActivityGoodStock uint `json:"activity_good_stock" binding:"required"`
	ActivityLimitBuy uint `json:"activity_limit_buy" default:"1"`
	ActivityStartTime string `json:"activity_start_time" binding:"required"`
	ActivityEndTime string `json:"activity_end_time" binding:"required"`

}

/**
创建活动接口
 */
func (a *ActivityController) CreateActivity(c *gin.Context)  {
	param:=&CreateActivityParam{}
	if err:=c.ShouldBindJSON(param);err!=nil{
		response.WriteResponse(c, errors.WithCode(code.ErrParamNotValid, err.Error()), nil)
		return
	}

	activityData:=&model.Activity{
		ActivityName: param.ActivityName,
		ActivityGoodId: param.ActivityGoodId,
		ActivityGoodStock: param.ActivityGoodStock,
		StartTime: util.TimeStrToTime(param.ActivityStartTime),
		EndTime: util.TimeStrToTime(param.ActivityEndTime),
		LimitBuy: param.ActivityLimitBuy,

	}
	err:=a.service.Activity().CreateActivity(c,activityData)
	if err != nil {
		response.WriteResponse(c,err,nil)
		return
	}
	response.WriteResponse(c,nil,nil)
	return

}
