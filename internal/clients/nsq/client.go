package nsq

import (
	"encoding/json"
	"example.com/m/v2/internal/models"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

// Client клиент для взаимодействия c nsq.
type Client struct {
	consumer *nsq.Consumer
	topic    string
	target   string
}

// NewClient конструктор клиента nsq.
// topic - топик, откуда будет производится чтение.
// target - эндпоинт для подключения к nsq.
func NewClient(topic, target string) (*Client, error) {
	if topic == "" {
		return nil, fmt.Errorf("empty topic")
	}
	if target == "" {
		return nil, fmt.Errorf("empty target")
	}

	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, "", cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot create consumer for %s: %w", topic, err)
	}

	err = consumer.ConnectToNSQD(target)
	if err != nil {
		return nil, fmt.Errorf("cannot connet to nsq by %s: %w", target, err)
	}

	client := &Client{
		consumer: consumer,
		topic:    topic,
		target:   target,
	}

	return client, nil
}

// Stop ...
func (c *Client) Stop() error {
	c.consumer.Stop()
	err := c.consumer.DisconnectFromNSQD(c.target)
	if err != nil {
		return fmt.Errorf("cannot disconnect from nsq by %s: %w", c.target, err)
	}
	return nil
}

// StartConsume ...
func (c *Client) StartConsume(fn func(event *models.Event)) {
	c.consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var event models.Event
		err := json.Unmarshal(message.Body, &event)
		if err != nil {
			log.Printf("StartConsume: cannot unmarshal msg: %s\n", err)
			return err
		}
		fn(&event)
		return nil
	}))
}
