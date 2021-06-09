package util

import (
	"reflect"
	"strconv"
)

type defaultParser interface {
	ParseDefault(string) error
}

func SetDefaults(data interface{}) {
	value := reflect.ValueOf(data).Elem()

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := t.Field(i).Tag
		tagVal := tag.Get("default")
		if len(tagVal) > 0 {
			if field.CanInterface() {
				if parser, ok := field.Interface().(defaultParser);ok{
					if err := parser.ParseDefault(tagVal); err != nil{
						panic(err)
					}
					continue
				}
			}

			if field.CanAddr() && field.Addr().CanInterface(){
				if parse,ok := field.Addr().Interface().(defaultParser); ok {
					if err := parse.ParseDefault(tagVal); err != nil{
						panic(err)
					}
					continue
				}
			}
			switch field.Interface().(type) {
			case string:
				field.SetString(tagVal)

			case int, uint32, int32, int64, uint64:
				i, err := strconv.ParseInt(tagVal, 10, 64)
				if err != nil {
					panic(err)
				}
				field.SetInt(i)

			case float64, float32:
				i, err := strconv.ParseFloat(tagVal, 64)
				if err != nil {
					panic(err)
				}
				field.SetFloat(i)

			case bool:
				field.SetBool(tagVal == "true")

			case []string:
				// We don't do anything with string slices here. Any default
				// we set will be appended to by the XML decoder, so we fill
				// those after decoding.

			default:
				panic(field.Type())
			}
		}else if field.CanSet() && field.Kind() == reflect.Struct && field.CanAddr() {
			if addr := field.Addr(); addr.CanInterface() {
				SetDefaults(addr.Interface())
			}
		}
	}
}
