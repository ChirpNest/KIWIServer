using System;

namespace ChirpNestCommunication.Models
{
    public class KellerDevice
    {
        public string Name { get; set; }

        public string Description { get; set; }

        public string EUI { get; set; }

        public string SerialNumber { get; set; }

        public string UniqueSerialNumber { get; set; }

        /// <summary>
        /// Looks like this: "5.5" OR "10.5" OR "5.10" ....
        /// </summary>
        public string DeviceType { get; set; }

        public DateTime? FirstMeasurement { get; set; }

        public DateTime? LastMeasurement { get; set; }

        public int NumberOfMeasurements { get; set; }
    }
}
