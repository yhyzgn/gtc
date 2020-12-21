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
// time   : 2020-11-25 15:43
// version: 1.0.0
// desc   : 地域字典表

package common

type Region string

const (
	Bangkok       Region = "ap-bangkok"        // 曼谷
	Beijing       Region = "ap-beijing"        // 北京
	Chengdu       Region = "ap-chengdu"        // 成都
	Chongqing     Region = "ap-chongqing"      // 重庆
	Guangzhou     Region = "ap-guangzhou"      // 广州
	GuangzhouOpen Region = "ap-guangzhou-open" // 广州Open
	HongKong      Region = "ap-hongkong"       // 中国香港
	Mumbai        Region = "ap-mumbai"         // 孟买
	Seoul         Region = "ap-seoul"          // 首尔
	Shanghai      Region = "ap-shanghai"       // 上海
	Nanjing       Region = "ap-nanjing"        // 南京
	ShanghaiFSI   Region = "ap-shanghai-fsi"   // 上海金融
	ShenzhenFSI   Region = "ap-shenzhen-fsi"   // 深圳金融
	Singapore     Region = "ap-singapore"      // 新加坡
	Tokyo         Region = "ap-tokyo"          // 东京
	Frankfurt     Region = "eu-frankfurt"      // 法兰克福
	Moscow        Region = "eu-moscow"         // 莫斯科
	Ashburn       Region = "na-ashburn"        // 阿什本
	SiliconValley Region = "na-siliconvalley"  // 硅谷
	Toronto       Region = "na-toronto"        // 多伦多
)
