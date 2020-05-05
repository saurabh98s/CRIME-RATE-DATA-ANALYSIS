package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AutoCrimeData []struct {
	AreaName                   string `json:"Area_Name"`
	Year                       int    `json:"Year"`
	GroupName                  string `json:"Group_Name"`
	SubGroupName               string `json:"Sub_Group_Name"`
	AutoTheftCoordinatedTraced int    `json:"Auto_Theft_Coordinated/Traced"`
	AutoTheftRecovered         int    `json:"Auto_Theft_Recovered"`
	AutoTheftStolen            int    `json:"Auto_Theft_Stolen"`
}

func servingLoginPage(w http.ResponseWriter, r *http.Request) {
	loginPage, err := template.ParseFiles("login.html")
	if err != nil {
		log.Println(err)
	}
	err = loginPage.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func postMethodPage(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("Auto_theft.json")
	var dataFromS3 AutoCrimeData
	err := json.Unmarshal(file, &dataFromS3)
	if err != nil {
		log.Fatalln(err)
	}

	for key, val := range dataFromS3 {
		fmt.Println(key, val)

	}
	mainPage, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
	}
	err = mainPage.Execute(w, dataFromS3)
	if err != nil {
		log.Println(err)
	}
}

func main() {

	router := mux.NewRouter()
	fmt.Println("Starting server for testing")
	router.HandleFunc("/", servingLoginPage).Methods("Get")
	router.HandleFunc("/data", postMethodPage).Methods("Post")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println(err)
		return
	}

}
