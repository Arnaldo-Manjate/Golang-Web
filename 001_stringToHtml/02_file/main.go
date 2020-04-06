package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Arnaldo"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1> My name is ` +
		name +
		`</h1>
		</body>
		</html>`)

	newFile, err := os.Create("someFile.html")
	if err != nil {
		log.Fatal("Error creating File", err)
	}
	defer newFile.Close()
	// copy string to the file
	io.Copy(newFile, strings.NewReader(str))

}
