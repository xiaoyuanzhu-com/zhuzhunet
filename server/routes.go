package server

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetManifest(c *gin.Context) {
	proxy := httputil.NewSingleHostReverseProxy(s.cloudURL)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = s.cloudURL.Host
		req.URL.Scheme = s.cloudURL.Scheme
		req.URL.Host = s.cloudURL.Host
		req.URL.Path = c.Request.URL.Path
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/manifest", s.GetManifest)
}
