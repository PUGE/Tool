package myslice

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
