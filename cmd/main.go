package main

import (
	"amqplogreader/internal/config"
	"amqplogreader/internal/domain"
	"encoding/json"
	"fmt"
	"log"

	amqplog "github.com/wilfridterry/contact-list/pkg/amqp_log"
)

func main() {
	cf, err := config.NewConfig();
	if err != nil {
		log.Panic(err)
	}

	amqpClient, err := amqplog.New(&amqplog.ConfigOptions{
		Host:     cf.Rabbitmq.Host,
		Port:     int(cf.Rabbitmq.Port),
		Username: cf.Rabbitmq.Username,
		Password: cf.Rabbitmq.Password,
		Queue:    cf.Rabbitmq.Queue,
	})
	if err != nil {
		log.Panic(err)
	}
	defer amqpClient.Close()

	messages, err := amqpClient.GetLogs()
	if err != nil {
		log.Panic(err)
	}
	for msg := range messages {
		var msgLog domain.MessageLog 
		if err := json.Unmarshal(msg.Body, &msgLog); err != nil {
			log.Panic(err)
		}

		fmt.Println(msgLog)
	}

	forever := make(chan struct{})

	<-forever
}