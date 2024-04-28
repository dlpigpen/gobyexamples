package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Exit with status 3")
	os.Exit(3)
	fmt.Println("Exit with status 3 - never called")
}
