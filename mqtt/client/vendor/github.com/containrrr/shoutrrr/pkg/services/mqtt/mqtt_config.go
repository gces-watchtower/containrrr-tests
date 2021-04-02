package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for use within the mqtt
type Config struct {
	Scheme       string       
	Host       	 string    	 `key:"broker" default:"" desc:"MQTT broker server hostname or IP address"`
	Port         int64       `key:"port" default:"1883" desc:"TCP Port"`
	Topic        string    	 `key:"topic" default:"" desc:"Topic where the message is sent"`
	ClientId     string      `key:"clientid" default:"" desc:"client's id from the message is sent"`
	Username     string      `key:"username" default:"" desc:"username for auth"`
	Password     string      `key:"password" default:"" desc:"password for auth"`
	DisableTLS   bool        `key:"disabletls" default:"No"`
	ParseMode    parseMode   `key:"parsemode" default:"None" desc:"How the text message should be parsed"`
}

// Enums returns the fields that should use a corresponding EnumFormatter to Print/Parse their values
func (config *Config) Enums() map[string]types.EnumFormatter {
	return map[string]types.EnumFormatter{
		"ParseMode": parseModes.Enum,
	}
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL {
	resolver := format.NewPropKeyResolver(config)
	return config.getURL(&resolver)
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error {
	resolver := format.NewPropKeyResolver(config)
	return config.setURL(&resolver, url)
}

func (config *Config) getURL(resolver types.ConfigQueryResolver) *url.URL {
	return &url.URL{
		Host:       fmt.Sprintf("%s:%d", config.Host, config.Port),
		Scheme:     Scheme,
		ForceQuery: true,
		RawQuery:   format.BuildQuery(resolver),
	}

}

func setTlsConfig() *tls.Config {
    certpool := x509.NewCertPool()
    ca, err := ioutil.ReadFile("certs/ca.crt")
	
	if err != nil {
        log.Fatalln(err.Error())
    }
    certpool.AppendCertsFromPEM(ca)

    clientKeyPair, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
    if err != nil {
        panic(err)
    }
    return &tls.Config{
        RootCAs: certpool,
        ClientAuth: tls.NoClientCert,
        ClientCAs: nil,
        InsecureSkipVerify: true,
        Certificates: []tls.Certificate{clientKeyPair},
    }
}

func Split(r rune) bool {
    return r == '/' || r == '?' || r == ':'
}

func getTopic(r rune) bool {
    return r == '&' || r == '?' || r == '='
}

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}


func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {

	u := strings.FieldsFunc(url.String(), Split)
	topic := strings.FieldsFunc(url.String(), getTopic)

	port, err := strconv.ParseInt(u[2], 10, 64)

	if err != nil {
		return err
	}

	tls, err := strconv.ParseBool(topic[4])
	
	if err != nil {
		return err
	}

	k, found := Find(topic, "clientId")
	if found && len(topic) >= k + 1 {
		config.ClientId = topic[k+1]
	}

	k, found = Find(topic, "username")
	if found && len(topic) >= k + 1 {
		config.Username = topic[k+1]
	}

	k, found = Find(topic, "password")
	if found && len(topic) >= k + 1 {
		config.Password = topic[k+1]
	}

	if len(u) > 4 {
		config.Scheme = "tcp"
		config.Host = u[1]
		config.Port = port
		config.Topic = topic[2]
		config.DisableTLS = tls		
	}

	return err
}

// Scheme is the identifying part of this service's configuration URL
const (
	Scheme = "mqtt"
)

