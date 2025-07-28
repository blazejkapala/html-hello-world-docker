package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Pobierz port z zmiennej środowiskowej lub użyj domyślnego
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Middleware do logowania requestów
	loggingHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		})
	}

	// Serwowanie plików statycznych
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", loggingHandler(fs))

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	// Dodatkowy endpoint z informacjami o aplikacji
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"name": "HTML Hello World Docker",
			"version": "1.0.0",
			"description": "Statyczna strona HTML serwowana przez Go w Docker",
			"author": "Docker & Go Example",
			"timestamp": "%s"
		}`, time.Now().Format(time.RFC3339))
	})

	log.Printf("🚀 Serwer uruchamia się na porcie %s", port)
	log.Printf("📱 Otwórz http://localhost:%s w przeglądarce", port)
	log.Printf("🔍 Health check: http://localhost:%s/health", port)
	log.Printf("ℹ️  Info endpoint: http://localhost:%s/info", port)

	// Uruchom serwer
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("❌ Błąd uruchamiania serwera: %v", err)
	}
}
