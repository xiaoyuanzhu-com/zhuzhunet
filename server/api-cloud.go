package server

import (
	"net/http"
	"strings"

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

func (s *Server) GetWebsiteList(c *gin.Context) {
	websiteList, err := s.cloud.GetWebsiteList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, websiteList)
}

func (s *Server) GetIPInfo(c *gin.Context) {
	ips := c.Param("ips")
	ipInfo, err := s.cloud.GetIPInfo(strings.Split(ips, ","))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ipInfo)
}
