package infrastructure

import (
	"backgammonclient/config"
	"backgammonclient/domain"
	"backgammonclient/utils"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

type UserRegistrator struct {
	client  WebClient
	storage *AuthStorage
}

func NewUserRegistrator(storage *AuthStorage) *UserRegistrator {
	return &UserRegistrator{
		client:  NewWebClient(),
		storage: storage,
	}
}

func (ur *UserRegistrator) RegisterUser() {
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, false, false, 10)
	response := domain.UserRegistrationResponseDTO{}
	output := ur.client.PostRequest(config.UrlBase+"/register", domain.UserAuthRequestDTO{Username: username, Password: password}, &response)
	ur.storage.AddUser(username, password, "")
	fmt.Println("In RegisterUser", username, password, output)
}

type UserAuthorizer struct {
	client  WebClient
	storage *AuthStorage
}

func NewUserAuthorizer(storage *AuthStorage) *UserAuthorizer {
	return &UserAuthorizer{
		client:  NewWebClient(),
		storage: storage,
	}
}

func (ua *UserAuthorizer) AuthorizeUser() {
	user, index := ua.storage.GetRandomUser()

	response := domain.UserAuthorizationResponseDTO{}
	output := ua.client.PostRequest(config.UrlBase+"/authorize", domain.UserAuthRequestDTO{Username: user.Username, Password: user.Password}, &response)

	ua.storage.UpdateToken(index, response.Token)

	fmt.Println("In AuthorizeUser", user.Username, user.Password, response.Token, output)

}

func (ua *UserAuthorizer) AuthorizeRandomUser() {
	user := UserData{
		Index:    utils.RandomInt(len(ua.storage.storage)),
		Username: gofakeit.Username(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Token:    "",
	}
	index := user.Index
	response := domain.UserAuthorizationResponseDTO{}
	output := ua.client.PostRequest(config.UrlBase+"/authorize", domain.UserAuthRequestDTO{Username: user.Username, Password: user.Password}, &response)

	ua.storage.UpdateToken(index, response.Token)

	fmt.Println("In AuthorizeUser", user.Username, user.Password, response.Token, output)
}

type WebSocketConnector struct {
	client  WebClient
	storage *AuthStorage
}

func NewWebSocketConnector(storage *AuthStorage) *WebSocketConnector {
	return &WebSocketConnector{
		client:  NewWebClient(),
		storage: storage,
	}
}

func (wsc *WebSocketConnector) ConnectWebSocket() {
	user, _ := wsc.storage.GetRandomUser()

	output := wsc.client.PingWebSocket(config.UrlWSBase + "/ws?token=" + user.Token)

	fmt.Println("In ConnectWebSocket", user.Username, user.Password, user.Token, output)
}
