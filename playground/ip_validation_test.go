package main

import (
	"testing"
)

func TestIPValidation(t *testing.T) {
	ip1 := "192.168.1.300"
	ip2 := "192.168.1.0"
	ip3 := "123.168.1.03"
	ip4 := "123.168.-1.3"
	res1 := ipValidation(ip1)
	res2 := ipValidation(ip2)
	res3 := ipValidation(ip3)
	res4 := ipValidation(ip4)
	if res1 {
		t.Fatalf("ip: %v, expected return: %v, return: %v", ip1, false, res1)
	}
	if !res2 {
		t.Fatalf("ip: %v, expected return: %v, return: %v", ip2, true, res2)
	}
	if res3 {
		t.Fatalf("ip: %v, expected return: %v, return: %v", ip3, false, res3)
	}
	if res4 {
		t.Fatalf("ip: %v, expected return: %v, return: %v", ip4, false, res4)
	}
}
