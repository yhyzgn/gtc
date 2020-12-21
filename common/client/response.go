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
// time   : 2020-12-01 10:52
// version: 1.0.0
// desc   : 

package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

type Response struct {
}

func (r *Response) Bytes(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return bs, nil
}

func (r *Response) Decoder(bs []byte) *json.Decoder {
	return json.NewDecoder(bytes.NewBuffer(bs))
}

func (r *Response) Decode(body io.ReadCloser, value interface{}) error {
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return r.Decoder(bs).Decode(value)
}

func NewResponse() *Response {
	return new(Response)
}