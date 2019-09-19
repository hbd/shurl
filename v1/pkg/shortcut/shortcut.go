package shortcut

import (
	"fmt"

	"github.com/hbd/shurl/pkg/help"
	"github.com/pkg/errors"
)

// Shortcuts is a global shortcut cache.
var Shortcuts map[string]Shortcut

// Shortcut contains args.
type Shortcut struct {
	Args []string
}

func init() {
	Shortcuts = map[string]Shortcut{}
}

// Bind expects the arguments to a bind request (excluding the 'bind' arg).
// First string in slice must be the name of the shortcut.
func Bind(args []string) (string, Shortcut, error) {
	if len(args) < 2 {
		return "", Shortcut{}, errors.New("invalid bind format")
	}
	return args[0], Shortcut{
		Args: args[1:],
	}, nil
}

func Unbind(args []string) {
	name := args[0]
	delete(Shortcuts, name)
}

// PrintShortcuts prints all shortcuts.
func PrintShortcuts() {
	fmt.Printf("Your shortcuts: \n")
	for name, args := range Shortcuts {
		fmt.Printf("%s: %s\n", name, args)
	}
}

// Handle any shortcut operations given the unedited arguments.
func Handle(args []string) ([]string, bool, error) {
	firstArg := args[0]

	// Check for a shortcut bind.
	if firstArg == "bind" {
		name, scArgs, err := Bind(args[1:])
		if err != nil {
			help.PrintHelpFor("bind")
			return args, false, errors.Wrap(err, "error binding shortcut")
		}
		Shortcuts[name] = scArgs
		return args, true, nil
	}

	// Check for a shortcut bind.
	if firstArg == "unbind" {
		Unbind(args[1:])
		return args, true, nil
	}

	// List shortcuts.
	if firstArg == "lbind" {
		PrintShortcuts()
		return args, true, nil
	}

	// Replace args with shortcuts if exists. Note: Destructive.
	if sc, ok := Shortcuts[firstArg]; ok {
		return sc.Args, false, nil
	}

	return args, false, nil
}
