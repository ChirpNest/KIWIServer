using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.IO;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Input;
using ChirpNestCommunication;
using ChirpNestCommunication.Models;
using GalaSoft.MvvmLight.CommandWpf;
using KellerAg.Shared.Entities.FileFormat;
using KIWIDesktop.Annotations;
using KIWIDesktop.Models;
using KIWIDesktop.Navigation;
using KIWIDesktop.Services;
using KIWIDesktop.ViewModels.TemplateViewModels;
using KIWIDesktop.Views.Templates;
using MaterialDesignThemes.Wpf;
using Microsoft.Win32;
using NLog;
using OxyPlot;

namespace KIWIDesktop.ViewModels
{ 
    public class VisualizeViewModel : IViewModel
    {
        private readonly IKellerFileService _fileService;
        private readonly IMeasurementService _measurementService;
        private readonly IViewNavigationService _navigationService;
        private readonly IExportService _exportService;

        public VisualizeViewModel(IKellerFileService fileService, 
            IMeasurementService measurementService, 
            IViewNavigationService navigationService,
            IExportService exportService)
        {
            _fileService = fileService;
            _measurementService = measurementService;
            _navigationService = navigationService;
            _exportService = exportService;
            SelectedRecord.Instance.SelectedRecordChanged += HandleChangedRecord;
        }
        public ICommand ExportCommand => new RelayCommand(Export);

        public ICommand DeleteLocalCommand => new RelayCommand(DeleteLocal);

        public ICommand DownloadCommand => new RelayCommand(Download);

        public ICommand DownloadAnDeleteCommand => new RelayCommand(DownloadAnDelete);

        public ICommand DeleteOnRemoteCommand => new RelayCommand(DeleteOnRemote);

        public string PageTitle { get; private set; }

        public MeasurementFileFormat File { get; private set; }

        public KellerDevice KellerDevice{ get; private set; }

        public Gateway Gateway{ get; private set; }

        public bool IsRecordSavedLocally { get; private set; }

        public ObservableCollection<OxyPlotSeries> DataSeries { get; private set; }

        private static readonly Logger Logger = LogManager.GetCurrentClassLogger();

        private void HandleChangedRecord(object sender, EventArgs e)
        {
            File = SelectedRecord.Instance.SelectedFile;
            KellerDevice = SelectedRecord.Instance.SelectedKellerDevice;
            Gateway = SelectedRecord.Instance.SelectedGateway;
            IsRecordSavedLocally = SelectedRecord.Instance.IsSelectedFileLocal;
            DataSeries = new ObservableCollection<OxyPlotSeries>(OxyPlotSeriesGenerator.GenerateOxyPlotSeries(File));
            PageTitle = File.Header.DeviceName;
            OnPropertyChanged(nameof(PageTitle));
            OnPropertyChanged(nameof(File));
            OnPropertyChanged(nameof(Gateway));
            OnPropertyChanged(nameof(KellerDevice));
            OnPropertyChanged(nameof(IsRecordSavedLocally));
            OnPropertyChanged(nameof(DataSeries));
        }

        public void NavigatedTo()
        {
        }

        public void NavigatedFrom()
        {
        }

        public event PropertyChangedEventHandler PropertyChanged;

        [NotifyPropertyChangedInvocator]
        protected virtual void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }

        private void Export()
        {
            var dialog = new SaveFileDialog
            {
                InitialDirectory = Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments),
                Filter = "KOLIBRI Format|*.json|Excel|*.xlsx|CSV|*.csv",
                DefaultExt = "json",
                FileName = File.Header.RecordId

            };
            var result = dialog.ShowDialog();
            if (result.HasValue && result.Value)
            {
                switch (Path.GetExtension(dialog.FileName))
                {
                    case ".json":
                        _exportService.ExportKolibriFormat(File, dialog.FileName);
                        break;
                    case ".xlsx":
                        _exportService.ExportExcel(File, dialog.FileName);
                        break;
                    case ".csv":
                        _exportService.ExportCsv(File, dialog.FileName);
                        break;
                }
            }
        }

        private async void DeleteLocal()
        {
            var result = await DialogHost.Show(new ConfirmDeleteDialog(), "ContentDialog");

            if (!(bool) result)
            {
                // delete got canceled
                return;
            }

            _fileService.RemoveFileFormat(File.Header.RecordId);
            _navigationService.GoBack();

        }

        private void Download()
        {
            DownloadData();
        }

        private void DownloadAnDelete()
        {
            DownloadData();
            DeleteOnRemote();
        }

        private void DeleteOnRemote()
        {
            _measurementService.RemoveMeasurements(Gateway, KellerDevice);
        }

        private async void DownloadData()
        {
            var existingFiles = _fileService.FindFilesFromDevice(KellerDevice.UniqueSerialNumber);
            if (existingFiles != null && existingFiles.Count > 0)
            {
                var view = new CombineFilesDialog
                {
                    DataContext = new CombineFilesDialogViewModel
                    {
                        FilesToCombine = existingFiles
                    }
                };
                await DialogHost.Show(view, "ContentDialog", DownloadDialogClosingHandler);
            }
            else
            {
                FetchFromGateway(false);
            }
        }

        private void DownloadDialogClosingHandler(object sender, DialogClosingEventArgs eventArgs)
        {
            var combineMeasurements = (bool)eventArgs.Parameter;
            var view = (CombineFilesDialog)eventArgs.Session.Content;
            var viewModel = (CombineFilesDialogViewModel)view.DataContext;
            FetchFromGateway(combineMeasurements, viewModel.FilesToCombine);
        }

        private void FetchFromGateway(bool combineMeasurements, List<MeasurementFileFormatHeader> files = null)
        {
            try
            {
                var measurement = _measurementService.GetMeasurements(Gateway, KellerDevice);
                if (combineMeasurements && files != null)
                {
                    _fileService.WriteFileCombinedWith(measurement, files);
                }
                else
                {
                    _fileService.WriteFileFormat(measurement);
                }
                SelectedRecord.Instance.SelectLocalRecord(File);
            }
            catch (Exception e)
            {
                Logger.Warn(e, "Failed to fetch measurements from the ChirpNest");
            }
        }
    }
}
