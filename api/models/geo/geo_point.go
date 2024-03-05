package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// GeoPoint represents a geographical point using longitude and latitude.
type GeoPoint struct {
	Lat float64
	Lng float64
}

// Value makes the GeoPoint struct implement the driver.Valuer interface.
// This method returns the string representation of the GeoPoint, which
// PostgreSQL understands as a geography type.
func (g GeoPoint) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", g.Lng, g.Lat), nil
}

// Scan implements the sql.Scanner interface for GeoPoint.
func (g *GeoPoint) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// Convert to string
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("GeoPoint must be a string, got %T instead", value)
	}

	// Remove the 'POINT(' prefix and the ')' suffix
	val = strings.TrimPrefix(val, "POINT(")
	val = strings.TrimSuffix(val, ")")

	// Split the string by space to get longitude and latitude
	parts := strings.Split(val, " ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid POINT format")
	}

	// Parse the longitude and latitude
	var lng, lat float64
	_, err := fmt.Sscanf(parts[0]+" "+parts[1], "%f %f", &lng, &lat)
	if err != nil {
		return err
	}

	g.Lng = lng
	g.Lat = lat

	return nil
}
