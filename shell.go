package shellhook

type Shell interface {
	Name() string
	Hook() (string, error)
	Export(vars map[string]string) string
}
