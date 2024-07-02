package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	validcmds := []string{"echo", "exit", "type", "ls"}
	validLocalcmds:=[]string{"abcd"}
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
			if slices.Contains(validcmds, inputarr[1]) {
				fmt.Fprint(os.Stdout, inputarr[1]+" is"+ filepath.Join(dirs[0],inputarr[1])+"\n")
			} else if slices.Contains(validLocalcmds,inputarr[1]) {
				fmt.Fprint(os.Stdout, inputarr[1]+" is"+ filepath.Join(dirs[1],inputarr[1])+"\n")
			}else{
				fmt.Fprint(os.Stdout,inputarr[1]+": not found\n")
			}

		default:
			fmt.Fprint(os.Stdout, command+": command not found\n")
		}
	}
}
