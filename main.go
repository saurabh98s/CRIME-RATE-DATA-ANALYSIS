package main

import (
	"fmt"
	"github.com/geeks/miniproject/logger"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
var db *mgo.Database

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	logger.Log.Info("Connecting to mongodb...")
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		logger.Log.Fatal(err)
	}
	db = session.DB("miniproject")

	logger.Log.Info("Connected to mongodb successfully...")

}

func DB() *mgo.Database {
	return db
}

func servingLoginPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		logger.Log.Info(err)
	}

}

func postMethodPage(w http.ResponseWriter, r *http.Request) {
	var data []*AutoCrimeData
	c:=db.C("autocrime")
	if err := r.ParseForm(); err != nil {
		logger.Log.Info(w, "ParseForm() err: %v", err)
		return
	}
	userInput := r.FormValue("userInput")
	query := bson.M{"Year": userInput}
	err:=c.Find(query).All(&data)
	if err == nil {
		logger.Log.Info(err)

	}
	logger.Log.Print(data)
	err= tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		logger.Log.Info(err)
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
