package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

const (
	address = "localhost:9092"
	topic   = "test_topic"
	network = "tcp"
)

func main() {

	var message string

	conn, err := kafka.DialLeader(context.Background(), network, address, topic, 0)

	if err != nil {
		fmt.Println("*** Unable to connect to Kafka ***")
		os.Exit(1)
	}

	fmt.Printf("*** Connected to Kafka on %s, messages will be published on %s ***\n", address, topic)

	//Taking input from prompt
	for {

		fmt.Println("*** Enter a message: ***")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message = scanner.Text() // Println will add back the final '\n'
		publish(message, conn)

	}
}

func publish(message string, connection *kafka.Conn) {

	_, err := connection.Write([]byte(message))
	if err != nil {
		fmt.Println("*** Unable to publish the message in the topic ***")
	}

}
