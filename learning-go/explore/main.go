package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Point Hit")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	fmt.Println("Update Successfully")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	fmt.Println("Delete Successfully")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Check")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/all/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/all/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("REST API MUX Router")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article", Content: "Content"},
		Article{Id: "2", Title: "Article", Desc: "Article", Content: "Test"},
	}
	handleRequest()
}
