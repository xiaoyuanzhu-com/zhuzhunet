package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetManifest(c *gin.Context) {
	manifest, err := s.cloud.GetManifest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, manifest)
}

func (s *Server) GetBrands(c *gin.Context) {
	brands, err := s.cloud.GetBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/manifest", s.GetManifest)
	r.GET("/api/brands", s.GetBrands)
}
