package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	validcmds:=[]string{"echo","exit","type"}
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputarr := strings.Split(strings.TrimSpace(input), " ")
		command := inputarr[0]
		switch command {
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(inputarr[1:], " ") + "\n")
		case "exit":
			os.Exit(0)
		case "type":
			if slices.Contains(validcmds,inputarr[1]){
				fmt.Fprint(os.Stdout,command+" is a shell builtin")
			}else{
				fmt.Fprint(os.Stdout,command+" :not found")
			}

		default:
			fmt.Fprint(os.Stdout, command + ": command not found\n")
		}
	}
}
