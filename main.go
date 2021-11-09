package main

import (
	"context"
	"fmt"
	"log"

	"github.com/platzily/email-consumer/config"
	"github.com/platzily/email-consumer/drivers/mongodb"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	mongoConfig := config.ReadMongoDBConfig()
	rabbitMQconfig := config.ReadRabbitMQConfig()

	client := mongodb.NewConnection(mongoConfig.URL)
	defer client.Disconnect(context.TODO())

	eventRepository := mongodb.New(client)

	fmt.Println("Starting consumer...")
	conn, err := amqp.Dial(rabbitMQconfig.URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//Event repository
	result, _ := eventRepository.GetById(0)
	fmt.Printf("This is the result: %v", result)

	eventRepository.UpdateById(0, "Updated")

	updatedResult, _ := eventRepository.GetById(0)
	fmt.Printf("This is the updated result: %v", updatedResult)

	fmt.Println("Hello world, from notification email consumer")

	//RabbitMQ queue declaration
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	//Consumer reciving messages
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
