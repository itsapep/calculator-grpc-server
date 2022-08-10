package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/itsapep/calculator-grpc-server/api"
	"github.com/itsapep/calculator-grpc-server/calculator"
	"google.golang.org/grpc"
)

func main() {
	url := os.Getenv("GRPC_URL")
	serverInfo := fmt.Sprint(url)

	listen, err := net.Listen("tcp", serverInfo)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}

	s := calculator.Server{}
	grpcServer := grpc.NewServer()

	api.RegisterCalculatorServer(grpcServer, &s)
	log.Println("server run on ", serverInfo)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln("failed to serve ...", err)
	}
}
