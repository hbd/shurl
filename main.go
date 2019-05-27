package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		// path := os.Getenv("PATH")
		for {
			var input string
			fmt.Printf("> ")
			scanner.Scan()
			input = scanner.Text()
			splitCmdLine := strings.Split(input, " ") // Split line into command and its args.
			cmd := exec.Command(splitCmdLine[0], splitCmdLine[1:]...)
			out, err := cmd.Output()
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			fmt.Printf("%s", out)
		}
	}
}
