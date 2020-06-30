package db

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unicode"

	"example.org/luksam/kiwi-server/config"
	"example.org/luksam/kiwi-server/db/structs"

	// import Go postgres driver for the database/sql package
	_ "github.com/lib/pq"
)

// GetMeasurements gets the measurements from the database
func GetMeasurements(deviceEUI string, start *time.Time, end *time.Time, config config.Database) ([]structs.Measurement, error) {

	deviceEUIBytes, err := hex.DecodeString(deviceEUI)
	if err != nil {
		return nil, err
	}

	db, err := openConnection(config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var data structs.Measurement
	var dataString string
	var time time.Time

	rows, err := db.Query(sqlStatementGetMeasurements, deviceEUIBytes, start, end)
	if err != nil {
		return nil, err
	}

	var measurements []structs.Measurement

	for rows.Next() {
		data = structs.Measurement{}
		rows.Scan(&dataString, &time)

		json.Unmarshal([]byte(dataString), &data)
		// Channels contains the same fields gain, probably filter for only capital letters
		json.Unmarshal([]byte(dataString), &data.Channels)

		data.Time = time
		// This is a workaround: removes entries that start lowercease, since channels all start uppercase
		for key := range data.Channels {
			if !unicode.IsUpper(rune(key[0])) {
				delete(data.Channels, key)
			}
		}

		// If there are measurements, add it to the list
		if len(data.Channels) > 0 {
			measurements = append(measurements, data)
		}

	}

	return measurements, nil
}

// DeleteMeasurements deletes the measurements from the database
func DeleteMeasurements(deviceEUI string, start *time.Time, end *time.Time, config config.Database) error {

	deviceEUIBytes, err := hex.DecodeString(deviceEUI)
	if err != nil {
		return err
	}

	db, err := openConnection(config)
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec(sqlStatementDeleteMeasurements, deviceEUIBytes, start, end)
	if err != nil {
		return err
	}

	blub, err := result.RowsAffected()
	if blub <= 0 {
		return errors.New("No measurements were found to delete")
	}

	return nil
}

// GetNewestDeviceInfo queries the database to get the newest device info
func GetNewestDeviceInfo(deviceEUI string, config config.Database) (*structs.DeviceInfo, error) {

	deviceEUIBytes, err := hex.DecodeString(deviceEUI)
	if err != nil {
		return nil, err
	}

	db, err := openConnection(config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var data *structs.DeviceInfo = nil

	rows, err := db.Query(sqlStatementNewestDeviceInfo, deviceEUIBytes)

	if err != nil {
		return nil, err
	}

	// scan the single result
	if rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			return nil, err
		}
	}

	// throw if there is more than one result
	if rows.Next() {
		return nil, errors.New("the sql statement to retrieve device unexpectedly returned more than 1 result")
	}

	return data, nil
}

// GetMeasurementOverviewDetails queries the database to get the time of the oldest and the newest measurement and the overall number of measurements.
func GetMeasurementOverviewDetails(deviceEUI string, config config.Database) (*structs.MeasurementOverviewDetails, error) {

	deviceEUIBytes, err := hex.DecodeString(deviceEUI)
	if err != nil {
		return nil, err
	}

	db, err := openConnection(config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(sqlStatementMeasurementOverviewDetails, deviceEUIBytes)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		// no measurements found: return nil
		return nil, nil
	}

	var oldestMeasurementTime, newestMeasurementTime time.Time
	var numberOfMeasurements int
	err = rows.Scan(&oldestMeasurementTime, &newestMeasurementTime, &numberOfMeasurements)
	if err != nil {
		return nil, err
	}

	result := structs.MeasurementOverviewDetails{
		OldestMeasurementTime: oldestMeasurementTime,
		NewestMeasurementTime: newestMeasurementTime,
		NumberOfMeasurements:  numberOfMeasurements,
	}

	return &result, nil
}

// TryConnect tries to connect to the database to figure out if the configuration works.
func TryConnect(config config.Database) error {
	db, err := openConnection(config)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

// CheckTableExists checks if a table with the given name exists on the database.
func CheckTableExists(tableName string, config config.Database) (*bool, error) {

	db, err := openConnection(config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(sqlStatementTableExists, tableName)

	var tableExists bool
	err = row.Scan(&tableExists)
	if err != nil {
		return nil, err
	}

	return &tableExists, nil
}

// RunSQLStatementForUnitTest runs the provided sql statement and is only used for testing.
func RunSQLStatementForUnitTest(sqlStatement string, config config.Database) error {

	db, err := openConnection(config)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	return err
}

func openConnection(config config.Database) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	// this does not actually create a connection but only checks the parameters
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// this does open up a connection to check if it works
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
