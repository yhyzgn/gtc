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
// time   : 2020-12-05 11:19
// version: 1.0.0
// desc   : 

package util

import "testing"

type Flt struct {
	String  string
	Int     int
	Int64   int64
	Uint    uint
	Uint64  uint64
	Float32 float32
	Float64 float64
	Slice   []string
}

func TestFlatParams(t *testing.T) {
	flt := Flt{
		String:  "测试",
		Int:     1,
		Int64:   2,
		Uint:    3,
		Uint64:  4,
		Float32: 5.5,
		Float64: 6.6,
		Slice:   []string{"1", "2"},
	}

	form := FlatParams(flt, true)
	t.Log(form)

	jsn := FlatParams(flt, false)
	t.Log(jsn)
}
