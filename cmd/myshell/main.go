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
		fmt.Fprintf(os.Stdout, input[:len(input)-1]+": command not found\n")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

}
