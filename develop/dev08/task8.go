package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "\\quit" {
			break
		}
		commands := strings.Split(input, "|")
		executePipeline(commands)
	}
}

func executePipeline(commands []string) {
	var prevOutput io.Reader
	for _, cmdStr := range commands {
		args := strings.Fields(strings.TrimSpace(cmdStr))
		switch args[0] {
		case "cd", "pwd", "echo", "kill":
			executeCommand(args)
		default:
			cmd := exec.Command(args[0], args[1:]...)
			if prevOutput != nil {
				cmd.Stdin = prevOutput
			}
			var output []byte
			output, err := cmd.Output()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "Output fail:", err)
				return
			}
			fmt.Print(string(output))
			prevOutput = bytes.NewReader(output)
		}
	}
}

func executeCommand(args []string) {
	if len(args) == 0 {
		return
	}
	command := args[0]
	switch command {
	case "cd":
		if len(args) < 2 {
			fmt.Println("Using: cd <directory>")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error when changing the directory:", err)
		}
	case "pwd":
		cwd, err := os.Getwd()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error getting the current directory:", err)
			return
		}
		fmt.Println(cwd)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			fmt.Println("Using: kill <pid>")
			return
		}
		pid := args[1]
		cmd := exec.Command("kill", pid)
		err := cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error when killing the process:", err)
		}
	default:
		cmd := exec.Command(command, args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error when executing the command:", err)
		}
	}
}
