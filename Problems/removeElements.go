package main

import "fmt"

func main() {
	var array1 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	const value = 2
	fmt.Println(array1)
	removeElements(array1, value)
	fmt.Println(array1)
}

func removeElements(array []int, value int) {
	var dynamicLenghtArray int = len(array)
	var lenghtArray int = len(array)
	for i := 0; i < dynamicLenghtArray; i++ {
		var valueArray int = array[i]
		for valueArray == value {
			dynamicLenghtArray--
			for x := i; x < lenghtArray-1; x++ {
				array[x] = array[x+1]
			}
			array[lenghtArray-1] = -1
			valueArray = array[i]
		}
	}
	fmt.Println(dynamicLenghtArray)
}
