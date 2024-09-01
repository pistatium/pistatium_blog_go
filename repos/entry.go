package repos

import (
	"cloud.google.com/go/datastore"
	"context"
	"net/url"
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
	Updated    *time.Time `datastore:"updated"`
	Public     bool       `datastore:"public"`
	IsMarkdown bool       `datastore:"is_markdown,noindex"`
	Thumbnail  string     `datastore:thumbnail,noindex`
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
	q := datastore.NewQuery("Blog")

	if publicOnly {
		q = q.Filter("public =", true)
	}
	q = q.Order("-datetime").
		Limit(limit).
		Offset(offset)
	entries := make([]*Entry, 0, limit)
	return entries, nil
	//keys, err := client.GetAll(ctx, q, &entries)
	//if err != nil && err != err.(*datastore.ErrFieldMismatch) {
	//	return nil, err
	//}
	//
	//for i, key := range keys {
	//	entries[i].Id = strconv.FormatInt(key.ID, 10)
	//}
	//return entries, nil
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
	return nil, nil
	//k := datastore.IDKey("Blog", int64(iid), nil)
	//e := new(Entry)
	//err = client.Get(ctx, k, e)
	//if err != nil {
	//	if err == datastore.ErrNoSuchEntity {
	//		//entityが存在しない場合
	//		return e, nil
	//	}
	//	if err == err.(*datastore.ErrFieldMismatch) {
	//		// datastore.ErrFieldMismatch は存在しないカラムをマッピングしようとするとでるので無視
	//	} else {
	//		return nil, err
	//	}
	//}
	//e.Id = id
	//return e, nil
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
	iid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	now := time.Now()

	if entry.Datetime == nil {
		entry.Datetime = &now
	}
	entry.Updated = &now

	entry.IsMarkdown = true
	if entry.Thumbnail == "" {
		titleEnc := url.PathEscape(entry.Title)
		entry.Thumbnail = "https://kimihiro-n.firebaseapp.com/uniqueogp?title=" + titleEnc + "&brand=Pistatium&mode=white"
	}

	key := datastore.IDKey("Blog", int64(iid), nil)
	if _, err := client.Put(ctx, key, &entry); err != nil {
		return err
	}
	return nil
}
