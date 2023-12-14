package shellhook

import (
	"fmt"
	"strings"
)

type Unsetter interface {
	Unset(key string) string
}

func Unset(u Unsetter, keys []string) string {
	sb := strings.Builder{}
	for _, key := range keys {
		result := u.Unset(key)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}
