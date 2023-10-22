package rabbit

import (
	"fmt"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	cfg     config.RabbitConfig
}

func NewConsumer(cfg config.RabbitConfig) *Consumer {
	return &Consumer{
		cfg: cfg,
	}
}

func (c *Consumer) Connect() error {
	conStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.cfg.User, c.cfg.Password, c.cfg.Host, c.cfg.Port)
	conn, err := amqp.Dial(conStr)
	if err != nil {
		return fmt.Errorf("unable to open connect to RabbitMQ server. Error: %w", err)
	}
	c.conn = conn

	channel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create a channel. Error:%w", err)
	}
	c.channel = channel

	queue, err := channel.QueueDeclare(
		c.cfg.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create a channel. Error:%w", err)
	}

	c.queue = queue
	return nil
}

func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	messages, err := c.channel.Consume(
		c.queue.Name, // queue name
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *Consumer) Close() error {
	if c.conn == nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("failed to close rabbit consumer's connection. Error: %w", err)
		}
	}
	return nil
}
