package logging

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ContextKey defines the context key type.
type ContextKey string

// ContextIDKey holds the key of the context ID.
const ContextIDKey ContextKey = "ctx_id"

// UnaryServerCtxIDInterceptor adds the ContextIDKey to the context and sets
// it as a log field.
func UnaryServerCtxIDInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// generate unique id
	ctxID, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "new uuid error")
	}

	// set id to context and add as logrus field
	ctx = context.WithValue(ctx, ContextIDKey, ctxID)
	ctxlogrus.AddFields(ctx, log.Fields{
		"ctx_id": ctxID,
	})

	// set id as response header
	header := metadata.Pairs("ctx-id", ctxID.String())
	grpc.SendHeader(ctx, header)

	// execute the handler
	return handler(ctx, req)
}
