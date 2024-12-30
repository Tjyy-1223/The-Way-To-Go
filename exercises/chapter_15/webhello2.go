package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HelloServer1(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler - name")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[len("/hello/"):])
}

func HelloServer2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler - NAME")
	fmt.Fprintf(w, "Hello,"+strings.ToUpper(req.URL.Path[len("/shouthello/"):]))
}

func main() {
	http.HandleFunc("/hello/", HelloServer1)
	http.HandleFunc("/shouthello/", HelloServer2)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}
