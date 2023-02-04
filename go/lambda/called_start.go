package lambda

import (
	"context"
	"encoding/json"
)

var (
	// CurrentContext is the last create lambda context object.
	CurrentContext context.Context
)

type (
	// HandlerListener is a point where listener logic can be injected into a handler
	HandlerListener interface {
		HandlerStarted(ctx context.Context, msg json.RawMessage) context.Context
		HandlerFinished(ctx context.Context, err error)
	}
)

// WrapHandlerWithListeners wraps a lambda handler, and calls listeners before and after every invocation.
// ref: https://github.com/DataDog/datadog-lambda-go/blob/1cdcde5b6f3e4c2d3d237e380fab7f6d9f1c1294/internal/wrapper/wrap_handler.go#L44-L70
func WrapHandlerWithListeners(handler interface{}, listeners ...HandlerListener) interface{} {
	err := validateHandler(handler)
	if err != nil {
		// This wasn't a valid handler function, pass back to AWS SDK to let it handle the error.
		// logger.Error(fmt.Errorf("handler function was in format ddlambda doesn't recognize: %v", err))
		return handler
	}
	// NOTE: 戻り値のクロージャがLambdaのフレームワークで実行されるので、
	// 同じインスタンスの中ならば変数が使い回されるのでコールドスタートが検知出来る。
	coldStart := true

	// Return custom handler, to be called once per invocation
	return func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
		//nolint
		ctx = context.WithValue(ctx, "cold_start", coldStart)
		for _, listener := range listeners {
			ctx = listener.HandlerStarted(ctx, msg)
		}
		CurrentContext = ctx
		result, err := callHandler(ctx, msg, handler)
		for _, listener := range listeners {
			// 			ctx = context.WithValue(ctx, extension.DdLambdaResponse, result)
			listener.HandlerFinished(ctx, err)
		}
		coldStart = false
		CurrentContext = nil
		return result, err
	}
}

func validateHandler(handler interface{}) error {
	return nil

}

func callHandler(ctx context.Context, msg json.RawMessage, handler interface{}) (interface{}, error) {
	return nil, nil
}
