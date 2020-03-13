package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var service Service

type Service struct {
	name string
	Host
}

type Host struct {
	name string
	addr string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	s := &service

	// w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,
		"Hello from behind Envoy (service %v)! hostname: %s "+
			"resolvedhostname: %s\n",
		s.name, s.Host.name, s.addr)
}

func TraceHandler(w http.ResponseWriter, r *http.Request) {

	s := &service

	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(5 * time.Second)}

	callService1From2(s, r, client)

	// w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,
		"Hello from behind Envoy (service %v)! hostname: %s "+
			"resolvedhostname: %s\n",
		s.name, s.Host.name, s.addr)
}

func callService1From2(s *Service, inReq *http.Request, client *http.Client) {
	if s.name == "1" {
		// Call Service 2
		req, _ := http.NewRequest("GET", "http://localhost:9000/trace/2", nil)
		req = setRequestHeaders(inReq, req)
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
	}
}

func setRequestHeaders(inReq *http.Request, req *http.Request) *http.Request {

	TRACE_HEADERS_TO_PROPAGATE := [8]string{
		"X-Ot-Span-Context",
		"X-Request-Id",
		"X-B3-TraceId",
		"X-B3-SpanId",
		"X-B3-ParentSpanId",
		"X-B3-Sampled",
		"X-B3-Flags",
		"uber-trace-id"}

	fmt.Println("Setting Request Headers")
	for _, header := range TRACE_HEADERS_TO_PROPAGATE {
		if val, ok := inReq.Header[header]; ok {
			req.Header.Set(header, val[0])
		}
	}
	req.Header.Set("Connection", "close")
	return req
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/service/{service_number}", HelloHandler)
	r.HandleFunc("/trace/{service_number}", TraceHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {

	// Get the Service name from the Environment variables
	sname := getEnv("SERVICE_NAME", "2")

	// Get the Hostname
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	// Get the Resolved Hostname (IP Address)
	addr, err := net.LookupHost(name)
	if err != nil {
		log.Fatal(err)
	}

	host := Host{name, addr[0]}

	// Set the global service variable
	service = Service{sname, host}
}
