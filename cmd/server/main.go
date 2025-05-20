package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"url-shortner/internal/url"
	"url-shortner/internal/url/repo"

	_ "github.com/lib/pq"
)

// Entry point
func main() {
	// TODO: Move to env
	connStr := "user=urlshortener password=urlshortener dbname=urlshortener host=db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	url.Repo = repo.NewPostgresRepo(db)

	// http handler (router)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/shorturl":
			url.ShortUrlHandler(w, r)
		default:
			url.RedirectUrl(w, r)
		}
	})

	fmt.Println("Server starts at 8000...")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
