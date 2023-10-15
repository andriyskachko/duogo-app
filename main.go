package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	mux := http.NewServeMux()
	userHandler := http.HandlerFunc(UserServer)
	listenAddr := flag.String("listenAddress", os.Getenv("SERVER_PORT"), "the Server address")

	flag.Parse()
	mux.Handle("/users", userHandler)
	log.Printf("Server is running at %v", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, mux))

}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
