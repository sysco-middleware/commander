package commander

// Consumer describes a message consumer
type Consumer interface {
	// Subscribe creates a new topic subscription that will receive
	// messages consumed by the consumer of the given topic. This method
	// will return a message channel and a close function.
	// Once a message is consumed should the marked channel be called. Pass a nil for a successfull consume and
	// a error if a error occured during processing.
	Subscribe(...Topic) (subscription <-chan *Message, marked chan<- error, err error)

	// Unsubscribe unsubscribes the given channel subscription from the given topic.
	// A boolean is returned that represents if the channel successfully got unsubscribed.
	Unsubscribe(subscription <-chan *Message) error

	// Close closes the kafka consumer, all topic subscriptions and event channels.
	Close() error
}
