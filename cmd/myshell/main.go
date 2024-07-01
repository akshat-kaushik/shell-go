package main

import (
	"bufio"

	"fmt"
	"os"
)

func main() {

	for {

		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command:=input[:len(input)-1]
		if command=="exit 0"{
			os.Exit(0)
			break;
		}
		fmt.Fprintf(os.Stdout,command+": command not found\n")

	}

}
