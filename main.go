package main

import (
	"getting-married/gen/guest/v1/guestv1connect"
	"getting-married/guest"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	addr   = "0.0.0.0:8080"
	dbFile = "getting_married.db"
)

func main() {
	boltDB, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	server, err := guest.NewServer(boltDB)
	if err != nil {
		log.Fatalf("failed to create guest server: %v", err)
	}
	path, handler := guestv1connect.NewGuestServiceHandler(server)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	fs := http.FileServer(http.Dir("./out"))
	mux.Handle("/", fs)

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
