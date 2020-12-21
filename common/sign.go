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
// time   : 2020-12-01 14:14
// version: 1.0.0
// desc   : 

package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func Sign(method, secretKey, src string) string {
	hashed := hmac.New(sha1.New, []byte(secretKey))
	if method == SignMethod.HmacSHA256 {
		hashed = hmac.New(sha256.New, []byte(secretKey))
	}
	hashed.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

func PreSignString(httpMethod, domain, path string, params map[string]interface{}) string {
	var sb strings.Builder
	sb.WriteString(httpMethod)
	sb.WriteString(domain)
	sb.WriteString(path)

	if params != nil {
		sb.WriteString("?")
		keys := make([]string, 0, len(params))
		for key, _ := range params {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for i, key := range keys {
			if i > 0 {
				sb.WriteString("&")
			}
			sb.WriteString(key)
			sb.WriteString("=")
			sb.WriteString(func(val interface{}) string {
				if val == nil {
					return ""
				}
				return fmt.Sprintf("%v", val)
			}(params[key]))
		}
	}

	return sb.String()
}

func Sha256Hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func HmacSha256(src, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(src))
	return string(hashed.Sum(nil))
}
