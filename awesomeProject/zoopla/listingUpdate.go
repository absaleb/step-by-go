package zoopla

type ListingUpdate struct {
	BranchReference     string                 `json:"branch_reference"`
	Category            string                 `json:"category"`
	ListingReference    string                 `json:"listing_reference"`
	Location            Location               `json:"location"`
	Pricing             Pricing                `json:"pricing"`
	PropertyType        string                 `json:"property_type"`
	Areas               Areas                  `json:"areas"`
	BillsIncluded       []string               `json:"bills_included"`
	DetailedDescription []*DetailedDescription `json:"detailed_description"`
	FurnishedState      string                 `json:"furnished_state"`
	LifeCycleStatus     string                 `json:"life_cycle_status"`
}

// Area
type Area struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}

// Areas
type Areas struct {
	External *MinMaxArea `json:"external,omitempty"`
	Internal *MinMaxArea `json:"internal,omitempty"`
}

// Content
type Content struct {
	Caption string `json:"caption,omitempty"`
	Type    string `json:"type"`
	Url     string `json:"url"`
}

// DetailedDescription
type DetailedDescription struct {
	Dimensions string `json:"dimensions,omitempty"`
	Heading    string `json:"heading,omitempty"`
	Text       string `json:"text,omitempty"`
}

// Dimension
type Dimension struct {
	Length float64 `json:"length"`
	Units  string  `json:"units"`
	Width  float64 `json:"width"`
}

// EpcRatings
type EpcRatings struct {
	EerCurrentRating   string `json:"eer_current_rating,omitempty"`
	EerPotentialRating string `json:"eer_potential_rating,omitempty"`
	EirCurrentRating   string `json:"eir_current_rating,omitempty"`
	EirPotentialRating string `json:"eir_potential_rating,omitempty"`
}

// GoogleStreetView
type GoogleStreetView struct {
	Coordinates *Coordinates `json:"coordinates"`
	Heading     float64      `json:"heading"`
	Pitch       float64      `json:"pitch"`
}

// LeaseExpiry
type LeaseExpiry struct {
	ExpiryDate     string `json:"expiry_date,omitempty"`
	YearsRemaining int    `json:"years_remaining,omitempty"`
}

