package shellhook

import (
	"fmt"
	"sort"
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

	// sort the keys for determinstic order
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := vars[k]
		result := fmt.Sprintf(`$env:%s="%s";`, k, v)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}

func (sh powershell) Unset(vars []string) string {
	sb := strings.Builder{}
	for _, v := range vars {
		result := fmt.Sprintf("Remove-Item Env:\\%s;", v)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}
