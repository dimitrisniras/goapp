package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	goapp "goapp/internal/app/server"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmsgprefix | log.Lshortfile)
}

func main() {
	// Debug.
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	// Register signal handlers for exiting
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Parse the number of parallel websocket connections
	numOfConnections := flag.Int("n", 1, "Number of parallel WebSocket connections")
	flag.Parse()

	// Start.
	if err := goapp.Start(exitChannel, true, *numOfConnections); err != nil {
		log.Fatalf("fatal: %+v\n", err)
	}
}
