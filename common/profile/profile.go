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
// time   : 2020-12-01 9:48
// version: 1.0.0
// desc   : client profile

package profile

import (
	"github.com/yhyzgn/gtc/common"
	"time"
)

type Profile struct {
	Timeout         time.Duration // http 请求超时时间
	SignMethod      string        // 签名方法，可选【V1(HmacSHA1, HmacSHA256), V3(TC3-HMAC-SHA256)】
	UnsignedPayload bool
	Language        string // 方言，可选【zh-CN, en-US】
	Debug           bool
}

func New() *Profile {
	return &Profile{
		Timeout:         60 * time.Second,
		SignMethod:      common.SignMethod.TC3HmacSha256,
		UnsignedPayload: false,
		Language:        common.Language.Chinese,
		Debug:           true,
	}
}
