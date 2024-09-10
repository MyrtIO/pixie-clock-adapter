package mqtt

import (
	"log"
	"pixie_adapter/pkg/job"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Router routes MQTT messages
type Router struct {
	cancel  chan struct{}
	updates []*job.Interval
	routes  map[string]mqtt.MessageHandler
	qos     byte
	client  mqtt.Client
}

func newRouter(c mqtt.Client) *Router {
	r := &Router{
		routes: make(map[string]mqtt.MessageHandler),
		qos:    0,
		client: c,
	}
	return r
}

// Stop stops the router
func (r *Router) Stop() {
	r.cancel <- struct{}{}
}

// OnTopicUpdate registers a handler for a topic
func (r *Router) OnTopicUpdate(topic string, handler mqtt.MessageHandler) {
	r.routes[topic] = handler
}

// Report registers a handler for a topic with a given interval
func (r *Router) Report(handler func(mqtt.Client), interval time.Duration) {
	r.updates = append(r.updates, job.NewInterval(func() {
		handler(r.client)
	}, interval))
}

// Start starts the router
func (r *Router) Start() {
	for topic, handler := range r.routes {
		token := r.client.Subscribe(topic, 0, handler)
		if token.Wait() && token.Error() != nil {
			log.Printf("Error subscribing to %s: %s\n", topic, token.Error())
		}
	}

	for _, update := range r.updates {
		update.Start(r.cancel)
	}
}
