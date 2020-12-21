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
// time   : 2020-12-04 15:17
// version: 1.0.0
// desc   : 短信服务

package sms

import (
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/client"
	"github.com/yhyzgn/gtc/common/profile"
	"net/http"
)

type SMS struct {
	profile    *profile.Profile
	credential *common.Credential
	request    *client.Request
	config     *Config
}

type Config struct {
	AppID          string // 短信SdkAppId，在 短信控制台 添加应用后生成的实际SdkAppId
	TemplateID     string // 模板 ID，必须填写已审核通过的模板 ID
	Sign           string // 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名
	CountryCode    int    // 国家码，如 中国：86
	ExtendCode     string // 短信码号扩展号，默认未开通，如需开通请联系 请联系smsHelper
	SessionContext string // 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回
	SenderId       string // 国内短信无senderId，无需填写该项；若需开通国际/港澳台短信senderId，请联系smsHelper
}

func New(region common.Region, credential *common.Credential) *SMS {
	prof := profile.New()
	prof.SignMethod = common.SignMethod.TC3HmacSha256

	req := client.
		NewRequest(region).
		Option(
			client.Service("sms"),
			client.Version("2019-07-11"),
			client.HttpMethod(http.MethodPost),
			client.ContentType(common.ContentType.JSON),
		)

	return &SMS{
		profile:    prof,
		credential: credential,
		request:    req,
	}
}

func (s *SMS) Config(config *Config) *SMS {
	s.config = config
	return s
}

func (s *SMS) Client() *client.Client {
	return client.NewWithRequest(s.request).
		Profile(s.profile).
		Credential(s.credential)
}
