package shellhook

import (
	"fmt"
)

const bashTemplate = `
_{{.Name}}_hook(){
  local previous_exit_status=$?;
  trap -- '' SIGINT;
  eval "$("{{.Executable}}" {{ .Args | join " " }})";
  trap - SIGINT;
  return $previous_exit_status;
};
if ! [[ "${PROMPT_COMMAND:-}" =~ _{{.Name}}_hook ]]; then
  PROMPT_COMMAND="_{{.Name}}_hook${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
`

type bash struct{}

func NewBash() Shell {
	return &bash{}
}

func (sh bash) Name() string {
	return "bash"
}

// original https://github.com/direnv/direnv/blob/master/internal/cmd/shell_bash.go
func (sh bash) Hook() (string, error) {
	return bashTemplate, nil
}

func (sh bash) Export(vars map[string]string) string {
	results := ""
	for k, v := range vars {
		result := fmt.Sprintf("export %s=%s;", k, v)
		results = fmt.Sprintln(result)
	}
	return results
}
