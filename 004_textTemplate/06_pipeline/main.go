package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("Christmas.temp")
	if err != nil {
		log.Fatalln(err)
	}

	friends := []string{"Arnaldo", "Touch", "Keke"}
	if err = temp.Execute(os.Stdout, friends); err != nil {
		log.Fatalln(err)
	}

}
