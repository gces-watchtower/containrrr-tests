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

	// url := "telegram://1630241163:AAFQxAFGFyERgAXOYKalQbxd4zKQY8J0Y_c@telegram?channels=-100145003121"
	
	err := shoutrrr.Send(url, "test mqtt")

	fmt.Println(err);
}

