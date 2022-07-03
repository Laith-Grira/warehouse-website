package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	//get the value of the ADDR environment variable
	addr := os.Getenv("ADDR")

	// Set default port to 6000 if addr doesn't exist
	if len(addr) == 0 {
		addr = ":6000"
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(map[string]string{"data": "Testing GET method"})
	}).Methods("GET")

	fmt.Printf("Starting server at port%v\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
