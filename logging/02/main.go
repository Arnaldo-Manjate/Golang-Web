package main

import (
	"log"
	"os"
)

func init() {
	newFile, err := os.Create("log.text")
	if err != nil {
		log.Println(err)
	}

	// the output file that the errors will be written to
	log.SetOutput(newFile)
}

func main() {
	// deleberately causing an error
	_, err := os.Open("bomb.text")
	if err != nil {
		log.Println(err)
	}

}
