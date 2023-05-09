package db

import (
	"context"
	"os"
	"path/filepath"

	utils "github.com/ByteGum/go-ssrc/utils"
	ds "github.com/ipfs/go-datastore"
)

func Key(key string) ds.Key {
	return ds.NewKey(key)
}

func New(mainCtx *context.Context, keyStore string) *Datastore {
	ctx, cancel := context.WithCancel(*mainCtx)
	defer cancel()
	cfg, ok := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	if !ok {

	}
	dir := filepath.Join(cfg.DataDir, keyStore)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	ds, err := NewDatastore(dir, &DefaultOptions)
	if err != nil {
		panic(err)
	}
	return ds

	// err = db.View(func(txn *badger.Txn) error {
	// 	_, err := txn.Get([]byte("key"))
	// 	// We expect ErrKeyNotFound
	// 	fmt.Println("Error", err)
	// 	return nil
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// txn := db.NewTransaction(true) // Read-write txn
	// err = txn.SetEntry(badger.NewEntry([]byte("key"), []byte("value set successfully")))
	// if err != nil {
	// 	panic(err)
	// }
	// err = txn.Commit()
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte("key"))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	val, err := item.ValueCopy(nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("Valueeeee %s\n", string(val))
	// 	return nil
	// })

	// if err != nil {
	// 	panic(err)
	// }
}
