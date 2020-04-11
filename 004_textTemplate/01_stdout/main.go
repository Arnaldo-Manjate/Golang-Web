package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("spp.gohtml", "jjp.ohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "jjp.ohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
