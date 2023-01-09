package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"xoq/pkg/domain"
	qTable "xoq/pkg/q_table/in_memory"
	statistic "xoq/pkg/statistic/in_memory"
)

const (
	workers = 20_000
	rounds  = 2_000_000_000
)

func main() {
	games := make(chan Game)
	inMemoryStatistic := statistic.InMemoryStatistic{}
	inMemoryQTableA := qTable.NewQTable()
	inMemoryQTableB := qTable.NewQTable()
	var waiter sync.WaitGroup
	for i := 0; i < workers; i++ {
		go worker(games, &waiter)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go printStatistics(ctx, &inMemoryStatistic)
	for i := 0; i < rounds; i++ {
		agent := domain.Agent{QTable: inMemoryQTableA}
		agentB := domain.Agent{QTable: inMemoryQTableB}
		game := Game{
			Statistic: &inMemoryStatistic,
			Board:     domain.NewEmptyBoard(),
			AgentA:    agent,
			AgentB:    agentB,
		}
		games <- game
	}
	close(games)
	waiter.Wait()
	cancel()
	inMemoryQTableA.WriteToDisk("policy_a.json")
	inMemoryQTableB.WriteToDisk("policy_b.json")
}

func printStatistics(ctx context.Context, memoryStatistic domain.Statistic) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			fmt.Print("\033[H\033[2J")
			t, w, l, d := memoryStatistic.Get()
			fmt.Printf("Total: %d, Wins: %d, Losses: %d, Draws: %d\n", t, w, l, d)
		case <-ctx.Done():
			return
		}
	}
}

func worker(games chan Game, waiter *sync.WaitGroup) {
	waiter.Add(1)
	for game := range games {
		game.Run()
	}
	waiter.Done()
}

type Game struct {
	Statistic domain.Statistic
	Board     *domain.Board
	AgentA    domain.Agent
	AgentB    domain.Agent
}

func (g *Game) Run() {
	for {
		randomAction := g.AgentB.ChooseAction(*g.Board)
		g.Board[randomAction.Row][randomAction.Column] = 'P'
		if g.isWin() {
			return
		}
		agentAction := g.AgentA.ChooseAction(*g.Board)
		g.Board[agentAction.Row][agentAction.Column] = 'A'
		if g.isWin() {
			return
		}
	}
}

func (g *Game) isWin() bool {
	winner := g.Board.Winner()
	if winner != domain.SymbolNone {
		if winner == domain.SymbolAgent {
			g.AgentA.Reward(1)
			g.AgentB.Reward(0)
			g.Statistic.Won()
		} else {
			g.AgentA.Reward(0)
			g.AgentB.Reward(1)
			g.Statistic.Lost()
		}
		return true
	} else {
		if g.Board.IsDraw() {
			g.AgentA.Reward(0.5)
			g.AgentB.Reward(0.3)
			g.Statistic.Draw()
			return true
		}
	}
	return false
}
