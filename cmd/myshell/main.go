package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputarr := strings.Split(strings.TrimSpace(input), " ")
		command := inputarr[0]
		switch command {
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(inputarr[1:], " ") + "\n")
		default:
			fmt.Fprint(os.Stdout, command + ": command not found\n")
		}
	}
}
