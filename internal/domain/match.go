package domain

import "time"

type Match struct {
	ID        string
	Duration  time.Duration
	Queue     QueueType
	Players   []MatchPlayer
}

type MatchPlayer struct {
	PlayerID  PlayerID
	Champion  Champion
	KDA       KDA
	Win       bool
}

type PlayerID string

type KDA struct {
	Kills   int
	Deaths  int
	Assists int
}

func (k KDA) Ratio() float64 {
	if k.Deaths == 0 {
		return float64(k.Kills + k.Assists)
	}
	return float64(k.Kills+k.Assists) / float64(k.Deaths)
}

type QueueType int

const (
	QueueUnknown QueueType = iota
	QueueRankedSolo
	QueueRankedFlex
	QueueARAM
)

func QueueFromRiotID(id int) QueueType {
	switch id {
	case 420:
		return QueueRankedSolo
	case 440:
		return QueueRankedFlex
	case 450:
		return QueueARAM
	default:
		return QueueUnknown
	}
}
