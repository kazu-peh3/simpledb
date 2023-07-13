package db

import (
	"os"

	"github.com/kazu-peh3/toydb/storage"
)

type ToyDb struct {
	exit chan int
	// contexts    map[string]*dbSession
	storage *storage.Storage
	// catalog     *storage.Catalog
	// tranManager *storage.TransactionManager
	home string
}

// type dbSession struct {
// 	tran *storage.Transaction
// }

func NewToyDb() (*ToyDb, error) {
	home, ok := os.LookupEnv("TOY_HOME")
	if !ok {
		home = ".toy/"
		if _, err := os.Stat(home); os.IsNotExist(err) {
			err := os.Mkdir(home, 0777)
			if err != nil {
				panic(err)
			}
		}
	}

	// catalog, err := storage.LoadCatalog(home)
	// if err != nil {
	// 	return nil, err
	// }

	return &ToyDb{
		// catalog:     catalog,
		storage: storage.NewStorage(home),
		// tranManager: storage.NewTransactionManager(),
		// contexts:    make(map[string]*dbSession),
		exit: make(chan int, 1),
		home: home,
	}, nil
}
