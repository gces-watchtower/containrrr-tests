package main

import (
	"fmt"

	"github.com/containrrr/shoutrrr"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	broker := "192.168.0.102"
	port := 8883
	topic := "topic/test"
	disableTls := "false"
	clientId := "1"
	username := "eugenio"
	password := "12345"
	verbose := "true"

	// Config
	mqtt.NewClientOptions()

	url := fmt.Sprintf("mqtt://%s:%d?topic=%s&disableTls=%s&clientId=%s&username=%s&password=%s&verbose=%s", broker, port, topic, disableTls, clientId, username, password, verbose)

	err := shoutrrr.Send(url, "Implementing TLS")

	fmt.Println(err)
}
