package models

type Cfg struct {
	Local  map[string][]string
	Remote map[string]map[string][]string
}
