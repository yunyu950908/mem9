package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/qiffang/mnemos/server/internal/domain"
	"github.com/qiffang/mnemos/server/internal/service"
)

type ingestRequest struct {
	Messages  []service.IngestMessage `json:"messages"`
	SessionID string                  `json:"session_id"`
	AgentID   string                  `json:"agent_id"`
	Mode      service.IngestMode      `json:"mode,omitempty"`
}

func (s *Server) ingestMemories(w http.ResponseWriter, r *http.Request) {
	var req ingestRequest
	if err := decode(r, &req); err != nil {
		s.handleError(w, err)
		return
	}
	if len(req.Messages) == 0 {
		s.handleError(w, &domain.ValidationError{Field: "messages", Message: "required"})
		return
	}
	auth := authInfo(r)
	svc := s.resolveServices(auth)
	agentID := req.AgentID
	if agentID == "" {
		agentID = auth.AgentName
	}
	ingestReq := service.IngestRequest{
		Messages:  req.Messages,
		SessionID: req.SessionID,
		AgentID:   agentID,
		Mode:      req.Mode,
	}
	go func() {
		result, err := svc.ingest.Ingest(context.Background(), auth.AgentName, ingestReq)
		if err != nil {
			slog.Error("async ingest failed", "agent", agentID, "session", req.SessionID, "err", err)
			return
		}
		slog.Info("async ingest complete", "agent", agentID, "session", req.SessionID,
			"status", result.Status, "memories_changed", result.MemoriesChanged)
	}()
	respond(w, http.StatusAccepted, map[string]string{"status": "accepted"})
}
