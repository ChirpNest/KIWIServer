using System;
using System.IO;
using System.Text;
using ClosedXML.Excel;
using KellerAg.Shared.Entities.Channel;
using KellerAg.Shared.Entities.FileFormat;
using Newtonsoft.Json;

namespace KIWIDesktop.Services
{
    public class ExportService : IExportService
    {
        public void ExportKolibriFormat(MeasurementFileFormat file, string filePath)
        {
            try
            {
                var output = JsonConvert.SerializeObject(file);
                File.WriteAllText(filePath, output);
            }
            catch (DirectoryNotFoundException e)
            {
                throw new DirectoryNotFoundException("Export directory not found.", e);
            }
            catch (UnauthorizedAccessException e)
            {
                throw new UnauthorizedAccessException("Permission denied for export directory", e);
            }
        }

        public void ExportExcel(MeasurementFileFormat file, string filePath)
        {
            using (var workbook = new XLWorkbook())
            {
                var worksheet = workbook.Worksheets.Add("MeasurementData");
                var table = CreateTable(file);
                for(int i = 0; i < table.GetLength(0); i++)
                {
                    for (int j = 0; j < table.GetLength(1); j++)
                    {
                        worksheet.Cell(i+1, j+1).Value = table[i, j];
                    }
                }
                workbook.SaveAs(filePath);
            }
        }


        public void ExportCsv(MeasurementFileFormat file, string filePath)
        {
            var stringBuilder = new StringBuilder();
            var table = CreateTable(file);
            for (var i = 0; i < table.GetLength(0); i++)
            {
                stringBuilder.Append(table[i, 0]);

                for (var j = 1; j < table.GetLength(1); j++)
                {
                    stringBuilder.Append($",{table[i,j]}");
                }

                stringBuilder.AppendLine();
            }

            using (var streamWriter = new StreamWriter(filePath))
            {
                streamWriter.WriteLine(stringBuilder.ToString());
            }
        }

        private static object[,] CreateTable(MeasurementFileFormat file)
        {
            var channelAmount = file.Header.MeasurementDefinitionsInBody.Length;
            var measurementAmount = file.Body.Count;
            var dataTable = new object[measurementAmount + 1, channelAmount + 1];
            
            for (var i = 0; i < channelAmount; i++)
            {
                dataTable[0, i+1] = ChannelInfo.GetMeasurementDefinitionName(file.Header.MeasurementDefinitionsInBody[i]);
            }
            for (var i = 0; i < measurementAmount; i++)
            {
                dataTable[i + 1, 0] = file.Body[i].Time;
            }

            for (var i = 1; i < measurementAmount; i++)
            {
                for (var j = 1; j < channelAmount; j++)
                {
                    dataTable[i, j] = file.Body[i].Values[j];
                }
            }
            return dataTable;

        }
    }
}