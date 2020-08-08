package main

import (
	"fmt"
)

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	tmp := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > tmp {
			tmp = vals[i]
		}
	}
	return tmp
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	tmp := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < tmp {
			tmp = vals[i]
		}
	}
	return tmp
}

func max2(val int, vals ...int) int {
	tmp := val
	for i := 0; i < len(vals); i++ {
		if vals[i] > tmp {
			tmp = vals[i]
		}
	}
	return tmp
}

func min2(val int, vals ...int) int {
	tmp := val
	for i := 0; i < len(vals); i++ {
		if vals[i] < tmp {
			tmp = vals[i]
		}
	}
	return tmp
}

func main() {
	fmt.Println(max(1, 2, 3))
	fmt.Println(max())
	fmt.Println(min(1, 2, 3))
	fmt.Println(min())
	fmt.Println(max2(1, 3, 2))
	fmt.Println(max2(1))
	fmt.Println(min2(2, 1, 3))
	fmt.Println(min2(1))
}
