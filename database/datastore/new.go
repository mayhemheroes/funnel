package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/ohsu-comp-bio/funnel/compute/scheduler"
	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/events"
	"google.golang.org/api/option"
)

// Datastore provides a task database on Google Cloud Datastore.
type Datastore struct {
	scheduler.UnimplementedSchedulerServiceServer
	events.UnimplementedEventServiceServer
	client *datastore.Client
}

// NewDatastore returns a new Datastore instance with the given config.
func NewDatastore(conf config.Datastore) (*Datastore, error) {
	ctx := context.Background()

	opts := []option.ClientOption{}
	if conf.CredentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(conf.CredentialsFile))
	}

	client, err := datastore.NewClient(ctx, conf.Project, opts...)
	if err != nil {
		return nil, err
	}
	return &Datastore{client: client}, nil
}

// Init is a noop action for Datastore.
func (db *Datastore) Init() error {
	return nil
}

func (db *Datastore) Close() {
	db.client.Close()
}
