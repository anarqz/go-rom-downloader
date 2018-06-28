package domains

type Rom struct {
	Name        string
	Console     string
	URL         string
	DownloadURL string
}

// SetDownloadURL from the rom
func (self *Rom) SetDownloadURL(url string) {
	self.DownloadURL = url
}

func CreateRom(name string, console string, url string, downloadUrl string) *Rom {
	return &Rom{name, console, url, downloadUrl}
}
