package main

import "fmt"

func 타입() {
	// 주로 사용하는 타입

	var isBool bool = false
	fmt.Printf("isBool: %t\n", isBool) // isBool: true (bool로 선언)

	var num1 int = 10
	fmt.Printf("num1: %d\n", num1) // numInt: 10 (int로 선언)
	// int8, int16, int32, int64
	// int8 : -128 ~ 127
	// int16 : -32768 ~ 32767
	// int32 : -2147483648 ~ 2147483647
	// int64 : -9223372036854775808 ~ 9223372036854775807

	var num2 uint = 20
	fmt.Printf("num2: %d\n", num2) // num2: 20 (uint로 선언)
	// uint8, uint16, uint32, uint64
	// uint8 : 0 ~ 255
	// uint16 : 0 ~ 65535
	// uint32 : 0 ~ 4294967295
	// uint64 : 0 ~ 18446744073709551615

	var num3 float32 = 30.1
	fmt.Printf("num3: %f\n", num3) // num3: 30.100000 (float32로 선언)
	// float32, float64
	// float32 : 32비트 부동소수점
	// float64 : 64비트 부동소수점

	var num4 byte = 40
	fmt.Printf("num4: %d\n", num4) // num4: 40 (byte로 선언)
	// 1 byte = 8 bit
	// 0000 0000 ~ 1111 1111 (0 ~ 255)

	var string1 string = "Hello, World!"
	fmt.Printf("string1: %s\n", string1) // string1: Hello, World! (string으로 선언)
	// 문자열은 ""로 감싸서 표현


	// 형변환
	var float1 float32 = 3.14
	fmt.Printf("float1: %f\n", float1) // float1: 3.140000 (float32로 선언)

	fmt.Println("int(float1):", int(float1)) // int(float1): 3 (int로 형변환)


	// %T : 타입
	fmt.Printf("isBool: %T\n", isBool) // isBool: bool
	// 타입을 출력할 때 사용

	// type : 타입 정의
	type myInt int
	var num5 myInt = 50
	fmt.Printf("num5: %d\n", num5) // num5: 50 (myInt로 선언)
}
