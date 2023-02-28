package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type comment struct {
	ID      uint8  `json:"id"`
	Avatar  string `json:"avatar"`
	Message string `json:"message"`
	Name    string `json:"name"`
}
type picture struct {
	ID          uint8     `json:"id"`
	URL         string    `json:"url"`
	Likes       uint16    `json:"likes"`
	Comments    []comment `json:"comments"`
	Description string    `json:"description"`
}

type pictures []picture

var data pictures

func loadData(file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 10)
	fmt.Println(w, r.MultipartForm)
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}

func main() {
	server := http.Server{
		Addr: "localhost:3010",
	}

	loadData("data.json")

	http.HandleFunc("/processForm", processForm)
	http.HandleFunc("/data", getData)
	server.ListenAndServe()
}
