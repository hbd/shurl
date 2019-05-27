package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hbd/shurl/pkg/help"
	"github.com/hbd/shurl/pkg/shortcut"
)

var builtins = map[string]func(){
	"help": help.PrintHelp,
}

func builtin(args []string) bool {
	if len(args) < 1 {
		return false
	}
	run, ok := builtins[args[0]]
	if !ok {
		return false
	}
	run()
	return true
}

func execCurl(args []string) {
	cmd := exec.Command("curl", args[0:]...)

	// out, err := cmd.Output() // Contains only stdout.
	out, err := cmd.CombinedOutput() // Contains stderr + stdout.
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("%s\n", out)
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			var input string
			fmt.Printf("shurl > ")
			scanner.Scan()
			input = scanner.Text()
			args := strings.Split(input, " ") // Split line into command and its args.

			// Check for built-ins.
			if builtin(args) {
				continue
			}

			// Handle shortcuts.
			args, err := shortcut.Handle(args)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				continue
			}

			execCurl(args)
		}
	}
}
