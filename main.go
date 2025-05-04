package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte //slice
}

// This is a method it takes as it's receiver a pointer to a page, it takes no parameters and returns an error.
// It saves the page's body to a text file. It returns an error because that is the return type of WriteFile.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// This is a function it loads a page, Readfile returns either a byte slice or an error.
// We are returning an error if it fails or a pointer to the new page struct by reference.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body, Title: title}, nil
}

func main() {
	p1 := &Page{Title: "Sample Page", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("Sample Page")
	fmt.Println(string(p2.Title) + ": " + string(p2.Body))
}
