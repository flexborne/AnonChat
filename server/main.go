package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var app = NewApp()
	app.runMatchmaking()

	go app.handleNewFindMatch()

	http.HandleFunc("/ws", app.HandleConnections)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
