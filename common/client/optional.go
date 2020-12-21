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
// time   : 2020-12-02 21:19
// version: 1.0.0
// desc   :

package client

import (
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/util"
	"net/http"
	"strings"
)

type Optional interface {
	apply(*Request)
}

type optional func(req *Request)

func (opt optional) apply(req *Request) {
	opt(req)
}

func Region(region string) Optional {
	return optional(func(req *Request) {
		req.region = region
	})
}

func Scheme(scheme string) Optional {
	return optional(func(req *Request) {
		scheme = strings.ToLower(scheme)
		if scheme != common.Scheme.HTTPS {
			scheme = common.Scheme.HTTP
		}
		req.scheme = scheme
	})
}

func Service(service string) Optional {
	return optional(func(req *Request) {
		req.service = service
	})
}

func RootDomain(rootDomain string) Optional {
	return optional(func(req *Request) {
		req.rootDomain = rootDomain
	})
}

func Path(path string) Optional {
	return optional(func(req *Request) {
		req.path = path
	})
}

func Version(version string) Optional {
	return optional(func(req *Request) {
		req.version = version
	})
}

func Domain(domain string) Optional {
	return optional(func(req *Request) {
		// 这里的 domain 参数可能包含 scheme，需要将其分离
		scheme, domain := util.HostSplit(domain)
		if scheme != "" {
			req.scheme = scheme
		}
		if domain != "" {
			req.domain = domain
		}
	})
}

func HttpMethod(method string) Optional {
	return optional(func(req *Request) {
		method = strings.ToUpper(method)
		if method != http.MethodPost {
			method = http.MethodGet
		}
		if method == http.MethodGet {
			req.contentType = common.ContentType.XForm
		}
		req.httpMethod = method
	})
}

func ContentType(contentType string) Optional {
	return optional(func(req *Request) {
		req.contentType = contentType
	})
}

func Action(action string) Optional {
	return optional(func(req *Request) {
		req.action = action
	})
}

func Params(params map[string]interface{}) Optional {
	return optional(func(req *Request) {
		req.params = params
	})
}

func Body(body interface{}) Optional {
	return optional(func(req *Request) {
		req.body = body
	})
}
