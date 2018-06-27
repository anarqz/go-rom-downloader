package domains

import "fmt"

type Rom struct {
	Name    string
	Console string
	URL     string
}

// Download the rom
func (self *Rom) Download() {
	fmt.Println(self.Name)
}

func CreateRom(name string, console string, url string) *Rom {
	return &Rom{name, console, url}
}