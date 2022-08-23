package tes

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/ohsu-comp-bio/funnel/tes/openapi"
)

func Proto2OpenApi(src proto.Message, dst interface{}) {

	sv := reflect.ValueOf(src).Elem()
	dv := reflect.ValueOf(dst).Elem()

	for i := 0; i < sv.NumField(); i++ {
		sf := sv.Type().Field(i)
		if df, ok := dv.Type().FieldByName(sf.Name); ok {
			switch df.Type.Kind() {
			case reflect.String, reflect.Int, reflect.Uint, reflect.Uint32, reflect.Int32, reflect.Float32, reflect.Float64:
				//fmt.Printf("Found string %s\n", sf.Name)
				if sf.Type == reflect.TypeOf(State_UNKNOWN) {
					state := sv.FieldByIndex(sf.Index).Interface().(State)
					dv.FieldByIndex(df.Index).Set(reflect.ValueOf(openapi.TesState(state.String())))
				} else if df.Type == reflect.TypeOf(FileType_FILE) {
					state := sv.FieldByIndex(sf.Index).Interface().(openapi.TesFileType)
					dv.FieldByIndex(df.Index).Set(reflect.ValueOf(FileType(FileType_value[string(state)])))
				} else {
					dv.FieldByIndex(df.Index).Set(sv.FieldByIndex(sf.Index))
				}
			case reflect.Pointer:
				dst := reflect.New(df.Type.Elem())
				v := dst.Interface()
				nv := v.(proto.Message)
				OpenApi2Proto(sv.FieldByIndex(sf.Index).Interface(), nv)
				dv.FieldByIndex(df.Index).Set(dst)
			case reflect.Slice:
				sfv := sv.FieldByIndex(sf.Index)
				if df.Type.Elem().Kind() == reflect.Pointer {
					dstSlice := reflect.MakeSlice(df.Type, sfv.Len(), sfv.Len())
					for i := 0; i < sfv.Len(); i++ {
						div := reflect.New(df.Type.Elem().Elem())
						dii := div.Interface()
						siv := sfv.Index(i).Interface().(proto.Message)
						Proto2OpenApi(siv, dii)
						reflect.Append(dstSlice, div)
					}
					dv.FieldByIndex(df.Index).Set(dstSlice)
				} else {
					//fmt.Printf("Slice dest %s: %s %s\n", sf.Name, sf.Type.Elem().Kind(), df.Type.Elem().Kind())
					dstSlice := reflect.MakeSlice(df.Type, sfv.Len(), sfv.Len())
					for i := 0; i < sfv.Len(); i++ {
						div := reflect.New(df.Type.Elem())
						dii := div.Interface()
						siv := sfv.Index(i).Interface().(proto.Message)
						Proto2OpenApi(siv, dii)
						reflect.Append(dstSlice, div.Elem())
					}
					//dv.FieldByIndex(df.Index).Set(dstSlice)
					//dst := reflect.New(df.Type.Elem())
					//v := dst.Interface()
					//fmt.Printf("Slice dest: %T %s\n", v, df.Type.Elem().Kind())
				}
			case reflect.Map:
				smv := sv.FieldByIndex(sf.Index)
				dmv := reflect.MakeMap(df.Type)
				for _, kv := range smv.MapKeys() {
					vv := smv.MapIndex(kv)
					dmv.SetMapIndex(kv, vv)
				}
				dv.FieldByIndex(df.Index).Set(dmv)
			case reflect.Struct:
				//fmt.Printf("Found a struct\n")
			default:
				fmt.Printf("Found unknown: %s\n", df.Type.Kind())
			}
		} else {
			if sf.Name != "state" && sf.Name != "sizeCache" && sf.Name != "unknownFields" {
				fmt.Printf("Missing %s\n", sf.Name)
			}
		}
	}
}
