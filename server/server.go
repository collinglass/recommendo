package main

import (
	"encoding/json"
	"fmt"
	"github.com/collinglass/recommendo/reco"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RecommendationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ureco, ureco2 := reco.UserRunner()
	ireco, ireco2 := reco.ItemRunner()

	slc := make([][2]interface{}, 0, 50)
	slcval := [2]interface{}{"Book", "Rating"}
	slc = append(slc, slcval)
	for _, val := range ureco {
		slcval = [2]interface{}{val.Book, val.Score}
		slc = append(slc, slcval)
	}

	slc2 := make([][2]interface{}, 0, 50)
	slcval = [2]interface{}{"Book", "Rating"}
	slc2 = append(slc2, slcval)
	for _, val := range ureco2 {
		slcval = [2]interface{}{val.Book, val.Score}
		slc2 = append(slc2, slcval)
	}

	slc3 := make([][2]interface{}, 0, 50)
	slcval = [2]interface{}{"Book", "Rating"}
	slc3 = append(slc3, slcval)
	for _, val := range ireco {
		slcval = [2]interface{}{val.Book, val.Score}
		slc3 = append(slc3, slcval)
	}

	slc4 := make([][2]interface{}, 0, 50)
	slcval = [2]interface{}{"Book", "Rating"}
	slc4 = append(slc4, slcval)
	for _, val := range ireco2 {
		slcval = [2]interface{}{val.Book, val.Score}
		slc4 = append(slc4, slcval)
	}

	metaslc := [4]interface{}{slc, slc2, slc3, slc4}

	json, err := json.Marshal(metaslc)
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
