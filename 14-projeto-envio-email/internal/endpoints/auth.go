package endpoints

import (
	"context"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            render.Status(r, 401)
            render.JSON(w, r, map[string]string{"error": "Not Authorized"})
            return
        }

        tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
        provider, err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOAK_URL"))
        if err != nil {
            render.Status(r, 500)
            render.JSON(w, r, map[string]string{"error": "error to connect to provider"})
        }

        verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
        _, err = verifier.Verify(r.Context(), tokenStr)

        if err != nil {
            render.Status(r, 401)
            render.JSON(w, r, map[string]string{"error": "Not Authorized"})
            return
        }

        token, _ := jwtgo.Parse(tokenStr, nil)
        claims := token.Claims.(jwtgo.MapClaims)
        email := claims["email"]

        ctx := context.WithValue(r.Context(), "email", email)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
