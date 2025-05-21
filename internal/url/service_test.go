package url

import (
	"reflect"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	original_url := "https://google.com"
	expected := Url{
		OriginalURL: original_url,
		ShortURL:    "99999ebc",
	}
	short_url := generateShortURL(original_url)
	if !reflect.DeepEqual(short_url, expected.ShortURL) {
		t.Errorf("Expected %v, got %v", expected.ShortURL, short_url)
	}
}
