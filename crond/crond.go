package main

import (
	"crond/shell"
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)

const (
	HOST = "127.0.0.1"
	PORT = "9999"
)

func version() {
	fmt.Println("\nThis crond version is 1.0 !\n")
	os.Exit(1)
}

func parseArgs() (string, string) {
	host := flag.String("host", "", "run host example 127.0.0.1")
	port := flag.String("port", "", "listen port example 9999")
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

func createServer(host string, port string) *net.TCPListener {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err.Error())
	}
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	return listener
}

func main() {
	host, port := parseArgs()
	listener := createServer(host, port)

	server := rpc.NewServer()
	server.Register(new(shell.ExecCron))

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
