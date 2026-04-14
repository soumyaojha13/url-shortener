package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLUsecase interface {
	Create(string) (string, error)
	Get(string) (string, error)
}

type Handler struct {
	u URLUsecase
}

func NewHandler(u URLUsecase) *Handler {
	return &Handler{u: u}
}

func (h *Handler) CreateURL(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}
	c.BindJSON(&req)

	code, _ := h.u.Create(req.URL)

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + code,
	})
}

func (h *Handler) Redirect(c *gin.Context) {
	code := c.Param("code")

	url, _ := h.u.Get(code)

	c.Redirect(http.StatusFound, url)
}
