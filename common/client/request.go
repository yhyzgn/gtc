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
// time   : 2020-11-25 16:26
// version: 1.0.0
// desc   : 腾讯云接口专用 request 封装对象

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yhyzgn/gtc/common"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// 适配器模式
type Request struct {
	region      string
	scheme      string
	service     string
	rootDomain  string
	path        string
	version     string
	domain      string
	httpMethod  string
	contentType string
	action      string
	params      map[string]interface{}
	body        interface{}
}

func NewRequest(region common.Region) *Request {
	return &Request{
		region:      string(region),
		scheme:      common.Scheme.HTTPS,
		rootDomain:  "tencentcloudapi.com",
		path:        "/",
		httpMethod:  http.MethodPost,
		contentType: common.ContentType.JSON,
		params:      make(map[string]interface{}),
	}
}

func (r *Request) Option(opts ...Optional) *Request {
	for _, opt := range opts {
		opt.apply(r)
	}
	return r
}

func (r *Request) GetScheme() string {
	return r.scheme
}

func (r *Request) GetRootDomain() string {
	return r.rootDomain
}

func (r *Request) GetPath() string {
	return r.path
}

func (r *Request) GetVersion() string {
	return r.version
}

func (r *Request) GetService() string {
	return r.service
}

func (r *Request) GetRegion() string {
	return r.region
}

func (r *Request) GetDomain() string {
	if r.domain != "" {
		return r.domain
	}

	var sb strings.Builder
	if r.GetService() != "" {
		sb.WriteString(r.GetService())
		sb.WriteString(".")
	}
	if r.GetRegion() != "" {
		sb.WriteString(r.GetRegion())
		sb.WriteString(".")
	}
	sb.WriteString(r.GetRootDomain())
	return sb.String()
}

func (r *Request) GetURI() string {
	return r.GetScheme() + "://" + r.GetDomain()
}

func (r *Request) GetEndpoint() string {
	endpoint := r.GetURI()
	if r.GetPath() != "" {
		endpoint += r.GetPath()
	}

	// 为 GET 方法拼接参数
	if r.GetHttpMethod() == http.MethodGet {
		params := QueryParams(r.GetParams())
		if params != "" {
			if !strings.Contains(endpoint, "?") {
				endpoint += "?" + params
			} else {
				if strings.HasSuffix(endpoint, "?") {
					endpoint += params
				} else {
					endpoint += "&" + params
				}
			}
		}
	}

	return endpoint
}

func (r *Request) GetAction() string {
	return r.action
}

func (r *Request) GetHttpMethod() string {
	return r.httpMethod
}

func (r *Request) GetContentType() string {
	return r.contentType
}

func (r *Request) IsContentType(contentType string) bool {
	return strings.HasPrefix(r.contentType, contentType)
}

func (r *Request) GetParams() map[string]interface{} {
	return r.params
}

func (r *Request) SetParam(name string, value interface{}) *Request {
	r.params[name] = value
	return r
}

func (r *Request) SetParams(params map[string]interface{}) *Request {
	for name, value := range params {
		r.params[name] = value
	}
	return r
}

func (r *Request) GetBody() interface{} {
	return r.body
}

func (r *Request) SetBody(body interface{}) *Request {
	r.body = body
	return r
}

func (r *Request) GetReader() io.Reader {
	if r.httpMethod == http.MethodGet {
		return nil
	}
	if r.httpMethod == http.MethodPost {
		if r.IsContentType(common.ContentType.JSON) {
			// application/json
			if r.GetBody() != nil {
				bs, err := json.Marshal(r.GetBody())
				if err != nil {
					panic(err)
				}
				return bytes.NewBuffer(bs)
			}
			return bytes.NewBuffer([]byte{})
		} else {
			// application/x-www-form-urlencoded 或者 multipart/form-data
			params := QueryParams(r.GetParams())
			return strings.NewReader(params)
		}
	}
	panic("unsupported http method " + r.httpMethod)
}

func QueryParams(params map[string]interface{}) string {
	if params == nil {
		return ""
	}

	// 参数名称字典排序
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 重组参数
	values := url.Values{}
	for _, key := range keys {
		values[key] = []string{func(val interface{}) string {
			if val == nil {
				return ""
			}
			return fmt.Sprintf("%v", val)
		}(params[key])}
	}
	return values.Encode()
}
