package main

import (
	"flag"
	"fmt"
	"getting-married/gen/guest/v1/guestv1connect"
	"getting-married/guest"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	port   = flag.Int("port", 8080, "The port at the server listens.")
	dbPath = flag.String("db_path", "getting_married.db", "Path to the BoltDB database file.")
)

func main() {
	boltDB, err := bolt.Open(*dbPath, 0600, nil)
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

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
