package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoint_return_internal_error(t *testing.T) {
    assert := assert.New(t)
    endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
        return nil, 0, internalerrors.ErrInternal
    }

    handlerFunc := HandlerError(endpoint)
    req, _ := http.NewRequest("GET", "/", nil)

    res := httptest.NewRecorder()

    handlerFunc.ServeHTTP(res, req)

    assert.Equal(http.StatusInternalServerError, res.Code)
    assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
} 

func Test_HandlerError_when_endpoint_return_domain_error(t *testing.T) {
    assert := assert.New(t)
    endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
        return nil, 0, errors.New("Domain Error")
    }

    handlerFunc := HandlerError(endpoint)
    req, _ := http.NewRequest("GET", "/", nil)

    res := httptest.NewRecorder()

    handlerFunc.ServeHTTP(res, req)

    assert.Equal(http.StatusBadRequest, res.Code)
    assert.Contains(res.Body.String(), "Domain Error")
}
