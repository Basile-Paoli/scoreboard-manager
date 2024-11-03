package startgg

import (
	"context"
	"fmt"
	"scoreboard-manager/scoreboard"
	"scoreboard-manager/startgg/api"
)

type Startgg struct {
	ctx             context.Context
	StreamQueues    []api.GetStreamQueuesTournamentStreamQueue
	currentStreamId int
	currentSetId    int
	Scoreboard      *scoreboard.Scoreboard
}

func (s *Startgg) fetchStreamQueues(tournamentSlug string) error {
	res, err := api.GetStreamQueues(s.ctx, api.NewClient(), tournamentSlug)
	if err != nil {
		return fmt.Errorf("failed to fetch stream queues: %w", err)
	}

	s.StreamQueues = res.Tournament.StreamQueue
	return nil
}

func (s *Startgg) GetStreamQueues(tournamentSlug string) ([]api.GetStreamQueuesTournamentStreamQueue, error) {
	err := s.fetchStreamQueues(tournamentSlug)
	return s.StreamQueues, err
}

func (s *Startgg) Startup(ctx context.Context, sb *scoreboard.Scoreboard) {
	s.ctx = ctx
	s.Scoreboard = sb
}
