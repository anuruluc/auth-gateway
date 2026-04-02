package authgateway

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func encodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func decodeBase64(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err!= nil {
		return "", err
	}
	return string(data), nil
}

func extractBearerToken(r *http.Request) (string, bool) {
	bearerToken := r.Header.Get("Authorization")
	if!strings.HasPrefix(bearerToken, "Bearer ") {
		return "", false
	}
	return strings.TrimSpace(strings.TrimPrefix(bearerToken, "Bearer ")), true
}

func extractTokenFromHeader(r *http.Request) (string, bool) {
	token, ok := extractBearerToken(r)
	if!ok {
		return "", false
	}
	return token, true
}

func extractTokenFromQuery(r *http.Request) (string, bool) {
	token := r.URL.Query().Get("token")
	return token, token!= ""
}

func getHeader(r *http.Request, header string) (string, bool) {
	if value := r.Header.Get(header); value!= "" {
		return value, true
	}
	return "", false
}

func getHeaderSlice(r *http.Request, header string) ([]string, bool) {
	values := r.Header.Values(header)
	if len(values) > 0 {
		return values, true
	}
	return nil, false
}

func getCookie(r *http.Request, name string) (string, bool) {
	cookie, err := r.Cookie(name)
	if err!= nil {
		return "", false
	}
	return cookie.Value, true
}

func logRequest(r *http.Request) {
	log.Printf("Request: %s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
}

func logResponse(w http.ResponseWriter, r *http.Request, status int, duration time.Duration) {
	log.Printf("Response: %d %s %s %s\n", status, r.Method, r.URL.Path, duration)
}