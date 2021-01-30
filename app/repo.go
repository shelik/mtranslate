package app

// Repository ...
type Repository interface {
	Close() error
	Translate(srcLang string, desLang string, text string) string
}
