package storage

import (
	"context"
	"fmt"
	"os"
	"io"
	"strings"

	"github.com/ohsu-comp-bio/funnel/config"
	"google.golang.org/api/option"
	"google.golang.org/api/iterator"
	"cloud.google.com/go/storage"
)

// The gs url protocol
const gsProtocol = "gs://"

// GoogleCloud provides access to an GS object store.
type GoogleCloud struct {
	svc *storage.Client
}

// NewGoogleCloud creates an GoogleCloud client instance, give an endpoint URL
// and a set of authentication credentials.
func NewGoogleCloud(conf config.GoogleCloudStorage) (*GoogleCloud, error) {
	ctx := context.Background()
	var opts option.ClientOption = nil

	if conf.CredentialsFile != "" {
		// Pull the client configuration (e.g. auth) from a given account file.
		// This is likely downloaded from Google Cloud manually via IAM & Admin > Service accounts.
		opts = option.WithCredentialsFile(conf.CredentialsFile)
	}
	svc, cerr := storage.NewClient(ctx, opts)
	if cerr != nil {
		return nil, cerr
	}

	return &GoogleCloud{svc}, nil
}

// Stat returns information about the object at the given storage URL.
func (gs *GoogleCloud) Stat(ctx context.Context, url string) (*Object, error) {
	u, err := gs.parse(url)
	if err != nil {
		return nil, err
	}

	obj := gs.svc.Bucket(u.bucket).Object(u.path)
	attr, err := obj.Attrs(context.Background())
	if err != nil {
		return nil, err
	}

	return &Object{
		URL:          url,
		Name:         attr.Name,
		ETag:         attr.Etag,
		Size:         int64(attr.Size),
		LastModified: attr.Updated,
	}, nil
}

// List lists the objects at the given url.
func (gs *GoogleCloud) List(ctx context.Context, url string) ([]*Object, error) {
	u, err := gs.parse(url)
	if err != nil {
		return nil, err
	}

	var objects []*Object

	query := &storage.Query{Prefix: u.path}

	//var names []string
	it := gs.svc.Bucket(u.bucket).Objects(ctx, query)
	for {
    attrs, err := it.Next()
    if err == iterator.Done {
        break
    }
		if strings.HasSuffix(attrs.Name, "/") {
			continue
		}
		objects = append(objects, &Object{
			URL:          gsProtocol + attrs.Bucket + "/" + attrs.Name,
			Name:         attrs.Name,
			ETag:         attrs.Etag,
			Size:         int64(attrs.Size),
			LastModified: attrs.Updated,
		})
	}

		/*
	err = gs.svc.Objects.List(u.bucket).Prefix(u.path).Pages(ctx,
		func(objs *storage.Objects) error {

			for _, obj := range objs.Items {
				if strings.HasSuffix(obj.Name, "/") {
					continue
				}

				modtime, _ := time.Parse(time.RFC3339, obj.Updated)

				objects = append(objects, &Object{
					URL:          gsProtocol + obj.Bucket + "/" + obj.Name,
					Name:         obj.Name,
					ETag:         obj.Etag,
					Size:         int64(obj.Size),
					LastModified: modtime,
				})
			}
			return nil
		})
*/

	if err != nil {
		return nil, err
	}
	return objects, nil
}

// Get copies an object from GS to the host path.
func (gs *GoogleCloud) Get(ctx context.Context, url, path string) (*Object, error) {
	u, err := gs.parse(url)
	if err != nil {
		return nil, err
	}
	obj := gs.svc.Bucket(u.bucket).Object(u.path)
	attr, err := obj.Attrs(context.Background())
	if err != nil {
		return nil, err
	}
	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("googleStorage: getting object %s: %v", url, err)
	}
	dest, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("googleStorage: creating file %s: %v", path, err)
	}
	defer r.Close()
	if _, err := io.Copy(dest, r); err != nil {
		return nil, fmt.Errorf("googleStorage: copying file: %v", err)
	}
	return &Object{
		URL:          gsProtocol + attr.Bucket + "/" + attr.Name,
		Name:         attr.Name,
		ETag:         attr.Etag,
		Size:         int64(attr.Size),
		LastModified: attr.Updated,
	}, nil
}

// Put copies an object (file) from the host path to GS.
func (gs *GoogleCloud) Put(ctx context.Context, url, path string) (*Object, error) {
	u, err := gs.parse(url)
	if err != nil {
		return nil, err
	}

	reader, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("googleStorage: opening file: %v", err)
	}
	defer reader.Close()

	obj := gs.svc.Bucket(u.bucket).Object(u.path)

	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return nil, fmt.Errorf("googleStorage: copying file: %v", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("googleStorage: uploading object %s: %v", url, err)
	}
	return gs.Stat(ctx, url)
}

// Join joins the given URL with the given subpath.
func (gs *GoogleCloud) Join(url, path string) (string, error) {
	return strings.TrimSuffix(url, "/") + "/" + path, nil
}

// UnsupportedOperations describes which operations (Get, Put, etc) are not
// supported for the given URL.
func (gs *GoogleCloud) UnsupportedOperations(url string) UnsupportedOperations {
	_, err := gs.parse(url)
	if err != nil {
		return AllUnsupported(err)
	}
	//TODO: Check the bucket?
	//_, err = gs.svc.Buckets.Get(u.bucket).Do()
	//if err != nil {
	//	err = fmt.Errorf("googleStorage: failed to find bucket: %s. error: %v", u.bucket, err)
	//	return AllUnsupported(err)
	//}
	return AllSupported()
}

func (gs *GoogleCloud) parse(rawurl string) (*urlparts, error) {
	if !strings.HasPrefix(rawurl, gsProtocol) {
		return nil, &ErrUnsupportedProtocol{"googleStorage"}
	}

	path := strings.TrimPrefix(rawurl, gsProtocol)
	if path == "" {
		return nil, &ErrInvalidURL{"googleStorage"}
	}

	split := strings.SplitN(path, "/", 2)
	url := &urlparts{}
	if len(split) > 0 {
		url.bucket = split[0]
	}
	if len(split) == 2 {
		url.path = split[1]
	}
	return url, nil
}
