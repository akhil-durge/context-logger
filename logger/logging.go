package logger

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

//Info logs with context
func Info(ctx context.Context, msg string, v ...interface{}) {
	ctxString := fmt.Sprintf("[%v:%v:%v] %v", ReqID(ctx), ClientID(ctx), UserID(ctx), caller())
	log.Printf("%v INFO: %v", ctxString, fmt.Sprintf(msg, v...))
}

//Error logs with context
func Error(ctx context.Context, msg string, v ...interface{}) {
	ctxString := fmt.Sprintf("[%v:%v:%v] %v", ReqID(ctx), ClientID(ctx), UserID(ctx), caller())
	log.Printf("%v ERROR: %v", ctxString, fmt.Sprintf(msg, v...))
}

//Debug logs with context
func Debug(ctx context.Context, msg string, v ...interface{}) {
	ctxString := fmt.Sprintf("[%v:%v:%v] %v", ReqID(ctx), ClientID(ctx), UserID(ctx), caller())
	log.Printf("%v DEBUG: %v", ctxString, fmt.Sprintf(msg, v...))
}

//to give you name of the file and line number from where logs are emitted
func caller() string {
	_, file, no, ok := runtime.Caller(2)
	if ok {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return fmt.Sprintf("%v:%v ", file, no)
	}
	return "???:0 "
}
