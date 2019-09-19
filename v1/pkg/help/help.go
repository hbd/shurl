package help

import (
	"fmt"
	"log"
)

var help = map[string]string{
	"bind": "bind [name] [args...] | ex: bind lh localhost:8080",
}

// PrintHelpFor a given command.
func PrintHelpFor(cmd string) {
	helpText, ok := help[cmd]
	if !ok {
		log.Fatalf("Unrecognized cmd for help: %s", cmd)
	}
	fmt.Printf("%s: %s\n", cmd, helpText)
}

// PrintHelp prints all of help.
func PrintHelp() {
	for cmd, helpText := range help {
		fmt.Printf("%s: %s\n", cmd, helpText)
	}
}
