package axal

import (
	"testing"
)

func TestGame(t *testing.T) {
	g, err := NewGame(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	g.Loop()
}

func TestDuplicateCountries(t *testing.T) {
	ps := []*Player{
		&Player{Countries: map[CountryName]*Country{US: nil}},
		&Player{Countries: map[CountryName]*Country{USSR: nil}},
	}
	if duplicateCountries(ps) {
		t.Errorf("expected no duplicate countries, but found some")
	}
	ps = append(ps, &Player{Countries: map[CountryName]*Country{USSR: nil}})
	if !duplicateCountries(ps) {
		t.Errorf("expected duplicate countries, but found none")
	}
}
