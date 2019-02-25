package axal

func InitialState() *State {
	return &State{
		Territories: InitialTerritories(),
	}
}

type State struct {
	Territories map[TerritoryName]*Territory
}
