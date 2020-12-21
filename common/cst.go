// Copyright 2020 yhyzgn gotc
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
// time   : 2020-12-01 14:08
// version: 1.0.0
// desc   : 

package common

const (
	InternalRequestID = "00000000" // 内部错误时请求码
)

var (
	// SDK 信息
	SDK = struct {
		Version string
	}{
		Version: "SDK_GO_1.0.0",
	}

	// 签名方式
	SignMethod = struct {
		HmacSHA1      string
		HmacSHA256    string
		TC3HmacSha256 string
	}{
		HmacSHA1:      "HmacSHA1",        // v1
		HmacSHA256:    "HmacSHA256",      // v1
		TC3HmacSha256: "TC3-HMAC-SHA256", // v3
	}

	// 语言
	Language = struct {
		Chinese string
		English string
	}{
		Chinese: "zh-CN", // 中文
		English: "en-US", // 英文
	}

	// scheme
	Scheme = struct {
		HTTP  string
		HTTPS string
	}{
		HTTP:  "http",
		HTTPS: "https",
	}

	// Content-Type
	ContentType = struct {
		XForm    string
		JSON     string
		FormData string
	}{
		XForm:    "application/x-www-form-urlencoded",
		JSON:     "application/json",
		FormData: "multipart/form-data",
	}
)
