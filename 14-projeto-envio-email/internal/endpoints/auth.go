package endpoints

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
    oidc "github.com/coreos/go-oidc/v3/oidc"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            render.Status(r, 401)
            render.JSON(w, r, map[string]string{"error": "Not Authorized"})
            return
        }

        token = strings.Replace(token, "Bearer ", "", 1)
        provider, err := oidc.NewProvider(r.Context(), "http://localhost:8081/realms/provider")
        if err != nil {
            render.Status(r, 500)
            render.JSON(w, r, map[string]string{"error": "error to connect to provider"})
        }

        verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
        _, err = verifier.Verify(r.Context(), token)

        if err != nil {
            render.Status(r, 401)
            render.JSON(w, r, map[string]string{"error": "Not Authorized"})
            return
        }

        next.ServeHTTP(w, r)
    })
}
