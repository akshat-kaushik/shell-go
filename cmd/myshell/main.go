package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type Shell struct {
	validCmds []string
	builtin   []string
	pathDirs  []string
}

func NewShell() *Shell {
	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")
	return &Shell{
		validCmds: []string{"echo", "exit", "type"},
		builtin:   []string{"echo", "exit", "type","pwd"},
		pathDirs:  dirs,
	}
}

func (sh *Shell) ExecuteCommand(input string) {
	input = strings.TrimSpace(input)
	params := strings.Split(input, " ")

	if len(params) == 0 {
		return
	}

	command := params[0]

	switch command {
	case "echo":
		sh.handleEcho(params[1:])
	case "exit":
		os.Exit(0)
	case "type":
		sh.handleType(params)
	case "pwd":
		sh.handlePwd()
	case "cd":
		sh.handleCd(params[1])
	default:
		sh.handleExternalCommand(command, params[1:])
	}
}

func (sh *Shell) handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func (sh *Shell) handleType(params []string) {
	if len(params) < 2 {
		fmt.Println("type: missing operand")
		return
	}

	cmd := params[1]
	if slices.Contains(sh.builtin, cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else {
		sh.findCommandInPath(cmd)
	}
}

func (sh *Shell) findCommandInPath(cmd string) {
	for _, dir := range sh.pathDirs {
		fp := filepath.Join(dir, cmd)
		if _, err := os.Stat(fp); err == nil {
			fmt.Printf("%s is %s\n", cmd, fp)
			return
		}
	}
	fmt.Printf("%s: not found\n", cmd)
}

func (sh *Shell) handleExternalCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: command not found\n", command)
	}
}
 func (sh *Shell) handlePwd(){
	dir,err:=os.Getwd()
	if err==nil{
		fmt.Printf("%s\n",dir)
	}
 }

 func (sh *Shell) handleCd(path string){
	if slices.Contains(sh.pathDirs,path){
		os.Chdir(path)
	}
	
 }

func main() {
	shell := NewShell()
	scanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			continue
		}
		shell.ExecuteCommand(input)
	}
}
