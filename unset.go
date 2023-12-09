package shellhook

type Unsetter interface {
	Unset(vars []string) string
}
