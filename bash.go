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
if ! [[ "${PROMPT_COMMAND[*]:-}" =~ _{{.Name}}_hook ]]; then
  if [[ "$(declare -p PROMPT_COMMAND 2>&1)" == "declare -a"* ]]; then
    PROMPT_COMMAND=(_{{.Name}}_hook "${PROMPT_COMMAND[@]}")
  else
    PROMPT_COMMAND="_{{.Name}}_hook${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
  fi
fi
`

const (
	Bash = "bash"
)

type bash struct{}

func NewBash() Shell {
	return &bash{}
}

func (sh bash) Name() string {
	return Bash
}

// original https://github.com/direnv/direnv/blob/master/internal/cmd/shell_bash.go
func (sh bash) Hook() (string, error) {
	return bashTemplate, nil
}

func (sh bash) Export(key string, value string) string {
	return fmt.Sprintf("export %s='%s';", key, value)
}

func (sh bash) Unset(key string) string {
	return fmt.Sprintf("unset %s;", key)
}
