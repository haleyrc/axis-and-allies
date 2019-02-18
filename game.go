package axal // import "github.com/haleyrc/axal"

import (
	"errors"
	"fmt"
)

func NewPlayer(countries []CountryName) (*Player, error) {
	var p Player
	for _, c := range countries {
		p.Countries[c] = NewCountry(c)
	}
	if err := p.Valid(); err != nil {
		return nil, err
	}
	return &p, nil
}

type Player struct {
	Countries map[CountryName]*Country
}

func (p *Player) Valid() error {
	if len(p.Countries) < 1 {
		return nil
	}
	var axis, allies bool
	for _, v := range p.Countries {
		axis = v.Allegiance == Axis
		allies = v.Allegiance == Allies
		if axis && allies {
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

type Turn struct {
	Player  *Player
	Country *Country
}

func createTurnForCountry(country CountryName, ps []*Player) (*Turn, bool) {
	for _, p := range ps {
		c, ok := p.Countries[country]
		if ok {
			return &Turn{
				Player:  p,
				Country: c,
			}, true
		}
	}
	return nil, false
}

func NewGame(players []*Player) (*Game, error) {
	if duplicateCountries(players) {
		return nil, errors.New("duplicate countries assigned")
	}
	var turns []*Turn
	for _, country := range TurnOrder {
		t, found := createTurnForCountry(country, players)
		if found {
			turns = append(turns, t)
		}
	}
	return &Game{
		Turns: turns,
		Turn:  0,
	}, nil
}

type Game struct {
	Turns []*Turn
	Turn  int
}

const MaxTurns = 100

func (g *Game) Loop() error {
	for ; g.Turn < MaxTurns; g.Turn++ {
		curr := TurnOrder[g.Turn%len(TurnOrder)]
		fmt.Printf("Current turn: %s\n", curr)
	}
	return nil
}

func duplicateCountries(ps []*Player) bool {
	seen := make(map[CountryName]bool)
	for _, p := range ps {
		for c := range p.Countries {
			_, found := seen[c]
			if found {
				return true
			}
			seen[c] = true
		}
	}
	return false
}
