package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("spp.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
