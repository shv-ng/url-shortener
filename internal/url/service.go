package url

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var URLDB = make(map[string]Url)

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
	id := generateShortURL(original_url)
	url := Url{
		Id:          id,
		OriginalURL: original_url,
		ShortURL:    id,
		CreatedAt:   time.Now(),
	}
	URLDB[url.Id] = url
	return url
}

func getURL(short_url string) (Url, error) {
	url, ok := URLDB[short_url]
	if !ok {
		return Url{}, errors.New("URL not found")
	}
	return url, nil
}
