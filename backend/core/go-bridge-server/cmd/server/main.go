package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "go-bridge-server/internal"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/eth/relay", internal.HandleEthRelay).Methods("POST")
    r.HandleFunc("/sol/relay", internal.HandleSolRelay).Methods("POST")

    fmt.Println("Bridge server listening on : http://localhost:8084")
    http.Handle("/", r)
    if err := http.ListenAndServe(":8084", nil); err != nil {
        panic(err)
    }
}