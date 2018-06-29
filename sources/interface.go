package sources

import "github.com/alcmoraes/go-rom-downloader/domains"

var RomSources = map[string]interface{}{
	"Coolrom": NewCoolromSource(),
	"Emuparadise": NewEmuparadiseSource(),
}

type Source interface {
	Lookup(name string) []domains.Rom
	GetDownloadLink(rom *domains.Rom) string
}

func LoadSource(s string, o map[string]interface{}) Source {
	return RomSources[s].(Source)
}
