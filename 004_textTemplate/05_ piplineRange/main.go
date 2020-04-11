package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	temp, err := template.ParseFiles("list.html")
	if err != nil {
		log.Fatalln(err)
	}

	list := []string{"Arnaldo", "Andile"}

	temp.Execute(os.Stdout, list)
}
