package model

import (
	"context"
	"fmt"
)

type (
	contextKeyRequestID struct{}
)

func SetRequestIDToContext(ctx context.Context, reqID string) (context.Context, error) {
	if ctx == nil {
		return nil, fmt.Errorf("parent context must be non-nil")
	}
	return context.WithValue(ctx, contextKeyRequestID{}, reqID), nil
}

func GetRequestIDFromContext(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", fmt.Errorf("parent context must be non-nil")
	}
	v := ctx.Value(contextKeyRequestID{})
	reqID, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("cannot get claimes error")
	}
	return reqID, nil
}
