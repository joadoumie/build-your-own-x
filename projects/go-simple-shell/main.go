package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Create a set for command history
	var history []string

	for {
		// Read the keyboard input
		fmt.Print("Jordi's Shell > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "read error:", err)
		}
		// Execute the command
		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, "exec error:", err)
		}
		history = append(history, input)
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	// get args from input
	args := strings.Fields(input)

	switch args[0] {
	case "exit":
		fmt.Println("Exiting shell...")
		os.Exit(0)
	case "cd":
		if len(args) < 2 {
			return os.Chdir(os.Getenv("HOME"))
		}
		return os.Chdir(args[1])
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()	
}