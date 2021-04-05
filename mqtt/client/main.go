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

	err := shoutrrr.Send(url, "mqtt Lorem ipsum dolor sit amet, consectetuer adipiscing elit.\nAenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequa")
	
	fmt.Println(err)
}
