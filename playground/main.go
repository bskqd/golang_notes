package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	arr := []int{1, 0, 3, 0, 5, 6, 7, 8, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3, 0, 1, 2, 3}
	ip := "192.168.1.300"
	fmt.Printf("Largest range of %v is: %v\n", arr, largestRange(arr))
	fmt.Printf("IP %v is correct: %v\n", ip, ipValidation(ip))
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Printf("%v\n", Comp(a1, a2))
}
