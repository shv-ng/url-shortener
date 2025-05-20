package main

import (
	"fmt"
	"net/http"

	"url-shortner/internal/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			url.RootUrlHandler(w, r)
		case "/shorturl":
			url.ShortUrlHandler(w, r)
		default:
			url.RedirectUrl(w, r)
		}
	})

	fmt.Println("Server starts at 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
}
