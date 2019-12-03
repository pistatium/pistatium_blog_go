package repos

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"strconv"
	"time"
)

type Entry struct {
	Id         string     `datastore:"-"`
	Title      string     `datastore:"title,noindex"`
	Body       string     `datastore:"body,noindex"`
	More       string     `datastore:"more,noindex"`
	Category   string     `datastore:"category"`
	Datetime   *time.Time `datastore:"datetime"`
	Updated    *time.Time `datestore:"datetime"`
	Public     bool       `datastore:"public"`
	IsMarkdown bool       `datastore:"is_markdown,noindex"`
	//ModifyUser string            `datastore:"modify_user"`
	//CreateUser *datastore.Entity `datastore:"create_user"`
}

type Entries struct {
	Entries []*Entry `json:"entries"`
}

type EntryRepo interface {
	GetEntries(ctx context.Context, offset int, limit int, publicOnly bool) ([]*Entry, error)
	GetEntry(ctx context.Context, id string) (*Entry, error)
	// CreateEntry(ctx context.Context, entry Entry) error
	UpdateEntry(ctx context.Context, id string, entry Entry) error
}

// ======================================================================

type DatastoreEntryRepoImpl struct {
	projectID string
}

func NewDatastoreEntryRepoImpl(projectID string) EntryRepo {
	return &DatastoreEntryRepoImpl{projectID: projectID}
}

func (d DatastoreEntryRepoImpl) getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	client, err = datastore.NewClient(ctx, d.projectID)
	return
}

func (d DatastoreEntryRepoImpl) GetEntries(ctx context.Context, offset int, limit int, publicOnly bool) ([]*Entry, error) {
	client, err := d.getDatastoreClient(ctx)
	if err != nil {
		return nil, err
	}

	// 最新10件取得
	q := datastore.NewQuery("Blog").
		Filter("public =", true).
		Order("-datetime").
		Limit(limit).
		Offset(offset)
	entries := make([]*Entry, 0, limit)
	keys, err := client.GetAll(ctx, q, &entries)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		return nil, err
	}

	for i, key := range keys {
		entries[i].Id = strconv.FormatInt(key.ID, 10)
	}
	return entries, nil
}

func (d DatastoreEntryRepoImpl) GetEntry(ctx context.Context, id string) (*Entry, error) {
	client, err := d.getDatastoreClient(ctx)
	if err != nil {
		return nil, err
	}
	iid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	k := datastore.IDKey("Blog", int64(iid), nil)
	e := new(Entry)
	err = client.Get(ctx, k, e)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		return nil, err
	}
	e.Id = id
	return e, nil
}

//func (d DatastoreEntryRepoImpl) CreateEntry(ctx context.Context, entry Entry) error {
//	client, err := d.getDatastoreClient(ctx)
//	if err != nil {
//		return err
//	}
//
//	if entry.Datetime == nil {
//		now := time.Now()
//		entry.Datetime = &now
//	}
//
//	id := fmt.Sprintf("%d", time.Now().UnixNano() / 1000000)
//	key := datastore.NameKey("Blog", id, nil)
//	if _, err := client.Put(ctx, key, &entry); err != nil {
//		return err
//	}
//	return nil
//}

func (d DatastoreEntryRepoImpl) UpdateEntry(ctx context.Context, id string, entry Entry) error {
	client, err := d.getDatastoreClient(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	if entry.Datetime == nil {
		entry.Datetime = &now
	}
	entry.Updated = &now

	key := datastore.IDKey("Blog", id, nil)
	fmt.Println(key, entry)
	if _, err := client.Put(ctx, key, &entry); err != nil {
		return err
	}
	return nil
}
