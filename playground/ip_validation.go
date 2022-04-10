package main

import (
	"strconv"
	"strings"
)

func ipValidation(ip string) bool {
	splitIP := strings.Split(ip, ".")
	if len(splitIP) != 4 {
		return false
	}
	for _, element := range splitIP {
		intElement, err := strconv.Atoi(element)
		if err != nil ||
			(strings.HasPrefix(element, "0") && intElement != 0) ||
			intElement > 255 ||
			intElement < 0 {
			return false
		}
	}
	return true
}
