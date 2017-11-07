package generation

import (
	"github.com/andy-zhangtao/gojsonschema/parse"
	"strings"
)

// Generate 生成Schema对应的Struct.go文件
func Generate(name string, mapStruct map[string]parse.GOStruct, mapTag map[string]string) string {

	data := ""
	for _, value := range mapStruct {
		data += generate(value, mapTag)
	}

	return data
}

func generate(gos parse.GOStruct, mapTag map[string]string) string {
	data := "type " + gos.Name + " struct{ " + "\n"
	for _, props := range gos.Props {
		if strings.Contains(props.Name, "-") {
			props.Name = strings.Replace(props.Name, "-", "_", -1)
		}
		data += "    " + strings.Title(props.Name) + " "
		if props.Vtype == "object" {
			data += mapTag[props.Name]
		} else if props.Vtype == "array" {
			data += "[]" + gos.Name + "_" + props.Name
		} else if props.Vtype == "number" {
			data += "int"
		} else if props.Vtype == "boolean" {
			data += "bool"
		} else {
			data += strings.ToLower(props.Vtype)
		}

		data += " "
		data += "`json:\"" + strings.ToLower(props.Name) + "\"`"
		data += "\n"
	}

	data += "} \n"
	return data
}
