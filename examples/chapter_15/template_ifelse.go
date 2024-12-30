package main

import (
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	// if 是一个空管道时的内容
	tEmpty = template.Must(tEmpty.
		Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	// 如果条件满足，则为非空管道
	tWithValue = template.Must(tWithValue.
		Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	// 如果条件满足，则为非空管道
	tIfElse = template.Must(tIfElse.
		Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
