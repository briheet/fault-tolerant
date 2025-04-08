package server

import (
	"net/http"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Server struct {
	Addr   string
	Logger *zap.Logger
	Redis  *redis.Client
}

func NewServer(port string, logger *zap.Logger, redis *redis.Client) *Server {
	return &Server{
		Addr:   port,
		Logger: logger,
		Redis:  redis,
	}
}

func (s *Server) Serve(mux *http.ServeMux) {

	mux.HandleFunc("GET /health", s.Health)

}
