package util

func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveDuplicateUntilLimit[T string | int](sliceList []T, limit int) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
		if len(list) == limit {
			break
		}
	}
	return list
}

func GetAllNextUniqueValues[T string | int](sliceList []T) []T {
	uniqueValue := sliceList[0]
	list := []T{uniqueValue}
	for _, item := range sliceList {
		if item != uniqueValue {
			uniqueValue = item
			list = append(list, item)
		}
	}
	return list
}
