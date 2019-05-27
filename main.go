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
			fmt.Printf("curl > ")
			scanner.Scan()
			input = scanner.Text()
			splitCmdLine := strings.Split(input, " ") // Split line into command and its args.
			cmd := exec.Command("curl", splitCmdLine[0:]...)
			// out, err := cmd.Output() // Contains only stdout.
			out, err := cmd.CombinedOutput() // Contains stderr + stdout.
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			fmt.Printf("%s\n", out)
		}
	}
}
