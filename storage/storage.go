package storage

type Storage struct {
	buffer *bufferPool
	disk   *diskManager
	prefix string
}

func NewStorage(home string) *Storage {
	return &Storage{
		buffer: newBufferPool(),
		disk:   newDiskManager(),
		prefix: home,
	}
}
