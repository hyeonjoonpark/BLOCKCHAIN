package main

import "fmt"

func ex_pointer() {
	// 포인터
	num := 1
	numPointer := &num
	*numPointer = 2 // numPointer가 가리키는 주소에 2를 대입
	fmt.Println(num) // 2
}
