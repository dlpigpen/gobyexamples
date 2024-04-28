package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("Failed executing command:", err)
		case *exec.ExitError:
			fmt.Println("Command exit rc=", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

}
