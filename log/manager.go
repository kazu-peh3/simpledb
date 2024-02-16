package log

import (
	"sync"

	"github.com/kazu-peh3/simpledb/file"
)

type LogMgr struct {
	fileMgr      *file.Mgr
	logFile      string
	logPage      *file.Page
	currentBlock file.BlockID
	latestLSN    int
	lastSavedLSN int
	sync.Mutex
}

func NewLogMgr(fileMgr *file.Mgr, logFile string) *LogMgr {
	b := make([]byte, fileMgr.BlockSize())
	logSize := fileMgr.Size(logFile)
	logPage := file.NewPageWithByteSlice(b)

	logMgr := &LogMgr{
		fileMgr:      fileMgr,
		logFile:      logFile,
		logPage:      logPage,
		latestLSN:    0,
		lastSavedLSN: 0,
	}

	if logSize == 0 {
		logMgr.currentBlock = logMgr.appendNewBlock()
	} else {
		logMgr.currentBlock = file.NewBlockID(logFile, logSize-1)
		fileMgr.Read(logMgr.currentBlock, logPage)
	}

	return logMgr
}

func (logMgr *LogMgr) flush() {
	logMgr.fileMgr.Write(logMgr.currentBlock, logMgr.logPage)
	logMgr.lastSavedLSN = logMgr.latestLSN
}

func (logMgr *LogMgr) Flush(lsn int) {
	if lsn >= logMgr.lastSavedLSN {
		logMgr.flush()
	}
}

func (logMgr *LogMgr) Iterator() *Iterator {
	logMgr.flush()
	return newIterator(logMgr.fileMgr, logMgr.currentBlock)
}

func (logMgr *LogMgr) Append(records []byte) int {
	logMgr.Lock()
	defer logMgr.Unlock()

	spaceLeft := logMgr.logPage.GetInt(0)

	recsize := len(records)

	bytesneeded := recsize + file.IntBytes

	if bytesneeded+file.IntBytes > spaceLeft {
		logMgr.flush()
		logMgr.currentBlock = logMgr.appendNewBlock()
		spaceLeft = logMgr.logPage.GetInt(0)
	}

	recpos := spaceLeft - bytesneeded
	logMgr.logPage.SetBytes(recpos, records)
	logMgr.logPage.SetInt(0, recpos) // the new boundary
	logMgr.latestLSN++
	return logMgr.latestLSN

}

func (logMgr *LogMgr) appendNewBlock() file.BlockID {
	block := logMgr.fileMgr.Append(logMgr.logFile)
	logMgr.logPage.SetInt(0, logMgr.fileMgr.BlockSize())

	// write the logpage into the newly created block
	logMgr.fileMgr.Write(block, logMgr.logPage)

	return block
}
