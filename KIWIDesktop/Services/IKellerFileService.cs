using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using KellerAg.Shared.Entities.FileFormat;

namespace KIWIDesktop.Services
{
    public interface IKellerFileService
    {
        MeasurementFileFormat ReadFileFormat(string uniqueId);

        List<MeasurementFileFormatHeader> FindFilesFromDevice(string uniqueDeviceId);

        List<MeasurementFileFormatHeader> ReadTableOfContent();

        void WriteFileFormat(MeasurementFileFormat file);

        void WriteFileCombinedWith(MeasurementFileFormat file, List<MeasurementFileFormatHeader> measurementsToCombine);

        void RemoveFileFormat(string uniqueId);
    }
}
