package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/alcmoraes/go-cr-scraper/sources"
	"github.com/cavaliercoder/grab"
)

func main() {
	source := sources.LoadSource("Coolrom", nil)
	fmt.Println("Enter the game you would like to search:")
	gameInputScanner := bufio.NewScanner(os.Stdin)
	for gameInputScanner.Scan() {
		gameToLookAt := gameInputScanner.Text()
		fmt.Println("Searching for " + gameToLookAt + " roms, wait a second...")
		roms := source.Lookup(url.QueryEscape(gameToLookAt))
		fmt.Printf("A total of %d roms found.\n", len(roms))
		if len(roms) == 0 {
			fmt.Println("Exiting...")
			os.Exit(1)
		}
gamechose:
		fmt.Println("============================")
		for i, rom := range roms {
			fmt.Printf("[%d] %s (%s)\n", i+1, rom.Name, rom.Console)
		}
		fmt.Printf("============================\n")
		fmt.Printf("Type the number of the game you want do download (eg.: 1) and press enter.\n")
		gameChosenScanner := bufio.NewScanner(os.Stdin)
		for gameChosenScanner.Scan() {
			gameChosen, err := strconv.ParseInt(gameChosenScanner.Text(), 10, 32)
			if err != nil || gameChosen < 1 || gameChosen > int64(len(roms) + 1) {
				fmt.Println("Invalid option.")
				goto gamechose
			}
			rom := roms[gameChosen-1]
			client := grab.NewClient()
			req, _ := grab.NewRequest(".", rom.URL)
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
				fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Download saved to ./%v \n", resp.Filename)
			os.Exit(1)
		}
	}
	os.Exit(1)
}
