package shellhook_test

import (
	"testing"

	"github.com/patrickhuber/go-shellhook"
	"github.com/stretchr/testify/require"
)

func TestBashCanExport(t *testing.T) {
	sh := shellhook.NewBash()
	result := shellhook.Export(sh, map[string]string{
		"TEST":  "VALUE",
		"TEST2": "VALUE2",
	})
	expected := "export TEST='VALUE';\nexport TEST2='VALUE2';\n"
	require.Equal(t, expected, result)
}

func TestBashCanUnset(t *testing.T) {
	sh := shellhook.NewBash()
	actual := shellhook.Unset(sh, []string{"ONE", "TWO"})
	expected := "unset ONE;\nunset TWO;\n"
	require.Equal(t, expected, actual)
}
