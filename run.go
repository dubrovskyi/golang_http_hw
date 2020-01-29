package main

import (
    "fmt"
    "net/http"
    "sync/atomic"
)

var requests int64 = 0

func incRequests() int64 {
    return atomic.AddInt64(&requests, 1)
}

func getRequests() int64 {
    return atomic.LoadInt64(&requests)
}

func handler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "%s", incRequests())

}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8082", nil)
}
