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
// time   : 2020-11-25 15:47
// version: 1.0.0
// desc   : SDK错误

package errors

import "fmt"

type TCError struct {
	Code      string
	RequestID string
	Message   string
}

func (er *TCError) Error() string {
	return fmt.Sprintf("[GTC.Error] Code = %s, RequestID = %s, Message = %s", er.Code, er.RequestID, er.Message)
}

func New(code, requestId, message string) *TCError {
	return &TCError{
		Code:      code,
		RequestID: requestId,
		Message:   message,
	}
}
