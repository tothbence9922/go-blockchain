package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/tothbence9922/go-blockchain/internal/chain"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

type HttpServer struct {
	Port int
}

func handleGetBlocks(w http.ResponseWriter, req *http.Request) {
	blockChain := chain.GetInstance()
	outJson, _ := json.Marshal(blockChain)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow", http.MethodGet)
	fmt.Fprintf(w, string(outJson))
}

func handlePostBlock(w http.ResponseWriter, req *http.Request) {

	var newContent content.Content

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "An error occured while reading the POST-ed Data.")
	}
	json.Unmarshal(reqBody, &newContent)

	blockChain := chain.GetInstance()
	fmt.Println(newContent.Value)
	blockChain.AddBlock(newContent)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow", http.MethodGet)
	outJson, _ := json.Marshal(blockChain)
	fmt.Fprintf(w, string(outJson))
}

func handleBlocks(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handleGetBlocks(w, req)
	case "POST":
		handlePostBlock(w, req)
	}
}

func (hs HttpServer) Serve(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		server := new(http.Server)
		server.ReadTimeout = 5 * time.Second
		server.WriteTimeout = 5 * time.Second
		defer wg.Done()
		http.HandleFunc("/blocks", handleBlocks)

		portString := fmt.Sprintf(":%d", hs.Port)
		http.ListenAndServe(portString, nil)
	}()
	fmt.Println("HTTP Server started")
}
