package apicontext

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type contextKey int

const (
	reqIDKey    contextKey = 0
	clientIDKey contextKey = 1
	userIDKey   contextKey = 2
)

// WithReqID returns a new Context that carries request ID
func WithReqID(ctx context.Context) context.Context {
	return context.WithValue(ctx, reqIDKey, uuid.NewV4().String())
}

// WithClientID returns a new Context that carries client ID
func WithClientID(ctx context.Context, clientID uint) context.Context {
	return context.WithValue(ctx, clientIDKey, clientID)
}

// WithUserID returns a new Context that carries user ID
func WithUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

//ReqID returns the requestID from context.
//Returns 0 if requestID is not set in context
func ReqID(ctx context.Context) string {
	reqID := ctx.Value(reqIDKey)
	if reqID != nil {
		return reqID.(string)
	}
	return ""
}

//ClientID returns the client_id of authorized client from context.
//Returns 0 if client_id is not set in context
func ClientID(ctx context.Context) uint {
	clientID := ctx.Value(clientIDKey)
	if clientID != nil {
		return clientID.(uint)
	}
	return 0
}

//UserID returns user_id of authenticated user of a client from context.
//Returns 0 if user_id is not set in context
func UserID(ctx context.Context) uint {
	userID := ctx.Value(userIDKey)
	if userID != nil {
		return userID.(uint)
	}
	return 0
}
