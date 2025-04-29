package main

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte //slice
}

// This method takes pointer to a a page a page, it takes no parameters and returns an error.
// It saves the page's body to a text file. It returns an error because that is the return type of WriteFile
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func main() {
	p := &Page{Title: "Hydrogen", Body: []byte("H is the first element.")}
	err := p.save()
	if err != nil {
		panic(err)
	}
}
