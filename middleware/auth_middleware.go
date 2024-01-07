package middleware

// import (
// 	"net/http"
// )

// // OAuth2Middleware is a middleware that handles OAuth2 authentication.
// func OAuth2Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Perform OAuth2 authentication logic here

// 		// Example: Check if the user is authenticated
// 		if !isAuthenticated(r) {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Call the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }

// // isAuthenticated checks if the user is authenticated.
// func isAuthenticated(r *http.Request) bool {
// 	// Implement your authentication logic here
// 	// Example: Check if the user has a valid access token
// 	accessToken := r.Header.Get("Authorization")
// 	// Perform token validation logic here

// 	return true // Return true if the user is authenticated, false otherwise
// }
