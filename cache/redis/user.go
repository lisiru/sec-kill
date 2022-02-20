package redis

import (
	"sec-kill/cache"
	"context"
	"fmt"
	"time"
)

type users struct {
	client *redisStore
}

func genSendPhoneCodeKey(phone string) string  {
	return fmt.Sprintf(cache.UserCacheSendPhoneCodeKey,phone)
}

func (u *users) SetSendPhoneCodeCache(ctx context.Context, phone string, code string) error {
	key:=genSendPhoneCodeKey(phone)
	return u.client.Set(ctx,key,code,2000*time.Second)
}

func (u *users) GetSendPhoneCodeFromCache(ctx context.Context, phone string) (interface{}, error) {
	key:=genSendPhoneCodeKey(phone)
	return u.client.Get(ctx,key)
}

func (u *users) GetUserCache(ctx context.Context,uid int64) (interface{},error)  {
	key:=fmt.Sprintf(cache.UserCacheInfoKey,uid)
	return u.client.Get(ctx,key)
}

func NewUsers(ch *redisStore) *users {
	return &users{
		ch,
	}
}
