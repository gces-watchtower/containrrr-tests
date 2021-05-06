package mqtt

import (
	"fmt"
	"github.com/containrrr/shoutrrr/pkg/format"
	"log"
	"net/url"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	logger  = log.New(GinkgoWriter, "Test", log.LstdFlags)
	service *Service
	config  *Config
)

func TestMqtt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MQTT Suite")
}

var _ = Describe("the MQTT service", func() {
	BeforeEach(func() {
		service = &Service{}
		service.SetLogger(logger)
	})

	When("parsing a custom URL", func() {
		It("should strip generic prefix before parsing", func() {
			broker := "localhost"
			port := 8883
			disableTLS := "false"
			//customURL, err := url.Parse(fmt.Sprintf("mqtt://%s:%d?disableTLS=%s", broker, port, disableTLS))
			//fmt.Println(customURL)
			//fmt.Println(err)
			//Expect(err).NotTo(HaveOccurred())

			config, _ := testServiceURL(fmt.Sprintf("mqtt://%s:%d?disableTLS=%s", broker, port, disableTLS))
			whURL := config.Host
			fmt.Println(whURL)

			//Expect(err).NotTo(HaveOccurred())
			// _, expectedURL := testCustomURL("https://test.tld")
			// Expect(actualURL.String()).To(Equal(expectedURL.String()))
		})

	})

	When("a MQTTS URL is provided", func() {
		It("should disable TLS", func() {
			broker := "localhost"
			port := 8883
			disableTLS := "false"
			config, _ := testServiceURL(fmt.Sprintf("mqtt://%s:%d?disableTLS=%s", broker, port, disableTLS))
			postURL := config.MqttURL()
			//fmt.Println("Aqui:", postURL)

			Expect(postURL).To(Equal(fmt.Sprintf("mqtts://%s:%d", broker, port)))
		})
	})

	When("a MQTT URL is provided", func() {
		It("should enable TLS", func() {
			broker := "localhost"
			port := 1883
			disableTLS := "true"
			config, _ := testServiceURL(fmt.Sprintf("mqtt://%s:%d?disableTLS=%s", broker, port, disableTLS))
			postURL := config.MqttURL()
			//fmt.Println("Aqui:", postURL)

			Expect(postURL).To(Equal(fmt.Sprintf("mqtt://%s:%d", broker, port)))
		})
	})

	When("a HTTP URL is provided", func() {
		It("should disable TLS", func() {
			broker := "localhost"
			port := 1883
			disableTLS := "true"
			config, _ := testServiceURL(fmt.Sprintf("mqtt://%s:%d?disableTLS=%s", broker, port, disableTLS))
			config.MqttURL()
			Expect(config.DisableTLS).To(BeTrue())
		})
	})

	When("a HTTPS URL is provided", func() {
		It("should enable TLS", func() {
			broker := "localhost"
			port := 8883
			disableTLS := "false"
			config, _ := testServiceURL(fmt.Sprintf("mqtt://%s:%d?disableTLS=%s", broker, port, disableTLS))
			config.MqttURL()
			Expect(config.DisableTLS).To(BeFalse())
		})
	})

	Describe("creating a config", func() {
		When("creating a default config", func() {
			It("should not return an error", func() {
				config := &Config{}
				pkr := format.NewPropKeyResolver(config)
				err := pkr.SetDefaultProps(config)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("generating a config object", func() {
			mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test&disableTls=true&clientId=1&username=testUser&password=password")
			config := &Config{}
			err := config.SetURL(mqttURL)
			It("should not have caused an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should set host", func() {
				Expect(config.Host).To(Equal("localhost"))
			})
			It("should set Port", func() {
				Expect(config.Port).To(Equal(uint16(1883)))
			})

			It("should set topic", func() {
				Expect(config.Topic).To(Equal("topic/test"))
			})


			It("should set client", func() {
				Expect(config.ClientId).To(Equal("1"))
			})

			It("should set username", func() {
				Expect(config.Username).To(Equal("testUser"))
			})

			It("should set password", func() {
				Expect(config.Password).To(Equal("password"))
			})

			It("should not set DisableTLS", func() {
				Expect(config.DisableTLS).Should(BeTrue())
			})
		})

		When("generating a config object", func() {
			mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
			config := &Config{}
			err := config.SetURL(mqttURL)
			It("should not have caused an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should not set client", func() {
				Expect(config.ClientId).To(Equal(""))
			})

			It("should set username", func() {
				Expect(config.Username).To(Equal(""))
			})

			It("should set password", func() {
				Expect(config.Password).To(Equal(""))
			})

			It("should set DisableTLS", func() {
				Expect(config.DisableTLS).Should(BeFalse())
			})
		})


	})

	// When("a HTTP URL is provided", func() {
	//  It("should disable TLS", func() {
	//   config, _ := testCustomURL("http://example.com")
	//   Expect(config.DisableTLS).To(BeTrue())
	//  })
	// })
	// When("a HTTPS URL is provided", func() {
	//  It("should enable TLS", func() {
	//   config, _ := testCustomURL("https://example.com")
	//   Expect(config.DisableTLS).To(BeFalse())
	//  })
	// })
	// It("should escape conflicting custom query keys", func() {
	//  expectedURL := "generic://example.com/?__template=passed"
	//  config, srvURL := testCustomURL("https://example.com/?template=passed")
	//  Expect(config.Template).NotTo(Equal("passed")) // captured
	//  whURL := config.WebhookURL().String()
	//  Expect(whURL).To(Equal("https://example.com/?template=passed"))
	//  Expect(srvURL.String()).To(Equal(expectedURL))

	// })
	// It("should handle both escaped and service prop version of keys", func() {
	//  config, _ := testServiceURL("generic://example.com/?__template=passed&template=captured")
	//  Expect(config.Template).To(Equal("captured"))
	//  whURL := config.WebhookURL().String()
	//  Expect(whURL).To(Equal("https://example.com/?template=passed"))
	// })

	// var _ = Describe("the  service", func() {
	//  BeforeSuite(func() {
	//   service = &Service{}
	//   //envMqttURL, _ = url.Parse(os.Getenv("SHOUTRRR_MQTT_URL"))
	//   var envMqttURL = "mqtt://%s:%d?topic=%s&disableTLS=%s&clientID=%s&username=%s&password=%s";
	//   print(envMqttURL);
	//  })
	//  When("running integration tests", func() {
	//   It("should work without errors", func() {
	//    if envMqttURL.String() == "" {
	//     return
	//    }
	//    serviceURL, _ := url.Parse(envMqttURL.String())
	//    service.Initialize(serviceURL, util.TestLogger())
	//    err := service.Send(
	//     "this is an integration test",
	//     nil,
	//    )
	//    Expect(err).NotTo(HaveOccurred())
	//   })
	//  })

	// Describe("the mqtt config", func() {
	//  When("generating a config object", func() {
	//   mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
	//   config := &Config{}
	//   err := config.SetURL(mqttURL)
	//   It("should not have caused an error", func() {
	//    Expect(err).NotTo(HaveOccurred())
	//   })
	//   It("should set host", func() {
	//    Expect(config.Host).To(Equal("mqtt.my-domain.com"))
	//   })
	//   It("should set token", func() {
	//    Expect(config.Token).To(Equal("thisshouldbeanapitoken"))
	//   })
	//   It("should not set channel or username", func() {
	//    Expect(config.Host).To(BeEmpty())
	//    Expect(config.TLS).To(BeEmpty())
	//   })
	//  })
	//  When("generating a new config with url, that has no token", func() {
	//   mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
	//   config := &Config{}
	//   err := config.SetURL(mqttURL)
	//   It("should return an error", func() {
	//    Expect(err).To(HaveOccurred())
	//   })
	//  })
	// Describe("Sending messages", func() {
	//  When("sending a message completely without parameters", func() {
	//   mqttURL, _ := url.Parse("mqtt://localhost:1883?topic=topic/test")
	//   config := &Config{}
	//   config.SetURL(mqttURL)

	//   It("should generate the correct url to call", func() {
	//    generatedURL := buildURL(config)
	//    Expect(generatedURL).To(Equal("mqtt://localhost:1883?topic=topic/test"))
	//   })
	//   It("should generate the correct JSON body", func() {
	//    err := mqtt.Send(mqttURL, "this is a message")
	//    Expect(err).NotTo(HaveOccurred())
	//    })
	//   })
	//  })
	// })
})

func testServiceURL(testURL string) (*Config, *url.URL) {

	serviceURL, err := url.Parse(testURL)
	Expect(err).NotTo(HaveOccurred())
	config, pkr := DefaultConfig()
	err = config.setURL(&pkr, serviceURL)
	Expect(err).NotTo(HaveOccurred())
	//test := MqttURL()

	return config, config.getURL(&pkr)
}

/*
func test(){
	test, a:= DefaultConfig()
	fmt.Println(test)
	fmt.Println(a)
}*/
