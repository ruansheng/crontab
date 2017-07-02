package main

import (
	"cront/model"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	HOST = "127.0.0.1"
	PORT = "9999"
)

func version() {
	fmt.Println("\nThis cront version is 1.0 !\n")
	os.Exit(1)
}

func parseArgs() (string, string) {
	host := flag.String("host", "", "crond host example 127.0.0.1")
	port := flag.String("port", "", "crond port example 9999")
	flag.Parse()
	var cmd string = flag.Arg(0)
	if strings.EqualFold(cmd, "version") || strings.EqualFold(cmd, "Version") {
		version()
	}
	if strings.EqualFold(*host, "") {
		*host = HOST
	}
	if strings.EqualFold(*port, "") {
		*port = PORT
	}
	return *host, *port
}

func main() {
	var wg sync.WaitGroup

	host, port := parseArgs()

	results := model.GetTaskData()
	for _, result := range results {
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go model.PushToCrond(&wg, host, port, result)
		}
	}
	wg.Wait()
}
