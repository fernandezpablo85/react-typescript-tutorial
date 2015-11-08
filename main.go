package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var appPath = os.Getenv("GOPATH") + "/src/github.com/fernandezpablo85/react-typescript-base"

func HomePage(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile(appPath + "/views/index.html")
	if err != nil {
		log.Panicf("error reading index.html: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

func main() {
	address := ":8080"
	mux := http.NewServeMux()

	statics := os.Getenv("GOPATH") + "/src/github.com/fernandezpablo85/react-typescript-base/public"
	fs := http.FileServer(http.Dir(statics))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", HomePage)
	log.Printf("server started on %s", address)
	http.ListenAndServe(address, mux)
}
