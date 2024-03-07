package vehicle_attributes

type BaseAircraftAttributes struct {
	AircraftType     string `json:"aircraft_type"`
	OperationalRange int    `json:"operational_range"` // Measured in kilometers (km)
}

type DefenseAircraftAttributes struct {
	BaseAircraftAttributes
	Armament         []string `json:"armament"`
	Communication    []string `json:"communication"`
	Surveillance     []string `json:"surveillance"`
	SpecialEquipment []string `json:"special_equipment"`
}
