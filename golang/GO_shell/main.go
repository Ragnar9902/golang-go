package main
import (
	"fmt"
	"os"
	"bufio"
)

type builtin struct {
	id_func string
	f func(string)int
}

var ch1 chan error = make(chan error)

func main(){
	var command string
	status := 1
	var built_in func(...string)int

	for status != 0 {
		fmt.Printf("Ragnar>>")
		scanner := bufio.NewScanner(os.Stdin)
    		scanner.Scan() // use `for scanner.Scan()` to keep reading
    		command = scanner.Text()
		args := tokenize(command)
		fmt.Print(len(args))
		if args == nil{
			continue
		}
		fmt.Print(args)
		built_in = get_builtin(args[0])
		if built_in != nil {
			built_in(args...)
			continue
		}

		go execute(args, ch1)

		err := <-ch1
		if err != nil {
        		fmt.Println(err.Error())
    		}
		// Print the output
		//fmt.Printf("the number of bytes readed was %d", num)
		//fmt.Println("the number of bytes readed was ", err)
		if command == "exit" {
			status = 0
			break	
		}

	}
}
