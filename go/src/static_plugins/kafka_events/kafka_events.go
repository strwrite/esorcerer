package main

import (
	"flag"
	"os"
	"log"
	"strings"
	"os/signal"
	"gopkg.in/Shopify/sarama.v1"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type pluginVersion struct {
	Githash, Buildstamp, API string
}

var (
	Buildstamp = ""
	Githash = ""
	API = "v1"

	brokers = flag.String("brokers",
		getEnv("KAFKA_PEERS", "kafka:9092"),
		"The Kafka brokers to connect to, as a comma separated list")
	brokerList []string
)

func Plugin_init() {
	Version = pluginVersion{Githash, Buildstamp, API}
	flag.Parse()

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList = strings.Split(*brokers, ",")
	log.Printf("Kafka brokers : %s", strings.Join(brokerList, ", "))

}

func Spawn_event_loop() {

	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		log.Panicf("Error during kafka consumer spawn in %s: %v", brokerList, err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	kafka_topic := "my_topic"
	partitionConsumer, err := consumer.ConsumePartition(kafka_topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Panicf("Error during kafka consumer spawn on topic %s in %s: %v",
				kafka_topic, brokerList, err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)

}


