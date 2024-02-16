package file

import (
	"io"
	"os"
	"path"
	"strings"
	"sync"
)

type Manager struct {
	dbDirectory string
	blockSize   int
	isNew       bool
	openFiles   map[string]*os.File
	sync.Mutex
}

func NewManager(dbDirectory string, blockSize int) *Manager {
	_, err := os.Stat(dbDirectory)
	isNew := os.IsNotExist(err)

	if isNew {
		err := os.MkdirAll(dbDirectory, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	entries, err := os.ReadDir(dbDirectory)
	if err != nil {
		panic(err)
	}

	for _, v := range entries {
		if strings.HasPrefix(v.Name(), "tmp") {
			os.Remove(v.Name())
		}
	}

	return &Manager{
		dbDirectory: dbDirectory,
		blockSize:   blockSize,
		isNew:       isNew,
		openFiles:   make(map[string]*os.File),
	}
}

func (manager *Manager) BlockSize() int {
	return manager.blockSize
}

func (manager *Manager) Read(blockID BlockID, page *Page) {
	manager.Lock()
	defer manager.Unlock()

	f := manager.getFile(blockID.FileName())

	if _, err := f.ReadAt(page.contents(), int64(blockID.BlkNum())*int64(manager.blockSize)); err != io.EOF && err != nil {
		panic(err)
	}
}

func (manager *Manager) Write(blockID BlockID, p *Page) {
	manager.Lock()
	defer manager.Unlock()

	f := manager.getFile(blockID.fileName)
	f.WriteAt(p.contents(), int64(blockID.blkNum)*int64(manager.blockSize))
}

func (manager *Manager) getFile(fileName string) *os.File {
	f, ok := manager.openFiles[fileName]
	if !ok {
		p := path.Join(manager.dbDirectory, fileName)
		table, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR|os.O_SYNC, 0755)
		if err != nil {
			panic(err)
		}
		manager.openFiles[fileName] = table
		return table
	}
	return f
}

func (manager *Manager) Size(fileName string) int {
	f := manager.getFile(fileName)
	finfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return int(finfo.Size() / int64(manager.blockSize))
}

func (manager *Manager) Append(fileName string) BlockID {

	newBlkNum := manager.Size(fileName)
	block := NewBlockID(fileName, newBlkNum)
	buf := make([]byte, manager.blockSize)

	f := manager.getFile(fileName)
	f.WriteAt(buf, int64(block.blkNum)*int64(manager.blockSize))
	return block
}
