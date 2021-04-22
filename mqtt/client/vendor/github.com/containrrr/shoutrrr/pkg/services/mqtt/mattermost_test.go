package mqtt

import (
	"github.com/containrrr/shoutrrr/pkg/types"
	"github.com/containrrr/shoutrrr/pkg/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/url"
	"os"
	"testing"
)

var (
	service          *Service
	envMqttURL *url.URL
)

func TestMqtt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shoutrrr MQTT Suite")
}

var _ = Describe("the mattermost service", func() {
	BeforeSuite(func() {
		service = &Service{}
		envMattermostURL, _ = url.Parse(os.Getenv("SHOUTRRR_MQTT_URL"))
	})
	When("running integration tests", func() {
		It("should work without errors", func() {
			if envMqttURL.String() == "" {
				return
			}
			serviceURL, _ := url.Parse(envMqttURL.String())
			service.Initialize(serviceURL, util.TestLogger())
			err := service.Send(
				"this is an integration test",
				nil,
			)
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("the mqtt config", func() {
		When("generating a config object", func() {
			mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
			config := &Config{}
			err := config.SetURL(mattermostURL)
			It("should not have caused an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should set host", func() {
				Expect(config.Host).To(Equal("mqtt.my-domain.com"))
			})
			It("should set token", func() {
				Expect(config.Token).To(Equal("thisshouldbeanapitoken"))
			})
			It("should not set channel or username", func() {
				Expect(config.Host).To(BeEmpty())
				Expect(config.TLS).To(BeEmpty())
			})
		})
		When("generating a new config with url, that has no token", func() {
			mattermostURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
			config := &Config{}
			err := config.SetURL(mattermostURL)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	Describe("Sending messages", func() {
		When("sending a message completely without parameters", func() {
			mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
			config := &Config{}
			config.SetURL(mqttURL)
			
			It("should generate the correct url to call", func() {
				generatedURL := buildURL(config)
				Expect(generatedURL).To(Equal("mqtt://localhost:1883?topic=topic/test"))
			})
			It("should generate the correct JSON body", func() {
				err := mqtt.Send(mqttURL, "this is a message")
				Expect(err).NotTo(HaveOccurred())
				})
		})
