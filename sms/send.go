// Copyright 2020 yhyzgn gtc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-12-04 14:48
// version: 1.0.0
// desc   : 短信发送
//
// https://cloud.tencent.com/document/product/382/38778

package sms

import (
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/client"
	"github.com/yhyzgn/gtc/common/errors"
	"strconv"
	"strings"
)

// 单条发送
func (s *SMS) Single(phone string, args ...interface{}) (sr *ResultSend, err error) {
	return s.Send([]string{phone}, args...)
}

// 批量发送
func (s *SMS) Send(phones []string, args ...interface{}) (res *ResultSend, err error) {
	if s.config == nil {
		err = errors.New("SMS.ConfigError", common.InternalRequestID, "Must set 'Config' of SMS at first.")
		return
	}
	if phones == nil || len(phones) == 0 {
		err = errors.New("SMS.PhoneError", common.InternalRequestID, "There is no any phone to send sms.")
		return
	}

	// 默认国家码
	if s.config.CountryCode <= 0 {
		s.config.CountryCode = 86
	}

	// 参数
	mp := map[string]interface{}{
		"SmsSdkAppid": s.config.AppID,
		"TemplateID":  s.config.TemplateID,
		"Sign":        s.config.Sign,
	}
	if s.config.ExtendCode != "" {
		mp["ExtendCode"] = s.config.ExtendCode
	}
	if s.config.SessionContext != "" {
		mp["SessionContext"] = s.config.SessionContext
	}
	if s.config.SenderId != "" {
		mp["SenderId"] = s.config.SenderId
	}

	// 手机号码
	countryCode := strconv.Itoa(s.config.CountryCode)
	mobiles := make([]string, len(phones))
	for i, ph := range phones {
		if !strings.HasPrefix(ph, "+"+countryCode) {
			ph = "+" + countryCode + ph
		}
		mobiles[i] = ph
	}
	// 实测 PhoneNumberSet.N 形式参数报错
	mp["PhoneNumberSet"] = mobiles

	// 模板参数
	// 实测 TemplateParamSet.N 形式参数报错
	mp["TemplateParamSet"] = args

	// 设置 request
	s.request.Option(client.Action("SendSms")).SetBody(mp)

	// 请求
	res = new(ResultSend)
	err = s.Client().Do(res)
	return
}
