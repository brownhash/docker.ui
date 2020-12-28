package main

import (
	"github.com/sharma1612harshit/docker.ui/api/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/images", handler.ImagesHandler)
	http.HandleFunc("/pullimage", handler.PullImageHandler)
	http.HandleFunc("/containers", handler.ContainersHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
