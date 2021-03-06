package network

import (
	"github.com/ThingiverseIO/thingiverseio/config"
	"github.com/ThingiverseIO/uuid"
	"github.com/joernweissenborn/eventual2go"
)

type Transport interface {
	eventual2go.Shutdowner

	// Init initializes a transports incoming channel.
	Init(cfg *config.Config) error

	// Connect connectes to peer using the given details.
	Connect(details Details, uuid uuid.UUID) (Connection, error)

	// Details returns the details of the incoming connection. This will be advertised to other peers.
	Details() Details

	// Packages returns a stream of incoming messages.
	Packages() *PackageStream
}
