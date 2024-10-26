package scoreboard

import (
	"fmt"
	"log"
	"scoreboard-manager/output"
	"strconv"
	"sync"
)

type Scoreboard struct {
	BestOf    int
	RoundName string
	Players   []*Player
	mu        sync.Mutex
}

type Player struct {
	Name  string
	Score int
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

func (s *Scoreboard) GetScoreBoard() *Scoreboard {
	return s
}

func (s *Scoreboard) SetScore(playerIndex, score int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	player := s.Players[playerIndex]

	player.Score = score

	return s.writeScore(playerIndex, score)
}

func (s *Scoreboard) writeScore(index, score int) error {
	if err := output.WriteFile("debug/score/players/"+strconv.Itoa(index)+"/score.txt", strconv.Itoa(score)); err != nil {
		return fmt.Errorf("failed to write player score: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetPlayerName(playerIndex int, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	player := s.Players[playerIndex]

	player.Name = name

	return s.writePlayerName(playerIndex, name)
}

func (s *Scoreboard) writePlayerName(index int, name string) error {
	if err := output.WriteFile("debug/score/players/"+strconv.Itoa(index)+"/name.txt", name); err != nil {
		return fmt.Errorf("failed to write player name: %w", err)
	}
	return nil
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

func (s *Scoreboard) SaveEntireScoreboard() {
	s.mu.Lock()
	defer s.mu.Unlock()
	wg := sync.WaitGroup{}
	wg.Add(2 + len(s.Players)*2)

	for i, player := range s.Players {
		go func(i int, player *Player) {
			defer wg.Done()
			if err := s.writePlayerName(i, player.Name); err != nil {
				log.Println(err)
			}
		}(i, player)
		go func(i int, player *Player) {
			defer wg.Done()
			if err := s.writeScore(i, player.Score); err != nil {
				log.Println(err)
			}
		}(i, player)
	}

	go func() {
		defer wg.Done()
		if err := s.writeBestOf(s.BestOf); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := s.writeRoundName(s.RoundName); err != nil {
			log.Println(err)
		}
	}()

	wg.Wait()
}
