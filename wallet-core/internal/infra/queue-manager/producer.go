package qm

import (
	"encoding/json"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

func NewProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{ConfigMap: configMap}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}

	bs, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          bs,
		Key:            key,
	}

	err = producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
