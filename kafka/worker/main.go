package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	
	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers,config)

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// criando a mensagem que vai ser enviada
	message := &sarama.ProducerMessage{
		Topic: "meu-topico",
		Value: sarama.StringEncoder("dasdasdadas"),
	}

	partition ,offset, err := producer.SendMessage(message)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensagem enviada para o topico meu-topico, na particao %d, no offset %d \n",partition,offset)
}