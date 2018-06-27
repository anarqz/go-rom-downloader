package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/alcmoraes/go-cr-scraper/sources"
	"github.com/alcmoraes/go-cr-scraper/utils"
)

func main() {
	source := sources.LoadSource("Coolrom", nil)

	var holdTmp string
	var outputFile string

	fmt.Println("Enter the game you would like to search:")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		gameToLookAt := scanner.Text()

		outputFile = utils.Slugify(gameToLookAt, true) + ".json"

		roms := source.Lookup(url.QueryEscape(gameToLookAt))

		if output, err := json.MarshalIndent(roms, "", "\t"); err == nil {
			if err := ioutil.WriteFile(outputFile, output, 0644); err == nil {
				fmt.Println("Download links stored in " + outputFile + " file. You can close this window now")
				fmt.Scanf("%s", &holdTmp)
			}
		}

	}

	os.Exit(1)

}
