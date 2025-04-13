package main

import (
	"log"
	"os"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello, World!")
	//})

	//http.ListenAndServe(":8080", nil)

	// Setup MongoDB
	//store := eshop.NewMongoStore("mongodb://localhost:27017", "yourdbname", "eshops")

	// Setup routes
	//mux := http.NewServeMux()
	//mux.HandleFunc("/api/eshops", eshop.MakeHandler(store))

	// Optional: Wrap with middleware
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
