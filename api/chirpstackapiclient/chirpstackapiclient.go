package chirpstackapiclient

import (
	pb "github.com/brocaar/chirpstack-api/go/as/external/api"
)

// GetDevices returns all devices as returned from chirpstack-application-server
func GetDevices() []*pb.DeviceListItem {
	return getDevices()
}
