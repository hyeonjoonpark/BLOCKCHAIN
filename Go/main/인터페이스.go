package main

/*
* 인터페이스

* type 인터페이스명 interface {
* 	 메서드 집합
* }
 */

type Ducky interface {
	DuckySound() string
	DuckyWalk() string
	isSwim() string
}

type Bird struct {
	name string
}

func (b Bird) DuckySound() string {
	return "꽥꽥"
}

func (b Bird) DuckyWalk() string {
	return "뒤뚱뒤뚱"
}

func (b Bird) isSwim() string {
	if b.name == "오리" {
		return "헤엄칠 수 있다"
	} else {
		return "헤엄칠 수 없다"
	}
}

func ex_interface() {

}
