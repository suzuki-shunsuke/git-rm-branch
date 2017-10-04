package models

type Cfg struct {
	Confirm bool
	Local   []string
	Remote  map[string][]string
}
