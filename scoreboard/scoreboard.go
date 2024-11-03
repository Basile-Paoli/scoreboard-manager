package scoreboard

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"scoreboard-manager/events"
	"scoreboard-manager/output"
	"strconv"
	"sync"
)

type Scoreboard struct {
	BestOf    int
	RoundName string
	Players   []*Player
	mu        sync.Mutex
	ctx       context.Context
}

func NewScoreboard(bestOf int, roundName string, playerNames []string) *Scoreboard {
	players := make([]*Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = &Player{Name: name}
	}
	return &Scoreboard{
		BestOf:    bestOf,
		RoundName: roundName,
		Players:   players,
	}
}

func (s *Scoreboard) Startup(ctx context.Context) {
	s.ctx = ctx
	err := s.SaveEntireScoreboard()
	if err != nil {
		runtime.LogError(ctx, err.Error())
		runtime.EventsEmit(ctx, string(events.Error), err.Error())
	}
}

func (s *Scoreboard) GetScoreBoard() *Scoreboard {
	return s
}

func (s *Scoreboard) SetBestOf(bestOf int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.BestOf = bestOf
	return s.writeBestOf(bestOf)
}

func (s *Scoreboard) writeBestOf(bestOf int) error {
	if err := output.WriteFile("debug/score/bestof.txt", strconv.Itoa(bestOf)); err != nil {
		return fmt.Errorf("failed to write best of: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetRoundName(roundName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.RoundName = roundName
	return s.writeRoundName(roundName)
}

func (s *Scoreboard) writeRoundName(roundName string) error {
	if err := output.WriteFile("debug/score/roundName.txt", roundName); err != nil {
		return fmt.Errorf("failed to write round name: %w", err)
	}
	return nil
}

func (s *Scoreboard) SaveEntireScoreboard() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	var errs []error

	for i, player := range s.Players {
		if err := s.writePlayerName(i, player.Name); err != nil {
			errs = append(errs, err)
		}
		if err := s.writeScore(i, player.Score); err != nil {
			errs = append(errs, err)
		}
	}
	if err := s.writeBestOf(s.BestOf); err != nil {
		errs = append(errs, err)
	}
	if err := s.writeRoundName(s.RoundName); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to save entire scoreboard: %w", errors.Join(errs...))
	}
	return nil
}
