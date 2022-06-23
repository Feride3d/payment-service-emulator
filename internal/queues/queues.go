package queues

// TODO сбор статистики в очередь
/*
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Feride3d/payment-service-emulator/internal/storage/models"
	"github.com/streadway/amqp"
)

type RMQConnection interface {
	Channel() (*amqp.Channel, error)
}

type Publisher struct {
	name string
	conn RMQConnection
}

func NewPublisher(name string, conn RMQConnection) *Publisher {
	return &Publisher{
		name: name,
		conn: conn,
	}
}

// Publish publishes messages (events) to the exchange.
func (p *Publisher) Publish(ctx context.Context, message models.Events) error {
	if ctx.Err() == context.Canceled {
		return errors.New("messages publication canceled")
	}

	ch, err := p.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(p.name, false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("failed to create queue: %w", err)
	}

	if ch != nil {
		bytes, err := json.Marshal(message)
		if err != nil {
			return fmt.Errorf("failed to marshall message: %w", err)
		}

		err = ch.Publish(
			"",
			p.name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        bytes,
			})
		if err != nil {
			return fmt.Errorf("failed to publish message: %w", err)
		}

		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}
*/
