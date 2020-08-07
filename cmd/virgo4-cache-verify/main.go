package main

import (
	"log"
	"os"
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
	loader, err := NewRecordLoader(cfg.InputFile)
	fatalIfError(err)

	// validate the file and ensure each item appears in the cache
	err = loader.Validate(cacheProxy)
	loader.Done()

	if err != nil {
		log.Printf("INFO: terminating with error: %s", err.Error())
		os.Exit(1)
	} else {
		log.Printf("INFO: terminating normally")
		os.Exit(0)
	}
}

//
// end of file
//
