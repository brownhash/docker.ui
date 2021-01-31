package main

import (
	"flag"
	"fmt"
	"github.com/sharma1612harshit/docker.ui/api/handler"
	"github.com/sharma1612harshit/docker.ui/pkg/config"
	"github.com/sharma1612harshit/gomuf/logs"
	"net/http"
)

func main() {
	c := flag.Bool("c", false, "use -c to start dui controller node")
	control := flag.Bool("controler", false, "use -control to start dui controller node")

	w := flag.Bool("w", false, "use -w to start dui worker node")
	work := flag.Bool("worker", false, "use -work to start dui worker node")

	logLevel := flag.String("logLevel", "INFO", "use -logLevel to indicate logging level. DEBUG | INFO | WARNING | ERROR")

	flag.Parse()

	logs.SetLogLevel(*logLevel)

	if *w || *work {
		logs.Success(config.DuiLogo)
		logs.Info("Initiating docker.ui worker node...")

		http.HandleFunc("/images", handler.ImagesHandler)
		http.HandleFunc("/pull_image", handler.PullImageHandler)

		http.HandleFunc("/containers", handler.ContainersHandler)
		http.HandleFunc("/run_container", handler.LaunchContainerHandler)

		err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.NodeAddr, config.WorkerPort), logs.LogRequest(http.DefaultServeMux))
		logs.Error(err)

	} else if *c || *control {
		logs.Success(config.DuiLogo)
		logs.Info("Server not yet coded...")
	}
}
