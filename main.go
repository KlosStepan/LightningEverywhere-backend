package main

import (
	"log"
	"net/http"
	"os"

	"lightningeverywhere_backend/internal/eshop"
)

func main() {
	// Choose store implementation — memory for now
	store := eshop.NewMemoryStore()

	// Setup routes
	mux := http.NewServeMux()
	mux.Handle("/api/eshops", eshop.MakeHandler(store)) // all methods (GET, POST, etc.)

	// Optional: Middleware (basic logger)
	//handler := loggingMiddleware(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s", port)
	//err := http.ListenAndServe(":"+port, handler)
	//if err != nil {
	//	log.Fatal("Server failed:", err)
	//}
}
