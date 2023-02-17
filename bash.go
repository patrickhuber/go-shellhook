package shellhook

const bashTemplate = `
_{{.Name}}_hook(){
  local previous_exit_status=$?;
  trap -- '' SIGINT;
  eval "$("{{.SelfPath}}" export bash)";
  trap - SIGINT;
  return $previous_exit_status;
};
if ! [[ "${PROMPT_COMMAND:-}" =~ _{{.Name}}_hook ]]; then
  PROMPT_COMMAND="_{{.Name}}_hook${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
`
