package services

import "github.com/IBM/sarama"

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(EventHandler EventHandler) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler: EventHandler}
}

func (obj consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (obj consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
func (obj consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		obj.eventHandler.Handle(message.Topic, message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}
