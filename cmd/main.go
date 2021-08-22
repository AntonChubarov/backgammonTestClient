package main

import (
	"backgammonclient/domain"
	"backgammonclient/infrastructure"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

var urlBase = "http://localhost:8080"
var urlWSBase = "ws://localhost:8080"

var client = infrastructure.NewWebClient()

func main() {
	//ServerRequetScript()

	TryToKillServer()

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
	output14 := client.GetRequest(urlBase + "/rooms?token=" + token1, &response14)
	fmt.Println(output14)
}

func TryToKillServer() {
	storage := infrastructure.NewAuthStorage()

	for {
		go func() {
			username := gofakeit.Username()
			password := gofakeit.Password(true, true, true, false, false, 10)

			fmt.Println("to register:", username, password)

			RegisterUser(username, password)
			storage.AddUser(username, password, "")
		}()

		go func() {
			user := storage.GetRandomUser()
			token := AuthorizeUser(user.Username, user.Password)
			storage.UpdateToken(user.Username, token)

			fmt.Println("authorized:", user.Username, user.Password, token)
		}()

		go func() {
			user := storage.GetRandomUser()

			fmt.Println("to websocket:", user)

			ConnectWebSocket(user.Token)
		}()
	}
}

func RegisterUser(username, password string) {
	response := domain.UserRegistrationResponseDTO{}
	output := client.PostRequest(urlBase + "/register", domain.UserAuthRequestDTO{Username: username, Password: password}, &response)
	fmt.Println(output)
}

func AuthorizeUser(username, password string) (token string) {
	response := domain.UserAuthorizationResponseDTO{}
	output := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: username, Password: password}, &response)
	fmt.Println(output)
	token = response.Token
	return
}

func ConnectWebSocket(token string) {
	output := client.PingWebSocket(urlWSBase + "/ws?token=" + token)
	fmt.Println(output)
}