package main

import (
	"fmt"

	"github.com/containrrr/shoutrrr"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	broker := "localhost"
	port := 1883
	topic := "topic/test"

	// Config
	mqtt.NewClientOptions()

	url := fmt.Sprintf("mqtt://%s:%d?topic=%s", broker, port, topic)

	err := shoutrrr.Send(url, "test mqtt")

	fmt.Println(err)
}
