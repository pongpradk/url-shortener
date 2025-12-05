package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pongpradk/url-shortener/internal/repository"
	"github.com/pongpradk/url-shortener/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

type ShortenRequest struct {
	LongURL string `json:"longUrl" binding:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortUrl"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// HandleShorten creates a short URL
func (h *URLHandler) HandleShorten(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request body",
		})
		return
	}

	shortURL, err := h.service.ShortenURL(c.Request.Context(), req.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to shorten URL",
		})
		return
	}

	c.JSON(http.StatusOK, ShortenResponse{
		ShortURL: shortURL,
	})
}

// HandleRedirect redirects to original URL
func (h *URLHandler) HandleRedirect(c *gin.Context) {
	shortURL := c.Param("shortUrl")

	longURL, err := h.service.GetLongURL(c.Request.Context(), shortURL)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error: "Short URL not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Internal server error",
		})
		return
	}

	// 301 redirect
	c.Redirect(http.StatusMovedPermanently, longURL)
}
