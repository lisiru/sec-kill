package smsutil

import (
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"sec-kill/pkg/options"
)

type SmsRequest struct {
	request sms.SendSmsRequest


}



func SecretKey(options *options.SmsOptions)  {

}

func NewSmsRequest()  {
	request:=sms.NewSendSmsRequest()
}
type RequestOption func(*SmsRequest)

func PhoneNumberSet() RequestOption {
	return func(request *SmsRequest) {
		request.request.PhoneNumberSet=
	}
}
