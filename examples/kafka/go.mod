module github.com/jeroenrinzema/commander/examples/kafka

replace github.com/jeroenrinzema/commander => ../../

replace github.com/jeroenrinzema/commander/dialects/kafka => ../../dialects/kafka

require (
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/jeroenrinzema/commander v1.0.0-rc.25
	github.com/jeroenrinzema/commander/dialects/kafka v0.0.0-20181217103823-01d74b882250
)
