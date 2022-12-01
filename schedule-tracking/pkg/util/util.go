package util

import "fmt"

func Pop[T comparable](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
func PopForInterfaces(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}
func GetIndex(item interface{}, s ...interface{}) int {
	for index, v := range s {
		if v == item {
			return index
		}
	}
	return -1
}
func ConvertArgsToInterface[T comparable](args []T) []interface{} {
	var arr []interface{}
	for _, v := range args {
		arr = append(arr, v)
	}
	return arr
}

func ConvertInterfaceArgsToString(args []interface{}) []string {
	var arr []string
	for _, v := range args {
		arr = append(arr, fmt.Sprintf(`%v`, v))
	}
	return arr
}
