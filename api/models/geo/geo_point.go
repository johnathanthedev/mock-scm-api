package models

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

// GeoPoint represents a geographical point using longitude and latitude.
type GeoPoint struct {
	Lat float64
	Lng float64
}

// Value makes the GeoPoint struct implement the driver.Valuer interface.
// This method converts the GeoPoint to a string representation in WKT (Well-Known Text) format.
func (g GeoPoint) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", g.Lng, g.Lat), nil
}

// Scan implements the sql.Scanner interface for GeoPoint.
// It decodes WKB-encoded data provided as a hexadecimal string into the GeoPoint struct.
func (g *GeoPoint) Scan(value interface{}) error {
	if value == nil {
		return nil // No data to scan
	}

	// Assume all values are hexadecimal strings representing WKB data.
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("GeoPoint requires a hex string, got %T instead", value)
	}

	// Log the attempt to parse the WKB hex string for debugging.
	log.Printf("Attempting to parse WKB hex string: %s", v)

	// Convert the hexadecimal string to a []byte before unmarshalling.
	wkbBytes, err := hex.DecodeString(v)
	if err != nil {
		return fmt.Errorf("error decoding hex string to binary: %v", err)
	}

	// Unmarshal the binary WKB data into an orb.Point.
	geom, err := wkb.Unmarshal(wkbBytes)
	if err != nil {
		return fmt.Errorf("error decoding WKB data from hex string: %v", err)
	}

	point, ok := geom.(orb.Point)
	if !ok {
		return fmt.Errorf("decoded geometry from hex string is not of type orb.Point, got %T", geom)
	}

	// Update the GeoPoint struct with the decoded longitude and latitude.
	g.Lng = point.Lon()
	g.Lat = point.Lat()

	return nil
}
