package main

import (
	"context"
	"log-context/apicontext"
	"log-context/logger"
)

func main() {
	ctx := apicontext.WithReqID(context.Background())
	logger.Info(ctx, "test info log formatted with var %v, var %v ", 1, "two")
	logger.Error(ctx, "test error log formatted with var %v, var %v ", 1, "two")
	logger.Debug(ctx, "test debug log formatted with var %v, var %v ", 1, "two")
}
