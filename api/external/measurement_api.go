package external

import (
	"context"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"example.org/luksam/kiwi-server/api/helpers"
	"example.org/luksam/kiwi-server/apidefinition/go/external"
	"example.org/luksam/kiwi-server/db"
)

// MeasurementServerAPI implements the external.MeasurementServerServer interface.
type MeasurementServerAPI struct {
}

// NewMeasurementServerAPI returns a new MeasurementServerAPI.
func NewMeasurementServerAPI() *MeasurementServerAPI {
	return &MeasurementServerAPI{}
}

func mapInterfaceToMapFloat64(m map[string]interface{}) map[string]float64 {
	mapFloat := make(map[string]float64, len(m))
	for key, value := range m {

		mapFloat[key] = value.(float64)
	}
	return mapFloat
}

// Get returns the requested measurements.
func (a *MeasurementServerAPI) Get(ctx context.Context, req *external.GetMeasurementsRequest) (*external.GetMeasurementsResponse, error) {

	// prepare start time
	var startTime *time.Time = nil
	if req.Start != nil {
		startTimeConverted, err := ptypes.Timestamp(req.Start)
		if err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
		}
		startTime = &startTimeConverted
	}

	// prepare end time
	var endTime *time.Time = nil
	if req.End != nil {
		endTimeConverted, err := ptypes.Timestamp(req.End)
		if err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
		}
		endTime = &endTimeConverted
	}

	measurementsFromDb, err := db.GetMeasurements(req.DevEui, startTime, endTime, conf.Database)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	measurements := make([]*external.MeasurementListItem, len(measurementsFromDb))
	for i, element := range measurementsFromDb {
		timeProto, err := ptypes.TimestampProto(element.Time)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		channel, err := strconv.ParseInt(element.MeasuredChannels, 2, 64)
		measurements[i] = &external.MeasurementListItem{
			Time:          timeProto,
			Port:          int64(element.Port),
			Channel:       uint32(channel),
			ChannelCount:  int64(element.ChannelCount),
			Ct:            int64(element.ChannelType),
			Func:          int64(element.FuncNum),
			ChannelValues: mapInterfaceToMapFloat64(element.Channels),
		}
	}

	resp := external.GetMeasurementsResponse{
		NumberOfMeasurements: int64(len(measurements)),
		DevEui:               req.DevEui,
		Measurements:         measurements,
	}

	return &resp, nil
}

// Delete deletes the specified measurements.
func (a *MeasurementServerAPI) Delete(ctx context.Context, req *external.DeleteMeasurementsRequest) (*empty.Empty, error) {

	// prepare start time
	var startTime *time.Time = nil
	if req.Start != nil {
		startTimeConverted, err := ptypes.Timestamp(req.Start)
		if err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
		}
		startTime = &startTimeConverted
	}

	// prepare end time
	var endTime *time.Time = nil
	if req.End != nil {
		endTimeConverted, err := ptypes.Timestamp(req.End)
		if err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
		}
		endTime = &endTimeConverted
	}

	err := db.DeleteMeasurements(req.DevEui, startTime, endTime, conf.Database)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}
