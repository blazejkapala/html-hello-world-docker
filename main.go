package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Request logging middleware
	loggingHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		})
	}

	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", loggingHandler(fs))

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	// Additional endpoint with application information
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"name": "HTML Hello World Docker",
			"version": "1.0.0",
			"description": "Static HTML page served by Go in Docker",
			"author": "Docker & Go Example",
			"timestamp": "%s"
		}`, time.Now().Format(time.RFC3339))
	})

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📱 Open http://localhost:%s in your browser", port)
	log.Printf("🔍 Health check: http://localhost:%s/health", port)
	log.Printf("ℹ️  Info endpoint: http://localhost:%s/info", port)

	// Start server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("❌ Server startup error: %v", err)
	}
}
