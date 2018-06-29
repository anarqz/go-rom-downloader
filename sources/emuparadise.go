package sources

import (
	"fmt"

	"github.com/alcmoraes/go-rom-downloader/domains"
	"github.com/gocolly/colly"
)

type EmuparadiseSource struct {
	Endpoint  string
	UserAgent string
	LookupURL string
	c         *colly.Collector
}

func (self *EmuparadiseSource) Lookup(name string) []domains.Rom {

	roms := []domains.Rom{}

	// Find and visit all links
	self.c.OnHTML("#content .roms", func(e *colly.HTMLElement) {
		roms = append(roms, *domains.CreateRom(
			e.ChildText("a[data-filter]"),
			e.ChildText("a.sysname"),
			e.ChildAttr("a[data-filter]", "href"),
			"",
		))
	})

	// Do the first query
	self.c.Visit(fmt.Sprintf(self.Endpoint+self.LookupURL, name))

	self.c.Wait()

	return roms

}

func (self *EmuparadiseSource) GetDownloadLink(rom *domains.Rom) string {

	self.c.OnHTML("div.download-link a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	self.c.OnHTML("a#download-link", func(e *colly.HTMLElement) {
		rom.SetDownloadURL(self.Endpoint + e.Attr("href")[1:])
	})

	self.c.Visit(self.Endpoint + rom.URL)

	self.c.Wait()

	return rom.DownloadURL
}

func NewEmuparadiseSource() *EmuparadiseSource {
	return &EmuparadiseSource{
		Endpoint:  "https://m.emuparadise.me/",
		LookupURL: "roms/search.php?query=%s",
		c: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Linux; Android 6.0; SAMSUNG SM-G930F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36"),
		),
	}
}
