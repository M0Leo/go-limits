package main

import (
	"fmt"
	"go-limits/table"
	"go-limits/window"
	"net/http"
)

var tbl *table.Table
var fixedWindowlimiter *window.FixedWindowLimiter

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}

func bucket(w http.ResponseWriter, r *http.Request) {
	ip := getClientIpAddr(r)
	if tbl.HandleRequest(ip) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Request handled successfully.")
	} else {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, "Request declined.")
	}
	fmt.Fprintf(w, "limited %s", ip)
}

func fixedWindow(w http.ResponseWriter, r *http.Request) {
	ip := getClientIpAddr(r)
	if fixedWindowlimiter.HandleRequest(ip) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Request handled successfully.")
	} else {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, "Request declined.")
	}
	fmt.Fprintf(w, "limited %s", ip)
}

func main() {
	//tbl = table.NewTable()
	fixedWindowlimiter = window.NewFixedWindowLimiter(5, 10)
	http.HandleFunc("/limitedBucket", bucket)
	http.HandleFunc("/limitedFixedWindow", fixedWindow)


	http.ListenAndServe(":8080", nil)
}
