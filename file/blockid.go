package file

import "fmt"

const EOF = -1

type BlockID struct {
	fileName string
	blkNum   int
}

func NewBlockID(fileName string, blkNum int) BlockID {
	return BlockID{
		fileName: fileName,
		blkNum:   blkNum,
	}
}

func (bid BlockID) FileName() string {
	return bid.fileName
}

func (bid BlockID) BlkNum() int {
	return bid.blkNum
}

func (bid BlockID) Equals(obj BlockID) bool {
	return bid.fileName == obj.fileName && bid.blkNum == obj.blkNum
}

func (bid BlockID) ToString() string {
	return fmt.Sprintf("[file %s, block %d]", bid.fileName, bid.BlkNum())
}
