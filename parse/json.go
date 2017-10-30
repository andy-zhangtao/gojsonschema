/*
 *
 *  * Copyright (c) 2017. The gojsonschema Author ztao8607@gmail.com
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *
 *  *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package parse

import (
	"encoding/json"
	"errors"
)

func ParseJsonBytes(name string, jsons []byte)(map[string]GOStruct, map[string]string, error){
	return parseJson(name, string(jsons))
}

func ParseJson(name, jsons string) (map[string]GOStruct, map[string]string, error){
	return parseJson(name, jsons)
}

// ParseJson 解析Json字符串, 转换成GoJsonSchema可以理解的数据结构
// 返回两个MAP.
// map[string]GOStruct  保存解析后的Schema结构体, 结构体名称和包含的属性,属性如果是Object，则需要通过属性名在第二个map中查询对应的结构体名称
// map[string]string 保存每个object所对应的结构体名称
func parseJson(name, jsons string) (map[string]GOStruct, map[string]string, error) {
	var gjson GOJSON
	mapStruct := make(map[string]GOStruct)

	mapTag := make(map[string]string)

	err := json.Unmarshal([]byte(jsons), &gjson)
	if err != nil {
		return nil, nil, err
	}

	//第一次解析,获取最外层的数据结构
	err = parseProps(name, gjson, mapStruct, mapTag)
	if err != nil {
		return nil, nil, err
	}

	//log.Println(mapStruct)
	//for _,value := range mapStruct{
	//	log.Println(value)
	//}
	return mapStruct, mapTag, nil
}

func parseProps(name string, props GOJSON, mapStruct map[string]GOStruct, mapTag map[string]string) (error) {
	var gst GOStruct

	if props.Type == "object" {
		if n, ok := mapTag["NEEDSAVE"]; ok {
			mapTag[n] = name
			delete(mapTag, "NEEDSAVE")
		}
	}
	gst.Name = name
	for _, gj := range props.Required {
		if gj != "" {
			if v, ok := props.Properties[gj]; ok {
				gsp := GOStructProps{
					Name:  gj,
					Vtype: v.Type,
				}
				gst.Props = append(gst.Props, gsp)

				if gsp.Vtype == "object" {
					//	需要继续遍历解析
					if _, ok := mapTag[gj]; !ok {
						mapTag["NEEDSAVE"] = gj
						err := parseProps(name+"_"+gj, v, mapStruct, mapTag)
						if err != nil {
							return err
						}
					}
					continue
				}

				if gsp.Vtype == "array" {
					//	需要继续遍历解析
					if _, ok := mapTag[gj]; !ok {
						data, err := json.Marshal(v.Items)
						if err != nil {
							return err
						}

						var gjson GOJSON

						err = json.Unmarshal(data, &gjson)
						if err != nil {
							return err
						}

						err = parseProps(name+"_"+gj, gjson, mapStruct, mapTag)
						if err != nil {
							return err
						}
					}

					continue
				}
				//	到这里都是boolean, string和number类型的属性了
				//	没必要在遍历
			} else {
				return errors.New(gj + " not find in this struct")
			}
		}

	}

	mapStruct[gst.Name] = gst

	return nil
}
