package main

import (
	"fmt"

	"github.com/alcmoraes/go-cr-scraper/sources"
)

func main() {
	source := sources.LoadSource("Coolrom", nil)

	roms := source.Lookup("Castlevania")

	for _, rom := range roms {
		fmt.Printf("%s (%s): %s\n", rom.Name, rom.Console, rom.URL)
	}
}
