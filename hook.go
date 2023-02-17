package shellhook

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig/v3"
)

func Hook(shell Shell, metadata *Metadata) (string, error) {

	hookString, err := shell.Hook()
	if err != nil {
		return "", nil
	}

	hookTemplate, err := template.New("hook").Funcs(sprig.FuncMap()).Parse(hookString)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = hookTemplate.Execute(&buf, metadata)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
