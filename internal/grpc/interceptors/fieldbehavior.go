package interceptors

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValidateFieldBehaviorUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		callMeta := interceptors.NewServerCallMeta(info.FullMethod, nil, nil)
		if err := validate(callMeta, req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func ValidateFieldBehaviorStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{
			f: func(m any) error {
				callMeta := interceptors.NewServerCallMeta("", info, nil)
				return validate(callMeta, m)
			},
			ServerStream: ss,
		}
		return handler(srv, wrapper)
	}
}

func validate(info interceptors.CallMeta, req any) error {
	var msg protoreflect.ProtoMessage
	var ok bool
	if msg, ok = req.(protoreflect.ProtoMessage); !ok {
		return nil
	}

	if strings.HasPrefix(info.Method, "Create") || strings.HasPrefix(info.Method, "Update") {
		err := fieldbehavior.ValidateRequiredFields(msg)
		if err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
		fieldbehavior.ClearFields(msg, annotations.FieldBehavior_OUTPUT_ONLY)
	}

	return nil
}
