package properties

import (
	"os"
	"strconv"
)

type AppProperties struct {
	Token string
	Rabbit *RabbitProperties
}

type RabbitProperties struct {
	Host string
	Port int
	Username string
	Password string
	Exchange string
}

func LoadApplicationProperties() (*AppProperties, error) {
	var err error
	props := new(AppProperties)
	props.Token = os.Getenv("token")
	rabbitProps := new(RabbitProperties)
	rabbitProps.Host = os.Getenv("rabbit_host")
	rabbitProps.Port, err = strconv.Atoi(os.Getenv("rabbit_port"))
	rabbitProps.Username = os.Getenv("rabbit_username")
	rabbitProps.Password = os.Getenv("rabbit_password")
	rabbitProps.Exchange = os.Getenv("rabbit_exchange")
	props.Rabbit = rabbitProps
	return props, err
}

