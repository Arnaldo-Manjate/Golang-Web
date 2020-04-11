package main

import (
	"log"
	"os"
	"text/template"
)

// declare here in the package level scope to be able to use anywhere
var tpl *template.Template

func init() {
	// must handles error for you
	tpl = template.Must(template.ParseGlob("../01_stdout/*.gohtml"))
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}
}

func main() {

	if err := tpl.ExecuteTemplate(os.Stdout, "jjp.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "spp.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

}
