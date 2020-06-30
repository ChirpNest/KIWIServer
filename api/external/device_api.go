package external

import (
	"context"

	"example.org/luksam/kiwi-server/api/chirpstackapiclient"
	"example.org/luksam/kiwi-server/api/helpers"
	"example.org/luksam/kiwi-server/apidefinition/go/external"
	"example.org/luksam/kiwi-server/db"
	"example.org/luksam/kiwi-server/db/structs"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DeviceServerAPI implements the external.DeviceServerServer interface.
type DeviceServerAPI struct {
}

// NewDeviceServerAPI returns a new ApplicationServerAPI.
func NewDeviceServerAPI() *DeviceServerAPI {
	return &DeviceServerAPI{}
}

// List returns the available devices.
func (a *DeviceServerAPI) List(ctx context.Context, in *external.ListDeviceRequest) (*external.ListDeviceResponse, error) {

	devicesFromChirpStack := chirpstackapiclient.GetDevices()

	devicesResp := make([]*external.DeviceListItem, len(devicesFromChirpStack))
	for i, element := range devicesFromChirpStack {

		var deviceInfo *structs.DeviceInfo = nil
		deviceInfo, err := db.GetNewestDeviceInfo(element.DevEui, conf.Database)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}

		var serialNumber int64
		var deviceType string
		deviceInfoAvailable := deviceInfo != nil

		if deviceInfoAvailable {
			serialNumber = int64(deviceInfo.SerialNumber)
			deviceType = deviceInfo.ClassGroupText
		}

		measurementOverviewDetails, err := db.GetMeasurementOverviewDetails(element.DevEui, conf.Database)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}

		var oldestMeasurementTimeProto, newestMeasurementTimeProto *timestamppb.Timestamp
		var numberOfMeasurements int64

		if measurementOverviewDetails != nil {
			oldestMeasurementTimeProto, err = ptypes.TimestampProto(measurementOverviewDetails.OldestMeasurementTime)
			if err != nil {
				return nil, helpers.ErrToRPCError(err)
			}

			newestMeasurementTimeProto, err = ptypes.TimestampProto(measurementOverviewDetails.NewestMeasurementTime)
			if err != nil {
				return nil, helpers.ErrToRPCError(err)
			}

			numberOfMeasurements = int64(measurementOverviewDetails.NumberOfMeasurements)
		}

		devicesResp[i] = &external.DeviceListItem{
			DevEui:               element.DevEui,
			Name:                 element.Name,
			Description:          element.Description,
			SerialNumber:         serialNumber,
			DeviceType:           deviceType,
			DeviceInfoAvailable:  deviceInfoAvailable,
			FirstMeasurementTime: newestMeasurementTimeProto,
			LastMeasurementTime:  oldestMeasurementTimeProto,
			NumberOfMeasurements: numberOfMeasurements,
		}
	}

	resp := external.ListDeviceResponse{
		NumberOfDevices: int64(len(devicesResp)),
		Devices:         devicesResp,
	}

	return &resp, nil
}
