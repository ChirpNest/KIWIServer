package structs

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// DeviceInfo should be the content of the "object" field when the function code is "12"
type DeviceInfo struct {
	FuncNum                   int     `json:"func,omitempty"`
	Port                      int     `json:"port,omitempty"`
	Payload                   string  `json:"payload,omitempty"`
	SerialNumber              int     `json:"serial_number,omitempty"`
	BatteryVoltage            float64 `json:"battery_voltage,omitempty"`
	SwVersionText             string  `json:"sw_version_text,omitempty"`
	ClassGroupText            string  `json:"class_group_text,omitempty"`
	HumidityPercentage        int     `json:"humidity_percentage,omitempty"`
	DeviceLocalDatetime       string  `json:"device_local_datetime,omitempty"`
	BatteryCapacityPercentage int     `json:"battery_capacity_percentage,omitempty"`
}

// Value makes the DeviceInfo struct implement the driver.Value interface.
// This method simply returns the JSON-encoded representation of the struct.
func (a DeviceInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the DeviceInfo struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (a *DeviceInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
