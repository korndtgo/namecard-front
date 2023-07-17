package repository

import "errors"

var (
	ErrSchemaIsNotFound     = errors.New("schema is not found")
	ErrFundCodeIsNotFound   = errors.New("fundcode is not found")
	ErrImportAlreadyExisted = errors.New("file import already existed")
	ErrDuplicatedName       = errors.New("name is duplicated")
)
