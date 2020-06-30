using System;
using System.Collections.Generic;
using System.Linq;
using Kiwi.Api;
using ChirpNestCommunication.Mapping;
using Google.Protobuf.WellKnownTypes;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Shouldly;

namespace ChirpNestCommunication.Test.Mapping
{
    [TestClass]
    public class GetMeasurementsResponseMappingTest
    {
        private GetMeasurementsResponseMapping _testee;


        [TestInitialize]
        public void Initialize()
        {
            _testee = new GetMeasurementsResponseMapping();
        }

        [TestMethod]
        public void Map_WhenTwoDevicesAreRegistered_ThenTheDevicesAreReturnedAsAList()
        {
            var testObject = new GetMeasurementsResponse
            {
                NumberOfMeasurements = 2,
                DevEui = "123",
                Measurements =
                {
                    new MeasurementListItem
                    {
                        ChannelCount = 2,
                        Ct = 2,
                        Func = 1,
                        Port = 1,
                        Time = new Timestamp
                        {
                            Seconds = 60
                        },
                        Channel = 1,
                        ChannelValues = { new Dictionary<string, double>
                        {
                            { "TOB1", 15 },
                            { "P1", 0.9 }
                        }}
                    },
                    new MeasurementListItem
                    {
                        ChannelCount = 2,
                        Ct = 2,
                        Func = 1,
                        Port = 1,
                        Time = new Timestamp
                        {
                            Seconds = 120
                        },
                        Channel = 1,
                        ChannelValues = { new Dictionary<string, double>
                        {
                            { "TOB1", 16 },
                            { "P1", 1 }
                        }}
                    }
                }
            };
            var result = _testee.Map(testObject);
            result.Body.Count.ShouldBe(2);
            var resultFirstMeasurement = result.Body.FirstOrDefault();
            resultFirstMeasurement.ShouldNotBeNull();
            resultFirstMeasurement.Time.ShouldBe(new DateTime(1970, 1, 1, 0, 1, 0));
        }
    }
}
