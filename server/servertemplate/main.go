package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/temp", ServeTemplate)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server Error %v", err)
	}
}

func foo(res http.ResponseWriter, req *http.Request) {
	// write to the response
	io.WriteString(res, "Claudia's website ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func ServeTemplate(res http.ResponseWriter, req *http.Request) {
	temp := template.Must(template.ParseFiles("../../004_textTemplate/02_file/index.html"))
	// accessing the response headers to change the content type
	res.Header().Set("content-type", "text/html;charset-utf8")
	if err := temp.Execute(res, temp); err != nil {
		log.Fatalln("Error Executing Template ,", err)
	}
}
