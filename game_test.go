package axal_test

import (
	"testing"

	"github.com/haleyrc/axal"
)

func TestGame(t *testing.T) {
	g := axal.NewGame(nil)
	g.Loop()
}
