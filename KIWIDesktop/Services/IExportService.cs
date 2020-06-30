using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using KellerAg.Shared.Entities.FileFormat;

namespace KIWIDesktop.Services
{
    public interface IExportService
    {
        void ExportKolibriFormat(MeasurementFileFormat file, string filePath);

        void ExportExcel(MeasurementFileFormat file, string filePath);

        void ExportCsv(MeasurementFileFormat file, string filePath);
    }
}
