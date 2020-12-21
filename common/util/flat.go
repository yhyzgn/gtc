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
// time   : 2020-12-05 10:06
// version: 1.0.0
// desc   : 

package util

import (
	"fmt"
	"reflect"
	"strconv"
)

func FlatParams(arg interface{}, isForm bool) map[string]interface{} {
	if arg == nil {
		return nil
	}
	value := reflect.ValueOf(arg)
	if value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return flat(value, isForm, "")
}

func flat(value reflect.Value, isForm bool, prefix string) map[string]interface{} {
	tp := value.Type()
	params := make(map[string]interface{}, 0)

	for i := 0; i < tp.NumField(); i++ {
		tpf := tp.Field(i)
		tag := tpf.Tag
		name, has := tag.Lookup("name")
		if !has {
			name = tpf.Name
		}
		stringify, has := tag.Lookup("stringify")
		toString := false
		if has && stringify != "" {
			toString, _ = strconv.ParseBool(stringify)
		}

		field := value.Field(i)
		kind := field.Kind()

		if kind == reflect.Ptr && field.IsNil() {
			continue
		}

		// 指针类型
		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}

		key := prefix + name

		switch kind {
		case reflect.String:
			if val := field.String(); val != "" {
				params[key] = val
			}
			break
		case reflect.Bool:
			val := field.Bool()
			if toString {
				params[key] = strconv.FormatBool(val)
			} else {
				params[key] = val
			}
			break
		case reflect.Int, reflect.Int64:
			if val := field.Int(); val != 0 {
				if toString {
					params[key] = strconv.FormatInt(val, 10)
				} else {
					params[key] = val
				}
			}
			break
		case reflect.Uint, reflect.Uint64:
			if val := field.Uint(); val != 0 {
				if toString {
					params[key] = strconv.FormatUint(val, 10)
				} else {
					params[key] = val
				}
			}
			break
		case reflect.Float32:
			if val := field.Float(); val != 0 {
				if toString {
					params[key] = strconv.FormatFloat(val, 'f', -1, 32)
				} else {
					params[key] = val
				}
			}
			break
		case reflect.Float64:
			if val := field.Float(); val != 0 {
				if toString {
					params[key] = strconv.FormatFloat(val, 'f', -1, 64)
				} else {
					params[key] = val
				}
			}
			break
		case reflect.Slice:
			slc := value.Field(i)
			if slc.Interface() == nil || slc.Len() == 0 {
				break
			}

			if isForm {
				// form 形式
				for j := 0; j < slc.Len(); j++ {
					vj := slc.Index(j)
					slcKey := fmt.Sprintf("%s.%d", key, j)
					kind = vj.Kind()

					if kind == reflect.Ptr && field.IsNil() {
						continue
					}

					if kind == reflect.Ptr {
						vj = vj.Elem()
						kind = vj.Kind()
					}

					switch kind {
					case reflect.String:
						if val := vj.String(); val != "" {
							params[slcKey] = val
						}
						break
					case reflect.Bool:
						val := vj.Bool()
						if toString {
							params[key] = strconv.FormatBool(val)
						} else {
							params[key] = val
						}
						break
					case reflect.Int, reflect.Int64:
						if val := vj.Int(); val != 0 {
							if toString {
								params[slcKey] = strconv.FormatInt(val, 10)
							} else {
								params[key] = val
							}
						}
						break
					case reflect.Uint, reflect.Uint64:
						if val := vj.Uint(); val != 0 {
							if toString {
								params[slcKey] = strconv.FormatUint(val, 10)
							} else {
								params[key] = val
							}
						}
						break
					case reflect.Float32:
						if val := vj.Float(); val != 0 {
							if toString {
								params[slcKey] = strconv.FormatFloat(val, 'f', -1, 32)
							} else {
								params[key] = val
							}
						}
						break
					case reflect.Float64:
						if val := vj.Float(); val != 0 {
							if toString {
								params[slcKey] = strconv.FormatFloat(val, 'f', -1, 64)
							} else {
								params[key] = val
							}
						}
						break
					default:
						for k, v := range flat(vj, isForm, slcKey+".") {
							params[k] = v
						}
					}
				}
			} else {
				// json 形式，直接赋值
				params[name] = field.Interface()
			}
		default:
			for k, v := range flat(reflect.ValueOf(field.Interface()), isForm, key+".") {
				params[k] = v
			}
		}
	}

	return params
}
