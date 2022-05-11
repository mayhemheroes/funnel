package convert

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

func OpenApi2Proto(src interface{}, dst proto.Message) {
	fmt.Printf("%T\n", src)
	fmt.Printf("%T\n", dst)

	sv := reflect.ValueOf(src)
	dv := reflect.ValueOf(dst).Elem()
	for i := 0; i < sv.NumField(); i++ {
		sf := sv.Type().Field(i)
		fmt.Printf("%s\n", sf.Name)
		if df, ok := dv.Type().FieldByName(sf.Name); ok {
			fmt.Printf("%s -> %s\n", sf.Name, df.Name)
			switch df.Type.Kind() {
			case reflect.String, reflect.Int, reflect.Uint, reflect.Uint32, reflect.Int32:
				fmt.Printf("Found string\n")
				dv.FieldByIndex(df.Index).Set(sv.FieldByIndex(sf.Index))
			case reflect.Pointer:
				fmt.Printf("Found a pointer\n")
				dst := reflect.New(df.Type.Elem())
				v := dst.Interface()
				nv := v.(proto.Message)
				OpenApi2Proto(sv.FieldByIndex(sf.Index).Interface(), nv)
				dv.FieldByIndex(df.Index).Set(dst)
			case reflect.Struct:
				fmt.Printf("Found a struct\n")
			default:
				fmt.Printf("Found unknown: %s\n", df.Type.Kind())
			}
		} else {
			fmt.Printf("Missing %s\n", sf.Name)
		}
	}
}
