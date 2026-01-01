package main

import (
	"ai-homelab-tools/internal/httpserver"
	"log"
	"os"
)

func main() {
	addr := getEnv("ADDR", ":7070")
	app := httpserver.New()
	log.Printf("ai-homelab-tools listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
