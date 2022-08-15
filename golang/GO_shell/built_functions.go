package main

import (
	"fmt"
	"os"
)

func exit(args ...string)int{
	fmt.Printf("end the shell")
	if args != nil{
	}
	os.Exit(0)
	return 0
	
}

func version(args ...string)int{
	fmt.Println("version 4.2/0")
	if args != nil{
	}
	return 1
}
