package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"getting-married/gen/guest/v1/guestv1connect"
	"getting-married/guest"
	"log"
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	port              = flag.Int("port", 8080, "The port at the server listens.")
	dbPath            = flag.String("db_path", "getting_married.db", "Path to the BoltDB database file.")
	adminPassword     = flag.String("admin_password", "", "Password to access the admin interface.")
	allowedOriginsRaw = flag.String("allowed_origins", "http://localhost:*", "Comma-separated list of allowed origins for CORS.")
)

type loginRequest struct {
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
}

func handleLogin(w http.ResponseWriter, req *http.Request) {
	loginReq := &loginRequest{}
	if err := json.NewDecoder(req.Body).Decode(loginReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if loginReq.Password != *adminPassword {
		http.Error(w, "login failed", http.StatusUnauthorized)
		return
	}
	loginResp := &loginResponse{
		AccessToken: loginReq.Password,
	}
	if err := json.NewEncoder(w).Encode(loginResp); err != nil {
		http.Error(w, "encoding response failed", http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()

	if *adminPassword == "" {
		log.Fatal("--admin_password must be specified")
	}

	boltDB, err := bolt.Open(*dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	server, err := guest.NewServer(boltDB, *adminPassword)
	if err != nil {
		log.Fatalf("failed to create guest server: %v", err)
	}
	path, handler := guestv1connect.NewGuestServiceHandler(server)

	mux := http.NewServeMux()
	mux.Handle(path, handler)
	mux.HandleFunc("POST /login", handleLogin)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("listening on %s", addr)

	h2cHandler := h2c.NewHandler(mux, &http2.Server{})

	var allowedOrigins []string
	for _, allowedOrigin := range strings.Split(*allowedOriginsRaw, ",") {
		allowedOrigins = append(allowedOrigins, strings.TrimSpace(allowedOrigin)))
	}

	c := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedHeaders: []string{"connect-protocol-version", "content-type", "authorization"},
	})
	corsHandler := c.Handler(h2cHandler)

	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
