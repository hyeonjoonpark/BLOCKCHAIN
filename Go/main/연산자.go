package main

import "fmt"

func ex_operator() {
	// 산술연산자
	// +
	num1 := 10
	num2 := 20
	fmt.Printf("num1 + num2: %d\n", num1+num2) // num1 + num2: 30

	// -
	fmt.Printf("num1 - num2: %d\n", num1-num2) // num1 - num2: -10

	// *
	fmt.Printf("num1 * num2: %d\n", num1*num2) // num1 * num2: 200

	// / (몫)
	fmt.Printf("num1 / num2: %d\n", num1/num2) // num1 / num2: 0

	// % (나머지)
	fmt.Printf("num2 %% num1: %d\n", num2%num1) // num2 % num1: 0

	// 비교연산자
	// ==
	fmt.Printf("num1 == num2: %t\n", num1 == num2) // num1 == num2: false

	// !=
	fmt.Printf("num1 != num2: %t\n", num1 != num2) // num1 != num2: true

	// >
	fmt.Printf("num1 > num2: %t\n", num1 > num2) // num1 > num2: false

	// <
	fmt.Printf("num1 < num2: %t\n", num1 < num2) // num1 < num2: true

	// >= (이상)
	fmt.Printf("num1 >= num2: %t\n", num1 >= num2) // num1 >= num2: false

	// <= (이하)
	fmt.Printf("num1 <= num2: %t\n", num1 <= num2) // num1 <= num2: true


	// 대입연산자
	// = -> :=
	num3 := 10
	fmt.Printf("num3: %d\n", num3) // num3: 10

	// 논리연산자
	// &&
	isBool1 := true
	isBool2 := false
	fmt.Printf("isBool1 && isBool2: %t\n", isBool1 && isBool2) // isBool1 && isBool2: false

	// ||
	fmt.Printf("isBool1 || isBool2: %t\n", isBool1 || isBool2) // isBool1 || isBool2: true

	// !
	fmt.Printf("!isBool1: %t\n", !isBool1) // !isBool1: false

	// 증감연산자
	// ++
	num4 := 10
	num4++
	fmt.Printf("num4: %d\n", num4)

	// --
	num4--
	fmt.Printf("num4: %d\n", num4)

	// 비트연산자
	// & (AND)
	num5 := 10 // 1010
	num6 := 30 // 11110
	fmt.Printf("num5 & num6: %d\n", num5&num6) // num5 & num6: 10

	// | (OR)
	fmt.Printf("num5 | num6: %d\n", num5|num6) // num5 | num6: 30

	// ^ (XOR)
	fmt.Printf("num5 ^ num6: %d\n", num5^num6) // num5 ^ num6: 20
	fmt.Println(^num6) // 225 (30의 2의 보수)

	// &^ (AND NOT)
	fmt.Printf("num5 &^ num6: %d\n", num5&^num6) // num5 &^ num6: 0


	// 시프트연산자
	// << (왼쪽 시프트)
	num7 := 10
	fmt.Printf("num7 << 3: %d\n", num7<<3) // num7 << 3: 80
	fmt.Printf("num7 >> 3: %d\n", num7>>3) // num7 >> 3: 1
}
