package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Client interface {
	Connect(ctx context.Context, dsn string, topic string, partition int) error
	SendMessage(message []byte) error
}

func New() Client {
	return &client{}
}

type client struct {
	conn *kafka.Conn
}

func (c *client) Connect(ctx context.Context, dsn string, topic string, partition int) error {
	conn, err := kafka.DialLeader(ctx, "tcp", dsn, topic, partition)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) SendMessage(message []byte) error {
	_, err := c.conn.Write(message)
	return err
}
