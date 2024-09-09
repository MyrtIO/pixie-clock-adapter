package mqtt

import (
	"context"
	"crypto/tls"
	"log"
	"pixie_adapter/internal/config"
	"pixie_adapter/internal/interfaces"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Controller struct {
	serverURL string
	clientID  string
	username  string
	password  string

	handler Handler
	router  *Router
}

func New(c *config.Config, repos interfaces.Repositories) *Controller {
	handler := newHandler(repos)
	m := &Controller{
		serverURL: c.MQTTServerURL(),
		clientID:  c.ClientID,
		username:  c.Username,
		password:  c.Password,
		handler:   handler,
	}
	return m
}

func (m *Controller) handleConnect(c mqtt.Client) {
	log.Printf("Connected to %s\n", m.serverURL)
	m.router = m.handler.Router(c)
	m.router.Start()
}

func (m *Controller) handleDisconnect(_ mqtt.Client, err error) {
	m.router.Destroy()
	m.router = nil

	log.Printf("Disconnected from %s\n", m.serverURL)
	if err != nil {
		log.Printf("Reason: %s\n", err)
	}
}

func (m *Controller) Start(ctx context.Context) error {
	connOpts := mqtt.NewClientOptions().
		AddBroker(m.serverURL).
		SetClientID(m.clientID).
		SetCleanSession(true)
	if m.username != "" {
		connOpts.SetUsername(m.username)
		if m.password != "" {
			connOpts.SetPassword(m.password)
		}
	}
	connOpts.SetTLSConfig(&tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         tls.NoClientCert,
	})
	connOpts.OnConnect = m.handleConnect
	connOpts.OnConnectionLost = m.handleDisconnect

	client := mqtt.NewClient(connOpts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	<-ctx.Done()

	return nil
}
