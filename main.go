package main

import (
	"./internals"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/containers/getContainers", internals.ContainerHandler)
	http.HandleFunc("/images/deleteImages", internals.ImageDeletionHandler)
	http.HandleFunc("/images/getImages", internals.ImageHandler)
	http.HandleFunc("/", internals.DashboardHandler)

	err := http.ListenAndServe("0.0.0.0:9000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
