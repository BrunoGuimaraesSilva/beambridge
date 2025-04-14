package http

import (
	"fmt"
	"net/http"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/progress"
)

type ProgressHandler struct {
	usecase *progress.Progress
}

func NewProgressHandler(usecase *progress.Progress) *ProgressHandler {
	return &ProgressHandler{usecase: usecase}
}

func (h *ProgressHandler) StreamProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	uploadID := r.URL.Query().Get("uploadId")
	if uploadID == "" {
		domain.BadRequestError("Missing uploadId", "Must be a valid upload ID")
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		domain.InternalError("Streaming not supported", "Server does not support streaming")
		return
	}

	err := h.usecase.StreamProgress(r.Context(), uploadID, func(percent int) error {
		_, err := fmt.Fprintf(w, "data: {\"percent\": %d}\n\n", percent)
		flusher.Flush()
		return err
	})

	if err != nil {
		fmt.Fprintf(w, "data: {\"error\": %q}\n\n", err.Error())
		flusher.Flush()
	}
}
