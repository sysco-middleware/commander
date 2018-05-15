package websocket

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander/example/webservice/commands"
)

// Consume kafka event messages
func Consume() {
	consumer := commands.Commander.Consume("events")

	for msg := range consumer.Messages {
		event := commander.Event{}
		json.Unmarshal(msg.Value, &event)

		data, err := json.Marshal(event)

		if err != nil {
			continue
		}

		Broadcast(string(data))
	}
}