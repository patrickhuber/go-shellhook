package shellhook

import (
	"fmt"
	"sort"
	"strings"
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

func (sh bash) Export(vars map[string]string) string {
	sb := strings.Builder{}

	// sort the keys for determinstic order
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := vars[k]
		result := fmt.Sprintf("export %s='%s';", k, v)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}

func (sh bash) Unset(vars []string) string {
	sb := strings.Builder{}
	for _, v := range vars {
		result := fmt.Sprintf("unset %s;", v)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}
