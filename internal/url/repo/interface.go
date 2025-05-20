package repo

// Interface for database
type DBRepository interface {
	Save(original_url, short_url string) error
	GetURL(short_url string) (string, error)
}
