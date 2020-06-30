using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Kiwi.Api;
using AutoMapper;
using ChirpNestCommunication.Models;
using KellerAg.Shared.Entities.Channel;
using KellerAg.Shared.Entities.FileFormat;

namespace ChirpNestCommunication.Mapping
{
    public class GetMeasurementsResponseMapping : IMapping<GetMeasurementsResponse, MeasurementFileFormat>
    {

        private IMapper _mapper;

        public GetMeasurementsResponseMapping()
        {
            InitializeConfig();
        }

        public MeasurementFileFormat Map(GetMeasurementsResponse source)
        {
            var connectionType = source.Measurements.FirstOrDefault()?.Ct;
            RemoteTransmissionUnitInfo remoteTransmissionUnitInfo = null;
            if (connectionType.HasValue)
            {
                remoteTransmissionUnitInfo = new RemoteTransmissionUnitInfo
                {
                    ConnectionTypeId = (int)connectionType
                };

            }

            var kellerFileFormat = new MeasurementFileFormat();
            kellerFileFormat.Body = new List<Measurements>();

            var amountOfChannels = source.Measurements.FirstOrDefault()?.ChannelValues.Count ?? 0;

            ChannelInfo[] measurementDefinitionInBody = null;
            foreach (var measurement in source.Measurements)
            {
                if (measurementDefinitionInBody == null)
                {
                    var allChannels = ChannelInfo.GetChannels();
                    measurementDefinitionInBody = new ChannelInfo[amountOfChannels];
                    var counter = 0;
                    foreach (var measurementValue in measurement.ChannelValues)
                    {
                        measurementDefinitionInBody[counter] = allChannels.FirstOrDefault(x => x.Name == measurementValue.Key) ??
                                                               allChannels.Single(x => x.ChannelType == ChannelType.Undefined);
                        counter++;
                    }
                }

                var kellerMeasurement = new Measurements
                {
                    Time = measurement.Time.ToDateTime(),
                    Values = new double?[amountOfChannels]
                };
                var keyPositionMapping = new Dictionary<string, int>();
                foreach (var key in measurement.ChannelValues.Keys)
                {
                    var index = Array.IndexOf(measurementDefinitionInBody, measurementDefinitionInBody.FirstOrDefault(x => x.Name == key));
                    keyPositionMapping.Add(key, index);
                }

                foreach (var channels in measurement.ChannelValues)
                {
                    kellerMeasurement.Values[keyPositionMapping[channels.Key]] = channels.Value;
                }
                kellerFileFormat.Body.Add(kellerMeasurement);
            }

            kellerFileFormat.Body = kellerFileFormat.Body.OrderBy(x => x.Time).ToList();
            kellerFileFormat.Header = new MeasurementFileFormatHeader
            {
                MeasurementDefinitionsInBody = measurementDefinitionInBody?.Select(x => x.MeasurementDefinitionId).ToArray() ?? new int[0],
                CreationDateTimeUTC = DateTime.UtcNow,
                CreationDateTimeDeviceTime = DateTime.UtcNow,
                IanaTimeZoneName = "UTC",
                FirstMeasurementUTC = kellerFileFormat.Body.FirstOrDefault()?.Time ?? DateTime.MinValue,
                LastMeasurementUTC = kellerFileFormat.Body.LastOrDefault()?.Time ?? DateTime.MinValue,
                CreationOrigin = Origin.Script,
                IsBodyCompressed = false,
                RemoteTransmissionUnitInfo = remoteTransmissionUnitInfo,
            };


            return kellerFileFormat;
            //return _mapper.Map<MeasurementFileFormat>(source);
        }

        private void InitializeConfig()
        {
            var config = new MapperConfiguration(cfg =>
                {
                    cfg.CreateMap<GetMeasurementsResponse, MeasurementFileFormat>();
                });
            _mapper = config.CreateMapper();
        }
    }
}
