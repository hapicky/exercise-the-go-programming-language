package main

import "fmt"

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeSameNeighbor(data []string) []string {
	for i := len(data) - 1; i > 0; i-- {
		if data[i] == data[i-1] {
			data = remove(data, i)
		}
	}
	return data
}

func main()  {
	sample := []string{
		"abc",
		"abc",
		"def",
		"ghi",
		"ghi",
		"abc",
	}

	removed := removeSameNeighbor(sample)
	fmt.Printf("%v", removed)
}
