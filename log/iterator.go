package log

import "github.com/kazu-peh3/simpledb/file"

type Iterator struct {
	fileMgr    *file.Mgr
	block      file.BlockID
	page       *file.Page
	currentPos int
	boundary   int
}

// todo: this is a reader in go
func newIterator(fileMgr *file.Mgr, start file.BlockID) *Iterator {
	it := &Iterator{
		fileMgr: fileMgr,
		block:   start,
		page:    file.NewPageWithByteSlice(make([]byte, fileMgr.BlockSize())),
	}

	it.moveToBlock(start)
	return it
}

func (it Iterator) HasNext() bool {
	return it.currentPos < it.fileMgr.BlockSize() || it.block.BlkNum() > 0
}

func (it *Iterator) Next() []byte {
	if it.currentPos == it.fileMgr.BlockSize() {
		// we are at the end of the block, read the previous one
		it.block = file.NewBlockID(it.block.FileName(), it.block.BlkNum()-1)
		it.moveToBlock(it.block)
	}

	record := it.page.GetBytes(it.currentPos)
	// move the iterator position forward by the
	it.currentPos += len(record) + file.IntBytes
	return record
}

func (it *Iterator) moveToBlock(block file.BlockID) {
	it.fileMgr.Read(it.block, it.page)
	it.boundary = it.page.GetInt(0)
	it.currentPos = it.boundary
}
