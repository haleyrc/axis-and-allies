package axal

type CountryName string

const (
	US      CountryName = "US"
	UK                  = "UK"
	USSR                = "USSR"
	Germany             = "Germany"
	Japan               = "Japan"
)

type UnitType string

const (
	Infantry              UnitType = "Infantry"
	Artillery                      = "Artillery"
	Tank                           = "Tank"
	AntiAircraftArtillery          = "AntiAircraftArtillery"
	Fighter                        = "Fighter"
	Bomber                         = "Bomber"
	AircraftCarrier                = "AircraftCarrier"
	Battleship                     = "Battleship"
	Cruiser                        = "Cruiser"
	Destroyer                      = "Destroyer"
	Submarine                      = "Submarine"
	Transport                      = "Transport"
)

func UnitNameForCountry(country CountryName, unit UnitType) string {
	switch country {
	case US:
		switch unit {
		case Infantry:
			return "Infantry"
		case Artillery:
			return "105mm Howitzer"
		case Tank:
			return "Sherman"
		case AntiAircraftArtillery:
			return "90mm M1"
		case Fighter:
			return "P-38"
		case Bomber:
			return "B-17"
		case AircraftCarrier:
			return "Wasp"
		case Battleship:
			return "Iowa"
		case Cruiser:
			return "Portland"
		case Destroyer:
			return "Johnston"
		case Submarine:
			return "Ray"
		case Transport:
			return "Libery Ship"
		default:
			return ""
		}
	case UK:
		switch unit {
		case Infantry:
			return "Infantry"
		case Artillery:
			return "Ordnance QF 25 Pounder"
		case Tank:
			return "Matilda II"
		case AntiAircraftArtillery:
			return "3.7in QFAA"
		case Fighter:
			return "Spitfire"
		case Bomber:
			return "Halifax"
		case AircraftCarrier:
			return "Illustrious"
		case Battleship:
			return "Royal Oak"
		case Cruiser:
			return "Kent"
		case Destroyer:
			return "S Class"
		case Submarine:
			return "T Class"
		case Transport:
			return "Libery Ship"
		default:
			return ""
		}
	case USSR:
		switch unit {
		case Infantry:
			return "Infantry"
		case Artillery:
			return "105mm Howitzer"
		case Tank:
			return "T-34"
		case AntiAircraftArtillery:
			return "85mm M1939"
		case Fighter:
			return "Yak 3"
		case Bomber:
			return "Petlyakov PE-8"
		case AircraftCarrier:
			return "Illustrious"
		case Battleship:
			return "Gangut"
		case Cruiser:
			return "Kirov"
		case Destroyer:
			return "Gnevnyi"
		case Submarine:
			return "S Class"
		case Transport:
			return "Liberty Ship"
		default:
			return ""
		}
	case Germany:
		switch unit {
		case Infantry:
			return "Infantry"
		case Artillery:
			return "10.5cm LeFH 18"
		case Tank:
			return "Panther"
		case AntiAircraftArtillery:
			return "8.8cm Flak 36"
		case Fighter:
			return "BF-109"
		case Bomber:
			return "JU-88"
		case AircraftCarrier:
			return "Graf Zeppelin"
		case Battleship:
			return "Bismarck"
		case Cruiser:
			return "Hipper"
		case Destroyer:
			return "Zerstorer 1934"
		case Submarine:
			return "Type VII"
		case Transport:
			return "Hilfskruezer"
		default:
			return ""
		}
	case Japan:
		switch unit {
		case Infantry:
			return "Infantry"
		case Artillery:
			return "Type 92 Howitzer"
		case Tank:
			return "Type 95"
		case AntiAircraftArtillery:
			return "75mm Type 88"
		case Fighter:
			return "AGM2 Zero"
		case Bomber:
			return "G4M2E Betty"
		case AircraftCarrier:
			return "Shinano"
		case Battleship:
			return "Yamato"
		case Cruiser:
			return "Takao"
		case Destroyer:
			return "Fubuki"
		case Submarine:
			return "I-class"
		case Transport:
			return "Hakusan Maru"
		default:
			return ""
		}
	default:
		return ""
	}
}

func NewUnit(country CountryName, t UnitType) *Unit {
	return &Unit{
		Country: country,
		Type:    t,
		Name:    UnitNameForCountry(country, t),
	}
}

type Unit struct {
	Country CountryName
	Type    UnitType
	Name    string
}
