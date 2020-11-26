package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sharma1612harshit/docker.ui/api/docker"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/images", ImageHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	imageResponse, err := json.Marshal(docker.GetImages())

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		w.Write([]byte(imageResponse))
	}
}
