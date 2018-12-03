package sources

import (
	"fmt"

	"github.com/alcmoraes/go-rom-downloader/domains"
	"github.com/gocolly/colly"
)

type CoolromSource struct {
	Endpoint  string
	UserAgent string
	LookupURL string
	c         *colly.Collector
}

func (self *CoolromSource) Lookup(name string) []domains.Rom {

	roms := []domains.Rom{}

	// Find and visit all links
	self.c.OnHTML("ul[data-role=listview] a", func(e *colly.HTMLElement) {
		roms = append(roms, *domains.CreateRom(
			e.ChildText("h3"),
			e.ChildText("p"),
			e.Attr("href"),
			"",
		))
	})

	// Do the first query
	self.c.Visit(fmt.Sprintf(self.Endpoint+self.LookupURL, name))

	self.c.Wait()

	return roms

}

func (self *CoolromSource) GetDownloadLink(rom *domains.Rom) string {
	self.c.OnHTML("form[name=dlform] input[name=id]", func(e *colly.HTMLElement) {
		romURL := fmt.Sprintf("http://m.coolrom.com.au/download?id=%s&type=rom", e.Attr("value"))
		rom.SetDownloadURL(romURL)
	})

	self.c.Visit(self.Endpoint + rom.URL)

	self.c.Wait()

	return rom.DownloadURL
}

func NewCoolromSource() *CoolromSource {
	return &CoolromSource{
		Endpoint:  "http://m.coolrom.com.au/",
		LookupURL: "search/?q=%s",
		c: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Linux; Android 6.0; SAMSUNG SM-G930F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36"),
		),
	}
}
