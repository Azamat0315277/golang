package main

import (
	"custom-db/driver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const Version = "1.0.2"

// App holds the database driver dependency for our handlers
type App struct {
	db *driver.Driver
}

func main() {
	// Initialize the database driver
	db, err := driver.New("./", nil)
	if err != nil {
		log.Fatalf("Error creating database driver: %v", err)
	}

	// Create a new App instance with the database driver
	app := &App{db: db}
	port := "8080"

	fmt.Printf("Starting database server on port %s\n", port)
	fmt.Println("You can now send requests to http://localhost:" + port)

	// Register the master handler and start the server
	http.HandleFunc("/", app.handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handler routes all incoming requests based on method and URL path
func (a *App) handler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path to get collection and resource
	// Example: /users/jane -> ["", "users", "jane"]
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Malformed URL. Use /collection or /collection/resource.", http.StatusBadRequest)
		return
	}

	collection := parts[0]
	resource := ""
	if len(parts) > 1 {
		resource = parts[1]
	}

	// Route based on HTTP method
	switch r.Method {
	case http.MethodGet:
		if resource == "" {
			a.handleReadAll(w, r, collection)
		} else {
			a.handleRead(w, r, collection, resource)
		}
	case http.MethodPost:
		if resource == "" {
			http.Error(w, "Missing resource name in URL. Use POST /collection/resource.", http.StatusBadRequest)
			return
		}
		a.handleWrite(w, r, collection, resource)
	case http.MethodDelete:
		a.handleDelete(w, r, collection, resource)
	default:
		w.Header().Set("Allow", "GET, POST, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleWrite creates a new record
func (a *App) handleWrite(w http.ResponseWriter, r *http.Request, collection, resource string) {
	var record interface{}
	// Decode the JSON from the request body
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Write the record to the database
	if err := a.db.Write(collection, resource, record); err != nil {
		http.Error(w, "Failed to write record: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully wrote '%s' to '%s'\n", resource, collection)
}

// handleRead retrieves a single record
func (a *App) handleRead(w http.ResponseWriter, r *http.Request, collection, resource string) {
	var record interface{}
	if err := a.db.Read(collection, resource, &record); err != nil {
		http.Error(w, "Failed to read record: "+err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// handleReadAll retrieves all records in a collection
func (a *App) handleReadAll(w http.ResponseWriter, r *http.Request, collection string) {
	records, err := a.db.ReadAll(collection)
	if err != nil {
		http.Error(w, "Failed to read collection: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// The records are already JSON strings, so we wrap them in a JSON array
	var result []json.RawMessage
	for _, rec := range records {
		result = append(result, json.RawMessage(rec))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// handleDelete removes a record or an entire collection
func (a *App) handleDelete(w http.ResponseWriter, r *http.Request, collection, resource string) {
	if err := a.db.Delete(collection, resource); err != nil {
		http.Error(w, "Failed to delete: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if resource == "" {
		fmt.Fprintf(w, "Successfully deleted collection '%s'\n", collection)
	} else {
		fmt.Fprintf(w, "Successfully deleted record '%s' from '%s'\n", resource, collection)
	}
}
