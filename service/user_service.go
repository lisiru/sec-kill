package service

import (
	"context"
	"fmt"
	"github.com/marmotedu/errors"
	"math/rand"
	"sec-kill/cache"
	"sec-kill/global"
	"sec-kill/model"
	"sec-kill/pkg/code"
	"sec-kill/pkg/util/smsutil/tencenSms"
	"sec-kill/store"
	"time"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) (*model.UserList, error)
	SendPhoneCode(ctx context.Context, phone string) bool
	LoginByPhoneCode(ctx context.Context, phone string, code string) (*model.User, error)
}

type userService struct {
	store store.Factory
	cache cache.Factory
}

func (u *userService) LoginByPhoneCode(ctx context.Context, phone string, phoneCode string) (*model.User, error) {
	// 从缓存中获取该手机号对应的验证码是否匹配
	cacheCode, err := u.cache.UserCaches().GetSendPhoneCodeFromCache(ctx, phone)
	if err != nil {
		return nil, errors.WithCode(code.ErrUserPhoneCodeExpire, err.Error())
	}
	if cacheCode != phoneCode {
		return nil, errors.WithCode(code.ErrUserPhoneCodeMiss, "")
	}
	// 查询当前手机号是否在数据库中存在，若存在则更新登录次数，若不存在则生成一个新用户并插入到数据库中
	user, err := u.store.Users().GetUserByWhere(ctx, map[string]interface{}{"phone": phone})
	if err != nil {
		// 生成uuid
		//node, err := snowflake.NewNode(1)
		//if err != nil {
		//	return nil, err
		//}
		//uid := uint64(node.Generate().Int64())
		uid:=rand.Uint64()
		newUser := &model.User{
			Uid:      uid,
			Phone:    phone,
			LoginNum: 1,
		}

		err = u.store.Users().Create(ctx, newUser)
		if err != nil {
			return nil, errors.WithCode(code.ErrUserLoginFail, err.Error())
		}
		return newUser, nil
	}
	if user != nil {
		user.LoginNum += 1
		_ = u.store.Users().Update(ctx, user)
		return user, nil
	}
	return nil, nil

}

func newUser(s *service) *userService {
	return &userService{
		store: s.store,
		cache: s.cache,
	}
}

func (u *userService) Create(ctx context.Context, user *model.User) error {
	if err := u.store.Users().Create(ctx, user); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}
func (u *userService) Get(ctx context.Context, username string) (*model.User, error) {
	val, _ := u.cache.UserCaches().GetUserCache(ctx, 1)
	fmt.Println(val)
	user, err := u.store.Users().Get(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userService) Update(ctx context.Context, user *model.User) error {
	if err := u.store.Users().Update(ctx, user); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userService) List(ctx context.Context) (*model.UserList, error) {
	users, err := u.store.Users().List(ctx, 100, 2)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return users, nil
}

func (u *userService) SendPhoneCode(ctx context.Context, phone string) bool {

	// 获取配置参数
	smsSetting := global.TencenSmsSetting
	phoneSet := []string{phone}
	// 随机生成6位的验证码
	var randCode string = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	templateSet := []string{randCode, "60"}
	smsRequest := tencenSms.NewSmsRequest(smsSetting, tencenSms.WithPhoneNumberSet(phoneSet), tencenSms.WithTemplateParamSet(templateSet))
	smsClient := tencenSms.NewSmsClient(tencenSms.WithRequest(*smsRequest), tencenSms.WithCredential(*smsSetting))
	go smsClient.Send()
	// 将验证码和手机号保存到redis中
	_ = u.cache.UserCaches().SetSendPhoneCodeCache(ctx, phone, randCode)
	return true

}
