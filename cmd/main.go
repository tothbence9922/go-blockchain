package main

import (
	"sync"

	http "github.com/tothbence9922/go-blockchain/internal/server/http"
	p2p "github.com/tothbence9922/go-blockchain/internal/server/p2p"
)

var (
	wg sync.WaitGroup
)

func main() {

	p2p.GetInstance().Listen(&wg)
	http.GetInstance().Serve(&wg)

	wg.Wait()
}
