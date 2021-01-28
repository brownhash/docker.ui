package main

import (
	"flag"
	"fmt"
	"github.com/sharma1612harshit/docker.ui/api/handler"
	"github.com/sharma1612harshit/docker.ui/pkg/logger"
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

	// util.OpenLogFile(util.LogPath)

	if *w || *work {
		logger.Info("Initiating docker.ui worker node...")

		http.HandleFunc("/images", handler.ImagesHandler)
		http.HandleFunc("/pull_image", handler.PullImageHandler)

		http.HandleFunc("/containers", handler.ContainersHandler)
		http.HandleFunc("/run_container", handler.LaunchContainerHandler)

		err := http.ListenAndServe(fmt.Sprintf("%s:%s", nodeAddr, workerPort), logger.LogRequest(http.DefaultServeMux))
		logger.Error(err)

	} else if *c || *control {
		logger.Info("Server not yet coded...")
	}
}
