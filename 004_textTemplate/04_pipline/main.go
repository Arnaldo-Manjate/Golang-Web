package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("pipline.gohtml"))
}

func main() {
	// you could diplay the users data instaed od 26
	tpl.Execute(os.Stdout, 26)
}
