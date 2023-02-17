package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/patrickhuber/shellhook"
	"golang.org/x/exp/maps"
)

const ExportCommand = "export"
const HookCommand = "hook"

const Powershell = "powershell"
const Bash = "bash"

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

	vars := env()
	var err error

	sh, err := shell(args)
	if err != nil {
		fail(err)
	}

	switch command {
	case ExportCommand:
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
	result := sh.Export(vars)
	fmt.Println(result)
	return nil
}

func env() map[string]string {
	result := map[string]string{}
	for _, environ := range os.Environ() {
		split := strings.SplitN(environ, "=", 2)

		key := split[0]
		value := split[1]

		result[key] = value
	}
	return result
}

func shell(args []string) (shellhook.Shell, error) {
	shells := map[string]struct{}{
		Bash:       {},
		Powershell: {},
	}
	if len(args) == 2 {
		return nil, fmt.Errorf("missing shell parameter, expected %+q", maps.Keys(shells))
	}
	shell := args[2]
	if _, ok := shells[shell]; !ok {
		return nil, fmt.Errorf("invalid shell %s, expected %+q", shell, maps.Keys(shells))
	}
	switch shell {
	case Powershell:
		return shellhook.NewPowershell(), nil
	case Bash:
		return shellhook.NewBash(), nil
	default:
		return nil, fmt.Errorf("invalid shell")
	}
}
func fail(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
