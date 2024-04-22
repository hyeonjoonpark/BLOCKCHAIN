package main

import (
	"encoding/json"
	"fmt"
)

func ex_struct() {

	type Vertex struct {
		x int
		y string
	}

	// 구조체 (클래스)
	var v0 Vertex = Vertex{1, "Hello"};
	fmt.Println(v0);

	v0.x = 2;
	fmt.Println(v0); // 2

	// json 형식으로 출력
	type Vertex2 struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	v := Vertex2{10, 2};
	data, _ := json.Marshal(v);
	fmt.Println(string(data));
}
