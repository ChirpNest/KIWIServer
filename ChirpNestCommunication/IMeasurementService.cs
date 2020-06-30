using System.Collections.Generic;
using ChirpNestCommunication.Models;
using KellerAg.Shared.Entities.FileFormat;

namespace ChirpNestCommunication
{
    public interface IMeasurementService : IServiceBase
    {
        MeasurementFileFormat GetMeasurements(Gateway gateway, KellerDevice device);

        bool RemoveMeasurements(Gateway gateway, KellerDevice device);
    }
}
