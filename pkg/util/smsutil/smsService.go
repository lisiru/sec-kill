package smsutil

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type SmsClient struct {
	Credential common.Credential
	Region string
	Cpf profile.ClientProfile

}

type Option func(*SmsClient)

func NewSmsClient(options ...func(client *SmsClient)) *SmsClient  {
	client:=&SmsClient{
		Region: "ap-guangzhou",
	}
	for _,option:=range options{
		option(client)
	}
	return client

}

func (s *SmsClient) Send()  {
	sendClient, _ :=sms.NewClient(&s.Credential,s.Region,&s.Cpf)

}



