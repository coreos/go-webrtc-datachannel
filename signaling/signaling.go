package main

import (
	"flag"
	"fmt"
	"github.com/ccding/go-logging/logging"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
)

var (
	hostname   string
	port       int
	dir        string
	logger     *logging.Logger
	cpuprofile string
)

func init() {
	flag.StringVar(&hostname, "h", "0.0.0.0", "the host name of the signaling server")
	flag.IntVar(&port, "p", 7446, "the port number of the signaling server")
	flag.StringVar(&dir, "d", "/tmp/", "the directory to store log and snapshot")

	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")

	var err error
	logger, err = logging.RichLogger("signaling")
	if err != nil {
		logger.Critical(err)
		panic("creating logger error")
	}
	logger.SetLevel(logging.NOTSET)
}

func main() {
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			logger.Critical(err)
			panic("creating logger error")
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				fmt.Printf("captured %v, stopping profiler and exiting..", sig)
				pprof.StopCPUProfile()
				os.Exit(1)
			}
		}()

	}
	logger.Debug("server starting")
	startServer()
}

// start the signaling server
func startServer() {
	http.HandleFunc("/set/", SetHttpHandler)
	http.HandleFunc("/get/", GetHttpHandler)
	http.HandleFunc("/", VersionHttpHandler)

	fmt.Printf("signaling server [%s] listen on http port %v\n", hostname, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		logger.Critical(err)
		panic("signaling server unknown error")
	}
}
