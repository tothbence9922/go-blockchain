package main

import (
	"os"
	"strconv"
	"sync"

	http "github.com/tothbence9922/go-blockchain/internal/server/http"
	p2p "github.com/tothbence9922/go-blockchain/internal/server/p2p"
)

var (
	wg       sync.WaitGroup
	portEnv  int
	peersEnv string
)

func init() {
	portEnv, _ = strconv.Atoi(os.Getenv("WS_PORT"))
	peersEnv = os.Getenv("PEERS")
}
func main() {

	peerToPeerServer := p2p.PeerToPeerServer{}
	peerToPeerServer.Init()
	peerToPeerServer.Listen(&wg)

	server := http.HttpServer{Port: 8080}

	server.Serve(&wg)

	wg.Wait()
}
