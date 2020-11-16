package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ServerHeartBeat Function that will url as an argument and send the request to the server
func ServerHeartBeat() string {
	http.Get("registerClient")
	return "True"
}

// GetGameRoomServerAddress Function that takes two parameters bucket and gameMode
func GetGameRoomServerAddress(bucket string, gameMode string) string {
	resJSON, err := http.Get("registerClient")
	if err != nil {
		log.Fatalln("Failed To Fetch Server Address Due To Errror : ", err)
	}
	serverAddress, err := ioutil.ReadAll(resJSON.Body)
	if err != nil {
		log.Fatalln("Failed To Read Response Body Due to Error : ", err)
	}
	fmt.Println("serverAddress : ", serverAddress)
	return string(serverAddress)
}
