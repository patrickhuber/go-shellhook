package shellhook

type Exporter interface {
	Export(vars map[string]string) string
}
