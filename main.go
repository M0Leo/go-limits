package main

import (
	"fmt"
	"go-limits/table"
	"net/http"
)

var tbl *table.Table

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}

func limited(w http.ResponseWriter, r *http.Request) {
	ip := getClientIpAddr(r)
	if tbl.HandleRequest(ip) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Request handled successfully.")
	} else {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, "Request declined. Bucket is empty.")
	}
	fmt.Fprintf(w, "limited %s", ip)
}

func getAll(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/limited", limited)
  http.HandleFunc("/getAll", getAll)
	tbl = table.NewTable()
	http.ListenAndServe(":8080", nil)
}
