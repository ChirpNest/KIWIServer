package structs

import "time"

// MeasurementOverviewDetails is used to fill in measurement overview details about a device received from the database.
type MeasurementOverviewDetails struct {
	OldestMeasurementTime time.Time
	NewestMeasurementTime time.Time
	NumberOfMeasurements  int
}
