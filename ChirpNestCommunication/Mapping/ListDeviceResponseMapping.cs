using System;
using System.Collections.Generic;
using System.Text;
using Kiwi.Api;
using AutoMapper;
using ChirpNestCommunication.Models;
using KellerAg.Shared.Entities.FileFormat;

namespace ChirpNestCommunication.Mapping
{
    public class ListDeviceResponseMapping : IMapping<ListDeviceResponse, List<KellerDevice>>
    {

        private IMapper _mapper;

        public ListDeviceResponseMapping()
        {
            InitializeConfig();
        }

        public List<KellerDevice> Map(ListDeviceResponse source)
        {
            var devicesList = new List<KellerDevice>();
            foreach (var device in source.Devices)
            {
                devicesList.Add(_mapper.Map<KellerDevice>(device));
            }

            return devicesList;
            //return _mapper.Map<List<KellerDevice>>(source);
        }

        private void InitializeConfig()
        {
            var config = new MapperConfiguration(cfg =>
            {
                cfg.CreateMap<DeviceListItem, KellerDevice>()
                    .ForMember(dest => dest.Name, opt => opt.MapFrom(src => src.Name))
                    .ForMember(dest => dest.SerialNumber, opt => opt.MapFrom(src => src.SerialNumber.ToString()))
                    .ForMember(dest => dest.UniqueSerialNumber, opt => opt.MapFrom(src => MeasurementFileFormatHelper.GenerateUniqueSerialNumber(src.DeviceType, src.SerialNumber.ToString(), src.DevEui)))
                    .ForMember(dest => dest.DeviceType, opt => opt.MapFrom(src => src.DeviceType))
                    .ForMember(dest => dest.FirstMeasurement, opt =>
                    {
                        opt.MapFrom(src => src.FirstMeasurementTime.ToDateTime());
                        opt.NullSubstitute(DateTime.MinValue);
                    })
                    .ForMember(dest => dest.LastMeasurement, opt =>
                    {
                        opt.MapFrom(src => src.LastMeasurementTime.ToDateTime());
                        opt.NullSubstitute(DateTime.MaxValue);
                    })
                    .ForMember(dest => dest.Description, opt => opt.MapFrom(src => src.Description))
                    .ForMember(dest => dest.NumberOfMeasurements, opt => opt.MapFrom(src => (int)src.NumberOfMeasurements))
                    .ForMember(dest => dest.EUI, opt => opt.MapFrom(src => src.DevEui));
            });
            _mapper = config.CreateMapper();
        }
    }
}
