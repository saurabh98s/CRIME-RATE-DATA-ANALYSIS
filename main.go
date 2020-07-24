package main

import (
	"fmt"
	"github.com/geeks/miniproject/logger"
	"html/template"
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
	err := tpl.ExecuteTemplate(w, "index2.html", nil)
	if err != nil {
		logger.Log.Info(err)
	}

}

func postMethodPage(w http.ResponseWriter, r *http.Request) {

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
