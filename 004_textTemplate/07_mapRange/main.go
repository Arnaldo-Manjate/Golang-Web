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

	team := map[string]string{
		"Boss":    "Arnaldo",
		"Manager": "Jiendor",
		"Worker":  "Punk",
	}

	if err = temp.Execute(os.Stdout, team); err != nil {
		log.Fatalln(err)
	}

}
