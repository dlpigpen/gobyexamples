package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	person := person{name: "Alice", age: 30}

	fmt.Println("Initial value of person:", person)

}
