package client

import(
	"github.com/wagslane/go-rabbitmq"
	"telegram-callback-svc/properties"
	"fmt"
)

func CreateConnection(props *properties.RabbitProperties) (*rabbitmq.Conn, error) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d", 
		props.Username,
		props.Password, 
		props.Host,
		props.Port)
	return rabbitmq.NewConn(
		connectionString,
		rabbitmq.WithConnectionOptionsLogging,
	)
}

func Publish(conn *rabbitmq.Conn, exchange string, headers map[string]interface{}, data []byte) error {

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDurable,
		rabbitmq.WithPublisherOptionsExchangeKind("headers"),
		rabbitmq.WithPublisherOptionsExchangeName(exchange),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if (nil != err) {
		return err
	}
	return publisher.Publish(
		data,
		[]string{"my_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsHeaders(headers),
		rabbitmq.WithPublishOptionsExchange(exchange))
}