package main

import (
     "os"
     "log"
     "github.com/streadway/amqp"
)


func failOnError(err error, msg string) {

     if err != nil {
          log.Fatalf("%s: %s", msg, err)
     }

}


func main() {
     conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
     failOnError(err, "Failed to connect to RabbitMQ")
     defer conn.Close()

     ch, err := conn.Channel()
     failOnError(err, "Failed to open a channel")
     defer ch.Close()

     err = ch.ExchangeDeclare(
	"pobrebox",   // name
	"topic", // type
	true,     // durable
	false,    // auto-deleted
	false,    // internal
	false,    // no-wait
	nil,      // arguments
     )
     failOnError(err, "Failed to declare an exchange")

     q, err := ch.QueueDeclare(
	"",    // name
	false, // durable
	false, // delete when unused
	true,  // exclusive
	false, // no-wait
	nil,   // arguments
     )
     failOnError(err, "Failed to declare a queue")

     for _, s := range os.Args[1:] {

	err = ch.QueueBind(
		q.Name, // queue name
		s,     // routing key
		"pobrebox", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
     }

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

     forever := make(chan bool)

     go func() {
	for d := range msgs {
             log.Println("File received")
             final, _ := os.Create(os.Args[2])
             final.Write(d.Body)
             defer final.Close()
	}
     }()

     log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
     <-forever
}
