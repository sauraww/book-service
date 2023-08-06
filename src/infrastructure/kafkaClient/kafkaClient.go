package kafkaclient

type KafkaPort interface {
	ProduceMessage(topic string, message []byte)
	ConsumeMessage(topic string) ([]byte, error)
}

type KafkaClient interface {
	SendMessage(topic string, message []byte) error
	ReadMessage(topic string) ([]byte, error)
}

type KafkaService struct {
	client KafkaClient
}

func (ks *KafkaService) ProduceMessage(topic string, message []byte) error {
	return ks.client.SendMessage(topic, message)
}

func (ks *KafkaService) ConsumeMessage(topic string) ([]byte, error) {
	return ks.client.ReadMessage(topic)
}
