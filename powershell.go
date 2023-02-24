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

type powershell struct {
}

func NewPowershell() Shell {
	return &powershell{}
}

func (sh powershell) Name() string {
	return "powershell"
}

func (sh powershell) Hook() (string, error) {
	return powershellTemplate, nil
}

func (sh powershell) Export(vars map[string]string) string {
	results := ""
	for k, v := range vars {
		result := fmt.Sprintf(`$env:%s="%s";`, k, v)
		results = fmt.Sprintln(result)
	}
	return results
}
