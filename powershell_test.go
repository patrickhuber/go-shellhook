package shellhook_test

import (
	"fmt"
	"testing"

	"github.com/patrickhuber/go-shellhook"
	"github.com/stretchr/testify/require"
)

func TestPowershellCanExport(t *testing.T) {
	sh := shellhook.NewPowershell()
	result := sh.Export(map[string]string{
		"TEST": "VALUE",
	})
	require.Equal(t, result, fmt.Sprintln(`$env:TEST="VALUE";`))
}
