package vehicles_service

import (
	"encoding/json"
	"errors"
	"log"
	vehicle_types "scm-api/types/vehicles"
	vehicle_attributes "scm-api/types/vehicles/attributes"
	vehicle_requests "scm-api/types/vehicles/requests"
)

func ValidateAttributes(req *vehicle_requests.CreateVehicleRequest) error {
	switch req.VehicleType {
	case vehicle_types.Aircraft:
		return validateAircraft(req.Attributes)
	default:
		return errors.New("unsupported vehicle type")
	}
}

func validateAircraft(attributes interface{}) error {
	// Convert attributes interface to JSON bytes.
	attributesJSON, err := json.Marshal(attributes)

	if err != nil {
		return errors.New("failed to process aircraft attributes")
	}

	// Attempt to unmarshal into BaseAircraftAttributes to check if it's a base aircraft.
	var baseAttributes vehicle_attributes.BaseAircraftAttributes

	err = json.Unmarshal(attributesJSON, &baseAttributes)
	if err != nil || baseAttributes.AircraftType == "" {
		log.Printf("Failed to unmarshal into base aircraft attributes or missing aircraft type: %v, JSON: %s\n", err, attributesJSON)
		return errors.New("invalid aircraft base attributes")
	}

	// Base attribute validation.
	if err := validateBaseAircraftAttributes(baseAttributes); err != nil {
		return err
	}

	if baseAttributes.AircraftType == "Defense" {
		var defenseAttributes vehicle_attributes.DefenseAircraftAttributes

		err = json.Unmarshal(attributesJSON, &defenseAttributes)
		if err != nil {
			log.Printf("Failed to unmarshal into defense aircraft attributes: %v, JSON: %s\n", err, attributesJSON)
			return errors.New("invalid defense aircraft attributes")
		}
		return validateDefenseAircraftAttributes(defenseAttributes)
	}

	return errors.New("unknown aircraft type")
}

// validateBaseAircraftAttributes validates common attributes for all aircraft.
func validateBaseAircraftAttributes(attributes vehicle_attributes.BaseAircraftAttributes) error {
	if attributes.AircraftType == "" {
		return errors.New("aircraft type is required")
	}

	if attributes.OperationalRange <= 0 {
		return errors.New("operational range must be positive")
	}

	return nil
}

// validateDefenseAircraftAttributes validates attributes specific to defense aircraft.
func validateDefenseAircraftAttributes(attributes vehicle_attributes.DefenseAircraftAttributes) error {
	// Example validation: ensure that armament is not empty for defense aircraft
	// if len(attributes.Armament) == 0 {
	// 	return errors.New("defense aircraft must have armament")
	// }
	// Additional validations for defense aircraft can be added here
	return nil
}
