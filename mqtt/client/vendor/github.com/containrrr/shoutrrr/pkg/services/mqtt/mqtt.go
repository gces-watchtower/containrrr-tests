package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/util"
	"github.com/containrrr/shoutrrr/pkg/format"
	
	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

const (
	maxLength = 4096
)

// Service sends notifications to mqtt topic
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}



// Send notification to mqtt
func (service *Service) Send(message string, params *types.Params) error {

	items, omitted := util.MessageItemsFromLines(message, types.MessageLimit{ maxLength, maxLength, 2})
	

	if omitted > 0 {
		message = items[0].Text
        service.Logf("omitted %v character(s) from the message", omitted)
    }

	config := *service.config
	if err := service.pkr.UpdateConfigFromParams(&config, params); err != nil {
		return err
	}

	return publishMessageToTopic(message, &config)
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger *log.Logger) error {
	service.Logger.SetLogger(logger)
	service.config = &Config{
		DisableTLS:    false,
	}

	if err := service.pkr.SetDefaultProps(service.config); err != nil {
		return err
	}
	service.pkr = format.NewPropKeyResolver(service.config)
	if err := service.config.setURL(&service.pkr, configURL); err != nil {
		return err
	}

	
	return nil
}

// GetConfig returns the Config for the service
func (service *Service) GetConfig() *Config {
	return service.config
}


func (service *Service) LostConnection(err string) {
	service.Logf("Connect lost: %v", err)
}


func (service *Service)lost() {
	fmt.Println("Em meu lugar o que Eugenio faria?")
}

// Handle Connection Lost
var connectLostHandler mqtt.ConnectionLostHandler =  func(client mqtt.Client, err error) {
	//test("ok")
	fmt.Printf("Connect lost: %v", err)
	if err != nil{

	}
}


// Publish to topic
func publish(client mqtt.Client, topic string, data []byte) {
	token := client.Publish(topic, 0, false, data)
	token.Wait()
}



// Publish payload
func publishMessageToTopic(message string, config *Config) error {
	postURL := fmt.Sprintf("mqtt://%s:%d", config.Host, config.Port)
	payload := createSendMessagePayload(message, config.Topic, config)
	
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Config
    opts := mqtt.NewClientOptions()

    opts.AddBroker(postURL)

	opts.OnConnectionLost = connectLostHandler
    
	// Start client
	client := mqtt.NewClient(opts)
    
	token := client.Connect();

	if token.Error() != nil {
		return token.Error()
	}

	token.Wait()

    publish(client, config.Topic, jsonData)

    client.Disconnect(1)

	return nil
}




