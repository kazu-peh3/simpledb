package server

import (
	"github.com/kazu-peh3/simpledb/file"
)

type SimpleDB struct {
	FileMgr file.Mgr
}

func NewSimpleDBForDebugging(dirName string, blockSize int, bufferSize int) *SimpleDB {
	return &SimpleDB{
		FileMgr: *file.NewFileMgr(dirName, blockSize),
	}
}
