package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
)

func main() {
    http.HandleFunc("/downstreamGo", handler) // each request calls handler
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handler Called 1")
    resp, err := http.Get("http://upstreamNode:8080/upstreamNode")
    check(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "%v", string(body))
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}