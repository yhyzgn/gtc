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
// time   : 2020-12-05 9:33
// version: 1.0.0
// desc   : 

package sms

import "github.com/yhyzgn/gtc/common/client"

type ResultSend struct {
	Response struct {
		client.TCR
		SendStatusSet []struct {
			SerialNo       string `json:"SerialNo"`       // 发送流水号
			PhoneNumber    string `json:"PhoneNumber"`    // 手机号码
			Fee            int    `json:"Fee"`            // 计费条数
			SessionContext string `json:"SessionContext"` // 用户Session内容
			Code           string `json:"Code"`           // 短信请求错误码
			Message        string `json:"Message"`        // 短信请求错误码描述
			IsoCode        string `json:"IsoCode"`        // 国家码或地区码，例如CN,US等
		} `json:"SendStatusSet"`
	} `json:"Response"`
}
