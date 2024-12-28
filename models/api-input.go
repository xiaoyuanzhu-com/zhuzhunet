package models

type APIInput struct {
	ClientIP string              `json:"client_ip"`
	Method   string              `json:"method"`
	Path     string              `json:"path"`
	Headers  map[string][]string `json:"headers"`
	Body     string              `json:"body"`
	Query    map[string][]string `json:"query"`
}
