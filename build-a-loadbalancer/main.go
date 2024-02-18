package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	server          []Server
}

func NewLoadBalancer(port string, server []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		server:          server,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error :%v\n", err)
		os.Exit(1)
	}
}
func (s *simpleServer) Address() string {
	return s.addr
}
func (s *simpleServer) IsAlive() bool {
	return true
}
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}
func (lb *LoadBalancer) getNexAvailableServer() Server {
	server := lb.server[lb.roundRobinCount%len(lb.server)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.server[lb.roundRobinCount%len(lb.server)]
	}
	lb.roundRobinCount++
	return server
}
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNexAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, r)
}
func main() {
	servers := []Server{
		newSimpleServer("https://www.duckduckgo.com"),
		newSimpleServer("https://trendyol.com"),
		newSimpleServer("https://youtube.com"),
	}
	lb := NewLoadBalancer("8001", servers)
	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
