package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getMainPage).Methods("Get")
	router.HandleFunc("/predictions", getPredictionPage).Methods("Get")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Panic(err)
	}

}

func getMainPage(w http.ResponseWriter, r *http.Request) {
	getLoginPage, err := template.ParseFiles("index2.html")
	if err != nil {
		log.Println(err)
	}
	getLoginPage.Execute(w, nil)

}

func getPredictionPage(w http.ResponseWriter, r *http.Request) {
	getLoginPage, err := template.ParseFiles("prediction.html")
	if err != nil {
		log.Println(err)
	}
	getLoginPage.Execute(w, nil)

}
