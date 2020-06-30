using System;
using System.Collections.Generic;
using Kiwi.Api;
using ChirpNestCommunication.Mapping;
using ChirpNestCommunication.Models;
using Grpc.Core;
using KellerAg.Shared.Entities.FileFormat;

namespace ChirpNestCommunication
{
    public class DeviceService : IDeviceService
    {
        private readonly IMapping<ListDeviceResponse, List<KellerDevice>> _mapper;

        public DeviceService(IMapping<ListDeviceResponse, List<KellerDevice>> mapper)
        {
            _mapper = mapper;
        }

        public string ApiUrl => "/api/devices";

        public List<KellerDevice> GetRegisteredDevices(Gateway gateway)
        {
            var uriBuilder = new UriBuilder("", gateway.GatewayIp, gateway.GatewayPort, ApiUrl);
            var channel = new Channel(uriBuilder.ToString(), ChannelCredentials.Insecure);

            var client = new Kiwi.Api.DeviceService.DeviceServiceClient(channel);

            var reply = client.List(new ListDeviceRequest());
            channel.ShutdownAsync().Wait();




            return _mapper.Map(reply);
        }
    }
}