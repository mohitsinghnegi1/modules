package modules

import (
	"io/ioutil"
	"log"
	"net/http"
)

// SendServerHeartBeat Function that will url as an argument and send the request to the server
func SendServerHeartBeat(serverAddress string, activeConnections int, gameMode string) {

	resJSON, err := http.Get("http://localhost:9000/loadbalancer/updateServerInformation?address=http://localhost:9004&activeConnections=5&gameMode=dice")
	if err != nil {
		log.Fatalln("Failed To Fetch Server Address Due To Errror : ", err)
	}
	serverMessage, err := ioutil.ReadAll(resJSON.Body)
	if err != nil {
		log.Fatalln("Failed To Read Response Body Due to Error : ", err)
	}
	println("ssd ", string(serverMessage))

}

// GetGameRoomServerAddress Function that takes two parameters bucket and gameMode
func GetGameRoomServerAddress(bucket string, gameMode string, c chan string) {
	println("Inside GetGameRoomServer Address Function :")
	resJSON, err := http.Get("http://localhost:9000/loadbalancer/registerClient?bucket=hard&gameMode=cardParty")
	if err != nil {
		log.Fatalln("Failed To Fetch Server Address Due To Errror : ", err)
	}
	serverAddress, err := ioutil.ReadAll(resJSON.Body)
	if err != nil {
		log.Fatalln("Failed To Read Response Body Due to Error : ", err)
	}
	c <- string(serverAddress)
}
