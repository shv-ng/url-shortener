package url

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(data.URL) == 0 {
		http.Error(w, "URL is empty", http.StatusBadRequest)
		return
	}

	url := createURL(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: url.ShortURL,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // Trim leading slash
	log.Printf("Redirect requested for: %s", shortURL)

	originalURL, err := getURL(shortURL)
	if err != nil {
		log.Printf("URL not found: %v", err)
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL.OriginalURL, http.StatusFound)
}
