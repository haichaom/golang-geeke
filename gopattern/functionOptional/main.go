package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) WithMaxConns(maxconns int) *ServerBuilder {
	sb.Server.Maxconns = maxconns
	return sb
}

func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server.TLS = tls
	return sb
}

func (sb *ServerBuilder) Build() Server {
	return sb.Server
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.Maxconns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		Maxconns: 100,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func main() {
	sb := ServerBuilder{}
	server := sb.Create("127.0.0.1", 8080).WithProtocol("udp").WithMaxConns(1024).WithTimeout(30 * time.Second).Build()
	fmt.Println("server Creation by builder:", server)
	server1, err := NewServer("127.0.0.1", 8080, Timeout(20*time.Second), Protocol("udp"))
	fmt.Println("server creation by Functioal Options:", server1, "error:", err)
}
