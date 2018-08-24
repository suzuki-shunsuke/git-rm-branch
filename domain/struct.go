package domain

type (
	// Cfg represents a configuration.
	Cfg struct {
		Local  map[string][]string
		Remote map[string]map[string][]string
	}
)
