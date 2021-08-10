package main

import (
	"backgammonclient/domain"
	"backgammonclient/infrastructure"
	"fmt"
)

func main() {
	urlBase := "http://localhost:8080"

	client := infrastructure.NewWebClient()

	response1 := domain.UserRegistrationResponseDTO{}
	output1 := client.PostRequest(urlBase + "/register", domain.UserAuthRequestDTO{Login: "admin3", Password: "admin3password"}, &response1)
	fmt.Println(output1)

	response2 := domain.UserAuthorizationResponseDTO{}
	output2 := client.PostRequest(urlBase + "/login", domain.UserAuthRequestDTO{Login: "admin3", Password: "admin3password"}, &response2)
	fmt.Println(output2)

	response3 := domain.UserAuthorizationResponseDTO{}
	output3 := client.PostRequest(urlBase + "/login", domain.UserAuthRequestDTO{Login: "admin3", Password: "admin3passwor"}, &response3)
	fmt.Println(output3)

	response4 := domain.UserAuthorizationResponseDTO{}
	output4 := client.PostRequest(urlBase + "/login", domain.UserAuthRequestDTO{Login: "admin", Password: "admin3passwor"}, &response4)
	fmt.Println(output4)
}
