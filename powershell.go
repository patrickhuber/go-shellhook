package shellhook
const powershellTemplate = `
global.__{{.NAME}}OriginalPrompt=(Get-Item function:prompt).ScriptBlock

function global:prompt{
	$result = $global:__{{.Name}}OriginalPrompt.Invoke()
	return $result
}`