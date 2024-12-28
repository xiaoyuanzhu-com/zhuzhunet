package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/url"
)

func PrintJSON(v interface{}) {
	bs, _ := json.MarshalIndent(v, "", "    ")
	println(string(bs))
}

func GetDomain(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return ""
	}
	return parsedURL.Hostname()
}

func Hash(v interface{}) string {
	bs, _ := json.Marshal(v)
	return fmt.Sprintf("%x", sha256.Sum256(bs))
}
