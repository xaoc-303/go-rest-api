package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("homePage")
    fmt.Fprintf(w, "homePage")
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    fmt.Println("headers")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", homePage)
    router.HandleFunc("/hello", hello).Methods("GET")
    router.HandleFunc("/headers", headers).Methods("GET")

    articleHandleRequests(router)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
    handleRequests()
}
