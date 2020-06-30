using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using KellerAg.Shared.Entities.FileFormat;
using Newtonsoft.Json;
using NLog;
using NLog.Fluent;

namespace KIWIDesktop.Services
{
    public class KellerFileService : IKellerFileService
    {
        private static readonly string DefaultSavingPath =
            Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData), "ChirpNest", "data");

        private static readonly Logger Logger = LogManager.GetCurrentClassLogger();

        public MeasurementFileFormat ReadFileFormat(string uniqueId)
        {
            var file = FindFileFromRecordId(uniqueId);
            var fileFormat = GetObjectFromFile(file);

            return fileFormat;
        }

        public List<MeasurementFileFormatHeader> FindFilesFromDevice(string uniqueDeviceId)
        {
            return ReadTableOfContent().Where(x => x.UniqueSerialNumber == uniqueDeviceId).ToList();
        }

        public List<MeasurementFileFormatHeader> ReadTableOfContent()
        {
            var files = Directory.EnumerateFiles(DefaultSavingPath, "*.json");

            return files.Select(file => GetObjectFromFile(file).Header).ToList();
        }

        public void WriteFileFormat(MeasurementFileFormat file)
        {
            var json = FileFormatToJson(file);
            WriteJsonToFile(Path.Combine(DefaultSavingPath, $"KIWIDesktop_{file.Header.RecordId}.json"), json);
        }

        public void WriteFileCombinedWith(MeasurementFileFormat file, List<MeasurementFileFormatHeader> measurementsToCombine)
        {
            foreach (var measurement in measurementsToCombine)
            {
                if (measurement.UniqueSerialNumber != file.Header.UniqueSerialNumber || 
                    file.Header?.RemoteTransmissionUnitInfo?.ConnectionTypeId != measurement.RemoteTransmissionUnitInfo?.ConnectionTypeId) 
                {
                    continue;
                }
                file.Body.AddRange(ReadFileFormat(measurement.RecordId).Body);
                RemoveFileFormat(measurement.RecordId);
            }

            file.Body = file.Body.GroupBy(x => x.Time).Select(x => x.First()).OrderBy(x => x.Time).ToList();
            file.Header.FirstMeasurementUTC = file.Body.FirstOrDefault()?.Time ?? DateTime.MinValue;
            file.Header.FirstMeasurementUTC = file.Body.LastOrDefault()?.Time ?? DateTime.MaxValue;

            WriteFileFormat(file);
        }

        public void RemoveFileFormat(string uniqueId)
        {
            if (string.IsNullOrWhiteSpace(uniqueId))
            {
                return;
            }
            var file = FindFileFromRecordId(uniqueId);
            if (file != string.Empty)
            {
                try
                {
                    File.Delete(file);
                }
                catch (Exception e)
                {
                    Logger.Warn(e, "Failed to remove file with measurements");
                }
            }
        }

        private string FileFormatToJson(MeasurementFileFormat fileFormat)
        {
            return JsonConvert.SerializeObject(fileFormat);
        }
        private static void WriteJsonToFile(string filePath, string json)
        {
            CreateFolder(filePath);
            using (var file = new FileStream(filePath, FileMode.Create, FileAccess.Write, FileShare.None))
            {
                using (var writer = new StreamWriter(file, Encoding.Unicode))
                {
                    writer.Write(json);
                }
            }
        }

        private static void CreateFolder(string path)
        {
            var dir = Path.GetDirectoryName(path);
            if (!string.IsNullOrEmpty(dir) && !Directory.Exists(dir))
            {
                Directory.CreateDirectory(dir);
            }
        }
        private string FindFileFromRecordId(string recordId)
        {
            var files = Directory.EnumerateFiles(DefaultSavingPath, "*.json");
            var requestedFile = string.Empty;
            foreach (var file in files)
            {
                try
                {
                    var content = File.ReadAllText(file);
                    if (content.Contains(recordId))
                    {
                        if (JsonConvert.DeserializeObject<MeasurementFileFormat>(content).Header.RecordId == recordId)
                        {
                            requestedFile = file;
                            break;
                        }
                    }
                }
                catch (Exception e)
                {
                    Logger.Warn(e, "Failed to read file with measurements");
                }
            }

            return requestedFile;
        }
        private MeasurementFileFormat GetObjectFromFile(string filePath)
        {
            var fileFormat = new MeasurementFileFormat();
            try
            {
                var content = File.ReadAllText(filePath);
                fileFormat = JsonConvert.DeserializeObject<MeasurementFileFormat>(content);
            }
            catch (Exception e)
            {
                Logger.Warn(e, "Failed to read file with measurements");
            }

            return fileFormat;
        }
    }
}