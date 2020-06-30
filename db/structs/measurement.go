package structs

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Measurement should be the content of the "object" field when the function code is "1"
type Measurement struct {
	FuncNum          int    `json:"func,omitempty"`
	Port             int    `json:"port,omitempty"`
	ChannelType      int    `json:"ct,omitempty"`
	Payload          string `json:"payload,omitempty"`
	ChannelCount     int    `json:"channelCount,omitempty"`
	MeasuredChannels string `json:"channel,omitempty"`
	Time             time.Time
	Channels         Channels
}

// Channels contains the channels of that are in the "object" field when the funciton code is "1"
type Channels map[string]interface{}

// Value makes the Measurement struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (a Measurement) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the Measurement struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (a *Measurement) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Value makes the Channels struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (a Channels) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the Channels struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (a *Channels) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
