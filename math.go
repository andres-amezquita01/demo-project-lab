package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DecimalToBinary(num int) string {
	var binary []int
	var result strings.Builder

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}
	if len(binary) == 0 {
		fmt.Printf("%d\n", 0)
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			result.WriteString(strconv.Itoa(binary[i]))
			fmt.Printf("%d", binary[i])
		}
		fmt.Println()
	}
	return result.String()
}
