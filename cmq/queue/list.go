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
// time   : 2020-12-05 12:20
// version: 1.0.0
// desc   : 枚举队列
//
// https://cloud.tencent.com/document/product/406/42624

package queue

import (
	"github.com/yhyzgn/gtc/common/client"
	"github.com/yhyzgn/gtc/common/util"
)

func (q *Queue) All() (res *ResultList, err error) {
	return q.List(nil)
}

func (q *Queue) List(arg *ArgList) (res *ResultList, err error) {
	// 设置 request
	q.request.Option(client.Action("DescribeQueueDetail"))

	if arg != nil {
		// 重新组装参数
		q.request.SetBody(util.FlatParams(arg, false))
	}

	// 请求
	res = new(ResultList)
	err = q.Client().Do(res)
	return
}
