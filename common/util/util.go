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
// time   : 2020-12-02 15:51
// version: 1.0.0
// desc   : 

package util

import "regexp"

// 从 host 分离出 scheme 和 domain
//
// 如
// http://localhost:8080  -->  [scheme = http, domain = localhost:8080]
func HostSplit(host string) (scheme, domain string) {
	if rep, err := regexp.Compile("^(https?)://"); err == nil && rep.MatchString(host) {
		scheme = rep.FindStringSubmatch(host)[1] // [https:// https]
		domain = rep.ReplaceAllString(host, "")
		return
	}
	domain = host
	return
}
