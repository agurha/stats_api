package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

type Series struct {
}

type Events struct {
}

func PostSeriesMetric(w http.ResponseWriter, r *http.Request) {

	//Series Post has to be a json array (as to enable customers past multiple metrics at the same time) with following required fields
	/*
		api_key should be part of HTTP header

		So post to https://api.statsbox.io/v/series with following Data

		{
			"series" : [
				{
					"metricname" : "cpu"
					"points" : [[1223232, 10], [2332323, 13]]	  -- this should be json array of points with timestamp and value
					"type" : "gauge" -- It can be gauge or counter -- this is optional
					"host" : "some_prod_server" -- this is optional
					"tags" : "None" -- this is optional

				}
			]
		}

	*/

	fmt.Println("Post Series")
}

func GetSeriesMetric(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET Series")
}

func PostEvents(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Post Some Events")
}

func GetEvents(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get Some Events")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, Response{"success": true, "message": "Welcome to the STATS API SERVER!", "method": r.Method})
	})

	s := r.PathPrefix("/v1/series").Subrouter()
	s.HandleFunc("/", PostSeriesMetric).Methods("POST").Schemes("http")
	s.HandleFunc("/", GetSeriesMetric).Methods("GET").Schemes("http")

	e := r.PathPrefix("/v1/events").Subrouter()
	e.HandleFunc("/", PostEvents).Methods("POST").Schemes("http")
	e.HandleFunc("/{event_id}", GetEvents).Methods("GET").Schemes("http")

	log.Fatal(http.ListenAndServe(":8080", r))

}
