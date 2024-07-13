package main

import "fmt"
import "go_test/other"
import "math/rand"

func main(){
	fmt.Println("From main package")
	other.Hello()

	arr := gen(10)

	fmt.Println(arr)
	other.Sort(arr)
	fmt.Println(arr)

	max := other.GetMax(arr)
	fmt.Println(max)

	sum := other.GetSum(arr)
	fmt.Println(sum)
}

func gen(i int) []int {
	var arr = make([]int, i);

	for j := 0; j < i; j++ {
		arr[j] = rand.Intn(101)
	}
	return arr
}

