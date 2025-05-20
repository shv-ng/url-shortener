package url

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func RootUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root path")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	short_url := r.URL.Path[1:]
	fmt.Println("redirect", short_url)
	original_url, err := getURL(short_url)
	if err != nil {
		return
	}
	fmt.Println(original_url)
	http.Redirect(w, r, original_url.OriginalURL, http.StatusFound)
}
