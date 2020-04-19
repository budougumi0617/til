package datastore

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
	"google.golang.org/api/iterator"
)

type CloudDatastoreStruct struct {
	Test string
}

func TestCloudDatastore_Put(t *testing.T) {
	err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18081")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	cli, _ := datastore.NewClient(ctx, "budougumi0617")
	client, err := clouddatastore.FromClient(ctx, cli)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer client.Close()
	// defer cleanUp() need if delete data after test

	key := client.IncompleteKey("CloudDatastoreStruct", nil)
	key, err = client.Put(ctx, key, &CloudDatastoreStruct{"Hi!"})
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("key: %s", key.String())
}

// https://github.com/mercari/datastore/blob/5469cfe0447d4625defa1a89b59f53284652ff35/testbed/clouddatastore_test.go#L20
func cleanUp() error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "budougumi0617")
	if err != nil {
		return err
	}
	defer client.Close()

	q := datastore.NewQuery("__kind__").KeysOnly()
	iter := client.Run(ctx, q)
	var kinds []string
	for {
		key, err := iter.Next(nil)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		kinds = append(kinds, key.Name)
	}

	for _, kind := range kinds {
		q := datastore.NewQuery(kind).KeysOnly()
		keys, err := client.GetAll(ctx, q, nil)
		if err != nil {
			return err
		}
		err = client.DeleteMulti(ctx, keys)
		if err != nil {
			return err
		}
	}

	return nil
}

type Title string

func (t Title) ToPropertyValue(ctx context.Context) (interface{}, error) {
	return string(t), nil
}
func (t Title) FromPropertyValue(ctx context.Context, p datastore.Property) (dst interface{}, err error) {
	ts := fmt.Sprintf("%v", p.Value)
	return Title(ts), nil
}

type Post struct {
	// https://godoc.org/go.mercari.io/datastore/boom#hdr-Key_handling
	// stringにすると自動で補完されない。
	ID      int64     `datastore:"-" boom:"id" json:"id"`
	Title   Title     `json:"title"`
	Body    string    `json:"body"`
	EntryAt time.Time `json:"entry_at"`
}

func TestBoom_PutAndGetAll(t *testing.T) {
	err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18081")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	cli, _ := datastore.NewClient(ctx, "budougumi0617")
	client, err := clouddatastore.FromClient(ctx, cli)
	if err != nil {
		t.Fatalf("fromClient failed: %v", err)
	}
	defer client.Close()
	defer cleanUp()

	bm := boom.FromClient(ctx, client)
	p := &Post{
		Title:   "Data",
		Body:    "ok?",
		EntryAt: time.Now(),
	}
	k, err := bm.Put(p)
	if err != nil {
		t.Fatalf("bm.Put failed: %v", err)
	}

	t.Logf("key: %q", k.String())
	// int64の場合は自動でセットされる。stringの場合は自動でセットされない
	t.Logf("ID: %d", p.ID)
	var dst []*Post
	// キャストしないと、datastore: bad query filter value type: invalid Value type datastore.Title になってしまう。
	q := bm.Client.NewQuery(bm.Kind(p)).Filter("Title =", string(p.Title))
	ks, err := bm.GetAll(q, &dst)
	if err != nil {
		t.Fatalf("bm.GetAll failed: %v", err)
	}
	for _, k := range ks {
		t.Logf("key: %s", k.String())
	}
	for k, v := range dst {
		t.Logf("dst[%d] %v", k, v)
	}
}

func TestBoom_GetByName(t *testing.T) {
	err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18081")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	cli, _ := datastore.NewClient(ctx, "budougumi0617")
	client, err := clouddatastore.FromClient(ctx, cli)
	if err != nil {
		t.Fatalf("fromClient failed: %v", err)
	}
	defer client.Close()
	defer cleanUp()

	bm := boom.FromClient(ctx, client)
	ps := []*Post{
		{Title: "MyTitle", Body: "blue", EntryAt: time.Now()},
		{Title: "YourTitle", Body: "red", EntryAt: time.Now()},
		{Title: "MyTitle", Body: "blue", EntryAt: time.Now()},
		{Title: "YourTitle", Body: "green", EntryAt: time.Now()},
		{Title: "MyTitle", Body: "red", EntryAt: time.Now()},
	}

	if _, err := bm.PutMulti(ps); err != nil {
		t.Fatalf("bm.Put failed: %v", err)
	}
	var dst []*Post
	// https://cloud.google.com/datastore/docs/concepts/queries
	q := bm.Client.NewQuery(bm.Kind(ps[0])).Filter("Title =", string(ps[0].Title))
	ks, err := bm.GetAll(q, &dst)
	if err != nil {
		t.Fatalf("bm.GetAll failed: %v", err)
	}
	if len(ks) != 3 {
		t.Errorf("result count = %d", len(ks))
	}

	q = bm.Client.NewQuery(bm.Kind(ps[0])).Filter("Title =", string(ps[1].Title)).Filter("Body =", "green")
	ks, err = bm.GetAll(q, &dst)
	if err != nil {
		t.Fatalf("bm.GetAll failed: %v", err)
	}
	if len(ks) != 1 {
		t.Errorf("result count = %d", len(ks))
	}
}

// TODO: 削除するテストを書く

// TODO: インターフェイスごしのテストを書く
