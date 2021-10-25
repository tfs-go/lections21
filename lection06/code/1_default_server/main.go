package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Message struct {
	Text string `json:"text"`
}

func WSHandler(ws *websocket.Conn) {
	fmt.Println("web socket handler function")
	defer ws.Close()

	var s string
	// works only with string and []byte
	err := websocket.Message.Receive(ws, &s)
	if err != nil {
		fmt.Printf("web socket Message.Receive error: %s\n", err.Error())
		return
	}

	var b = []byte(s)
	// works only with string and []byte
	err = websocket.Message.Send(ws, b)
	if err != nil {
		fmt.Printf("web socket Message.Send error: %s\n", err.Error())
		return
	}

	var m = Message{Text: "welcome"}
	// unmarshals json text
	err = websocket.JSON.Receive(ws, &m)
	if err != nil {
		fmt.Printf("web socket JSON.Receive error: %s\n", err.Error())
		return
	}

	// marshals json text
	err = websocket.JSON.Send(ws, m)
	if err != nil {
		fmt.Printf("web socket JSON.Send error: %s\n", err.Error())
		return
	}

	// *ws.Conn implements io.Reader
	_, err = io.ReadAll(ws)
	if err != nil {
		fmt.Printf("web socket io.ReadAll error: %s\n", err.Error())
		return
	}

	// *ws.Conn implements io.Writer
	err = json.NewEncoder(ws).Encode(m)
	if err != nil {
		fmt.Printf("web socket io.ReadAll error: %s\n", err.Error())
		return
	}
}

func main() {
	http.Handle("/stocks", websocket.Handler(WSHandler))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//nolint: unused
func client() (*websocket.Conn, error) {
	return websocket.Dial(":5000", "chat", "/")
}
