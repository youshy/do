package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("vim-go")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"hello":         "world",
			"are you bald?": "yes",
		})
	})

	log.Printf("server is running on port 9000")
	http.ListenAndServe(":9000", router)
}
