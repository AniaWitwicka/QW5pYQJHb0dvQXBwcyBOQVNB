package handlers

import (
	"github.com/valyala/fasthttp"
	"url-collector/internal/operations"
	"url-collector/internal/operations/reader_implementations"
)

type listUrlsRequestWrapper struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type listUrlsResponse struct {
	Urls []string `json:"urls"`
}

func ListUrlsNasa(ctx *fasthttp.RequestCtx) {
	request := listUrlsRequestWrapper{}
	if val, err := GetQueryParamAsString(ctx, "from"); err == nil {
		request.From = val
	} else {
		SetErrorResponse(ctx, 400, err)
		return
	}
	if val, err := GetQueryParamAsString(ctx, "to"); err == nil {
		request.To = val
	} else {
		SetErrorResponse(ctx, 400, err)
		return
	}

	reader := reader_implementations.NewNasaReader()
	urls, err := operations.ListUrls(request.From, request.To, reader)
	if err != nil {
		SetErrorResponse(ctx, 400, err)
		return
	}
	resp := listUrlsResponse{
		Urls: urls,
	}
	SetJsonResponse(ctx, 200, resp)
}
