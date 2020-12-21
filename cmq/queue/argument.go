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
// time   : 2020-12-05 9:36
// version: 1.0.0
// desc   : 

package queue

import "github.com/yhyzgn/gtc/cmq/com"

// 创建队列参数
type ArgCreate struct {
	QueueName           string // 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	MaxMsgHeapNum       int    // 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	PollingWaitSeconds  int    // 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	VisibilityTimeout   int    // 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	MaxMsgSize          int    // 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MsgRetentionSeconds int    // 消息保留周期。取值范围 60-1296000 秒（1min-15天），默认值 345600 (4 天)。
	RewindSeconds       int    // 队列是否开启回溯消息能力，该参数取值范围0-msgRetentionSeconds,即最大的回溯时间为消息在队列中的保留周期，0表示不开启。
	Transaction         int    // 1 表示事务队列，0 表示普通队列
	FirstQueryInterval  int    // 第一次回查间隔
	MaxQueryCount       int    // 最大回查次数
	DeadLetterQueueName string // 死信队列名称
	Policy              int    // 死信策略。0为消息被多次消费未删除，1为Time-To-Live过期
	MaxReceiveCount     int    // 最大接收次数 1-1000
	MaxTimeToLive       int    // policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间msgRetentionSeconds
	Trace               bool   // 是否开启消息轨迹追踪，当不设置字段时，默认为不开启，该字段为true表示开启，为false表示不开启
}

// 删除队列参数
type ArgDelete struct {
	QueueName string // 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
}

type ArgList struct {
	QueueName string       // 精确匹配QueueName
	Offset    int          // 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Limit     int          // 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	TagKey    int          // 标签搜索
	Filters   []com.Filter // 筛选参数，目前支持QueueName筛选，且仅支持一个关键字
}

//
//------------------------------------------------------------------------------------- message -------------------------------------------------------------------------------------
//

// 发送消息参数
type ArgMessageSend struct {
	QueueName    string // 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	MsgBody      string // 消息正文。至少1Byte，最大长度受限于设置的队列消息最大长度属性。
	DelaySeconds string // 单位为秒，表示该消息发送到队列后，需要延时多久用户才可见该消息。
}
