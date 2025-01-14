package endpoints

import (
	"context"
	"emailn/internal/infrastructure/credentials"
	"net/http"

	"github.com/go-chi/render"
)

type ValidateTokenFunc func(token string, ctx context.Context) (string, error)

var ValidateToken ValidateTokenFunc = credentials.ValidateToken

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "Not Authorized"})
			return
		}

		email, err := ValidateToken(tokenStr, r.Context())

		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "Not Authorized"})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
