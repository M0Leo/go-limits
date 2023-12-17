package main

import (
	"flag"
	"fmt"
	"net/http"
)

var limiter Limiter
var limiterType = flag.String("limiter", "bucket", "limiter type")

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip := getClientIpAddr(r)
	if limiter.HandleRequest(ip) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Request handled successfully.")
	} else {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, "Request declined.")
	}
	fmt.Fprintf(w, "limited %s", ip)
}

func main() {

	flag.Parse()
	limiter = getLimiter(*limiterType)

	http.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", nil)
}
