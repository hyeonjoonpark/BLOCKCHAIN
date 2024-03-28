package main

import "fmt"

func 변수() {
	// 변수선언 var
	var num1 int = 10
	fmt.Printf("num1: %d\n", num1) // num1: 10 (int로 선언)

	// 변수선언 타입추론
	var num2 = 20
	fmt.Printf("num2: %d\n", num2) // num2: 20 (자동으로 추론해서 int로 선언)

	// 변수선언 축약형
	num3 := 30
	fmt.Printf("num3: %d\n", num3) // num3: 30 (자동으로 추론해서 int로 선언) -> 왈루스

	// 변수선언 다중
	var num4, num5 int = 40, 50
	fmt.Printf("num4: %d, num5: %d\n", num4, num5) // num4: 40, num5: 50 (int로 선언)

	// 상수선언 const
	const num6 int = 40
	fmt.Printf("num6: %d\n", num6) // num4: 40 (int로 선언)
	// num6 = 10; 상수는 변경 불가능 -> 에러

	// 상수선언 다중
	const (
		num7 int = 50
		num8 int = 60
	)
	fmt.Printf("num7: %d, num8: %d\n", num7, num8) // num7: 50, num8: 60 (int로 선언)
}
