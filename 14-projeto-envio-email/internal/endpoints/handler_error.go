package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        obj, status, err := endpointFunc(w, r)
        if err != nil {
            err = internalerrors.ProccessErrorToReturn(err)
            render.Status(r, http.StatusBadRequest)

            if errors.Is(err, internalerrors.ErrInternal) {
                render.Status(r, http.StatusInternalServerError)
            } else if errors.Is(err, gorm.ErrRecordNotFound) {
                render.Status(r, http.StatusNotFound)
            }

            render.JSON(w, r, map[string]string{"error": err.Error()})
            return
        }

        if status == http.StatusNoContent {
            render.NoContent(w, r)
            return
        }

        render.Status(r, status)

        if obj != nil {
            render.JSON(w, r, obj)
        }
    })
}
