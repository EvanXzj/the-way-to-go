package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() (err error) {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

func (p *Page) load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(title)
	return err
}

func main() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("Page.md")
	fmt.Println(string(new_page.Body))
}
