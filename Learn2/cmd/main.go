package main

import (
	"Learn_02/internal/utils"
	"fmt"
	"math/rand"
)

func main() {

	arr := gen(10)

	fmt.Println(arr)
	utils.DoQuickSort(arr)
	fmt.Println(arr)


}

func gen(i int) []int {
	var arr = make([]int, i);

	for j := 0; j < i; j++ {
		arr[j] = rand.Intn(101)
	}
	return arr
}
