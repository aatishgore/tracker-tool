package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(ConnWS *websocket.Conn) {
	gsConnWS = append(gsConnWS, ConnWS)
	for {
		// Read message from browser
		msgType, msg, err := ConnWS.ReadMessage()
		gsMessageType = msgType
		if err != nil {
			for i, connection := range gsConnWS {
				if connection == ConnWS {
					gsConnWS = remove(gsConnWS, i)
					break
				}
			}
			return
		}

		var response map[string]interface{}
		json.Unmarshal(msg, &response)

		// var message map[string]interface{}
		// json.Unmarshal([]byte(response["message"]), &message)
		// message := response["data"].(map[string]interface{})

		var result map[string]interface{}
		for key, value := range response {
			// Each value is an interface{} type, that is type asserted as a string
			fmt.Println(key, value)

			json.Unmarshal([]byte(value.(string)), &result)

		}

		message := result["message"].(map[string]interface{})

		if result["channel"] == "wfh" {

			if message["start"] == true {
				trackingStart = true

			} else {
				trackingStart = false
				sendLogs()
			}
			token = message["token"].(string)
		}
		if err = ConnWS.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func remove(s []*websocket.Conn, i int) []*websocket.Conn {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
func socketInit() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		ws, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		if err != nil {
			fmt.Println(err)
			fmt.Println("Error with socket")
		} else {
			reader(ws)
		}

	})

	mydir, _ := os.Getwd()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(mydir+"/websocket/html/"))))

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ping Pong!!!!"))

	})
	http.ListenAndServe(":8252", nil)

}
func notify(message string) {
	msg := []byte(message)
	fmt.Println("Active connections %v", len(gsConnWS))
	for _, conn := range gsConnWS {
		muTx.Lock()
		err := conn.WriteMessage(1, msg)
		muTx.Unlock()
		if err != nil {
			fmt.Println("Error")
		}
	}
}

func sendNotification(code string, message string) {
	bytes, _ := json.Marshal(socketMessage{
		code,
		message,
	})

	notify(string(bytes))
}
