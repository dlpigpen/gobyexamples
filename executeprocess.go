package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binnary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println(binnary)

	args := []string{"ls", "-a", "-l", "-h"}
	fmt.Println(args)

	env := os.Environ()
	fmt.Println(env)

	execErr := syscall.Exec(binnary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
