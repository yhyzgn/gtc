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
// time   : 2020-12-05 15:09
// version: 1.0.0
// desc   : 

package queue

import "testing"

func TestQueue_SendMessage(t *testing.T) {
	q := New(prof, cred, request)
	res, err := q.SendMessage(&ArgMessageSend{
		QueueName: "dev-gtc-2020-1",
		MsgBody:   "test",
	})

	t.Logf("%+v", res)
	if err != nil {
		t.Error(err)
	}
}
