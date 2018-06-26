package main

import (
	"context"
	"log-context/logger"
)

func main() {
	// example to populate context with known request ID
	ctx := logger.WithReqID(context.Background(), "test-req-id-123")
	logger.Info(ctx, "test log formatted with var %v, var %v ", 1, "two")
	doSomethingError(ctx)
	doSomethingDebug(ctx)

	// example to populate context with newly generated request ID which is UUID v4
	ctx2 := logger.WithNewReqID(context.Background())
	logger.Info(ctx2, "test log formatted with var %v, var %v ", 1, "two")

	// example to populate context with request ID, client ID and user ID
	ctx3 := logger.WithNewReqID(context.Background())
	ctx3 = logger.WithClientID(ctx3, 654)
	ctx3 = logger.WithUserID(ctx3, 987)
	logger.Info(ctx3, "test log formatted with var %v, var %v ", 1, "two")
}

func doSomethingError(ctx context.Context) {
	logger.Error(ctx, "test log formatted with var %v, var %v ", 2, "three")
}

func doSomethingDebug(ctx context.Context) {
	logger.Debug(ctx, "test log formatted with var %v, var %v ", "one", 2)
}
