package main

import (
	"sync"

	"github.com/tothbence9922/go-blockchain/internal/server"
)

var (
	wg sync.WaitGroup
)

func main() {

	server := server.HttpServer{Port: 8080}

	server.Serve(&wg)

	wg.Wait()
}
