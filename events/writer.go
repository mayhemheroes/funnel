package events

import (
	"context"
)

// Writer provides write access to a task's events
type Writer interface {
	//WriteEvent(context.Context, *Event) error
	EventServiceServer
	Close()
}

// MultiWriter allows writing an event to a list of Writers.
// Writing stops on the first error.
type MultiWriter []Writer

func (mw *MultiWriter) mustEmbedUnimplementedEventServiceServer() {}

// WriteEvent writes an event to all the writers. Writing stops on the first error.
func (mw *MultiWriter) WriteEvent(ctx context.Context, ev *Event) (*WriteEventResponse, error) {
	for _, w := range *mw {
		e, err := w.WriteEvent(ctx, ev)
		if err != nil {
			return e, err
		}
	}
	return &WriteEventResponse{}, nil
}

func (mw *MultiWriter) Close() {
	for _, w := range *mw {
		w.Close()
	}
}

// Noop provides an event writer that does nothing.
type Noop struct {
	UnimplementedEventServiceServer
}

// WriteEvent does nothing and returns nil.
func (n Noop) WriteEvent(ctx context.Context, ev *Event) (*WriteEventResponse, error) {
	return &WriteEventResponse{}, nil
}

func (n Noop) Close() {}
