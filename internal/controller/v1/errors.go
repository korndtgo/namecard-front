package controller_v1

import "errors"

var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrCopyDataToProto = errors.New("something went wrong when copy data to proto")
)
