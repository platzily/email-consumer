package main

import (
	"context"
	"fmt"

	"github.com/platzily/email-consumer/config"
	"github.com/platzily/email-consumer/drivers/mongodb"
)

func main() {

	mongoConfig := config.ReadMongoDBConfig()

	client := mongodb.NewConnection(mongoConfig.URL)
	defer client.Disconnect(context.TODO())

	eventRepository := mongodb.New(client)

	result, _ := eventRepository.GetById(0)
	fmt.Printf("This is the result: %v", result)

	eventRepository.UpdateById(0, "Updated")

	updatedResult, _ := eventRepository.GetById(0)
	fmt.Printf("This is the updated result: %v", updatedResult)

	fmt.Println("Hello world, from notification email consumer")
}
