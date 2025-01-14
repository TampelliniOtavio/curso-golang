package endpoints

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Auth_WhenAuthorizationIsMissing_ReturnError(t *testing.T) {
	assert := assert.New(t)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")
	})

	handlerFunc := Auth(handler)
	req, _ := http.NewRequest("GET", "/", nil)

	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "Not Authorized")
}

func Test_Auth_WhenAuthorizationIsInvalid_ReturnError(t *testing.T) {
	assert := assert.New(t)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")
	})

	handlerFunc := Auth(handler)

	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return "", errors.New("Invalid Token")
	}

	req, _ := http.NewRequest("GET", "/", nil)

	req.Header.Add("Authorization", "Bearer invalid-token")

	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "Not Authorized")
}

func Test_Auth_WhenAuthorizationIsValid_CallNextHandler(t *testing.T) {
	assert := assert.New(t)
	validEmail := "email@teste.com"
	var email string
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email = r.Context().Value("email").(string)
	})

	handlerFunc := Auth(handler)

	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return validEmail, nil
	}

	req, _ := http.NewRequest("GET", "/", nil)

	req.Header.Add("Authorization", "Bearer valid-token")

	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Equal(email, validEmail)
}
