package mycheck

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"strings"
)

// Contain 判断obj是否在target中，target支持的类型arrary,slice,map
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

// CheckPortOccupy 检查端口是否被占用
func CheckPortOccupy(port int) bool {
	// 执行系统命令
	out, err := exec.Command("netstat", "-ano").Output()
	if err != nil {
		log.Fatalf("exec command error: %s", err)
	}
	// 判断端口是否被占用
	return strings.Contains(string(out), fmt.Sprintf(":%d", port))
}
