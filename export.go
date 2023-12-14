package shellhook

import (
	"fmt"
	"sort"
	"strings"
)

type Exporter interface {
	Export(key string, value string) string
}

func Export(e Exporter, vars map[string]string) string {
	sb := strings.Builder{}

	// sort the keys for determinstic order
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := vars[k]
		result := e.Export(k, v)
		sb.WriteString(result)
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}
