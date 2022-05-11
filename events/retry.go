package events

import (
	"context"

	"github.com/ohsu-comp-bio/funnel/util"
)

type retrier struct {
	*util.Retrier
	Writer Writer
}

func (r *retrier) WriteEvent(ctx context.Context, e *Event) error {
	return r.Retry(ctx, func() error {
		_, err := r.Writer.WriteEvent(ctx, e)
		return err
	})
}
