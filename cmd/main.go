package main

import (
	"backgammonclient/domain"
	"backgammonclient/infrastructure"
	"fmt"
	"time"
)

func main() {
	urlBase := "http://localhost:8080"

	client := infrastructure.NewWebClient()

	response1 := domain.UserRegistrationResponseDTO{}
	output1 := client.PostRequest(urlBase + "/register", domain.UserAuthRequestDTO{Username: "admin3", Password: "admin3password"}, &response1)
	fmt.Println(output1)

	response2 := domain.UserAuthorizationResponseDTO{}
	output2 := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: "admin1", Password: "admin1password"}, &response2)
	fmt.Println(output2)

	response3 := domain.UserAuthorizationResponseDTO{}
	output3 := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: "admin2", Password: "admin2password"}, &response3)
	fmt.Println(output3)

	response4 := domain.UserAuthorizationResponseDTO{}
	output4 := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: "admin3", Password: "admin3password"}, &response4)
	fmt.Println(output4)

	response5 := domain.UserAuthorizationResponseDTO{}
	output5 := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: "admin3", Password: "admin3passwor"}, &response5)
	fmt.Println(output5)

	response6 := domain.UserAuthorizationResponseDTO{}
	output6 := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: "admin", Password: "admin3passwor"}, &response6)
	fmt.Println(output6)

	time.Sleep(30 * time.Second)

	response7 := domain.UserAuthorizationResponseDTO{}
	output7 := client.PostRequest(urlBase + "/authorize", domain.UserAuthRequestDTO{Username: "admin1", Password: "admin1password"}, &response7)
	fmt.Println(output7)

	response8 := domain.UserRegistrationResponseDTO{}
	output8 := client.PostRequest(urlBase + "/register", domain.UserAuthRequestDTO{Username: "ad4", Password: "admin4password"}, &response8)
	fmt.Println(output8)

	response9 := domain.UserRegistrationResponseDTO{}
	output9 := client.PostRequest(urlBase + "/register", domain.UserAuthRequestDTO{Username: "admin5", Password: "ad5p"}, &response9)
	fmt.Println(output9)
}
