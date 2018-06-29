package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/alcmoraes/go-rom-downloader/domains"
	"github.com/alcmoraes/go-rom-downloader/sources"
	"github.com/alcmoraes/go-rom-downloader/utils"
	"github.com/cavaliercoder/grab"
)

func SourceInput() string {

	sourcesMap := make([]string, len(sources.RomSources))
	i := 0
	for k := range sources.RomSources {
		sourcesMap[i] = k
		i++
	}

source_input:
	fmt.Println("From which source you would like to search?")

	for i, source := range sourcesMap {
		fmt.Printf("[%d] %s\n", i+1, source)
	}

	sourceInputScanner := bufio.NewScanner(os.Stdin)
	var sourceChosen int
	for sourceInputScanner.Scan() {
		sourceChosen, err := strconv.ParseInt(sourceInputScanner.Text(), 10, 32)
		if err != nil || sourceChosen < 1 || sourceChosen > int64(len(sourcesMap)) {
			utils.CallClear()
			fmt.Println("\rInvalid option.")
			goto source_input
		}
		return sourcesMap[sourceChosen-1]
	}
	return sourcesMap[sourceChosen-1]
}

func RomQueryInput() string {
	fmt.Println("Enter the game you would like to search.")
rom_query_input:
	romQueryScanner := bufio.NewScanner(os.Stdin)
	var romQueryChosen string
	for romQueryScanner.Scan() {
		romQueryChosen = romQueryScanner.Text()
		if len(romQueryChosen) == 0 {
			utils.CallClear()
			fmt.Println("\rInvalid query. Type it again.")
			goto rom_query_input
		}
		return romQueryChosen
	}
	return romQueryChosen
}

func ChooseRomInput(roms []domains.Rom) domains.Rom {
	fmt.Println("============================")
	for i, rom := range roms {
		fmt.Printf("[%d] %s (%s)\n", i+1, rom.Name, rom.Console)
	}
	fmt.Printf("============================\n")
	fmt.Printf("Type the number of the game you want do download (eg.: 13) and press enter.\n")
rom_choose_input:
	romChooseScanner := bufio.NewScanner(os.Stdin)
	var romChosen int
	for romChooseScanner.Scan() {
		romChosen, err := strconv.ParseInt(romChooseScanner.Text(), 10, 32)
		if err != nil || romChosen < 1 || romChosen > int64(len(roms)+1) {
			fmt.Println("Invalid option. Try again.")
			goto rom_choose_input
		}
		return roms[romChosen-1]
	}
	return roms[romChosen-1]
}

func Download(rom domains.Rom) {
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", rom.DownloadURL)
	fmt.Printf("Downloading %v...\n", rom.Name)
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("\rTransferred %v / %v bytes (%.2f%%)",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			break Loop
		}
	}
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "\nDownload failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nDownload saved to ./%v \n", resp.Filename)
}

func main() {
	utils.CallClear()
	fmt.Println("== GO ROM DOWNLOADER v1.0 ==")
	source := sources.LoadSource(SourceInput(), nil)
start:
	roms := source.Lookup(url.QueryEscape(RomQueryInput()))
	if len(roms) == 0 {
		utils.CallClear()
		fmt.Println("No roms match your query. Try again.")
		goto start
	}
	fmt.Printf("A total of %d roms found.\n", len(roms))
	romChosen := ChooseRomInput(roms)
	source.GetDownloadLink(&romChosen)
	Download(romChosen)
	os.Exit(0)
}
