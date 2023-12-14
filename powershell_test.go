package shellhook_test

import (
	"testing"

	"github.com/patrickhuber/go-shellhook"
	"github.com/stretchr/testify/require"
)

func TestPowershellCanExport(t *testing.T) {
	sh := shellhook.NewPowershell()
	result := shellhook.Export(sh, map[string]string{
		"TEST":  "VALUE",
		"TEST2": "VALUE2",
	})
	expected := "$env:TEST=\"VALUE\";\n$env:TEST2=\"VALUE2\";\n"
	require.Equal(t, expected, result)
}

func TestPowershellCanUnset(t *testing.T) {
	sh := shellhook.NewPowershell()
	actual := shellhook.Unset(sh, []string{"ONE", "TWO"})
	expected := "Remove-Item Env:\\ONE;\nRemove-Item Env:\\TWO;\n"
	require.Equal(t, expected, actual)
}
