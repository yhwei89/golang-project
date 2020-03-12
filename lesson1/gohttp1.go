package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

//func handler1(w http.ResponseWriter,r *http.Request) {
//   fmt.Fprintf(w,"hello world")

func handler1(w http.ResponseWriter, r *http.Request) {
	num := math.Pow(10, 2)
	fmt.Fprintln(w, num)
}

func main() {
	log.Println("start http server on localhost:8080")
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}
