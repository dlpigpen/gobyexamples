package main

import (
	"fmt"
	"math"
)

const s string = "here is the constant"

func main() {
	fmt.Println(s)
	var sinA float64 = math.Sin(1)
	fmt.Println(sinA)

	var sinB float64
	sinB = math.Sin(1)
	fmt.Println(sinB)

	const PI = 3.14159
}
