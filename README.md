# Commander

Commander gives you a toolset for writing event driven applications with Kafka as it's event log. Commender encurages to use the CQRS pattern to seperate write and read operations.

## Usage and documentation

Please see [godoc](https://godoc.org/github.com/jeroenrinzema/commander) for detailed usage docs.

## Getting started

A "data set" is represented in Commander as a group. Every data set contains a group of topics which are used to write different type of messages (commands, events). Commander does not limit on how many topics could be defined for a single group.

```go
package main

import (
	"github.com/jeroenrinzema/commander"
	uuid "github.com/satori/go.uuid"
)

cart := commander.Group{
	Topics: []commander.Topic{
		commander.Topic{
			Name: "cart-commands",
			Type: commander.CommandTopic,
			Consume: true,
			Produce: false
		},
		commander.Topic{
			Name: "cart-events",
			Type: commander.EventTopic,
			Consume: false,
			Produce: true
		},
	}
}

func main() {
	config := commander.NewConfig()
	config.Brokers = []string{"..."}
	config.AddGroups(cart)

	cmdr := commander.New(&config)
	go cmdr.Consume()

	cart.CommandHandle("NewCart", func(command *commander.Command) *commander.Event {
		// ...

		return command.NewEvent("CartCreated", 1, uuid.NewV4(), nil)
	})
}
```

## GDPR

Commander offers various APIs to handle GDPR complaints. To keep the immutable ledger immutable, do we offer the plausibility to encrypt all data sensitive events. Once a "right to erasure" request needs to be preformed can all data be erased by simply throwing away the key.
