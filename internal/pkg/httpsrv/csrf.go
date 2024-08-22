package httpsrv

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

// GenerateCSRFToken creates a random CSRF token
func GenerateCSRFToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CSRFMiddleware validates CSRF tokens for state-changing requests
func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			csrfToken := r.Header.Get("X-CSRF-Token")
			cookie, err := r.Cookie("csrf_token")
			if err != nil || csrfToken != cookie.Value {
				http.Error(w, "Invalid CSRF token", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// SetCSRFCookie sets the CSRF token in a secure cookie
func SetCSRFCookie(w http.ResponseWriter, csrfToken string) {
	cookie := &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Assuming HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)
}
