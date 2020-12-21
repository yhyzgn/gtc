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
// time   : 2020-12-05 15:02
// version: 1.0.0
// desc   : 消息

package queue

import (
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/client"
	"github.com/yhyzgn/gtc/common/errors"
	"github.com/yhyzgn/gtc/common/util"
)

func (q *Queue) SendMessage(arg *ArgMessageSend) (res *ResultMessageSend, err error) {
	if arg == nil {
		err = errors.New("CMQ.Queue.ArgMessageSendError", common.InternalRequestID, "Must set argument 'ArgMessageSend' of Queue.SendMessage() at first.")
		return
	}
	if arg.QueueName == "" {
		err = errors.New("CMQ.Queue.NameError", common.InternalRequestID, "The queue name can not be empty.")
		return
	}

	// 设置 request
	// 重新组装参数
	q.request.Option(client.Action("SendMessage")).SetBody(util.FlatParams(arg, false))

	// 请求
	res = new(ResultMessageSend)
	err = q.Client().Do(res)
	return
}
