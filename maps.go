package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["k1"] = 7
	m["k2"] = 13

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

}
