package middleware

import (
	"forum/app/http/controllers"
	"forum/app/models"
	"net/http"
	"sync"
	"time"
)

var (
	rateLimitMap = make(map[int]time.Time)
	mu           sync.Mutex
)

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		userId, err := models.GetUserId(r)
		if err != nil {
			controllers.ErrorController(w, r, http.StatusInternalServerError, "Error getting user ID")
			return
		}

		// Check if user has recently made a request
		if lastRequest, exists := rateLimitMap[userId]; exists {
			if time.Since(lastRequest) < time.Second {
				controllers.ErrorController(w, r, http.StatusTooManyRequests, "Too many requests. Please wait.")
				return
			}
		}

		// Update last request time
		rateLimitMap[userId] = time.Now()

		next(w, r)
	}
}
