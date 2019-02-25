package axal

type TerritoryName string

const (
	Alaska        TerritoryName = "Alaska"
	Argentina                   = "Argentina"
	Brazil                      = "Brazil"
	CentralUS                   = "Central United States"
	EasternCanada               = "Eastern Canada"
	EasternUS                   = "Eastern United States"
	Greenland                   = "Greenland"
	Iceland                     = "Iceland"
	Mexico                      = "Mexico"
	Panama                      = "Panama"
	Peru                        = "Peru"
	Venezuela                   = "Venezuela"
	WesternCanada               = "Western Canada"
	WesternUS                   = "Western United States"
	WestIndies                  = "West Indies"
)

func InitialTerritories() map[TerritoryName]*Territory {
	return map[TerritoryName]*Territory{
		Greenland: NewTerritory(Greenland, US, nil),
		Iceland:   NewTerritory(Iceland, Uncontrolled, nil),
		EasternCanada: NewTerritory(EasternCanada, UK, []TerritoryName{
			WesternCanada, EasternUS, CentralUS,
		}),
		EasternUS: NewTerritory(EasternUS, US, []TerritoryName{
			CentralUS, EasternCanada,
		}),
		CentralUS: NewTerritory(CentralUS, US, []TerritoryName{
			Panama, WesternUS, EasternCanada,
		}),
		WestIndies: NewTerritory(WestIndies, US, nil),
		Venezuela: NewTerritory(Venezuela, Uncontrolled, []TerritoryName{
			Panama, Peru, Brazil,
		}),
		Brazil: NewTerritory(Brazil, US, []TerritoryName{
			Peru, Venezuela, Argentina,
		}),
		Peru: NewTerritory(Peru, Uncontrolled, []TerritoryName{
			Venezuela, Brazil, Argentina,
		}),
		Argentina: NewTerritory(Argentina, Uncontrolled, []TerritoryName{
			Peru, Brazil,
		}),
		WesternCanada: NewTerritory(WesternCanada, UK, []TerritoryName{
			Alaska, EasternCanada, WesternUS,
		}),
		WesternUS: NewTerritory(WesternUS, US, []TerritoryName{
			WesternCanada, CentralUS, Mexico,
		}),
		Mexico: NewTerritory(Mexico, US, []TerritoryName{
			WesternUS, Panama,
		}),
		Panama: NewTerritory(Panama, US, []TerritoryName{
			CentralUS, Mexico, Venezuela,
		}),
		Alaska: NewTerritory(Alaska, US, []TerritoryName{
			WesternCanada,
		}),
	}
}

func NewTerritory(name TerritoryName, cont CountryName, neighbors []TerritoryName) *Territory {
	return &Territory{
		Name:       name,
		Controller: cont,
		Neighbors:  neighbors,
	}
}

type Territory struct {
	Name       TerritoryName
	Controller CountryName
	Neighbors  []TerritoryName
}
