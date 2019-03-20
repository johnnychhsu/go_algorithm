package main

import (
	"fmt"
)

func binarySearch(t []int, a int, low int, high int) int {
	if low > high {
		return -1
	}
	median := low + int((high-low)/2)
	if t[median] == a {
		return median
	} else if t[median] > a {
		return binarySearch(t, a, low, median)
	} else {
		return binarySearch(t, a, median, high)
	}
}

func main() {
	test := []int{1, 3, 5, 7, 9, 14, 16, 18, 19}
	ans := binarySearch(test, 9, 0, len(test))
	fmt.Println(ans)
}
