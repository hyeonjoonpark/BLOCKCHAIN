package main

import "fmt"

type Calculator struct {
	X int
}

func (c Calculator) add(y int) int {
    return c.X + y
}

func (c *Calculator) addPointer(y int) {
	c.X += y
}

func method() {
	calc := Calculator{X: 10}
	val := calc.add(20)

	fmt.Println(val) // 30

	// 포인터 리시버
	calc.addPointer(5)
	fmt.Println(calc.X) // 15

}
