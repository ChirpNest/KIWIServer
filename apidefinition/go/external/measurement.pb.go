// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: measurement.proto

package external

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeleteMeasurementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Delete only measurements after this timestamp.
	Start *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	// Delete only measurements before this timestamp.
	End *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DeleteMeasurementsRequest) Reset() {
	*x = DeleteMeasurementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMeasurementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMeasurementsRequest) ProtoMessage() {}

func (x *DeleteMeasurementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMeasurementsRequest.ProtoReflect.Descriptor instead.
func (*DeleteMeasurementsRequest) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteMeasurementsRequest) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *DeleteMeasurementsRequest) GetStart() *timestamp.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *DeleteMeasurementsRequest) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type GetMeasurementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Get only measurements after this timestamp.
	Start *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	// Get only measurements before this timestamp.
	End *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetMeasurementsRequest) Reset() {
	*x = GetMeasurementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeasurementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasurementsRequest) ProtoMessage() {}

func (x *GetMeasurementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasurementsRequest.ProtoReflect.Descriptor instead.
func (*GetMeasurementsRequest) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{1}
}

func (x *GetMeasurementsRequest) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *GetMeasurementsRequest) GetStart() *timestamp.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetMeasurementsRequest) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type GetMeasurementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of measurements available within the result-set.
	NumberOfMeasurements int64 `protobuf:"varint,1,opt,name=number_of_measurements,json=numberOfMeasurements,proto3" json:"number_of_measurements,omitempty"`
	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,2,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// All devices.
	Measurements []*MeasurementListItem `protobuf:"bytes,3,rep,name=measurements,proto3" json:"measurements,omitempty"`
}

func (x *GetMeasurementsResponse) Reset() {
	*x = GetMeasurementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeasurementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasurementsResponse) ProtoMessage() {}

func (x *GetMeasurementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasurementsResponse.ProtoReflect.Descriptor instead.
func (*GetMeasurementsResponse) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{2}
}

func (x *GetMeasurementsResponse) GetNumberOfMeasurements() int64 {
	if x != nil {
		return x.NumberOfMeasurements
	}
	return 0
}

func (x *GetMeasurementsResponse) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *GetMeasurementsResponse) GetMeasurements() []*MeasurementListItem {
	if x != nil {
		return x.Measurements
	}
	return nil
}

type MeasurementListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time of the measurement.
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Port.
	Port int64 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Channel (in last 2 bytes, e.g. 00000000 00000000 10100000 00010010).
	Channel uint32 `protobuf:"fixed32,3,opt,name=channel,proto3" json:"channel,omitempty"`
	// Number of channels.
	ChannelCount int64 `protobuf:"varint,4,opt,name=channel_count,json=channelCount,proto3" json:"channel_count,omitempty"`
	// ct (as decoded)
	Ct int64 `protobuf:"varint,5,opt,name=ct,proto3" json:"ct,omitempty"`
	// func (as decoded)
	Func int64 `protobuf:"varint,6,opt,name=func,proto3" json:"func,omitempty"`
	// Measured value of each channel.
	ChannelValues map[string]float64 `protobuf:"bytes,7,rep,name=channel_values,json=channelValues,proto3" json:"channel_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *MeasurementListItem) Reset() {
	*x = MeasurementListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasurementListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementListItem) ProtoMessage() {}

func (x *MeasurementListItem) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementListItem.ProtoReflect.Descriptor instead.
func (*MeasurementListItem) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{3}
}

func (x *MeasurementListItem) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MeasurementListItem) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *MeasurementListItem) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *MeasurementListItem) GetChannelCount() int64 {
	if x != nil {
		return x.ChannelCount
	}
	return 0
}

func (x *MeasurementListItem) GetCt() int64 {
	if x != nil {
		return x.Ct
	}
	return 0
}

func (x *MeasurementListItem) GetFunc() int64 {
	if x != nil {
		return x.Func
	}
	return 0
}

func (x *MeasurementListItem) GetChannelValues() map[string]float64 {
	if x != nil {
		return x.ChannelValues
	}
	return nil
}

var File_measurement_proto protoreflect.FileDescriptor

var file_measurement_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x69, 0x77, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65,
	0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x55, 0x49,
	0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x22, 0x91, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x45, 0x55, 0x49, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x16, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x14, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75,
	0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x55, 0x49, 0x12,
	0x41, 0x0a, 0x0c, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x69, 0x77, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xd7, 0x02, 0x0a, 0x13, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x63, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x75, 0x6e, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x75, 0x6e,
	0x63, 0x12, 0x57, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x69, 0x77, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xf1, 0x01, 0x0a,
	0x12, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x6b, 0x69, 0x77,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6b,
	0x69, 0x77, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x5f,
	0x65, 0x75, 0x69, 0x7d, 0x12, 0x6a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23,
	0x2e, 0x6b, 0x69, 0x77, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d,
	0x42, 0x3a, 0x5a, 0x38, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x6b, 0x73, 0x61, 0x6d, 0x2f, 0x6b, 0x69, 0x77, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_measurement_proto_rawDescOnce sync.Once
	file_measurement_proto_rawDescData = file_measurement_proto_rawDesc
)

func file_measurement_proto_rawDescGZIP() []byte {
	file_measurement_proto_rawDescOnce.Do(func() {
		file_measurement_proto_rawDescData = protoimpl.X.CompressGZIP(file_measurement_proto_rawDescData)
	})
	return file_measurement_proto_rawDescData
}

var file_measurement_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_measurement_proto_goTypes = []interface{}{
	(*DeleteMeasurementsRequest)(nil), // 0: kiwi.api.DeleteMeasurementsRequest
	(*GetMeasurementsRequest)(nil),    // 1: kiwi.api.GetMeasurementsRequest
	(*GetMeasurementsResponse)(nil),   // 2: kiwi.api.GetMeasurementsResponse
	(*MeasurementListItem)(nil),       // 3: kiwi.api.MeasurementListItem
	nil,                               // 4: kiwi.api.MeasurementListItem.ChannelValuesEntry
	(*timestamp.Timestamp)(nil),       // 5: google.protobuf.Timestamp
	(*empty.Empty)(nil),               // 6: google.protobuf.Empty
}
var file_measurement_proto_depIdxs = []int32{
	5, // 0: kiwi.api.DeleteMeasurementsRequest.start:type_name -> google.protobuf.Timestamp
	5, // 1: kiwi.api.DeleteMeasurementsRequest.end:type_name -> google.protobuf.Timestamp
	5, // 2: kiwi.api.GetMeasurementsRequest.start:type_name -> google.protobuf.Timestamp
	5, // 3: kiwi.api.GetMeasurementsRequest.end:type_name -> google.protobuf.Timestamp
	3, // 4: kiwi.api.GetMeasurementsResponse.measurements:type_name -> kiwi.api.MeasurementListItem
	5, // 5: kiwi.api.MeasurementListItem.time:type_name -> google.protobuf.Timestamp
	4, // 6: kiwi.api.MeasurementListItem.channel_values:type_name -> kiwi.api.MeasurementListItem.ChannelValuesEntry
	1, // 7: kiwi.api.MeasurementService.Get:input_type -> kiwi.api.GetMeasurementsRequest
	0, // 8: kiwi.api.MeasurementService.Delete:input_type -> kiwi.api.DeleteMeasurementsRequest
	2, // 9: kiwi.api.MeasurementService.Get:output_type -> kiwi.api.GetMeasurementsResponse
	6, // 10: kiwi.api.MeasurementService.Delete:output_type -> google.protobuf.Empty
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_measurement_proto_init() }
func file_measurement_proto_init() {
	if File_measurement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_measurement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMeasurementsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_measurement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeasurementsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_measurement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeasurementsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_measurement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasurementListItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_measurement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_measurement_proto_goTypes,
		DependencyIndexes: file_measurement_proto_depIdxs,
		MessageInfos:      file_measurement_proto_msgTypes,
	}.Build()
	File_measurement_proto = out.File
	file_measurement_proto_rawDesc = nil
	file_measurement_proto_goTypes = nil
	file_measurement_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeasurementServiceClient is the client API for MeasurementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeasurementServiceClient interface {
	// Get returns the requested measurements.
	Get(ctx context.Context, in *GetMeasurementsRequest, opts ...grpc.CallOption) (*GetMeasurementsResponse, error)
	// Delete deletes the specified measurements.
	Delete(ctx context.Context, in *DeleteMeasurementsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type measurementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasurementServiceClient(cc grpc.ClientConnInterface) MeasurementServiceClient {
	return &measurementServiceClient{cc}
}

func (c *measurementServiceClient) Get(ctx context.Context, in *GetMeasurementsRequest, opts ...grpc.CallOption) (*GetMeasurementsResponse, error) {
	out := new(GetMeasurementsResponse)
	err := c.cc.Invoke(ctx, "/kiwi.api.MeasurementService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementServiceClient) Delete(ctx context.Context, in *DeleteMeasurementsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kiwi.api.MeasurementService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasurementServiceServer is the server API for MeasurementService service.
type MeasurementServiceServer interface {
	// Get returns the requested measurements.
	Get(context.Context, *GetMeasurementsRequest) (*GetMeasurementsResponse, error)
	// Delete deletes the specified measurements.
	Delete(context.Context, *DeleteMeasurementsRequest) (*empty.Empty, error)
}

// UnimplementedMeasurementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMeasurementServiceServer struct {
}

func (*UnimplementedMeasurementServiceServer) Get(context.Context, *GetMeasurementsRequest) (*GetMeasurementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMeasurementServiceServer) Delete(context.Context, *DeleteMeasurementsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterMeasurementServiceServer(s *grpc.Server, srv MeasurementServiceServer) {
	s.RegisterService(&_MeasurementService_serviceDesc, srv)
}

func _MeasurementService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeasurementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiwi.api.MeasurementService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementServiceServer).Get(ctx, req.(*GetMeasurementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMeasurementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiwi.api.MeasurementService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementServiceServer).Delete(ctx, req.(*DeleteMeasurementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeasurementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kiwi.api.MeasurementService",
	HandlerType: (*MeasurementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _MeasurementService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MeasurementService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "measurement.proto",
}