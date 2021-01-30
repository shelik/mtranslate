package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelik/mtranslate/app"
)

// Handler ...
type Handler struct {
	uc app.Usecase
}

// NewHandler ...
func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// HelloWorld ...
func (h *Handler) Translate(c *gin.Context) {
	text := c.Request.URL.Query().Get("text")
	fmt.Println(text)

	tranText := h.uc.Translate(c.Request.Context(), "en", "uk", "Hello world")

	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
		"text":   tranText,
	})
}
