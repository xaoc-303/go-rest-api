package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/gorilla/mux"
    "strconv"
)

type Article struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func getArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getArticle")

    vars := mux.Vars(r)
    id := vars["id"]

    fmt.Println("id: " + id)

    for _, article := range Articles {
        if article.Id == id {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("deleteArticle")

    vars := mux.Vars(r)
    id := vars["id"]

    fmt.Println("id: " + id)

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

func newArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("newArticle")
    reqBody, _ := ioutil.ReadAll(r.Body)

    var article Article

    json.Unmarshal(reqBody, &article)

    id_last, _ := strconv.Atoi(Articles[len(Articles) - 1].Id)
    article.Id = strconv.Itoa(id_last + 1)

    Articles = append(Articles, article)

    fmt.Printf("new: %+v\n", article)
    json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("updateArticle")

    vars := mux.Vars(r)
    id := vars["id"]

    var article_modify Article
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &article_modify)
    article_modify.Id = id

    for index, article := range Articles {
        if article.Id == id {
            Articles[index] = Article{Id: id, Title: article_modify.Title, Desc: article_modify.Desc, Content: article_modify.Content}
            fmt.Printf("old: %+v\nnew: %+v\n", article, article_modify)
        }
    }

    json.NewEncoder(w).Encode(article_modify)
}

func allArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("allArticles")
    json.NewEncoder(w).Encode(Articles)
}

func initArticles() {
   Articles = []Article{
       Article{Id: "1", Title: "Title1", Desc: "Desc1", Content: "Content1"},
       Article{Id: "2", Title: "Title2", Desc: "Desc2", Content: "Content2"},
   }
}

func articleHandleRequests(router *mux.Router) {
    initArticles()
    router.HandleFunc("/articles", allArticles).Methods("GET")
    router.HandleFunc("/article", newArticle).Methods("POST")
    router.HandleFunc("/article/{id}", getArticle).Methods("GET")
    router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
}
