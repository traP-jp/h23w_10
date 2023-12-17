package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	backend := getEnvOrDefault("BACKEND_URL", "http://localhost:8080")
	frontend := getEnvOrDefault("FRONTEND_URL", "http://localhost:5173")
	reverseProxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			if strings.HasPrefix(pr.In.URL.Path, "/api") {
				log.Println("api")
				backend, err := url.Parse(backend)
				if err != nil {
					panic(err)
				}
				pr.SetURL(backend)
			} else {
				frontend, err := url.Parse(frontend)
				log.Println("frontend")
				if err != nil {
					panic(err)
				}
				pr.SetURL(frontend)
			}
		},
	}
	log.Fatal(http.ListenAndServe(":8000", reverseProxy))
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
