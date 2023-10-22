package rabbit

import (
	"context"
	"fmt"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	cfg     config.RabbitConfig
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewProducer(cfg config.RabbitConfig) *Producer {
	return &Producer{
		cfg: cfg,
	}
}

func (p *Producer) Connect() error {
	conStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", p.cfg.User, p.cfg.Password, p.cfg.Host, p.cfg.Port)
	conn, err := amqp.Dial(conStr)
	if err != nil {
		return fmt.Errorf("unable to open connect to RabbitMQ server. Error: %w", err)
	}
	p.conn = conn

	channel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create a channel. Error:%w", err)
	}
	p.channel = channel
	return nil
}

func (p *Producer) Publish(ctx context.Context, data []byte) error {
	queue, err := p.channel.QueueDeclare(
		p.cfg.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create a channel. Error:%w", err)
	}
	return p.channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
}

func (p *Producer) Close() error {
	if p.conn != nil {
		if err := p.conn.Close(); err != nil {
			return fmt.Errorf("failed to close rabbit producer's connection. Error: %w", err)
		}
	}
	return nil
}
