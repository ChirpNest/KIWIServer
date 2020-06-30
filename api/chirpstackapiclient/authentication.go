package chirpstackapiclient

import (
	"context"

	"example.org/luksam/kiwi-server/config"
	pb "github.com/brocaar/chirpstack-api/go/as/external/api"
	"google.golang.org/grpc"
)

func getAuthenticationDialOption() grpc.DialOption {
	jwt := getJwt()
	cred := simpleJwtCredentials{
		jsonWebToken: jwt,
	}
	return grpc.WithPerRPCCredentials(cred)
}

func getJwt() string {

	config := config.GetConfiguration()

	conn, err := grpc.Dial(config.ChirpStackApplicationServer.Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewInternalServiceClient(conn)

	request := &pb.LoginRequest{
		Username: config.ChirpStackApplicationServer.Username,
		Password: config.ChirpStackApplicationServer.Password,
	}

	myContext := context.Background()

	resp, err := client.Login(myContext, request)
	if err != nil {
		panic(err)
	}
	resultJwt := resp.Jwt

	defer conn.Close()

	return resultJwt
}

// implements credentials.PerRPCCredentials
type simpleJwtCredentials struct {
	jsonWebToken string
}

func (cred simpleJwtCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer" + " " + cred.jsonWebToken,
	}, nil
}

func (simpleJwtCredentials) RequireTransportSecurity() bool {
	return false
}
