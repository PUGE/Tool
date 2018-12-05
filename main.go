package tool

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

// GetRandomString 生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandInt64 生成指定区间的随机数
func RandInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

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

// BytesToString Bytes转字符串
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

// Delete 从切片中删除制定项目
func Delete(slice *[]string, str string) {
	temp := make([]string, len(*slice))
	temp = *slice
	for k, v := range temp {
		if v == str {
			kk := k + 1
			*slice = append(temp[:k], temp[kk:]...)
		}
	}
}
