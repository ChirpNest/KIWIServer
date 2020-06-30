using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using KellerAg.Shared.Entities.Channel;
using KellerAg.Shared.Entities.FileFormat;
using KIWIDesktop.Models;
using OxyPlot;
using OxyPlot.Axes;

namespace KIWIDesktop.Services
{
    public static class OxyPlotSeriesGenerator
    {
        public static List<OxyPlotSeries> GenerateOxyPlotSeries(MeasurementFileFormat file)
        {
            var series = new List<OxyPlotSeries>();

            var counter = 0;
            foreach (var measurementDefinition in file.Header.MeasurementDefinitionsInBody)
            {
                series.Add(new OxyPlotSeries
                (
                    ChannelInfo.GetMeasurementDefinitionName(measurementDefinition),
                    ChannelInfo.GetUnitType(measurementDefinition),
                    file.Body.Select(x => new DataPoint(DateTimeAxis.ToDouble(x.Time), x.Values[counter].HasValue ? x.Values[counter].Value : 0)).ToList()
                    
                ));

                counter++;
            }


            return series;
        }
    }
}
