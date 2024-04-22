package main

import "fmt"

func ex_slice() {
	// 슬라이스 (동적배열)

	// 선언
	var slice []int
	fmt.Println(slice)
	fmt.Printf("%T\n", slice) // 슬라이스 타입

	// 초기화
	var slice2 []int = []int{1, 2, 3}
	fmt.Println(slice2)

	// make 함수
	slice3 := make([]int, 3) // make(슬라이스 타입, 길이, 용량)
	fmt.Println(slice3) // [0 0 0] -> 0으로 초기화된 슬라이스

	slice3[0] = 1
	fmt.Println(slice3[0]) // 1

	// 값 추가
	// append(슬라이스, 값)
	slice3 = append(slice3, 4)
	fmt.Println(slice3) // [1 0 0 4]

	// cap: 용량
	slice4 := make([]int, 10, 10)
	fmt.Println(len(slice4), cap(slice4)) // 10 10

	// 슬라이스 복사
	slice5 := []int{1, 2, 3, 4, 5}
	slice6 := make([]int, 10)
	copy(slice6, slice5) // copy(복사할 슬라이스, 복사될 슬라이스)
	fmt.Println(slice6) // [1 2 3 4 5 0 0 0 0 0] -> 1, 2, 3, 4, 5는 복사하고 나머지 0으로 초기화

	// 배열 슬라이싱
	arr := [5]int{1, 2, 3, 4, 5} // 배열 선언
	slice7 := arr[0:3] // [1 2 3] -> 0번째부터 3번째까지 슬라이싱
	fmt.Println(slice7)

	// type check
	fmt.Printf("%T\n", arr) // [5]int
	fmt.Printf("%T\n", slice7) // []int

	// 슬라이스 슬라이싱
	slice8 := []int{1, 2, 3, 4, 5}
	slice9 := slice8[0:3] // [1 2 3] -> 0번째부터 3번째까지 슬라이싱
	fmt.Println(slice9)

	slice10 := slice9[:3] // [1 2 3] -> 0번째부터 3번째까지 슬라이싱
	fmt.Println(slice10)

	slice11 := slice9[1:3:4] // [2 3] -> 1번째부터 3번째까지 슬라이싱, 용량은 4
	fmt.Println(slice11)
}
