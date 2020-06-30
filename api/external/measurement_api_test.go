package external

import (
	"context"
	"testing"

	"example.org/luksam/kiwi-server/db"
	"example.org/luksam/kiwi-server/test"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"

	pb "example.org/luksam/kiwi-server/apidefinition/go/external"
)

func (ts *APITestSuite) TestMeasurementApi() {
	assert := require.New(ts.T())

	conf = test.GetTestConfiguration()

	// insert measurements
	err := db.RunSQLStatementForUnitTest(sqlStatementInsertTestData, conf.Database)
	assert.NoError(err)

	api := NewMeasurementServerAPI()

	ts.T().Run("Get measurements", func(t *testing.T) {
		// get measurements via api
		assert := require.New(t)
		devEui := "009d6b0000c5d24f"
		measurementsResult, err := api.Get(context.Background(), &pb.GetMeasurementsRequest{
			DevEui: devEui,
			Start:  nil,
			End:    nil,
		})
		assert.NoError(err)

		// check api result

		assert.Equal(int64(4), measurementsResult.NumberOfMeasurements)
		assert.Equal(devEui, measurementsResult.DevEui)
		assert.Equal(4, len(measurementsResult.Measurements))

		measurement1 := measurementsResult.Measurements[0]
		measurement4 := measurementsResult.Measurements[3]

		measurement1Time, err := ptypes.Timestamp(measurement1.Time)
		assert.NoError(err)
		assert.Equal("2020-06-04 15:14:03.242229 +0000 UTC", measurement1Time.String())
		assert.Equal(int64(1), measurement1.Port)
		assert.Equal(uint32(0b0000000000010010), measurement1.Channel)
		assert.Equal(int64(2), measurement1.ChannelCount)
		assert.Equal(int64(3), measurement1.Ct)
		assert.Equal(int64(1), measurement1.Func)
		assert.Equal(2, len(measurement1.ChannelValues))
		assert.Equal(22.89892578125, measurement1.ChannelValues["TOB1"])
		assert.Equal(0.9291008114814758, measurement1.ChannelValues["P1"])

		measurement4Time, err := ptypes.Timestamp(measurement4.Time)
		assert.NoError(err)
		assert.Equal("2020-06-04 13:14:03.360548 +0000 UTC", measurement4Time.String())
		assert.Equal(int64(1), measurement4.Port)
		assert.Equal(uint32(0b0000000000010010), measurement4.Channel)
		assert.Equal(int64(2), measurement4.ChannelCount)
		assert.Equal(int64(3), measurement4.Ct)
		assert.Equal(int64(1), measurement4.Func)
		assert.Equal(2, len(measurement4.ChannelValues))
		assert.Equal(23.081787109375, measurement4.ChannelValues["TOB1"])
		assert.Equal(0.9293192028999329, measurement4.ChannelValues["P1"])
	})

	ts.T().Run("Delete measurements", func(t *testing.T) {
		assert := require.New(t)
		devEui := "009d6b0000c5d24f"

		// delete all measurements
		deleteResult, err := api.Delete(context.Background(), &pb.DeleteMeasurementsRequest{
			DevEui: devEui,
			Start:  nil,
			End:    nil,
		})
		assert.NoError(err)
		assert.Equal(&empty.Empty{}, deleteResult)

		// check if there are no measurements there anymore
		measurementsResult, err := api.Get(context.Background(), &pb.GetMeasurementsRequest{
			DevEui: devEui,
			Start:  nil,
			End:    nil,
		})
		assert.NoError(err)
		assert.Equal(int64(0), measurementsResult.NumberOfMeasurements)
	})
}

