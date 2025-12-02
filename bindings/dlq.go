package bindings

// DLQMetadata holds the name and metadata of the dead letter target output binding.
type DLQMetadata struct {
	// Name is the name of the output binding component configured to receive failed messages.
	Name string

	// Metadata is the optional metadata to pass to the output binding invoke request.
	Metadata map[string]string
}

// DLQInputBinding is an optional interface that input bindings can implement
// to provide a dead letter queue configuration.
type DLQInputBinding interface {
	// GetDeadLetterBinding returns the configured output binding name and metadata
	// that should be used as the Dead Letter Queue for failed events.
	GetDeadLetterBinding() DLQMetadata
}
