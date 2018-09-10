package pusher

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

const PusherEventPrefix = "pusher:"
const protocolVersion = "4.2.2"
const clientName = "go"
const protocol = "7"

type Message struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type ConnectionEstablished struct {
	SocketId        string `json:"socket_id"`
	ActivityTimeout int    `json:"activity_timeout"`
}

func Connect(appKey string, cluster string) (*websocket.Conn, error) {
	u := getUrlForAppKeyAndCluster(appKey, cluster)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		defer close(done)

		for {
			message := Message{}
			err = websocket.ReadJSON(conn, &message)
			if err != nil {
				log.Println(err)
			}

			if strings.HasPrefix(message.Event, PusherEventPrefix) {
				handleSystemEvent(message)
			}
		}
	}()

	return conn, nil
}

func handleSystemEvent(message Message) {
	// {"event":"pusher:connection_established","data":"{\"socket_id\":\"125035.700573\",\"activity_timeout\":120}"}
	fmt.Printf("Event type: %s\r\n", message.Event)
	fmt.Printf("Event data: %s\r\n", message.Data)
}

func getUrlForAppKeyAndCluster(appKey string, cluster string) url.URL {
	return url.URL{
		Scheme:   "ws",
		Host:     fmt.Sprintf("ws-%s.pusher.com", cluster),
		Path:     fmt.Sprintf("/app/%s", appKey),
		RawQuery: fmt.Sprintf("protocol=%s&client=%s&version=%s", protocol, clientName, protocolVersion),
	}
}
