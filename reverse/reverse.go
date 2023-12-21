package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:9090")
	rp := httputil.NewSingleHostReverseProxy(u)
	http.ListenAndServe(":9999", rp)
}
