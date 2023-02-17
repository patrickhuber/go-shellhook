package shellhook_test

import (
	"fmt"
	"testing"

	"github.com/patrickhuber/shellhook"
	"github.com/stretchr/testify/require"
)

func TestBashCanExport(t *testing.T) {
	sh := shellhook.NewBash()
	result := sh.Export(map[string]string{
		"TEST": "VALUE",
	})
	require.Equal(t, result, fmt.Sprintln("export TEST=VALUE;"))
}
