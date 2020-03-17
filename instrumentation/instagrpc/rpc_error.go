package instagrpc

import (
	ot "github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
)

func addRPCError(sp ot.Span, err interface{}) {
	var (
		logField   otlog.Field
		errMessage interface{}
	)

	switch err := err.(type) {
	case error:
		logField = otlog.Error(err)
		errMessage = err.Error()
	default:
		logField = otlog.Object("error", err)
		errMessage = err
	}

	sp.SetTag("rpc.error", errMessage)
	sp.SetTag("message", errMessage)
	sp.LogFields(logField)
}
