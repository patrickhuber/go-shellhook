package main

import (
	"fmt"
	"os"

	"github.com/patrickhuber/go-shellhook"
	"golang.org/x/exp/maps"
)

const ExportCommand = "export"
const HookCommand = "hook"

func main() {
	args := os.Args
	commands := map[string]struct{}{
		ExportCommand: {},
		HookCommand:   {},
	}
	if len(args) == 1 {
		fail(fmt.Errorf("missing commands, expected %+q", maps.Keys(commands)))
	}
	command := args[1]
	if _, ok := commands[command]; !ok {
		fail(fmt.Errorf("invalid command %s, expected %+q", command, maps.Keys(commands)))
	}

	var err error

	sh, err := shell(args)
	if err != nil {
		fail(err)
	}

	switch command {
	case ExportCommand:
		vars := map[string]string{
			"SHELLHOOK": "TEST",
		}
		err = export(sh, vars)
	case HookCommand:
		err = hook(sh)
	}

	if err != nil {
		fail(err)
	}
}

func hook(sh shellhook.Shell) error {
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	result, err := shellhook.Hook(sh, &shellhook.Metadata{
		Executable: executable,
		Name:       "test",
		Args:       []string{"export", sh.Name()},
	})
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func export(sh shellhook.Shell, vars map[string]string) error {
	result := shellhook.Export(sh, vars)
	fmt.Println(result)
	return nil
}

func shell(args []string) (shellhook.Shell, error) {
	shells := map[string]struct{}{
		shellhook.Bash:       {},
		shellhook.Powershell: {},
	}
	if len(args) == 2 {
		return nil, fmt.Errorf("missing shell parameter, expected %+q", maps.Keys(shells))
	}
	shell := args[2]
	if _, ok := shells[shell]; !ok {
		return nil, fmt.Errorf("invalid shell %s, expected %+q", shell, maps.Keys(shells))
	}
	switch shell {
	case shellhook.Powershell:
		return shellhook.NewPowershell(), nil
	case shellhook.Bash:
		return shellhook.NewBash(), nil
	default:
		return nil, fmt.Errorf("invalid shell")
	}
}

func fail(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
