package url

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"

	"url-shortner/internal/url/repo"
)

var Repo repo.DBRepository

func generateShortURL(original_url string, args ...int) string {
	length := 8
	if len(args) > 0 {
		length = args[0]
	}
	hasher := md5.New()
	hasher.Write([]byte(original_url))
	return fmt.Sprintf("%x", hasher.Sum(nil))[:length]
}

func createURL(original_url string) Url {
	short_url := generateShortURL(original_url)
	url := Url{
		OriginalURL: original_url,
		ShortURL:    short_url,
	}

	err := Repo.Save(original_url, short_url)
	if err != nil {
		log.Fatalln(err)
	}

	return url
}

func getURL(short_url string) (Url, error) {
	var u Url
	url, err := Repo.GetURL(short_url)
	u = Url{
		url,
		short_url,
	}
	if err != nil {
		return Url{}, errors.New("URL not found")
	}
	return u, nil
}
