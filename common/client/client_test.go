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
// time   : 2020-12-02 15:38
// version: 1.0.0
// desc   : 

package client

import (
	"github.com/yhyzgn/gog"
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/profile"
	"net/http"
	"os"
	"testing"
)

var (
	secretId  = os.Getenv("TcSecretId")
	secretKey = os.Getenv("TcSecretKey")
)

var (
	cnt *Client
)

func init() {
	req := NewRequest(common.Guangzhou)

	prof := profile.New()
	prof.SignMethod = common.SignMethod.TC3HmacSha256

	req.Option(
		Service("cvm"),
		Version("2017-03-12"),
		Action("DescribeInstances"),
		HttpMethod(http.MethodPost),
		ContentType(common.ContentType.JSON),
	).
		SetBody(map[string]interface{}{
			"InstanceIds.0": "ins-09dx96dg",
			"Limit":         20,
			"Offset":        0,
		})

	cnt = NewWithRequest(req).
		Profile(prof).
		Secret(secretId, secretKey).
		Request(req)
}

func TestNew(t *testing.T) {
	var r TestR
	err := cnt.Do(&r)

	if err != nil {
		gog.Error(err)
	} else {
		gog.Info(r)
	}
}

type TestR struct {
	Response struct {
		RequestId   string   `json:"RequestId"`
		TotalCount  int      `json:"TotalCount"`
		InstanceSet []string `json:"InstanceSet"`
	} `json:"Response"`
}
