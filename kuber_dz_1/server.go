package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, `{"status": "OK"}`)
}

func main() {

    http.HandleFunc("/health", hello)

    http.ListenAndServe(":8000", nil)
}