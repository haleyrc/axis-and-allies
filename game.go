package axal // import "github.com/haleyrc/axal"

import (
	"errors"
	"fmt"
)

func NewPlayer(countries []CountryName) (*Player, error) {
	var p Player
	for _, c := range countries {
		p.Countries = append(p.Countries, NewCountry(c))
	}
	if err := p.Valid(); err != nil {
		return nil, err
	}
	return &p, nil
}

type Player struct {
	Countries []*Country
}

func (p *Player) Valid() error {
	if len(p.Countries) < 2 {
		return nil
	}
	allegiance := p.Countries[0].Allegiance
	for i := 1; i < len(p.Countries); i++ {
		if p.Countries[i].Allegiance != allegiance {
			return errors.New("cannot have mixed allegiance")
		}
	}
	return nil
}

var TurnOrder = []CountryName{
	USSR,
	Germany,
	UK,
	Japan,
	US,
}

func NewGame(players []*Player) *Game {
	return &Game{
		Players: players,
		Turn:    0,
	}
}

type Game struct {
	Players []*Player
	Turn    int
}

const MaxTurns = 100

func (g *Game) Loop() error {
	for ; g.Turn < MaxTurns; g.Turn++ {
		curr := TurnOrder[g.Turn%len(TurnOrder)]
		fmt.Printf("Current turn: %s\n", curr)
	}
	return nil
}
