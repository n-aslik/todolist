package repository

import (
	"Todolist/errs"
	"errors"

	"gorm.io/gorm"
)

func translateErrors(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}
	return err
}
