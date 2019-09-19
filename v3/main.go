package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type shell struct {
	scanner *bufio.Scanner
}

// listen for input.
func (s *shell) listen() {
	for {
		s.handler()
	}
}

// handler handles input.
func (s *shell) handler() {
	for s.scanner.Scan() {
		in := s.scanner.Text()
		fmt.Fprintf(os.Stderr, "%q: ", in)
		ins := strings.Split(in, " ")
		if len(in) < 1 {
			continue
		}
		inPath, err := exec.LookPath(ins[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s\n", inPath)
		out, err := exec.Command(ins[0], ins[1:]...).Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			continue
		}
		fmt.Fprintf(os.Stdout, "output: \n%s\n", out)
	}
	if err := s.scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}

func main() {
	s := shell{scanner: bufio.NewScanner(os.Stdin)}
	s.listen()
}
