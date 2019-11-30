package repos

import (
	"cloud.google.com/go/datastore"
	"context"
)

type Conf struct {
	Secret string `datastore:"secret"`
}

type ConfRepo interface {
	GetConf(ctx context.Context) (*Conf, error)
}

// ======================================================================

type DatastoreConfRepoImpl struct {
	projectID string
}

func NewDatastoreConfRepoImpl(projectID string) ConfRepo {
	return &DatastoreConfRepoImpl{projectID: projectID}
}


func (d DatastoreConfRepoImpl) getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	client, err = datastore.NewClient(ctx, d.projectID)
	return
}


func (d DatastoreConfRepoImpl) GetConf(ctx context.Context) (*Conf, error) {
	client, err := d.getDatastoreClient(ctx)
	if err != nil {
		return nil, err
	}

	k := datastore.NameKey("Conf", "1", nil)
	e := new(Conf)
	err = client.Get(ctx, k, e)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		return nil, err
	}
	return e, nil
}

