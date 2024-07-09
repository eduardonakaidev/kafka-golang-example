# app example kafka consumer and producer golang
## necessary to execute the project
- docker compose 
- docker 
- golang

## installation links
- https://go.dev/doc/install
- https://docs.docker.com/engine/install/


## Run Docker compose 
```bash
  docker-compose up -d
```


## Run consumer
Run the consumer to see what will be sent to Kafka, it will see the messages from the "meu-topico" topic and print it to the console
```bash
  cd consumer && go run main.go
```
## create message in worker
```bash
message := &sarama.ProducerMessage{
		Topic: "meu-topico",
		Value: sarama.StringEncoder("dasdasdadas"),
	}
 ```
## Run worker
```bash
  cd worker && go run main.go
```

