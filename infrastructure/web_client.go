package infrastructure

import (
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WebClient struct {
	client *resty.Client
}

func NewWebClient() WebClient {
	return WebClient{
		resty.New(),
	}
}

func (w *WebClient) GetRequest(url string, result interface{}) interface{} {
	resp, err:= w.client.R().SetResult(&result).Get(url)
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

	return string(resp.Body())
}

func (w *WebClient) PingWebSocket(url string) interface{} {
	c, resp, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		return err
	}

	go ListenWebSocket(c)

	return resp.StatusCode
}

func ListenWebSocket(conn *websocket.Conn) {
	defer conn.Close()
	for {
		log.Println("reader running")
		time.Sleep(200 * time.Millisecond)
		mtype, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(mtype, msg)
	}
}