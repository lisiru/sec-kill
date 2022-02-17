package service

import (
	"sec-kill/cache"
	"sec-kill/model"
	"sec-kill/pkg/code"
	"sec-kill/store"
	"context"
	"fmt"
	"github.com/marmotedu/errors"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) (*model.UserList, error)
	SendPhoneCode(ctx context.Context,phone string) bool


}

type userService struct {
	store store.Factory
	cache cache.Factory
}

func newUser(s *service) *userService  {
	return &userService{
		store: s.store,
		cache: s.cache,
	}
}

func (u *userService) Create(ctx context.Context,user *model.User) error  {
	if err:=u.store.Users().Create(ctx,user);err!=nil{
		return errors.WithCode(code.ErrDatabase,err.Error())
	}
	return nil
}
func (u *userService) Get(ctx context.Context,username string) (*model.User,error)  {
	val,_:=u.cache.UserCaches().GetUserCache(ctx,1)
	fmt.Println(val)
	user,err:=u.store.Users().Get(ctx,username)
	if err!=nil{
		return nil,err
	}
	return user,nil
}
func (u *userService) Update(ctx context.Context,user *model.User) error  {
	if err:=u.store.Users().Update(ctx,user);err!=nil{
		return errors.WithCode(code.ErrDatabase,err.Error())
	}
	return nil
}

func (u *userService) List(ctx context.Context) (*model.UserList,error)  {
	users, err :=u.store.Users().List(ctx,100,2)
	if err!=nil{
		return nil,errors.WithCode(code.ErrDatabase,err.Error())
	}
	return users,nil
}

func (u *userService) SendPhoneCode(ctx context.Context,phone string) bool  {

}

