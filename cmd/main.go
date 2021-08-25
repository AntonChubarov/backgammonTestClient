package main

import (
	"backgammonclient/config"
	"backgammonclient/domain"
	"backgammonclient/infrastructure"
	"fmt"
	"time"
)

var client = infrastructure.NewWebClient()

func main() {
	//ServerRequetScript()

	go TryToKillServer()

	for {
		time.Sleep(1 * time.Second)
	}
}

func ServerRequetScript() {
	RegisterUser("admin3", "admin3password")

	token1 := AuthorizeUser("admin1", "admin1password")
	token2 := AuthorizeUser("admin2", "admin2password")
	token3 := AuthorizeUser("admin3", "admin3password")
	AuthorizeUser("admin3", "admin3passwor")
	AuthorizeUser("admin", "admin3passwor")

	time.Sleep(5 * time.Second)

	token1 = AuthorizeUser("admin1", "admin1password")

	RegisterUser("ad4", "admin4password")
	RegisterUser("admin5", "ad5p")

	ConnectWebSocket(token1)

	time.Sleep(2 * time.Second)

	ConnectWebSocket(token2)

	time.Sleep(2 * time.Second)

	ConnectWebSocket(token3)

	time.Sleep(2 * time.Second)

	ConnectWebSocket("dkjfklsdADDF1456")

	response14 := domain.RoomsInfoDTO{}
	output14 := client.GetRequest(config.UrlBase+"/rooms?token="+token1, &response14)
	fmt.Println(output14)
}

func TryToKillServer() {
	storage := infrastructure.NewAuthStorage()

	registrator := infrastructure.NewUserRegistrator(storage)
	authorizer := infrastructure.NewUserAuthorizer(storage)
	webSocketConnector := infrastructure.NewWebSocketConnector(storage)

	go func() {
		for {//i := 0; i < 100; i++ {
			registrator.RegisterUser()
		}
	}()

	time.Sleep(1 * time.Second)

	go func() {
		for {//i := 0; i < 100; i++ {
			authorizer.AuthorizeUser()
		}
	}()

	//time.Sleep(5 * time.Second)

	go func() {
		for {//i := 0; i < 100; i++ {
			webSocketConnector.ConnectWebSocket()
		}
	}()
}

func RegisterUser(username, password string) {
	response := domain.UserRegistrationResponseDTO{}
	output := client.PostRequest(config.UrlBase+"/register", domain.UserAuthRequestDTO{Username: username, Password: password}, &response)
	fmt.Println(output)
}

func AuthorizeUser(username, password string) (token string) {
	response := domain.UserAuthorizationResponseDTO{}
	output := client.PostRequest(config.UrlBase+"/authorize", domain.UserAuthRequestDTO{Username: username, Password: password}, &response)
	fmt.Println(output)
	token = response.Token
	return
}

func ConnectWebSocket(token string) {
	output := client.PingWebSocket(config.UrlWSBase + "/ws?token=" + token)
	fmt.Println(output)
}
