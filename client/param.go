package client

import (
	"fmt"
	"reflect"
)

// msg's member should only be number or string
func paramEncode(msg interface{}) string {
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
		fv := v.Field(i)
		buf += fmt.Sprintf("%s%s=%v", s, tv, fv)
		s = "&"
	}
	return buf
}
