package app

import "context"

// Usecase ...
type Usecase interface {
	Translate(ctx context.Context, srcLang string, desLang string, text string) string
}
