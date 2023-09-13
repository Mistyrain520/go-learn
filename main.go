package main

import "golearn/kafka"

func main() {
	// go kafka.SaramaProducer()
	go kafka.SaramaConsumer()
	select {}
}
