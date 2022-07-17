package main

import (
	"flag"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"url-collector/internal/handlers"
)

var (
	Port = flag.String("PORT", "8080", "Port the server is running on")
)

func main() {
	flag.Parse()
	router := fasthttprouter.New()
	router.GET("/pictures", handlers.ListUrlsNasa)

	log.Fatal(fasthttp.ListenAndServe(":"+*Port, router.Handler))
}
