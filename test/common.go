package test

import (
	"os"

	"github.com/kazu-peh3/simpledb/buffer"
	"github.com/kazu-peh3/simpledb/file"
	"github.com/kazu-peh3/simpledb/log"
)

const dbFolder = "../test_data"
const logfile = "testlog"
const blockfile = "testfile"
const blockSize = 200
const buffersAvaialble = 3

type Conf struct {
	DbFolder         string
	LogFile          string
	BlockFile        string
	BlockSize        int
	BuffersAvailable int
}

var DefaultConfig = Conf{
	DbFolder:         dbFolder,
	LogFile:          logfile,
	BlockFile:        blockfile,
	BlockSize:        blockSize,
	BuffersAvailable: buffersAvaialble,
}

func ClearTestFolder() {
	os.RemoveAll(dbFolder)
}

func MakeManagers() (*file.Mgr, *log.LogMgr, *buffer.Manager) {
	fm := file.NewFileMgr(dbFolder, blockSize)
	lm := log.NewLogMgr(fm, logfile)

	bm := buffer.NewBufferManager(fm, lm, buffersAvaialble)

	return fm, lm, bm
}

func MakeManagersWithConfig(conf Conf) (*file.Mgr, *log.LogMgr, *buffer.Manager) {
	fm := file.NewFileMgr(conf.DbFolder, conf.BlockSize)
	lm := log.NewLogMgr(fm, conf.LogFile)

	bm := buffer.NewBufferManager(fm, lm, conf.BuffersAvailable)

	return fm, lm, bm
}
