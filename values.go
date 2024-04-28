package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")        // cat two strings and print
	fmt.Println("1+1 =", 1+1)         // add two numbers and print
	fmt.Println("7.0/3.0 =", 7.0/3.0) // divide two numbers and print
	fmt.Println(true && false)        // print the result of a logical AND
	fmt.Println(true || false)        // print the result of a logical OR
	fmt.Println(!true)                // print the result of a logical NOT

	var a string = "initial" // declare a variable and assign a value
	fmt.Println(a)           // print the value of the variable

	var b, c int = 1, 2 // declare two variables and assign values
	fmt.Println(b, c)   // print the values of the variables

	var d = true   // declare a variable and assign a value
	fmt.Println(d) // print the value of the variable

	var e int      // declare a variable
	fmt.Println(e) // print the value of the variable

	f := "short"   // declare a variable and assign a value
	fmt.Println(f) // print the value of the variable

}
