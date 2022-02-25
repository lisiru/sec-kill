package activity

import (
	"sec-kill/cache"
	"sec-kill/service"
	"sec-kill/store"
)

type ActivityController struct {
	service service.Service

}

func NewActivityController(store store.Factory,cache cache.Factory) *ActivityController  {
	return &ActivityController{
		service: service.NewService(store,cache),
	}
}
