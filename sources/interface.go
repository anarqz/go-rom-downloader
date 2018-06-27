package sources

import "github.com/alcmoraes/go-cr-scraper/domains"

type Source interface {
	Lookup(name string) []domains.Rom
}

func LoadSource(s string, o map[string]interface{}) Source {

	source := map[string]interface{}{
		"Coolrom": NewCoolromSource(),
	}

	return source[s].(Source)

}
