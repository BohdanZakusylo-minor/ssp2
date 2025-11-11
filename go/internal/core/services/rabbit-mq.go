package services

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ConnectToRabbit(msg []byte) {
	log.Println(" Rabbit MQ Easy demo!")

	conn, err := amqp.Dial("amqp://user:password@rabbit-mq:5672/")
	if err != nil {
		log.Println("Error in connection")
		return
	}

	//close connection at end of program

	defer conn.Close()

	fmt.Println(" succesfulle connected to rabbit mq")

	//open the channel
	chl, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	//close connection at end of program
	defer chl.Close()

	q, err := chl.QueueDeclare(
		"southpark_messages", // name of the queue
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)

	if err != nil {
		fmt.Println("Error in queue declaration")
		panic(err)
	}

	fmt.Println(q)

	err = chl.Publish(
		"",                   // exchange
		"southpark_messages", // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)

	if err != nil {
		fmt.Println("Error in publishing message")
		panic(err)
	}

	fmt.Println(" Pulished message to queue")
}
