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

// Bind expects the arguments to a bind request (excluding the 'bind' arg).
func Bind(args []string) (string, Shortcut, error) {
	if len(args) < 2 {
		return "", Shortcut{}, errors.New("invalid bind format")
	}
	return args[0], Shortcut{
		Args: args[1:],
	}, nil
}

// PrintShortcuts prints all shortcuts.
func PrintShortcuts() {
	fmt.Printf("Your shortcuts: \n")
	for name, args := range Shortcuts {
		fmt.Printf("%s: %s", name, args)
	}
}

// Handle any shortcut operations given the unedited arguments.
func Handle(args []string) ([]string, error) {
	firstArg := args[0]

	// Check for a shortcut bind.
	if firstArg == "bind" {
		name, scArgs, err := Bind(args[1:])
		if err != nil {
			help.PrintHelpFor("bind")
			return args, errors.Wrap(err, "error binding shortcut")
		}
		Shortcuts[name] = scArgs
		return args, nil
	}

	// List shortcuts.
	if firstArg == "lbind" {
		PrintShortcuts()
		return args, nil
	}

	// Replace args with shortcuts if exists. Note: Destructive.
	if sc, ok := Shortcuts[firstArg]; ok {
		return sc.Args, nil
	}

	return args, nil
}
