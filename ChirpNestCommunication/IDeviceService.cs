using System.Collections.Generic;
using ChirpNestCommunication.Models;

namespace ChirpNestCommunication
{
    public interface IDeviceService : IServiceBase
    {
        List<KellerDevice> GetRegisteredDevices(Gateway gateway);
    }
}
