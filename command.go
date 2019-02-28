package commander

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gofrs/uuid"
)

// NewCommand constructs a new command
func NewCommand(action string, version int8, key uuid.UUID, data []byte) *Command {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	// Fix: unexpected end of JSON input
	if len(data) == 0 {
		data = []byte("null")
	}

	command := &Command{
		Key:     key,
		Headers: make(map[string]string),
		ID:      id,
		Action:  action,
		Version: version,
		Data:    data,
	}

	return command
}

// Command contains the information of a consumed command.
type Command struct {
	Key     uuid.UUID         `json:"key,omitempty"`
	Headers map[string]string `json:"headers"`
	ID      uuid.UUID         `json:"id"`
	Action  string            `json:"action"`
	Data    json.RawMessage   `json:"data"`
	Version int8              `json:"version"`
	Origin  Topic             `json:"-"`
}

// NewEvent creates a new acknowledged event as a response to this command.
func (command *Command) NewEvent(action string, version int8, data []byte) *Event {
	event := NewEvent(action, version, command.ID, command.Key, data)
	return event
}

// NewError creates a error event as a response to this command.
func (command *Command) NewError(action string, data []byte) *Event {
	event := NewEvent(action, 0, command.ID, command.Key, data)
	event.Status = StatusInternalServerError

	return event
}

// Populate populates the command struct with the given message
func (command *Command) Populate(message *Message) error {
	command.Headers = make(map[string]string)
	var throw error

headers:
	for _, header := range message.Headers {
		key := header.Key
		value := string(header.Value)

		switch key {
		case ActionHeader:
			command.Action = value
			continue headers
		case IDHeader:
			id, err := uuid.FromString(value)

			if err != nil {
				throw = err
			}

			command.ID = id
			continue headers
		case VersionHeader:
			version, err := strconv.ParseInt(string(header.Value), 10, 8)
			if err != nil {
				return err
			}

			command.Version = int8(version)
			continue headers
		}

		command.Headers[key] = value
	}

	id, err := uuid.FromString(string(message.Key))

	if err != nil {
		throw = err
	}

	if len(command.Action) == 0 {
		return errors.New("No command action is set")
	}

	command.Key = id
	command.Data = message.Value
	command.Origin = message.Topic

	return throw
}
