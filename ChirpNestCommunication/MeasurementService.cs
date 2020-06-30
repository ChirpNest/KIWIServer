using System;
using Kiwi.Api;
using ChirpNestCommunication.Mapping;
using ChirpNestCommunication.Models;
using Grpc.Core;
using KellerAg.Shared.Entities.FileFormat;

namespace ChirpNestCommunication
{
    public class MeasurementService : IMeasurementService
    {
        private readonly IMapping<GetMeasurementsResponse, MeasurementFileFormat> _mapper;

        public MeasurementService(IMapping<GetMeasurementsResponse, MeasurementFileFormat> mapper)
        {
            _mapper = mapper;
        }

        public string ApiUrl => "api/measurements";

        public MeasurementFileFormat GetMeasurements(Gateway gateway, KellerDevice device)
        {
            var reply = SendGetMeasurementRequest(gateway, device);
            var kellerFormat = _mapper.Map(reply);

            kellerFormat.Header.SerialNumber = device.SerialNumber;
            kellerFormat.Header.DeviceName = device.Name;
            kellerFormat.Header.DeviceType = device.DeviceType;
            kellerFormat.Header.UniqueSerialNumber = MeasurementFileFormatHelper.GenerateUniqueSerialNumber(kellerFormat.Header, device.EUI);
            kellerFormat.Header.RecordId = MeasurementFileFormatHelper.GenerateRecordId(kellerFormat.Header);
            kellerFormat.Header.CustomAttributes = new MeasurementFileFormatCustomAttributes
            {
                RecordName = MeasurementFileFormatHelper.GenerateDefaultRecordName(kellerFormat.Header),
                RecordNotes = string.Empty,
            };

            return kellerFormat;
        }

        public bool RemoveMeasurements(Gateway gateway, KellerDevice device)
        {
            SendDeleteMeasurementRequest(gateway, device);
            return true;
        }

        private GetMeasurementsResponse SendGetMeasurementRequest(Gateway gateway, KellerDevice device)
        {
            var uriBuilder = new UriBuilder("", gateway.GatewayIp, gateway.GatewayPort, ApiUrl);
            var channel = new Channel(uriBuilder.ToString(), ChannelCredentials.Insecure);

            var client = new Kiwi.Api.MeasurementService.MeasurementServiceClient(channel);
            var request = new GetMeasurementsRequest {DevEui = device.EUI};
            var reply = client.Get(request);
            channel.ShutdownAsync().Wait();
            return reply;
        }
        private void SendDeleteMeasurementRequest(Gateway gateway, KellerDevice device)
        {
            var uriBuilder = new UriBuilder("", gateway.GatewayIp, gateway.GatewayPort, ApiUrl);
            var channel = new Channel(uriBuilder.ToString(), ChannelCredentials.Insecure);

            var client = new Kiwi.Api.MeasurementService.MeasurementServiceClient(channel);
            var request = new DeleteMeasurementsRequest { DevEui = device.EUI };
            client.Delete(request);
            channel.ShutdownAsync().Wait();
        }
    }
}