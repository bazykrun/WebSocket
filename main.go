package main

import (
    "fmt"
    "net/http"
    "log"
) 

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "WebSocket Endpoint")
}

func setupRoutes() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
}

func main() {
    fmt.Println("who is better there? yes i'm, only who win witch myself")
    setupRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil))
}

