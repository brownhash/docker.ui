package main

import (
	"./internals"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/images", internals.ImageHandler)
	http.HandleFunc("/", internals.DashboardHandler)

	err := http.ListenAndServe("0.0.0.0:9000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
