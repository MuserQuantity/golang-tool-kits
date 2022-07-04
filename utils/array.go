package utils

import "fmt"

// InterfaceArray2StringArray 将interface数组转换为字符串数组
func InterfaceArray2StringArray(array1 []interface{}) (array2 []string) {
	for _, a := range array1 {
		array2 = append(array2, fmt.Sprintf("%s", a))
	}
	return
}

// InterfaceArray2UintArray 将interface数组转换成uint数组
func InterfaceArray2UintArray(array1 []interface{}) (array2 []uint, ok bool) {
	for _, a := range array1 {
		temp, ok := a.(uint)
		fmt.Println(temp, ok)
		if !ok {
			return nil, ok
		}
		array2 = append(array2, temp)
	}
	return array2, true
}

// ContainsString 字符串数组中是否包含某字符串
func ContainsString(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

// ContainsUint uint数组中是否包含某uint
func ContainsUint(items []uint, item uint) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}