const sqlStatementInsertTestData = `
INSERT INTO public.device_up (id, received_at, dev_eui, device_name, application_id, application_name, frequency, dr, adr, f_cnt, f_port, tags, data, rx_info, object) VALUES ('8e594a58-a9b8-4f4e-9032-424f195b17be', '2020-06-04 11:48:39.374757+00', '\x009d6b0000c5d24f', 'ADT-Sam', 1, 'my-fancy-app-yeah', 868100000, 2, false, 6116, 4, '', '\x0c011300132f0000005c266b9c16409ea0d1502b', '[{"name": "", "rssi": -68, "loRaSNR": 11.5, "location": null, "uplinkID": "660cbcd4-2501-4549-8280-d5291245970c", "gatewayID": "fcc23dfffe0a89f9"}]', '{"func": 12, "port": 4, "payload": "0C011300132F0000005C266B9C16409EA0D1502B", "serial_number": 92, "battery_voltage": 4.9571309089660645, "sw_version_text": "19.47", "class_group_text": "19.00", "humidity_percentage": 43, "device_local_datetime": "2020-06-04 11:48:38", "battery_capacity_percentage": 80}');
INSERT INTO public.device_up (id, received_at, dev_eui, device_name, application_id, application_name, frequency, dr, adr, f_cnt, f_port, tags, data, rx_info, object) VALUES ('b298a161-37fe-4739-a23f-00d81dbaa962', '2020-06-04 13:14:03.360548+00', '\x009d6b0000c5d24f', 'ADT-Sam', 1, 'my-fancy-app-yeah', 867300000, 2, false, 6129, 1, '', '\x010300123f6de7dd41b8a780', '[{"name": "", "rssi": -67, "loRaSNR": 12, "location": null, "uplinkID": "1978b447-ecf3-4f2d-a74b-01fa90543b13", "gatewayID": "fcc23dfffe0a89f9"}]', '{"P1": 0.9293192028999329, "ct": 3, "TOB1": 23.081787109375, "func": 1, "port": 1, "channel": "0000000000010010", "payload": "010300123F6DE7DD41B8A780", "channelCount": 2}');
INSERT INTO public.device_up (id, received_at, dev_eui, device_name, application_id, application_name, frequency, dr, adr, f_cnt, f_port, tags, data, rx_info, object) VALUES ('314a7842-5b03-4e21-812c-d96b99434977', '2020-06-04 13:44:03.33505+00', '\x009d6b0000c5d24f', 'ADT-Sam', 1, 'my-fancy-app-yeah', 867500000, 2, false, 6130, 1, '', '\x010300123f6de77541b85100', '[{"name": "", "rssi": -65, "loRaSNR": 8.2, "location": null, "uplinkID": "53f3f65a-9975-46b6-bb79-42b7d98bc89f", "gatewayID": "fcc23dfffe0a89f9"}]', '{"P1": 0.9293130040168762, "ct": 3, "TOB1": 23.03955078125, "func": 1, "port": 1, "channel": "0000000000010010", "payload": "010300123F6DE77541B85100", "channelCount": 2}');
INSERT INTO public.device_up (id, received_at, dev_eui, device_name, application_id, application_name, frequency, dr, adr, f_cnt, f_port, tags, data, rx_info, object) VALUES ('fe4e5198-3d98-4556-b394-d360445218e5', '2020-06-04 14:44:03.96321+00', '\x009d6b0000c5d24f', 'ADT-Sam', 1, 'my-fancy-app-yeah', 867100000, 2, false, 6132, 1, '', '\x010300123f6de4b641b78780', '[{"name": "", "rssi": -66, "loRaSNR": 12.2, "location": null, "uplinkID": "7c8a3876-1d54-4c96-88f0-832fc29f35c8", "gatewayID": "fcc23dfffe0a89f9"}]', '{"P1": 0.9292711019515991, "ct": 3, "TOB1": 22.941162109375, "func": 1, "port": 1, "channel": "0000000000010010", "payload": "010300123F6DE4B641B78780", "channelCount": 2}');
INSERT INTO public.device_up (id, received_at, dev_eui, device_name, application_id, application_name, frequency, dr, adr, f_cnt, f_port, tags, data, rx_info, object) VALUES ('5a9ac541-1574-483a-a28a-eac1cf6b377b', '2020-06-04 15:14:03.242229+00', '\x009d6b0000c5d24f', 'ADT-Sam', 1, 'my-fancy-app-yeah', 868100000, 2, false, 6133, 1, '', '\x010300123f6dd98d41b73100', '[{"name": "", "rssi": -61, "loRaSNR": 10.2, "location": null, "uplinkID": "e36d11b3-48c6-4e6f-bf6a-744b0a7d043e", "gatewayID": "fcc23dfffe0a89f9"}]', '{"P1": 0.9291008114814758, "ct": 3, "TOB1": 22.89892578125, "func": 1, "port": 1, "channel": "0000000000010010", "payload": "010300123F6DD98D41B73100", "channelCount": 2}');
`
