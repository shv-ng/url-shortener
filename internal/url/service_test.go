package url

import (
	"reflect"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	original_url := "https://google.com"
	expected := "99999ebc"
	short_url := generateShortURL(original_url)
	if !reflect.DeepEqual(short_url, expected) {
		t.Errorf("Expected %v, got %v", expected, short_url)
	}
}
