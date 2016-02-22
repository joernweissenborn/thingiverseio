package connection

import (
	"github.com/joernweissenborn/eventual2go"
	"github.com/joernweissenborn/thingiverse.io/config"
	"github.com/joernweissenborn/thingiverse.io/service"
	"github.com/joernweissenborn/thingiverse.io/service/messages"
)

//go:generate event_generator -t Message

type Message struct {
	Iface   string
	Sender  config.UUID
	Payload []string
}

func isMsgFromSender(sender config.UUID) MessageFilter {
	return func(m Message) bool {
		return sender == m.Sender
	}
}

func validMsg(m Message) bool {
	if len(m.Payload) < 2 {
		return false
	}
	p := []byte(m.Payload[0])[0]

	if p != service.PROTOCOLL_SIGNATURE {
		return false
	}
	return true
}

func transformToMessage(d eventual2go.Data) eventual2go.Data {
	m := d.(Message)
	return messages.Unflatten(m.Payload)
}

type outgoingMessage struct {
	sent    *eventual2go.Completer
	payload [][]byte
}
