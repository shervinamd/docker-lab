package main

import (
 "fmt"
 "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Hello, World!")
}

func main() {
 http.HandleFunc("/", helloHandler)

 // Listen on port 8080
 port := ":8080"
 fmt.Printf("Server listening on port %s...\n", port)
 err := http.ListenAndServe(port, nil)
 if err != nil {
  panic(err)
 }
}
