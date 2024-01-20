package shellhook_test

import (
	"testing"

	"github.com/patrickhuber/go-shellhook"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type test struct {
		name string
		err  bool
	}
	tests := []test{
		{shellhook.Powershell, false},
		{shellhook.Bash, false},
		{"none", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shell, err := shellhook.New(test.name)
			if test.err && err != nil {
				return
			}
			require.Equal(t, test.name, shell.Name())
		})
	}
}
