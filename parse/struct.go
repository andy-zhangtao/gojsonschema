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


type GOJSON struct{
	Type string `json:"type"`
	Required []string `json:"required"`
	Properties map[string]GOJSON `json:"properties"`
	Items interface{} `json:"items"`
}

//type GOArrayJSON struct{
//	Items GOJSON `json:"items"`
//}
type GOStruct struct{
	Name string `json:"name"`
	Props []GOStructProps `json:"props"`
}

type GOStructProps struct{
	Name string `json:"name"`
	Vtype string `json:"vtype"`
}