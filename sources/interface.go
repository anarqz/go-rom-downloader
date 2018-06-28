package sources

import "github.com/alcmoraes/go-rom-downloader/domains"

type Source interface {
	Lookup(name string) []domains.Rom
	GetDownloadLink(rom *domains.Rom) string
}

func LoadSource(s string, o map[string]interface{}) Source {

	source := map[string]interface{}{
		"Coolrom": NewCoolromSource(),
	}

	return source[s].(Source)

}
