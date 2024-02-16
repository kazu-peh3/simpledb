package file

import (
	"encoding/binary"
	"fmt"
)

const IntBytes = int(8)

type Page struct {
	byteBuffer []byte
	maxSize    int
}

func NewPageWithSize(size int) *Page {
	return &Page{
		byteBuffer: make([]byte, size),
		maxSize:    size,
	}
}

func NewPageWithByteSlice(byteSlice []byte) *Page {
	return &Page{
		byteBuffer: byteSlice,
		maxSize:    len(byteSlice),
	}
}

func (page *Page) contents() []byte {
	return page.byteBuffer
}

func (page *Page) SetString(offset int, s string) {
	page.SetBytes(offset, []byte(s))
}

func (page *Page) GetString(offset int) string {
	buf := page.GetBytes(offset)
	return string(buf)
}

func (page *Page) SetBytes(offset int, data []byte) {
	if offset+len(data) > page.maxSize {
		panic(fmt.Sprintf("data out of page bounds. offset: %d length: %d. Max page size is %d", offset, IntBytes, page.maxSize))
	}

	copy(page.byteBuffer[offset:], intToBytes(int64(len(data))))
	copy(page.byteBuffer[offset+IntBytes:], data)
}

func (page *Page) GetBytes(offset int) []byte {
	size := bytesToInt(page.byteBuffer[offset : offset+IntBytes])
	from := offset + IntBytes
	to := offset + IntBytes + int(size)
	return page.byteBuffer[from:to]
}

func (page *Page) SetInt(offset int, n int) {
	if offset+IntBytes > page.maxSize {
		panic(fmt.Sprintf("data out of page bounds. offset: %d length: %d. Max page size is %d", offset, IntBytes, page.maxSize))
	}

	lb := intToBytes(int64(n))
	copy(page.byteBuffer[offset:], lb[:])
}

func (page *Page) GetInt(offset int) int {
	v := bytesToInt(page.byteBuffer[offset : offset+IntBytes])
	return int(v)
}

func intToBytes(v int64) []byte {
	buf := make([]byte, IntBytes)
	binary.LittleEndian.PutUint64(buf, uint64(v))
	return buf
}

func bytesToInt(b []byte) int64 {
	v := binary.LittleEndian.Uint64(b)
	return int64(v)
}

func MaxLength(strlen int) int {
	return strlen + IntBytes
}

func (page *Page) ByteBuffer() []byte {
	return page.byteBuffer
}
