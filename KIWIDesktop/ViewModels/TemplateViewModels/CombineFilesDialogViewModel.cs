using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using ChirpNestCommunication.Models;
using KellerAg.Shared.Entities.FileFormat;

namespace KIWIDesktop.ViewModels.TemplateViewModels
{
    public class CombineFilesDialogViewModel
    {
        public CombineFilesDialogViewModel()
        {

        }

        public List<MeasurementFileFormatHeader> FilesToCombine { get; set; }

        public KellerDevice Device { get; set; }
    }
}
