package chirpstackapiclient

import (
	"context"

	"example.org/luksam/kiwi-server/config"
	pb "github.com/brocaar/chirpstack-api/go/as/external/api"
	"google.golang.org/grpc"
)

func getDevices() []*pb.DeviceListItem {

	auth := getAuthenticationDialOption()

	config := config.GetConfiguration()

	conn, err := grpc.Dial(config.ChirpStackApplicationServer.Address, auth, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewDeviceServiceClient(conn)

	request := &pb.ListDeviceRequest{
		Limit: 1000,
	}

	resp, err := client.List(context.Background(), request)
	if err != nil {
		panic(err)
	}
	resultDevices := resp.Result

	return resultDevices
}
