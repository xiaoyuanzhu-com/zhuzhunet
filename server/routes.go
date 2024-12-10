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

func (s *Server) GetBrandList(c *gin.Context) {
	brandList, err := s.cloud.GetBrandList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brandList)
}

func (s *Server) GetDNSList(c *gin.Context) {
	dnsList, err := s.cloud.GetDNSList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dnsList)
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/manifest", s.GetManifest)
	r.GET("/api/brands", s.GetBrandList)
	r.GET("/api/dns", s.GetDNSList)
}
