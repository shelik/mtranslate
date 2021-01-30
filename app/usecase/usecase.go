package usecase

import (
	"context"

	"github.com/shelik/mtranslate/app"
)

// Usecase ...
type Usecase struct {
	repo app.Repository
}

// NewUsecase ...
func NewUsecase(repo app.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// Translate ...
func (u *Usecase) Translate(ctx context.Context, srcLang string, desLang string, text string) string {
	println("Hello")
	return u.repo.Translate(srcLang, desLang, text)
}
