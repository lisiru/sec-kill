package redis

import (
	"sec-kill/cache"
	"context"
	"fmt"
)

type users struct {
	client *redisStore
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
