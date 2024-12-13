package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (pg *Page) save() (err error) {
	outputName := "./chapter_12/" + pg.Title
	return ioutil.WriteFile(outputName, pg.Body, 0666)
}

func (pg *Page) load(title string) (err error) {
	pg.Title = title
	pg.Body, err = ioutil.ReadFile("./chapter_12/" + title)
	return err
}

func mainWikiPart() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	var newPage Page
	newPage.load("Page.md")
	fmt.Println(newPage.Body)
}
