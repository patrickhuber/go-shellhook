package shellhook

type Metadata struct {
	// the path to the executable
	Executable string

	// the name of the hook
	Name string

	// the args to pass
	Args []string
}
