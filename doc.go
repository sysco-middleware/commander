/*
Package commander gives you a toolset for writing distributed applications using ideas from CQRS and Event Sourcing using Kafka as a event log. The commander pattern is inspired of talk and architecture given by Bobby Calderwood.

    https://github.com/capitalone/cqrs-manager-for-distributed-reactive-services/blob/master/doc/architecture.png
    https://www.youtube.com/watch?v=B1-gS0oEtYc&t

Getting started

To get started quickly download/fork the boilerplate project.

    https://github.com/sysco-middleware/commander-boilerplate

The pattern

All services inside the commander pattern are triggered by events and commands and are immutable. The commander pattern exists out of 4 layers. Every layer has it's own responsibilities and contain different parts of your application.

Web service

This layer is accessible from the outside. The main responsibility is to preform queries on states or to write commands to the event log. Once a command is received is the data not "yet" validated. Optionally could this layer authenticate incoming requests.

Event log

The event log is the communication layer in between the web service layer and the business logic layer. It's main responsibility is to communicate messages and "log" them in the process. Kafka is used in Commander as the event log.

Business logic

The logics layer consumes commands/events to process them. Two types of consumers could exists in the business logic layer. The command processor processes commands received from the "commands" topic and generates a resulting event. This event could be a error or the resulting generated data. The projector processes events received from the "events" topic. A projector creates a projection of the consumed events. This projection could be consumed by the web service layer. Command processes and projector processes should never share their states between one another. If a command process requires to preform a validation/query on the latest state should he do it on it's own.

Datastore and projections

This layer contains sets of states that could be used to query upon. Every service could have it's own projection created of the consumed commands/events.
*/
package commander