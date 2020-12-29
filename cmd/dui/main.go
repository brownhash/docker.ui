package main

import (
	"flag"
	"fmt"
	"github.com/sharma1612harshit/docker.ui/api/handler"
	"log"
	"net/http"
)

const(
	nodeAddr = "0.0.0.0"
	workerPort = "8080"
)

func main() {
	c := flag.Bool("c", false, "use -c to start dui controller node")
	control := flag.Bool("control", false, "use -control to start dui controller node")

	w := flag.Bool("w", false, "use -w to start dui worker node")
	work := flag.Bool("work", false, "use -work to start dui worker node")

	flag.Parse()

	if *w || *work {
		log.Print("Initiating docker.ui worker node...")
		http.HandleFunc("/images", handler.ImagesHandler)
		http.HandleFunc("/pullimage", handler.PullImageHandler)
		http.HandleFunc("/containers", handler.ContainersHandler)

		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", nodeAddr, workerPort), nil))
	} else if *c || *control {
		log.Print("Server not yet coded...")
	}
}
