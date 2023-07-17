package service

import "errors"

var (
	ErrQueryDatabase = errors.New("something went wrong in database")
	ErrCopyStruct    = errors.New("something went wrong when copy struct")
	ErrFileIsNotCSV  = errors.New("file is not csv type (accept only csv)")
)
