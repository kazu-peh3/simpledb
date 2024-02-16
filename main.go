package main

import (
	"github.com/kazu-peh3/simpledb/file"
	_ "github.com/kazu-peh3/simpledb/file"
	"github.com/kazu-peh3/simpledb/log"
	_ "github.com/kazu-peh3/simpledb/log"
)

// func server() {
// 	_, err := file.NewToyDb(
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// var (
// 	serverMode = flag.Bool("server", false, "boot the db server")
// )

func main() {
	fileMgr := file.NewFileMgr("./testdb", 1024)
	logMgr := log.NewLogMgr(fileMgr, "test.log")
	//record := "abc"
	//logMgr.Append([]byte(record))
	//logMgr.Append([]byte(record))
	//logMgr.Flush(5)
	//println(logMgr)
	// fm := simpleDB.FileMgr
	// blk := file.NewBlockID("db", 2)
	// p1 := file.NewPageWithSize(fm.BlockSize())
	// p1.SetString(88, "abcdefghijklm")
	// fm.Write(blk, p1)
	// fm.Read(blk, p1)
	// flag.Parse()

	// if *serverMode {
	// 	server()
	// 	return
	// }
}
