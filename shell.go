package shellhook

type Namer interface {
	Name() string
}

type Shell interface {
	Namer
	Exporter
	Hooker
}
