package server

import (
	"fmt"
	"time"
)

// Server - structure server related values
type Server struct {
	port string
	timeout time.Duration
	maxConnections int
}

// Option
type Option func(*Server)

// WithTimeout
func WithTimeout(timeout time.Duration) Option {
	fmt.Println("adding timeout to new server")
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithMaxConnections
func WithMaxConnections(maxConn int) Option {
	fmt.Println("adding max connections to new server")
	return func(s *Server) {
		s.maxConnections = maxConn
	}
}

// NewServer - returns a new server with optional options i.e timeout and max connections
func NewServer(port string, options ...Option) *Server {
	fmt.Println("creating new server")
	server := &Server{
		port: port,
	}

	for _, option := range options {
		option(server)
	}

	return server
}
