package shellhook

import (
	"fmt"
)

const powershellTemplate = `
$global:__{{.Name}}OriginalPrompt=(Get-Item function:prompt).ScriptBlock

function global:prompt{
	# export powershell environment variables
	iex $({{.Executable}} {{.Args | join " " }} | out-string)
	$result = $global:__{{.Name}}OriginalPrompt.Invoke()
	return $result
}`

const (
	Powershell = "powershell"
)

type powershell struct {
}

func NewPowershell() Shell {
	return &powershell{}
}

func (sh powershell) Name() string {
	return Powershell
}

func (sh powershell) Hook() (string, error) {
	return powershellTemplate, nil
}

func (sh powershell) Export(key string, value string) string {
	return fmt.Sprintf(`$env:%s="%s";`, key, value)
}

func (sh powershell) Unset(key string) string {
	return fmt.Sprintf("Remove-Item Env:\\%s;", key)
}
