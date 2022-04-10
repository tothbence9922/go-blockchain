package server

import (
	"encoding/json"
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

var (
	p2pServer *PeerToPeerServer
)

func New() *PeerToPeerServer {
	envPort := os.Getenv("WS_PORT")
	port, err := strconv.Atoi(envPort)
	if err != nil {
		fmt.Println("Error converting environment WS_PORT string to int")
		port = 20
	}
	peersString := os.Getenv("WS_PEERS")
	var peers []string
	if len(peersString) > 1 {
		peers = strings.Split(peersString, ",")
	}
	sockets := make([]*websocket.Conn, 0)
	blockchain := chain.GetInstance()
	return &PeerToPeerServer{Blockchain: blockchain, Port: port, Peers: peers, Sockets: sockets}
}

func GetInstance() *PeerToPeerServer {
	if p2pServer == nil {
		p2pServer = New()
	}

	if len(p2pServer.Sockets) == 0 {
		p2pServer.ConnectToPeers()
	}

	return p2pServer
}

func (ptps *PeerToPeerServer) ConnectToPeers() {
	go func() {
		// Websocket client
		for _, peer := range ptps.Peers {
			fmt.Println(peer)
			socket, _, err := websocket.DefaultDialer.Dial(peer, nil)
			if err != nil {
				fmt.Println("[DIAL]", err)
			} else {
				ptps.Sockets = append(ptps.Sockets, socket)
				ptps.SendChain(socket)
			}
		}
	}()
}

func handleWebsockets(w http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		fmt.Println("Upgrade failed: ", err)
		return
	}

	GetInstance().Sockets = append(GetInstance().Sockets, conn)
	GetInstance().Peers = append(GetInstance().Peers, "ws://localhost:22") // todo getting this from conn.LocalAddress or conn.RemoteAddress

	for {
		_, message, err := conn.ReadMessage()
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
		newChain := chain.Chain{}

		err = json.Unmarshal(message, &newChain)

		if err != nil {
			fmt.Println("Failed to unmarshal incoming blockchain from json.")
		}

		curChain := chain.GetInstance()

		curChain.ReplaceChain(&newChain)

		chainJson, err := json.Marshal(curChain)

		if err != nil {
			fmt.Println("Failed to marshal blockchain to json.")
		}

		err = conn.WriteMessage(websocket.TextMessage, chainJson)
		if err != nil {
			fmt.Println("write failed:", err)
			break
		}
	}
}

func (ptps *PeerToPeerServer) SendChain(socket *websocket.Conn) {
	fmt.Println("\t\tSendChain...")

	chainJson, err := json.Marshal(chain.GetInstance())

	if err != nil {
		fmt.Println("Failed to marshal blockchain to json.")
	}
	err = socket.WriteMessage(websocket.TextMessage, chainJson)
	if err != nil {
		fmt.Println("Failed to send blockchain using websocket.")
	}
}

func (ptps *PeerToPeerServer) SyncChains() {
	fmt.Println("SyncChains...")
	fmt.Println(ptps.Sockets)
	for _, socket := range ptps.Sockets {
		if socket != nil {
			ptps.SendChain(socket)
		}
	}
}

func (ptps PeerToPeerServer) closeConnections() {
	for _, conn := range ptps.Sockets {
		conn.Close()
	}
}

func (ptps PeerToPeerServer) CreateWebSocketServer(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer ptps.closeConnections()
		defer wg.Done()
		// Websocket Server
		server := new(http.Server)
		server.ReadTimeout = 5 * time.Second
		server.WriteTimeout = 5 * time.Second
		http.HandleFunc("/", handleWebsockets)

		portString := fmt.Sprintf(":%d", ptps.Port)
		http.ListenAndServe(portString, nil)
	}()
	fmt.Println("WS Server started")
}

func (ptps PeerToPeerServer) Listen(wg *sync.WaitGroup) {
	ptps.CreateWebSocketServer(wg)
}
