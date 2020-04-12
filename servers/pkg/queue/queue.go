package queue

import "context"

// Queue is the interface for queueing services.
type Queue interface {
	Send(body string, opts ...Opt) (SendMessageOutput, error)
	SendWithContext(ctx context.Context, body string, opts ...Opt) (SendMessageOutput, error)
}

// SendMessageInput is the format to send message to
// the queueing service.
type SendMessageInput interface {
	GoString() string
	String() string
	Validate() error
}

// SendMessageOutput is the format to recieve message
// from the queueing service.
type SendMessageOutput interface {
	GoString() string
	String() string
}

// Opt is the format that acts an option for `SendMessageInput`.
type Opt func(SendMessageInput)
