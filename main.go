package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/hbd/shurl/pkg/help"
	"github.com/hbd/shurl/pkg/shortcut"
)

var debugFlag = flag.Bool("debug", false, "If set, debug messages are printed.")

func init() {
	flag.Parse()
	if *debugFlag {
		fmt.Println("Debug mode on.")
	}
}

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

func debugPrint(format string, a ...interface{}) (n int, err error) {
	if *debugFlag {
		return fmt.Fprintf(os.Stdout, format, a...)
	}
	return 0, nil
}

func main() {
	// Handle signals.
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	debugPrint("PID: %d\n", os.Getpid())
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				os.Exit(0)
			}
			debugPrint("\nReceived signal: %s\n", sig)
		}
	}()

	sc := bufio.NewScanner(os.Stdin)
	for {
		var input string
		fmt.Printf("shurl > ")

		if scanned := sc.Scan(); !scanned { // Handle EOF.
			if err := sc.Err(); err != nil {
				fmt.Printf("%s\n", err)
			}
			os.Exit(0)
		}
		input = sc.Text()
		args := strings.Split(input, " ") // Split line into command and its args.
		if args[0] == "" {
			continue
		}

		// Check for built-ins.
		if builtin(args) {
			continue
		}

		// Handle shortcuts.
		args, known, err := shortcut.Handle(args)
		if known {
			continue
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		execCurl(args)
	}
}

// ]^ (up arrow?) ^[[A
// ^C
// history
// rc
// env var
