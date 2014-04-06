package main

import (
	"encoding/json"
	"fmt"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/reco"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RecommendationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ureco := make(map[int]data.Recommendation)
	ureco, _ = reco.UserRunner()
	//ireco, ireco2 := reco.ItemRunner()

	slc := make([][2]interface{}, 0, 50)
	for _, val := range ureco {
		slcval := [2]interface{}{val.Book, val.Score}
		slc = append(slc, slcval)
	}
	fmt.Println(slc)
	json, err := json.Marshal(slc)
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(json))
}

func main() {
	log.Println("Starting Server")

	r := mux.NewRouter()
	r.HandleFunc("/api/recommendation", RecommendationHandler).Methods("GET")
	http.Handle("/api/", r)

	http.Handle("/", http.FileServer(http.Dir("../app/")))

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
