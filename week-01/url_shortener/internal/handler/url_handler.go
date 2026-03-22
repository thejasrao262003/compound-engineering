package handler

import (
	"net/http"
	"url_shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

// POST /shorten

func(h *URLHandler) ShortenURL(c* gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err:=c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	code, err := h.service.CreateShortURL(req.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_code": code,
	})
}

// GET /:code
func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	longURL, err := h.service.GetLongURL(code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
	
}