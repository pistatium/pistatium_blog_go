package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"strconv"
	"time"
)

type Photo struct {
	Id       int64      `datastore:"-"`
	Comment  string     `datastore:"comment"`
	Datetime *time.Time `datastore:"datetime"`
	Image    []byte     `datastore:"image"`
	Title    string     `datastore:"title"`
}

type PhotoRepo interface {
	GetPhoto(ctx context.Context, id string) (*Photo, error)
}

// ======================================================================

type DatastorePhotoRepoImpl struct {
	projectID string
}

func NewDatastorePhotoRepoImpl(projectID string) PhotoRepo {
	return &DatastorePhotoRepoImpl{projectID: projectID}
}

func (d DatastorePhotoRepoImpl) getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	client, err = datastore.NewClient(ctx, d.projectID)
	return
}

func (d DatastorePhotoRepoImpl) GetPhoto(ctx context.Context, id string) (*Photo, error) {
	client, err := d.getDatastoreClient(ctx)
	if err != nil {
		return nil, err
	}
	iid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	k := datastore.IDKey("Photo", int64(iid), nil)
	e := new(Photo)
	err = client.Get(ctx, k, e)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		return nil, err
	}
	e.Id = int64(iid)
	return e, nil
}
