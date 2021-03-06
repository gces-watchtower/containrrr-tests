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
	disableTls := "true"
	clientId := "1"
	username := "eugenio"
	password := "12345"

	// Config
	mqtt.NewClientOptions()

	url := fmt.Sprintf("mqtt://%s:%d?topic=%s&disableTls=%s&clientId=%s&username=%s&password=%s", broker, port, topic, disableTls, clientId, username, password)

	err := shoutrrr.Send(url, "Implementing TLS")

	fmt.Println(err)
}
