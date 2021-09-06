package kafka_client

import (
	"bytes"
	"context"
	"github.com/segmentio/kafka-go"
)

type Client interface {
	Connect(ctx context.Context, dsn string, topic string, partition int) error
	SendMessage(message string) error
}

type client struct {
	conn *kafka.Conn
}

func NewKafkaClient() Client {
	return &client{}
}

func (c *client) Connect(ctx context.Context, dsn string, topic string, partition int) error {
	conn, err := kafka.DialLeader(ctx, "tcp", dsn, topic, partition)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) SendMessage(message string) error {
	_, err := c.conn.Write(bytes.NewBufferString(message).Bytes())
	return err
}
