package cache

import "context"

type UserCache interface {
	GetUserCache(ctx context.Context,uid int64) (interface{},error)
	GetSendPhoneCodeFromCache(ctx context.Context,phone string) (interface{},error)
	SetSendPhoneCodeCache(ctx context.Context,phone string,code string) error
}
