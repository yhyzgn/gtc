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
// time   : 2020-12-05 12:16
// version: 1.0.0
// desc   : 

package queue

import (
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/client"
	"github.com/yhyzgn/gtc/common/profile"
	"net/http"
	"os"
)

var (
	cred    = common.NewCredential(os.Getenv("TcSecretId"), os.Getenv("TcSecretKey"))
	prof    *profile.Profile
	request *client.Request
)

func init() {
	prof = profile.New()
	prof.SignMethod = common.SignMethod.TC3HmacSha256

	request = client.
		NewRequest(common.Guangzhou).
		Option(
			client.Service("cmq"),
			client.Version("2019-03-04"),
			client.HttpMethod(http.MethodPost),
			client.ContentType(common.ContentType.JSON),
		)
}
