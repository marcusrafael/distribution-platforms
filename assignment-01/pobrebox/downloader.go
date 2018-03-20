package main

import (
    "log"
    "io"
    "os"
    "net/http"
    "github.com/streadway/amqp"
    "path/filepath"
)


func ReadMessageQueue() {

        conn, err := amqp.Dial("amqp://cloud:cloud@ec2-XX-XXX-XXX-X.compute-1.amazonaws.com:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"pobrebox", // queue
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
                        fileName := filepath.Base(string(d.Body))
			log.Println(fileName)
                        name := "/home/ubuntu/pobrebox/" + fileName
                        DownloadFile(name, string(d.Body))
			//log.Println("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever


}

func DownloadFile(filepath string, url string) error {

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}
func main() {
    ReadMessageQueue()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
