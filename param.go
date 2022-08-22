package aptos

import (
	"fmt"
	"reflect"
	"strings"
)

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// msg's member should only be number or string
func encodeURLParam(msg interface{}) string {
	if msg == nil {
		return ""
	}
	t := reflect.TypeOf(msg)
	v := reflect.ValueOf(msg)
	k := v.Kind()

	if k != reflect.Struct {
		return ""
	}

	buf := ""
	n := v.NumField()
	s := ""
	for i := 0; i < n; i++ {
		tv := t.Field(i).Tag.Get("param")
		if len(tv) == 0 {
			continue
		}
		fv := v.Field(i)
		tvs := strings.Split(tv, ",")
		if len(tvs) == 2 {
			tv = strings.TrimSpace(tvs[0])
			if strings.TrimSpace(tvs[1]) == "omitempty" {
				if isEmptyValue(fv) {
					continue
				}
			}
		}
		buf += fmt.Sprintf("%s%s=%v", s, tv, fv)
		s = "&"
	}
	return buf
}

func endodePathParam(paramPath string, msg interface{}) string {
	if msg == nil {
		return paramPath
	}
	t := reflect.TypeOf(msg)
	v := reflect.ValueOf(msg)
	k := v.Kind()

	if k != reflect.Struct {
		return paramPath
	}

	param := paramPath
	n := v.NumField()
	for i := 0; i < n; i++ {
		tv := t.Field(i).Tag.Get("path")
		fv := v.Field(i)
		param = strings.ReplaceAll(param, fmt.Sprintf("{%s}", tv), fmt.Sprintf("%v", fv))
	}

	return param
}

func endodeURLPath(paramPath string, msg interface{}) string {
	rst := endodePathParam(paramPath, msg)
	param := encodeURLParam(msg)
	if len(param) > 0 {
		rst += "?" + param
	}
	return rst
}
