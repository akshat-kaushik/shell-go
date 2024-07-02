package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// validcmds := []string{"echo", "exit", "type", "ls"}

	var path = os.Getenv("PATH")
	var dirs = strings.Split(path, ":")

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputarr := strings.Split(strings.TrimSpace(input), " ")
		command := inputarr[0]
		switch command {
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(inputarr[1:], " ")+"\n")
		case "exit":
			os.Exit(0)
		case "type":
			exist := false
			for _, path := range dirs {
				fp := filepath.Join(path, inputarr[1])
				if _, err := os.Stat(fp); err == nil {
					fmt.Fprint(os.Stdout, inputarr[1]+" is "+fp+"\n")
					exist = true
					break
				}
			}
			if !exist {
				fmt.Fprint(os.Stdout, inputarr[1]+": not found\n")
			}

		default:
			fmt.Fprint(os.Stdout, command+": command not found\n")
		}
	}
}
