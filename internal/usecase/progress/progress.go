package progress

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/port"
)

type Progress struct {
	sessionRepo port.SessionRepository
}

func New(sessionRepo port.SessionRepository) *Progress {
	return &Progress{sessionRepo: sessionRepo}
}

func (p *Progress) StreamProgress(ctx context.Context, uploadID string, send func(int) error) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			percent, err := p.sessionRepo.GetProgress(ctx, uploadID)
			if err != nil {
				return err
			}
			if err := send(percent); err != nil {
				return err
			}
			if percent >= 100 {
				return nil
			}
		}
	}
}