// ListingUpdate
type ListingUpdateGen struct {
	Accessibility        bool                   `json:"accessibility,omitempty"`
	AdministrationFees   string                 `json:"administration_fees,omitempty"`
	AnnualBusinessRates  float64                `json:"annual_business_rates,omitempty"`
	Areas                *Areas                 `json:"areas,omitempty"`
	AvailableBedrooms    int                    `json:"available_bedrooms,omitempty"`
	AvailableFromDate    string                 `json:"available_from_date,omitempty"`
	Basement             bool                   `json:"basement,omitempty"`
	Bathrooms            int                    `json:"bathrooms,omitempty"`
	BillsIncluded        []string               `json:"bills_included,omitempty"`
	BranchReference      string                 `json:"branch_reference"`
	BurglarAlarm         bool                   `json:"burglar_alarm,omitempty"`
	BusinessForSale      bool                   `json:"business_for_sale,omitempty"`
	BuyerIncentives      []string               `json:"buyer_incentives,omitempty"`
	Category             string                 `json:"category"`
	CentralHeating       string                 `json:"central_heating,omitempty"`
	ChainFree            bool                   `json:"chain_free,omitempty"`
	CommercialUseClasses []string               `json:"commercial_use_classes,omitempty"`
	ConnectedUtilities   []string               `json:"connected_utilities,omitempty"`
	Conservatory         bool                   `json:"conservatory,omitempty"`
	ConstructionYear     int                    `json:"construction_year,omitempty"`
	Content              []*Content             `json:"content,omitempty"`
	CouncilTaxBand       string                 `json:"council_tax_band,omitempty"`
	DecorativeCondition  string                 `json:"decorative_condition,omitempty"`
	Deposit              float64                `json:"deposit,omitempty"`
	DetailedDescription  []*DetailedDescription `json:"detailed_description"`
	DisplayAddress       string                 `json:"display_address,omitempty"`
	DoubleGlazing        bool                   `json:"double_glazing,omitempty"`
	EpcRatings           *EpcRatings            `json:"epc_ratings,omitempty"`
	FeatureList          []string               `json:"feature_list,omitempty"`
	Fireplace            bool                   `json:"fireplace,omitempty"`
	FishingRights        bool                   `json:"fishing_rights,omitempty"`
	FloorLevels          []string               `json:"floor_levels,omitempty"`
	Floors               int                    `json:"floors,omitempty"`
	FurnishedState       string                 `json:"furnished_state,omitempty"`
	GoogleStreetView     *GoogleStreetView      `json:"google_street_view,omitempty"`
	GroundRent           float64                `json:"ground_rent,omitempty"`
	Gym                  bool                   `json:"gym,omitempty"`
	LeaseExpiry          *LeaseExpiry           `json:"lease_expiry,omitempty"`
	LifeCycleStatus      string                 `json:"life_cycle_status"`
	ListedBuildingGrade  string                 `json:"listed_building_grade,omitempty"`
	ListingReference     string                 `json:"listing_reference"`
	LivingRooms          int                    `json:"living_rooms,omitempty"`
	Location             *Location              `json:"location"`
	Loft                 bool                   `json:"loft,omitempty"`
	NewHome              bool                   `json:"new_home,omitempty"`
	OpenDay              string                 `json:"open_day,omitempty"`
	Outbuildings         bool                   `json:"outbuildings,omitempty"`
	OutsideSpace         []string               `json:"outside_space,omitempty"`
	Parking              []string               `json:"parking,omitempty"`
	PetsAllowed          bool                   `json:"pets_allowed,omitempty"`
	PorterSecurity       bool                   `json:"porter_security,omitempty"`
	Pricing              *Pricing               `json:"pricing"`
	PropertyType         string                 `json:"property_type"`
	RateableValue        float64                `json:"rateable_value,omitempty"`
	RentalTerm           string                 `json:"rental_term,omitempty"`
	Repossession         bool                   `json:"repossession,omitempty"`
	Retirement           bool                   `json:"retirement,omitempty"`
	SapRating            string                 `json:"sap_rating,omitempty"`
	ServiceCharge        *ServiceCharge         `json:"service_charge,omitempty"`
	Serviced             bool                   `json:"serviced,omitempty"`
	SharedAccommodation  bool                   `json:"shared_accommodation,omitempty"`
	SummaryDescription   string                 `json:"summary_description,omitempty"`
	SwimmingPool         bool                   `json:"swimming_pool,omitempty"`
	TenantEligibility    *TenantEligibility     `json:"tenant_eligibility,omitempty"`
	Tenanted             bool                   `json:"tenanted,omitempty"`
	TennisCourt          bool                   `json:"tennis_court,omitempty"`
	Tenure               string                 `json:"tenure,omitempty"`
	TotalBedrooms        int                    `json:"total_bedrooms,omitempty"`
	UtilityRoom          bool                   `json:"utility_room,omitempty"`
	Waterfront           bool                   `json:"waterfront,omitempty"`
	WoodFloors           bool                   `json:"wood_floors,omitempty"`
}

// MinMaxArea
type MinMaxArea struct {
	Maximum *Area `json:"maximum,omitempty"`
	Minimum *Area `json:"minimum,omitempty"`
}

// PricePerUnitArea
type PricePerUnitArea struct {
	Price float64 `json:"price"`
	Units string  `json:"units"`
}

// Pricing
type Pricing struct {
	Auction          bool              `json:"auction,omitempty"`
	CurrencyCode     string            `json:"currency_code"`
	Price            float64           `json:"price,omitempty"`
	PricePerUnitArea *PricePerUnitArea `json:"price_per_unit_area,omitempty"`
	PriceQualifier   string            `json:"price_qualifier,omitempty"`
	RentFrequency    string            `json:"rent_frequency,omitempty"`
	TransactionType  string            `json:"transaction_type"`
}

// ServiceCharge
type ServiceCharge struct {
	Charge           float64 `json:"charge"`
	Frequency        string  `json:"frequency,omitempty"`
	PerUnitAreaUnits string  `json:"per_unit_area_units,omitempty"`
}

// TenantEligibility
type TenantEligibility struct {
	Dss      string `json:"dss,omitempty"`
	Students string `json:"students,omitempty"`
}
