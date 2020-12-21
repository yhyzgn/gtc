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
// time   : 2020-12-05 9:33
// version: 1.0.0
// desc   : 

package queue

import (
	"github.com/yhyzgn/gtc/cmq/com"
	"github.com/yhyzgn/gtc/common/client"
)

// 创建结果
type ResultCreate struct {
	Response struct {
		client.TCR
		QueueId string `json:"QueueId"` // 消息队列ID
	} `json:"Response"`
}

// 删除结果
type ResultDelete struct {
	Response struct {
		client.TCR
	} `json:"Response"`
}

// 查询结果
type ResultList struct {
	Response struct {
		client.TCR
		TotalCount int        `json:"TotalCount"` // 总队列数
		QueueSet   []QueueSet `json:"QueueSet"`   // 队列详情列表。
	} `json:"Response"`
}

// 发送消息结果
type ResultMessageSend struct {
	Response struct {
		client.TCR
		MsgId string `json:"MsgId"` // 服务器生成消息的唯一标识 ID
	} `json:"Response"`
}

type QueueSet struct {
	QueueId             string             `json:"QueueId"`             // 消息队列ID
	QueueName           string             `json:"QueueName"`           // 消息队列名字
	Qps                 int                `json:"Qps"`                 // 每秒钟生产消息条数的限制，消费消息的大小是该值的1.1倍，此字段可能返回 null，表示取不到有效值
	Bps                 int                `json:"Bps"`                 // 带宽限制，此字段可能返回 null，表示取不到有效值
	MaxDelaySeconds     int                `json:"MaxDelaySeconds"`     // 飞行消息最大保留时间，此字段可能返回 null，表示取不到有效值
	MaxMsgHeapNum       int                `json:"MaxMsgHeapNum"`       // 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000，此字段可能返回 null，表示取不到有效值
	PollingWaitSeconds  int                `json:"PollingWaitSeconds"`  // 消息接收长轮询等待时间。取值范围0 - 30秒，默认值0，此字段可能返回 null，表示取不到有效值
	MsgRetentionSeconds int                `json:"MsgRetentionSeconds"` // 消息保留周期。取值范围60-1296000秒（1min-15天），默认值345600秒（4 天），此字段可能返回 null，表示取不到有效值
	VisibilityTimeout   int                `json:"VisibilityTimeout"`   // 消息可见性超时。取值范围1 - 43200秒（即12小时内），默认值30，此字段可能返回 null，表示取不到有效值
	MaxMsgSize          int                `json:"MaxMsgSize"`          // 消息最大长度。取值范围1024 - 1048576 Byte（即1K - 1024K），默认值65536，此字段可能返回 null，表示取不到有效值
	RewindSeconds       int                `json:"RewindSeconds"`       // 回溯队列的消息回溯时间最大值，取值范围0 - 43200秒，0表示不开启消息回溯，此字段可能返回 null，表示取不到有效值
	CreateTime          int                `json:"CreateTime"`          // 队列的创建时间。返回 Unix 时间戳，精确到秒，此字段可能返回 null，表示取不到有效值
	LastModifyTime      int                `json:"LastModifyTime"`      // 最后一次修改队列属性的时间。返回 Unix 时间戳，精确到秒，此字段可能返回 null，表示取不到有效值
	ActiveMsgNum        int                `json:"ActiveMsgNum"`        // 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值，此字段可能返回 null，表示取不到有效值
	InactiveMsgNum      int                `json:"InactiveMsgNum"`      // 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值，此字段可能返回 null，表示取不到有效值
	DelayMsgNum         int                `json:"DelayMsgNum"`         // 延迟消息数，此字段可能返回 null，表示取不到有效值
	RewindMsgNum        int                `json:"RewindMsgNum"`        // 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量，此字段可能返回 null，表示取不到有效值
	MinMsgTime          int                `json:"MinMsgTime"`          // 消息最小未消费时间，单位为秒，此字段可能返回 null，表示取不到有效值
	Transaction         bool               `json:"Transaction"`         // 事务消息队列。true表示是事务消息，false表示不是事务消息，此字段可能返回 null，表示取不到有效值
	CreateUin           int                `json:"CreateUin"`           // 创建者Uin，此字段可能返回 null，表示取不到有效值
	Trace               bool               `json:"Trace"`               // 消息轨迹。true表示开启，false表示不开启，此字段可能返回 null，表示取不到有效值
	DeadLetterSource    []DeadLetterSource `json:"DeadLetterSource"`    // 死信队列，此字段可能返回 null，表示取不到有效值
	DeadLetterPolicy    DeadLetterPolicy   `json:"DeadLetterPolicy"`    // 死信队列策略，此字段可能返回 null，表示取不到有效值
	TransactionPolicy   TransactionPolicy  `json:"TransactionPolicy"`   // 事务消息策略，此字段可能返回 null，表示取不到有效值
	Tags                []com.Tag          `json:"Tags"`                // 关联的标签，此字段可能返回 null，表示取不到有效值
}

type DeadLetterSource struct {
	QueueId   string `json:"QueueId"`   // 消息队列ID，此字段可能返回 null，表示取不到有效值
	QueueName string `json:"QueueName"` // 消息队列名字，此字段可能返回 null，表示取不到有效值
}

type DeadLetterPolicy struct {
	DeadLetterQueueName string `json:"DeadLetterQueueName"` // 死信队列名字，此字段可能返回 null，表示取不到有效值
	DeadLetterQueue     string `json:"DeadLetterQueue"`     // 死信队列，此字段可能返回 null，表示取不到有效值
	Policy              int    `json:"Policy"`              // 死信队列策略，此字段可能返回 null，表示取不到有效值
	MaxTimeToLive       int    `json:"MaxTimeToLive"`       // 最大未消费过期时间。Policy为1时必选。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds，此字段可能返回 null，表示取不到有效值
	MaxReceiveCount     int    `json:"MaxReceiveCount"`     // 最大接收次数，此字段可能返回 null，表示取不到有效值
}

type TransactionPolicy struct {
	FirstQueryInterval int `json:"FirstQueryInterval"` // 第一次回查时间，此字段可能返回 null，表示取不到有效值
	MaxQueryCount      int `json:"MaxQueryCount"`      // 最大查询次数，此字段可能返回 null，表示取不到有效值
}
