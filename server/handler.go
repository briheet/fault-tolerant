package server

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Running health handler")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "server healthy"}); err != nil {
		s.Logger.Error("Failed to encode the data", zap.Error(err))
		return
	}
}
