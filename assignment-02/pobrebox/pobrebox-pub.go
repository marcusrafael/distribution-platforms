package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "github.com/streadway/amqp"
)


func failOnError(err error, msg string) {

     if err != nil {
          log.Fatalf("%s: %s", msg, err)
          panic(fmt.Sprintf("%s: %s", msg, err))
     }

}


func main() {

     conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
     failOnError(err, " [x] Failed to connect to RabbitMQ")
     defer conn.Close()

     ch, err := conn.Channel()
     failOnError(err, " [x] Failed to open a channel")
     defer ch.Close()

     err = ch.ExchangeDeclare(
             "pobrebox", // name
             "topic",    // type
             true,       // durable
             false,      // auto-deleted
             false,      // internal
             false,      // no-wait
             nil,        // arguments
     )
     failOnError(err, " [x] Failed to declare an exchange")

     file, _ := ioutil.ReadFile(os.Args[2])

     err = ch.Publish(
             "pobrebox", // exchange
             os.Args[1], // routing key
             false,      // mandatory
             false,      // immediate
             amqp.Publishing{
                     ContentType: "application/text",
                     Body:        file,
             })
     failOnError(err, "Failed to publish a message")
     log.Println(" [*] Sent sent!")
}
