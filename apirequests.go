package modules

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/mohitsinghnegi1/Frontend-loadbalancer/config"
)

// Configuration Variables
var configuration config.Configuration = config.GetConfig("dev")
var loadbalancer config.LB = configuration.LoadBalancer
var basePath string = loadbalancer.BasePath
var endPoint config.EP = loadbalancer.EndPoints
var updateServerInformation string = endPoint.UpdateServerInformation
var registerClient string = endPoint.RegisterClient

// SendServerHeartBeat Function that will url as an argument and send the request to the server
func SendServerHeartBeat(serverAddress string, activeConnections int, gameMode string) bool {

	resJSON, err := http.Get(basePath + updateServerInformation +
		"?address=" + serverAddress + "&activeConnections=" +
		strconv.Itoa(activeConnections) + "&gameMode=" +
		gameMode)

	if err != nil {
		log.Fatalln("Failed To Fetch Server Response Due To Errror : ", err)
		return false
	}

	defer resJSON.Body.Close()
	serverMessage, err := ioutil.ReadAll(resJSON.Body)
	if err != nil {
		log.Fatalln("Failed To Read Response Body Due to Error : ", err)
		return false
	}

	println("Response : ", string(serverMessage))
	return true

}

// FetchGameRoomServerAddress Function that takes two parameters bucket and gameMode
func FetchGameRoomServerAddress(bucket string, gameMode string, channel chan string) {

	println("Inside GetGameRoomServer Address Function :")
	resJSON, err := http.Get(basePath + registerClient + "?bucket=" +
		bucket + "&gameMode=" + gameMode)

	if err != nil {
		log.Fatalln("Failed To Fetch Server Address Due To Errror : ", err)
	}

	defer resJSON.Body.Close()
	serverAddress, err := ioutil.ReadAll(resJSON.Body)

	if err != nil {
		log.Fatalln("Failed To Read Response Body Due to Error : ", err)
	}

	channel <- string(serverAddress)

}
