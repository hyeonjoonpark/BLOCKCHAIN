package main

import "fmt"

func ex_array() {
	// 배열
	var arr [3]int
	arr = [3]int{1, 2, 3}

	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	length := len(arr) // 배열 길이
	println(length)

	// 배열 순회
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	// 다중 배열
	var arr2 [3][3]int
	arr2 = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(arr2[1][2])
}
