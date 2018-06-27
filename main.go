package main

import (
	"fmt"
	"net/url"

	"github.com/alcmoraes/go-cr-scraper/sources"
)

func main() {
	source := sources.LoadSource("Coolrom", nil)

	var gameToLookAt string
	fmt.Println("Enter the game you would like to search:")
	fmt.Scanf("%s", &gameToLookAt)

	roms := source.Lookup(url.QueryEscape(gameToLookAt))

	for _, rom := range roms {
		fmt.Printf("%s (%s): %s\n", rom.Name, rom.Console, rom.URL)
	}
}
