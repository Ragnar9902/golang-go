package main

import (
	"os/exec"
	"errors"
	"fmt"
)

func execute(arg []string, ch chan error)int{
	cmd := exec.Command(arg[0],arg[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		ch <- errors.New("a error in the execution of the command")
		return -1
	}
	fmt.Print(string(out))
	ch <- nil
	return 0

}
