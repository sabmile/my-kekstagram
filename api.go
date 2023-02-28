package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 10)
	fmt.Println(w, r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3010",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
