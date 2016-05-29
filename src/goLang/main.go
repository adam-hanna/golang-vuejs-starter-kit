package main

import (
	"io"
	"net/http"
	"log"
)

var host = "localhost";
var port = "8082";

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Printf("goLang server listening on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(host + ":" + port, nil))
}