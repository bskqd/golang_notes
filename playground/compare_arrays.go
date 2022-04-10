package main

//
//func Comp(array1 []int, array2 []int) bool {
//	if array1 == nil || array2 == nil {
//		return false
//	}
//	var r []int
//	for _, s := range array1 {
//		r = append(r, s*s)
//	}
//	sort.Ints(r)
//	sort.Ints(array2)
//	return reflect.DeepEqual(r, array2)
//}

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, el := range array1 {
		m1[el]++
	}
	for _, el := range array2 {
		m2[el]++
	}
	for key, value := range m1 {
		doubledKey := key * key
		if m2[doubledKey] == 0 || m2[doubledKey] != value {
			return false
		}
	}
	return true
}
