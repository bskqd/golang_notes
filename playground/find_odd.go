package main

func FindOdd(seq []int) int {
	m := make(map[int]int)
	for _, el := range seq {
		m[el]++
	}
	for k, v := range m {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}
