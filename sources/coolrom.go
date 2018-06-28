package sources

import (
	"fmt"
	"regexp"

	"github.com/alcmoraes/go-cr-scraper/domains"
	"github.com/gocolly/colly"
)

type CoolromSource struct {
	Endpoint  string
	UserAgent string
	LookupURL string
}

func (self *CoolromSource) Lookup(name string) []domains.Rom {

	idExtractPatternFromDownload := regexp.MustCompile(`http://dfw.coolrom.com/dl/(.*?)/.*`)
	idExtractPatternFromList := regexp.MustCompile(`/roms/.*/(.*?)/`)

	c := colly.NewCollector(
		colly.UserAgent(self.UserAgent),
	)

	romsTemp := map[string]interface{}{}
	roms := []domains.Rom{}

	// Find and visit all links
	c.OnHTML("ul[data-role=listview] a", func(e *colly.HTMLElement) {
		gameID := idExtractPatternFromList.FindStringSubmatch(e.Attr("href"))[1]

		romsTemp[gameID] = map[string]string{
			"name":    e.ChildText("h3"),
			"console": e.ChildText("p"),
		}

		e.Request.Visit(e.Attr("href"))

	})

	c.OnHTML("form[name=dlform]", func(e *colly.HTMLElement) {
		gameID := idExtractPatternFromDownload.FindStringSubmatch(e.Attr("action"))[1]
		rom := romsTemp[gameID].(map[string]string)

		roms = append(roms, *domains.CreateRom(
			rom["name"],
			rom["console"],
			e.Attr("action"),
		))
	})

	// Do the first query
	c.Visit(fmt.Sprintf(self.Endpoint+self.LookupURL, name))

	c.Wait()

	return roms

}

func NewCoolromSource() *CoolromSource {
	return &CoolromSource{
		Endpoint:  "http://m.coolrom.com.au/",
		UserAgent: "Mozilla/5.0 (Linux; Android 6.0; SAMSUNG SM-G930F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36",
		LookupURL: "search/?q=%s",
	}
}
