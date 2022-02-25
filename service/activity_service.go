package service

import (
	"context"
	"github.com/marmotedu/errors"
	"sec-kill/cache"
	"sec-kill/model"
	"sec-kill/pkg/code"
	"sec-kill/store"
)

type ActivityService interface {
	CreateActivity(ctx context.Context,activity *model.Activity) error
	GetActivityInfo(ctx context.Context,where map[string]interface{}) (*model.Activity,error)
}

type activityService struct {
	store store.Factory
	cache cache.Factory
}

// 创建活动
func (a activityService) CreateActivity(ctx context.Context, activity *model.Activity)  error {

	err:=a.store.Activity().Create(ctx,activity)
	if err != nil {
		return errors.WithCode(code.ErrCreateActivityFail,err.Error())
	}
	return nil

}

// 获取活动的信息
func (a activityService) GetActivityInfo(ctx context.Context,where map[string]interface{}) (*model.Activity,error)  {
	activityInfo,err:=a.store.Activity().GetActivityByWhere(ctx,where)
	if err!=nil {
		return nil, err
	}
	return activityInfo,nil

}



func NewActivityService(s *service) *activityService  {
	return &activityService{
		store: s.store,
		cache: s.cache,
	}
}
