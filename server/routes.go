package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes(r *gin.Engine) {
	// cloud
	r.GET("/api/manifest", s.GetManifest)
	r.GET("/api/brands", s.GetBrandList)
	r.GET("/api/dns", s.GetDNSList)
	r.GET("/api/websites", s.GetWebsiteList)
	r.GET("/api/ip/:ips", s.GetIPInfo)

	// diagnose
	r.GET("/api/ping", s.DiagnosePingHandler)
}
