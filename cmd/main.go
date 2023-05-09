package cmd

import (
	// db "github.com/ByteGum/go-ssrc/pkg/core/db"
	// p2p "github.com/ByteGum/go-ssrc/pkg/core/p2p"
	// originatorRoutes "github.com/ByteGum/go-ssrc/pkg/core/rest/originator"
	// // ds "github.com/ipfs/go-ds-badger"
	// "github.com/gin-gonic/gin"
	"io/ioutil"

	badger "github.com/dgraph-io/badger"
)

func main() {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		panic(err)
	}
	//defer ioutil.removeDir(dir)
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//if 1 == 2 {
	// d := ds.Datastore{DB: db}

	// go p2p.Run()
	// r := gin.Default()
	// r = originatorRoutes.Init(r)
	// r.Run("localhost:8081")
	//}
}
