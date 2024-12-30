package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name                string
	nonExportedAgeField string
}

func main() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")
	p := Person{Name: "Mary", nonExportedAgeField: "31"} // data

	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
