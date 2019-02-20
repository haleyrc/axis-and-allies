package axal

func StartingIncomeForCountry(country CountryName) int {
	switch country {
	case US:
		return 42
	case USSR:
		return 24
	case Germany:
		return 40
	case UK:
		return 30
	case Japan:
		return 30
	default:
		return 0
	}
}

func ColorForCountry(country CountryName) string {
	switch country {
	case US:
		return "#336600"
	case USSR:
		return "#660000"
	case Germany:
		return "#222222"
	case UK:
		return "#CC9900"
	case Japan:
		return "#FF6600"
	default:
		return "#000000"
	}
}

type Allegiance string

const (
	Axis   Allegiance = "Axis"
	Allies            = "Allies"
)

func AllegianceForCountry(country CountryName) Allegiance {
	switch country {
	case Germany, Japan:
		return Axis
	default:
		return Allies
	}
}

// TODO (RCH): Implement this
func UnitsForCountry(country CountryName) []Unit {
	return nil
}

func NewCountry(name CountryName) *Country {
	return &Country{
		Name:       name,
		Income:     StartingIncomeForCountry(name),
		Color:      ColorForCountry(name),
		Allegiance: AllegianceForCountry(name),
		Units:      UnitsForCountry(name),
	}
}

type Country struct {
	Name       CountryName
	Income     int
	Color      string
	Allegiance Allegiance
	Units      []Unit
}
