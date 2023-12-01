package shellhook

import (
	"fmt"
	"strings"
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

func (sh powershell) Export(vars map[string]string) string {
	sb := strings.Builder{}
	for k, v := range vars {
		result := fmt.Sprintf(`$env:%s="%s";`, k, v)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}
