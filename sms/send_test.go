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
// time   : 2020-12-04 15:50
// version: 1.0.0
// desc   : 

package sms

import (
	"github.com/yhyzgn/gtc/common"
	"os"
	"testing"
)

var (
	secretId  = os.Getenv("TcSecretId")
	secretKey = os.Getenv("TcSecretKey")
)

func TestSMS_Send(t *testing.T) {
	s := New(common.Guangzhou, common.NewCredential(secretId, secretKey)).Config(&Config{
		AppID:       "1400435504",
		TemplateID:  "744696",
		Sign:        "测试",
		CountryCode: 86,
	})

	res, err := s.Send([]string{"18987526232"}, "668822", "6")

	t.Logf("%+v", res)
	t.Error(err)
}
