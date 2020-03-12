package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")

}

func main() {
	log.Println("start http server on localhost:8080")
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}
