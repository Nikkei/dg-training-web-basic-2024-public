package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	address := fmt.Sprintf(":%s", port)
	proxy := goproxy.NewProxyHttpServer()
	// proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			fmt.Println("Request to: " + r.URL.String())
			return r, nil
		})

	proxy.OnResponse().DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		if ctx.Req.Method == "CONNECT" {
			return r
		}
		fmt.Println(ctx.Req.Proto)
		if ctx.Req.Proto == "HTTP/1.1" {
			msg := []byte("<p style='font-size:200px;'>Hello! I'm laughingcat hahahaha</p>")
			r.Body = ioutil.NopCloser(bytes.NewBuffer(msg))
			r.Header.Set("content-length", fmt.Sprintf("%d", len(msg)))
		}

		return r
	})

	fmt.Println("Proxy is listening on " + address)
	err := http.ListenAndServe(address, proxy)
	if err != nil {
		log.Fatal(err)
	}
}
