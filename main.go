package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var arr []int
	arr = initArray(5)

	fmt.Println(arr)

	fmt.Println("min => ", Min(arr))
	fmt.Println("max => ", Max(arr))
	fmt.Println("sum => ", ArrSum(arr))
	fmt.Println("max sum => ", MaxNumSum(arr))
	fmt.Println("min sum => ", MinNumSum(arr))
}

func ArrSum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func Max(arr []int) int {
	max := arr[0]
	for _, n := range arr {
		if max < n {
			max = n
		}
	}
	return max
}

func Min(arr []int) int {
	min := arr[0]
	for _, n := range arr {
		if min > n {
			min = n
		}
	}
	return min
}

func MaxNumSum(arr []int) int {
	sum := ArrSum(arr)
	min := Min(arr)

	return sum - min
}

func MinNumSum(arr []int) int {
	sum := ArrSum(arr)
	max := Max(arr)

	return sum - max
}

func initArray(l int) []int {
	var arr []int
	for i := 0; i < l; i++ {
		n := rand.Intn(50)
		arr = append(arr, n)
	}
	return arr
}
