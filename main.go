package main

import (
	"flag"
	"fmt"
	"golangSampleCRUD/api"
	"golangSampleCRUD/storage"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "Listen address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Printf("server running on port %s...\n", *listenAddr)
	log.Fatal(server.Start())
}
