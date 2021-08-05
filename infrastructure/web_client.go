package infrastructure

import (
	"github.com/go-resty/resty/v2"
	"log"
	"github.com/gorilla/websocket"
)

type WebClient struct {
	client *resty.Client
}

func NewWebClient() *WebClient {
	return &WebClient{
		resty.New(),
	}
}

func (w *WebClient) GetRequest(url string, result *interface{}) interface{} {
	resp, err:= w.client.R().SetResult(result).Get(url)
	if err !=nil {
		log.Println(err)
	}
	if resp.IsSuccess() {
		return resp
	}
	return resp.RawResponse.Status
}

func (w *WebClient) PostRequest(url string, request interface{},  result interface{}) interface{} {
	resp, err:= w.client.R().SetResult(&result).SetBody(request).Post(url)
	if err != nil {
		log.Println(err)
		return err
	}

	if resp.IsError() {
		return resp.StatusCode()
	}

	return resp.String()
}

func (w *WebClient) PingWebSocket(url string) interface{} {
	c, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer c.Close()
	return resp.StatusCode
}