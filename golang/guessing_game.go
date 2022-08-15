package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var secret_num int = rand.Intn(100)
	var guess int
	var status int8 = 1

	fmt.Println("this is a guessing game")
	for status != 0 {
		fmt.Print("enter your guess:")
		num, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("that is'n a valid number try again")
			continue
		}

		fmt.Printf("the number of bytes readed was %d ", num)
		switch {
		case guess == secret_num:
			fmt.Println("you Win!")
			status = 0
			break
		case guess > secret_num:
			fmt.Println("too big try again \n")
		case guess < secret_num:
			fmt.Println("too small try again \n")
		}
	}

}
