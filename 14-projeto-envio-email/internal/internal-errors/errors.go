package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("Internal Server Error")

func ProccessErrorToReturn (err error) error {
    if err == nil {
        return nil
    }

    if !errors.Is(err, gorm.ErrRecordNotFound) {
        return ErrInternal
    }

    return err
}
