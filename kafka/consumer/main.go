package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Certifique-se de que o endereço do broker não inclui 'http://'
	brokers := []string{"localhost:9092"}

	fmt.Println("Iniciando consumidor Kafka...")

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}
	defer func() {
		fmt.Println("Fechando consumidor...")
		if err := consumer.Close(); err != nil {
			log.Fatalf("Erro ao fechar consumidor: %v", err)
		}
	}()

	fmt.Println("Consumidor criado com sucesso. Iniciando consumo de partição...")

	partitionConsumer, err := consumer.ConsumePartition("meu-topico", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Erro ao consumir partição: %v", err)
	}

	fmt.Println("Consumindo partição com sucesso.")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	messages := partitionConsumer.Messages()

	for {
		select {
		case msg := <-sigCh:
			fmt.Printf("Saindo da aplicação, código recebido %v", msg)
			return
		case msg := <-messages:
			fmt.Printf("Mensagem recebida: %s \n", string(msg.Value))
		}
	}
}
