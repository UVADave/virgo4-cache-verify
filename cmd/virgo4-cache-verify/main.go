package main

import (
	"log"
)

//
// main entry point
//
func main() {

	// Get config params and use them to init service context. Any issues are fatal
	cfg := LoadConfiguration()

	cacheProxy, err := NewCacheProxy(cfg)
	fatalIfError(err)

	// create a new loader
	loader, e := NewRecordLoader( cfg.InputFile )
	fatalIfError(e)

	// validate the file and ensure each item appears in the cache
	_ = loader.Validate(cacheProxy)
	loader.Done()

	log.Printf("INFO: terminating normally")
}

//
// end of file
//
