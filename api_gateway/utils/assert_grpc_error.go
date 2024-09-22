package utils

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AssertGrpcStatus(err error) error {
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.Unavailable:
			return echo.NewHTTPError(ErrUnprocessable.EchoFormatDetails(e.Message()))
		case codes.InvalidArgument:
			return echo.NewHTTPError(ErrBadRequest.EchoFormatDetails(e.Message()))
		case codes.FailedPrecondition:
			return echo.NewHTTPError(ErrBadRequest.EchoFormatDetails(e.Message()))
		case codes.NotFound:
			return echo.NewHTTPError(ErrNotFound.EchoFormatDetails(e.Message()))
		case codes.PermissionDenied:
			return echo.NewHTTPError(ErrUnauthorized.EchoFormatDetails(e.Message()))
		case codes.Unauthenticated:
			return echo.NewHTTPError(ErrUnauthorized.EchoFormatDetails(e.Message()))
		case codes.Internal:
			return echo.NewHTTPError(ErrInternalServer.EchoFormatDetails(e.Message()))
		}
	}

	return echo.NewHTTPError(ErrInternalServer.EchoFormatDetails(err.Error()))
}
