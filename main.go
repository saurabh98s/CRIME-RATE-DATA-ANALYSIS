package main

import (
	"encoding/json"
	"fmt"
	"github.com/geeks/miniproject/logger"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AutoCrimeData contains all the parsed data from Json file
type AutoCrimeData []struct {
	AreaName                   string `json:"Area_Name"`
	Year                       int    `json:"Year"`
	GroupName                  string `json:"Group_Name"`
	SubGroupName               string `json:"Sub_Group_Name"`
	AutoTheftCoordinatedTraced int    `json:"Auto_Theft_Coordinated/Traced"`
	AutoTheftRecovered         int    `json:"Auto_Theft_Recovered"`
	AutoTheftStolen            int    `json:"Auto_Theft_Stolen"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func servingLoginPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println(err)
	}

}

func postMethodPage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Log.Info(w, "ParseForm() err: %v", err)
		return
	}
	userInput:=r.FormValue("userInput")

	file, _ := ioutil.ReadFile("Auto_theft.json")
	var dataFromS3 AutoCrimeData
	err := json.Unmarshal(file, &dataFromS3)
	if err != nil {
		log.Fatalln(err)
	}

	for key, val := range dataFromS3 {
		fmt.Println(key, val)

	}
	if err != nil {
		log.Println(err)
	}
	err = tpl.ExecuteTemplate(w, "index.html", dataFromS3)
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
