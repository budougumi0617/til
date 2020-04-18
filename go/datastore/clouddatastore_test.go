package datastore

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	"go.mercari.io/datastore/clouddatastore"
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

// TODO: 検索するテストを書く

// TODO: 削除するテストを書く

// TODO: インターフェイスごしのテストを書く
