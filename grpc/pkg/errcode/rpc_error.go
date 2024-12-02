package errcode

import (
	"grpcproj/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(e *Error) error {
	s, _ := status.New(TogRPCCode(e.Code()), e.Msg()).WithDetails(&proto.Error{
		Code:    int32(e.Code()),
		Message: e.Msg(),
	})
	return s.Err()
}

func TogRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}

	return statusCode
}

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func ToRPCStatus(code int, msg string) *Status {
	s, _ := status.New(TogRPCCode(code), msg).WithDetails(&proto.Error{
		Code:    int32(code),
		Message: msg,
	})
	return &Status{s}
}
