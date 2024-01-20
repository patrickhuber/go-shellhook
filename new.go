package shellhook

import (
	"fmt"
	"io/fs"
)

var ErrNotExist = fs.ErrNotExist

func New(shell string) (Shell, error) {
	switch shell {
	case Powershell:
		return NewPowershell(), nil
	case Bash:
		return NewBash(), nil
	}
	return nil, fmt.Errorf("%w : %s", ErrNotExist, shell)
}
