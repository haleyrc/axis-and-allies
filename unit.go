package axal

type CountryName string

const (
	Uncontrolled CountryName = "Uncontrolled"
	US                       = "United States"
	UK                       = "United Kingdom"
	USSR                     = "Soviet Union"
	Germany                  = "Germany"
	Japan                    = "Japan"
)

type UnitType string

const (
	UnitInfantry              UnitType = "Infantry"
	UnitArtillery                      = "Artillery"
	UnitTank                           = "Tank"
	UnitAntiAircraftArtillery          = "AntiAircraftArtillery"
	UnitFighter                        = "Fighter"
	UnitBomber                         = "Bomber"
	UnitAircraftCarrier                = "AircraftCarrier"
	UnitBattleship                     = "Battleship"
	UnitCruiser                        = "Cruiser"
	UnitDestroyer                      = "Destroyer"
	UnitSubmarine                      = "Submarine"
	UnitTransport                      = "Transport"
)

type Unit interface {
	Name() string
	Type() UnitType
}

func NewInfantry(country CountryName) *Infantry {
	return &Infantry{Country: country}
}

type Infantry struct {
	Country CountryName
}

func (i *Infantry) Name() string {
	return "Infantry"
}

func (i *Infantry) Type() UnitType {
	return UnitInfantry
}

func NewArtillery(country CountryName) *Artillery {
	return &Artillery{Country: country}
}

type Artillery struct {
	Country CountryName
}

func (a *Artillery) Name() string {
	switch a.Country {
	case US:
		return "105mm Howitzer"
	case UK:
		return "Ordnance QF 25 Pounder"
	case USSR:
		return "105mm Howitzer"
	case Germany:
		return "10.5cm LeFH 18"
	case Japan:
		return "Type 92 Howitzer"
	default:
		return ""
	}
}

func (a *Artillery) Type() UnitType {
	return UnitArtillery
}

func NewTank(country CountryName) *Tank {
	return &Tank{Country: country}
}

type Tank struct {
	Country CountryName
}

func (t *Tank) Name() string {
	switch t.Country {
	case US:
		return "Sherman"
	case UK:
		return "Matilda II"
	case USSR:
		return "T-34"
	case Germany:
		return "Panther"
	case Japan:
		return "Type 95"
	default:
		return ""
	}
}

func (t *Tank) Type() UnitType {
	return UnitTank
}

func NewAntiAircraftArtillery(country CountryName) *AntiAircraftArtillery {
	return &AntiAircraftArtillery{Country: country}
}

type AntiAircraftArtillery struct {
	Country CountryName
}

func (a *AntiAircraftArtillery) Name() string {
	switch a.Country {
	case US:
		return "90mm M1"
	case UK:
		return "3.7in QFAA"
	case USSR:
		return "85mm M1939"
	case Germany:
		return "8.8cm Flak 36"
	case Japan:
		return "75mm Type 88"
	default:
		return ""
	}
}

func (a *AntiAircraftArtillery) Type() UnitType {
	return UnitAntiAircraftArtillery
}

func NewFighter(country CountryName) *Fighter {
	return &Fighter{Country: country}
}

type Fighter struct {
	Country CountryName
}

func (f *Fighter) Name() string {
	switch f.Country {
	case US:
		return "P-38"
	case UK:
		return "Spitfire"
	case USSR:
		return "Yak 3"
	case Germany:
		return "BF-109"
	case Japan:
		return "AGM2 Zero"
	default:
		return ""
	}
}

func (f *Fighter) Type() UnitType {
	return UnitFighter
}

func NewBomber(country CountryName) *Bomber {
	return &Bomber{Country: country}
}

type Bomber struct {
	Country CountryName
}

func (b *Bomber) Name() string {
	switch b.Country {
	case US:
		return "B-17"
	case UK:
		return "Halifax"
	case USSR:
		return "Petlyakov PE-8"
	case Germany:
		return "JU-88"
	case Japan:
		return "G4M2E Betty"
	default:
		return ""
	}
}

func (b *Bomber) Type() UnitType {
	return UnitBomber
}

func NewAircraftCarrier(country CountryName) *AircraftCarrier {
	return &AircraftCarrier{Country: country}
}

type AircraftCarrier struct {
	Country CountryName
}

func (a *AircraftCarrier) Name() string {
	switch a.Country {
	case US:
		return "Wasp"
	case UK:
		return "Illustrious"
	case USSR:
		return "Illustrious"
	case Germany:
		return "Graf Zeppelin"
	case Japan:
		return "Shinano"
	default:
		return ""
	}
}

func (a *AircraftCarrier) Type() UnitType {
	return UnitAircraftCarrier
}

func NewBattleship(country CountryName) *Battleship {
	return &Battleship{Country: country}
}

type Battleship struct {
	Country CountryName
}

func (b *Battleship) Name() string {
	switch b.Country {
	case US:
		return "Iowa"
	case UK:
		return "Royal Oak"
	case USSR:
		return "Gangut"
	case Germany:
		return "Bismarck"
	case Japan:
		return "Yamato"
	default:
		return ""
	}
}

func (b *Battleship) Type() UnitType {
	return UnitBattleship
}

func NewCruiser(country CountryName) *Cruiser {
	return &Cruiser{Country: country}
}

type Cruiser struct {
	Country CountryName
}

func (c *Cruiser) Name() string {
	switch c.Country {
	case US:
		return "Portland"
	case UK:
		return "Kent"
	case USSR:
		return "Kirov"
	case Germany:
		return "Hipper"
	case Japan:
		return "Takao"
	default:
		return ""
	}
}

func (c *Cruiser) Type() UnitType {
	return UnitCruiser
}

func NewDestroyer(country CountryName) *Destroyer {
	return &Destroyer{Country: country}
}

type Destroyer struct {
	Country CountryName
}

func (d *Destroyer) Name() string {
	switch d.Country {
	case US:
		return "Johnston"
	case UK:
		return "S Class"
	case USSR:
		return "Gnevnyi"
	case Germany:
		return "Zerstorer 1934"
	case Japan:
		return "Fubuki"
	default:
		return ""
	}
}

func (d *Destroyer) Type() UnitType {
	return UnitDestroyer
}

func NewSubmarine(country CountryName) *Submarine {
	return &Submarine{Country: country}
}

type Submarine struct {
	Country CountryName
}

func (s *Submarine) Name() string {
	switch s.Country {
	case US:
		return "Ray"
	case UK:
		return "T Class"
	case USSR:
		return "S Class"
	case Germany:
		return "Type VII"
	case Japan:
		return "I-class"
	default:
		return ""
	}
}

func (s *Submarine) Type() UnitType {
	return UnitSubmarine
}

func NewTransport(country CountryName) *Transport {
	return &Transport{Country: country}
}

type Transport struct {
	Country CountryName
}

func (t *Transport) Name() string {
	switch t.Country {
	case US:
		return "Libery Ship"
	case UK:
		return "Libery Ship"
	case USSR:
		return "Liberty Ship"
	case Germany:
		return "Hilfskruezer"
	case Japan:
		return "Hakusan Maru"
	default:
		return ""
	}
}

func (t *Transport) Type() UnitType {
	return UnitTransport
}
