package vehicle_attributes

type BaseAircraftAttributes struct {
	Model            string `json:"model"`
	AircraftType     string `json:"aircraftType"`
	OperationalRange int    `json:"operationalRange"` // Measured in kilometers (km)
}

type DefenseAircraftAttributes struct {
	BaseAircraftAttributes
	Armament         []string `json:"armament"`
	Communication    []string `json:"communication"`
	Surveillance     []string `json:"surveillance"`
	SpecialEquipment []string `json:"specialEquipment"`
}
