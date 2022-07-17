package handlers

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"sync"
)

var (
	jsonIterator       = jsoniter.ConfigCompatibleWithStandardLibrary
	responseBufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

type responseBuffer struct {
	buffer *bytes.Buffer
}

func (b *responseBuffer) Read(p []byte) (int, error) {
	return b.buffer.Read(p)
}

func (b *responseBuffer) Close() error {
	b.buffer.Reset()
	responseBufferPool.Put(b.buffer)
	return nil
}

func SetJsonResponse(ctx *fasthttp.RequestCtx, code int, v interface{}) {
	buf := responseBufferPool.Get().(*bytes.Buffer)
	enc := jsonIterator.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		ctx.SetStatusCode(500)
		ctx.SetBodyString(err.Error())
		return
	}

	ctx.SetStatusCode(code)
	ctx.SetContentType("application/json")
	ctx.SetBodyStream(&responseBuffer{buffer: buf}, -1)
}

type errorResponse struct {
	Error string `json:"error"`
}

func SetErrorResponse(ctx *fasthttp.RequestCtx, errCode int, errBody error) {
	SetJsonResponse(ctx, errCode, errorResponse{errBody.Error()})
}

func GetQueryParamAsString(ctx *fasthttp.RequestCtx, paramName string) (string, error) {
	paramString := string(ctx.QueryArgs().Peek(paramName))
	if paramString == "" {
		return "", fmt.Errorf("no param %s", paramName)
	}
	return paramString, nil
}
