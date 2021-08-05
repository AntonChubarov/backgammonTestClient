package main

import (
	"backgammonclient/domain"
	"backgammonclient/infrastructure"
	"fmt"
)

func main() {
	urlBase := "http://localhost:8080"

	client := infrastructure.NewWebClient()

	response := domain.UserRegistrationResponse{}
	output := client.PostRequest(urlBase + "/register", domain.UserRegistrationRequest{Login: "admin", Password: "admin"}, &response)
	fmt.Println(response)
	fmt.Println(output)
}
