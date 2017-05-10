package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mickelsonm/gcs-api/models/client"
)

func main() {
	router := mux.NewRouter()
	//just the default handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	//gets all clients
	router.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		clients, err := client.GetAllClients()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		toJSON(w, clients)
	})

	//gets a specific client by the id
	router.HandleFunc("/clients/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		vars := mux.Vars(r)
		id := vars["id"]
		cl, err := client.GetClientByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		toJSON(w, cl)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func toJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
