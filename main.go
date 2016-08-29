package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
		logger.WithField("template", "index").Errorf("template not found")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()

	handler := http.StripPrefix("/public", http.FileServer(http.Dir("./public")))
	router.PathPrefix("/public").Handler(handler)
	router.HandleFunc("/", HomePage).Methods("GET")

	address := ":8080"
	logger.Infof("server started on %s", address)
	logger.Fatal(http.ListenAndServe(address, router))
}
