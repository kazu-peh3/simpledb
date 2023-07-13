package server

import (
	"github.com/kazu-peh3/toydb/file"
)

const (
	blockSize  = 400
	bufferSize = 8
	logFile    = "toydb.log"
)

type ToyDB struct {
	Fm file.Manager
}

func NewToyDB(dirname string, blockSize int, bufferSize int) *ToyDB {
	return &ToyDB{
		Fm: *file.NewManager(dirname, blockSize),
	}
}
