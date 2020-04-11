package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("../01_stdout/jjp.ohtml")
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	if err = temp.Execute(newFile, nil); err != nil {
		log.Fatal(err)
	}

}
