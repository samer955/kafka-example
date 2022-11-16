package main

import (
	"fmt"
	"kafka-example/kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Println("Starting Kafka...")
	go kafka.StartKafka()
	fmt.Println("Waiting for incoming messages...")

	//kafka.WriteMessage()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	fmt.Println("Stopping the program ")

}
