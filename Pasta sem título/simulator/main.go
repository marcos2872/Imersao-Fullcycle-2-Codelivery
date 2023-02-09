package main

import (
	"fmt"
	"log"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/codmarcos/imersaofsfc2-simulator/application/kafka"
	"github.com/codmarcos/imersaofsfc2-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	cosumer := kafka.NewKafkaConsumer(msgChan)
	go cosumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}