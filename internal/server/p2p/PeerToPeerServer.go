package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tothbence9922/go-blockchain/internal/chain"
)

type PeerToPeerServer struct {
	Blockchain *chain.Chain
	Port       int
	Peers      []string
	Sockets    []*websocket.Conn
}

func (ptps PeerToPeerServer) ConnectToPeers(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		// Websocket client
		for _, peer := range ptps.Peers {
			if peer != "" {

				socket, _, err := websocket.DefaultDialer.Dial(peer, nil)
				if err != nil {
					fmt.Println("[DIAL]", err)
				}
				socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Hello, %s", peer)))
				ptps.Sockets = append(ptps.Sockets, socket)
			}
		}
	}()
	fmt.Println("WS Client started")
}

func handleWebsockets(w http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		fmt.Println("Upgrade failed: ", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		fmt.Println("Incoming message:\t", string(message))
		if err != nil {
			fmt.Println("read failed:", err)
			break
		}
		//input := string(message)
		//cmd := getCmd(input)
		//msg := getMessage(input)
		//if cmd == "add" {
		//	todoList = append(todoList, msg)
		//} else if cmd == "done" {
		//	updateTodoList(msg)
		//}
		//output := "Current Todos: \n"
		//for _, todo := range todoList {
		//	output += "\n - " + todo + "\n"
		//}
		//output += "\n----------------------------------------"
		//message = []byte(output)
		message = append(message, []byte(" :)")...)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write failed:", err)
			break
		}
	}
}

func (ptps *PeerToPeerServer) Init() {
	ptps.Blockchain = chain.GetInstance()
	envPort := os.Getenv("WS_PORT")
	port, err := strconv.Atoi(envPort)
	if err != nil {
		fmt.Println("Error converting environment WS_PORT string to int")
	}
	ptps.Port = port
	peersString := os.Getenv("WS_PEERS")
	ptps.Peers = strings.Split(peersString, ",")
	fmt.Println(ptps.Peers)
	ptps.Sockets = make([]*websocket.Conn, 0)
}

func (ptps PeerToPeerServer) CreateWebSocketServer(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		// Websocket Server
		server := new(http.Server)
		server.ReadTimeout = 5 * time.Second
		server.WriteTimeout = 5 * time.Second
		defer wg.Done()
		http.HandleFunc("/", handleWebsockets)

		portString := fmt.Sprintf(":%d", ptps.Port)
		http.ListenAndServe(portString, nil)
	}()
	fmt.Println("WS Server started")
}

func (ptps PeerToPeerServer) Listen(wg *sync.WaitGroup) {
	ptps.ConnectToPeers(wg)
	ptps.CreateWebSocketServer(wg)
}
