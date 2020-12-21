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
// time   : 2020-12-05 12:27
// version: 1.0.0
// desc   : 

package com

type Filter struct {
	Name   string   // 过滤参数的名字
	Values []string // 数值
}

type Tag struct {
	TagKey   string // 标签Key，此字段可能返回 null，表示取不到有效值。
	TagValue string // 标签值，此字段可能返回 null，表示取不到有效值。，此字段可能返回 null，表示取不到有效值。
}
