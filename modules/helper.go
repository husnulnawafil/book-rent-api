package modules

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func GetFunctionName(i interface{}) string {
	name := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	return fmt.Sprintf("%s-%s", name[len(name)-1], name[len(name)-2])
}

func GenerateRedisKey(f, id interface{}) string {
	fname := GetFunctionName(f)
	key := fname
	if id != nil {
		key = fmt.Sprintf("%s-%d", fname, id.(uint))
	}
	return key
}
