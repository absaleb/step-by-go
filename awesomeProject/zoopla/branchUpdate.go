package zoopla

type BranchUpdate struct {
	BranchName      string    `json:"branch_name"`
	BranchReference string    `json:"branch_reference"`
	Email           string    `json:"email,omitempty"`
	Location        *Location `json:"location"`
	Telephone       string    `json:"telephone,omitempty"`
	Website         string    `json:"website,omitempty"`
}

// Coordinates
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Location
type Location struct {
	Coordinates          *Coordinates `json:"coordinates,omitempty"`
	CountryCode          string       `json:"country_code"`
	County               string       `json:"county,omitempty"`
	Locality             string       `json:"locality,omitempty"`
	PafAddress           *PafAddress  `json:"paf_address,omitempty"`
	PafUdprn             string       `json:"paf_udprn,omitempty"`
	PostalCode           string       `json:"postal_code,omitempty"`
	PropertyNumberOrName interface{}  `json:"property_number_or_name,omitempty"`
	StreetName           string       `json:"street_name,omitempty"`
	TownOrCity           string       `json:"town_or_city"`
}

// PafAddress
type PafAddress struct {
	AddressKey      string `json:"address_key"`
	OrganisationKey string `json:"organisation_key"`
	PostcodeType    string `json:"postcode_type"`
}
