package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoErr := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	infoErr.Fatal(err)
}
