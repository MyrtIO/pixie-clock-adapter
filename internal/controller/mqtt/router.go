package mqtt

import (
	"log"
	"pixie_adapter/pkg/timing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Router struct {
	cancel  chan struct{}
	updates []*timing.Interval
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

func (r *Router) Destroy() {
	r.cancel <- struct{}{}
}

func (r *Router) OnTopicUpdate(topic string, handler mqtt.MessageHandler) {
	r.routes[topic] = handler
}

func (r *Router) Report(handler func(mqtt.Client), interval time.Duration) {
	r.updates = append(r.updates, timing.NewInterval(interval, func() {
		handler(r.client)
	}))
}

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
