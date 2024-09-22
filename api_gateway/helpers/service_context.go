package helpers

import (
	"api_gateway/utils"
	"context"
	"time"

	"github.com/labstack/echo/v4"
	grpcMetadata "google.golang.org/grpc/metadata"
)

func NewServiceContext() (context.Context, context.CancelFunc, error) {
	token, err := SignJWTForGRPC()
	if err != nil {
		return nil, nil, echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	return ctxWithAuth, cancel, nil
}
