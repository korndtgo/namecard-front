package error_wrapper

import (
	"errors"
	"fmt"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorCode string

type GqlErrorCode int

// GrpcErrorWrapper .
func GrpcErrorWrapper(code codes.Code, message string, err error, metadata map[string]string) error {
	newError := status.Newf(code, message+": %v", err)

	detail := &errdetails.ErrorInfo{
		Reason:   message,
		Domain:   code.String(),
		Metadata: metadata,
	}

	newError, err = newError.WithDetails(detail)
	if err != nil {
		return fmt.Errorf("[GrpcErrorWrapper] add detail error: %v", err)
	}
	return newError.Err()
}

// grpcErrorUnwrapDetail get error from grpc error in form of map[string]string
func GrpcErrorUnwrapDetail(grpcError error) map[string]interface{} {
	errorStatus := status.Convert(grpcError)
	for _, detail := range errorStatus.Details() {
		switch t := detail.(type) {
		case *errdetails.ErrorInfo:
			res := make(map[string]interface{}, len(t.Metadata))
			for k, v := range t.Metadata {
				res[k] = v
			}
			return res
		default:
			fmt.Printf("detail type %T not found \n", t)
		}
	}
	return nil
}

// CreateGrpcErrorDetail .
func CreateGrpcErrorDetail(code ErrorCode, metaData ...map[string]string) map[string]string {
	m := make(map[string]string)
	m["code"] = code.String()
	m["status"] = ErrorCodeStatus[code]

	if len(metaData) > 0 {
		for k, v := range metaData[0] {
			m[k] = v
		}
	}
	return m
}

// GetGrpcErrorCode .
func GetGrpcErrorCode(grpcError error) (string, error) {
	errorStatus := status.Convert(grpcError)
	for _, detail := range errorStatus.Details() {
		switch t := detail.(type) {
		case *errdetails.ErrorInfo:
			res := make(map[string]interface{}, len(t.Metadata))
			for k, v := range t.Metadata {
				res[k] = v
			}
			if res["code"] == nil {
				return "", errors.New("get code error: code not found")
			}
			return res["code"].(string), nil
		default:
			fmt.Printf("detail type %T not found \n", t)
		}
	}
	return "", nil
}

// GQL

// GqlErrorWrapper .
func GqlErrorWrapper(code GqlErrorCode, message string, err error, metadata map[string]interface{}) error {

	gqlerr := &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"technical_message": err.Error(),
		},
	}

	if metadata == nil {
		gqlerr.Extensions["code"] = "500"
		gqlerr.Extensions["status"] = "internal server error"
	}

	for k, v := range metadata {
		gqlerr.Extensions[k] = v
	}

	return gqlerr
}

// GqlErrorWrapperWithGrpcDetail .
func GqlErrorWrapperWithGrpcDetail(message string, err error) error {

	grpcErrDetail := GrpcErrorUnwrapDetail(err)
	gqlErr := &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"technical_message": err.Error(),
		},
	}

	_, ok := grpcErrDetail["code"]
	if !ok {
		gqlErr.Extensions["code"] = "500"
		gqlErr.Extensions["status"] = "INTERNAL ERROR"
	} else {
		metaData := make(map[string]interface{})
		for k, v := range grpcErrDetail {
			if k != "code" && k != "status" {
				metaData[k] = v
			} else {
				gqlErr.Extensions[k] = v
			}
		}
		gqlErr.Extensions["metadata"] = metaData
	}

	return gqlErr
}

func (e ErrorCode) String() string {
	return string(e)
}
