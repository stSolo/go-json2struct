package writers

import (
	"fmt"
	"reflect"

	"github.com/stSolo/go-json2struct/app/helpers"
)

func StringStruct(m map[string]interface{}, temp []string, num *int) []string {
	str := fmt.Sprintf("type Struct%d struct{\n", *num)
	*num++
	for k, v := range m {
		t := reflect.TypeOf(v)
		n := t.Name()
		if n != "" {
			str += fmt.Sprintf("%s %v `json:\"%s\"`\n", helpers.Normalize(k), t, k)
			continue
		}
		kind := t.Kind()
		switch kind {
		case reflect.Slice:
			val := reflect.ValueOf(v)
			elemType := val.Index(0).Elem().Type().Name()
			if elemType != "" {
				str += fmt.Sprintf("%s []%v `json:\"%s\"`\n", helpers.Normalize(k), val.Index(0).Elem().Kind(), k)
				continue
			}
			str += fmt.Sprintf("%s []Struct%d `json:\"%s\"`\n", helpers.Normalize(k), *num, k)
			t1 := StringStruct(val.Index(0).Elem().Interface().(map[string]interface{}), temp[len(temp):], num)
			temp = append(temp, t1...)
		case reflect.Map:
			val := reflect.ValueOf(v)
			str += fmt.Sprintf("%s Struct%d `json:\"%s\"`\n", helpers.Normalize(k), *num, k)
			t1 := StringStruct(val.Interface().(map[string]interface{}), temp[len(temp):], num)
			temp = append(temp, t1...)
		}
	}
	str += "}\n"
	temp = append(temp, str)
	return temp
}
