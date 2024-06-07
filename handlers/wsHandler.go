package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

type ReturnChatNames struct {
	Success   bool     `json:"success"`
	ChatNames []string `json:"chatNames"`
}

func GetChatNames(w http.ResponseWriter, r *http.Request, chatRooms map[string][]string) {
	if r.Method == "GET" {

		var rooms []string

		for k, _ := range chatRooms {
			rooms = append(rooms, k)

		}

		res := ReturnChatNames{
			Success:   true,
			ChatNames: rooms,
		}

		jsonRes, err := json.Marshal(res)

		if err != nil {
			ErrLog(err, w)
		}

		fmt.Println("json sent successfully")
		w.Header().Set("content-type", "application/json")
		w.Write(jsonRes)

	} else {
		ErrLog(errors.New("WongMethod"), w)
	}

}

///GET CAHT LOAGS NEEDS TO HEPPEN

func SocketHandler(w http.ResponseWriter, r *http.Request, chatRooms map[string][]string) {
	//check if the protocol can become a websocket one
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic("could not create a ws connection")
	}
	clients = append(clients, *wsConn)
	for {
		msgType, msg, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s send: %s\n", wsConn.RemoteAddr(), string(msg))
		//loop if found and sent to browser


		messageVals := strings.Split(string(msg), "::")
        if !strings.Contains(string(msg), "::change1234321") {

		    chatRooms[messageVals[0]] = append(chatRooms[messageVals[0]], messageVals[1])
        }


		for _, client := range clients {

			if err = client.WriteMessage(msgType, []byte(fmt.Sprint(chatRooms[messageVals[0]]))); err != nil {
				fmt.Println(err)
				return
			}

		}
	}

}
