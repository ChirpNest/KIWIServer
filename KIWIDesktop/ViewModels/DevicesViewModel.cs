using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Windows.Input;
using ChirpNestCommunication;
using ChirpNestCommunication.Models;
using GalaSoft.MvvmLight.CommandWpf;
using GalaSoft.MvvmLight.Views;
using KellerAg.Shared.Entities.FileFormat;
using KIWIDesktop.Annotations;
using KIWIDesktop.Navigation;
using KIWIDesktop.Services;
using KIWIDesktop.ViewModels.TemplateViewModels;
using KIWIDesktop.Views.Templates;
using MaterialDesignThemes.Wpf;
using NLog;

namespace KIWIDesktop.ViewModels
{ 
    public class DevicesViewModel : IViewModel
    {
        private readonly IDeviceService _deviceService;
        private readonly IMeasurementService _measurementService;
        private readonly IKellerFileService _fileService;
        private readonly IViewNavigationService _navigationService;

        public DevicesViewModel(IDeviceService deviceService, IMeasurementService measurementService, IKellerFileService fileService, IViewNavigationService navigationService)
        {
            _deviceService = deviceService;
            _measurementService = measurementService;
            _fileService = fileService;
            _navigationService = navigationService;
            LocalMeasurements = new ObservableCollection<MeasurementFileFormatHeader>();
            LoadLocal();
        }

        public string PageTitle => "Overview";

        public ICommand ConnectCommand => new RelayCommand(Connect);

        public ICommand DisconnectCommand => new RelayCommand(Disconnect);

        public ICommand ShowDeviceCommand => new RelayCommand<KellerDevice>(SelectData);

        public ICommand ShowLocalRecordCommand => new RelayCommand<MeasurementFileFormatHeader>(ShowLocalRecord);

        public ICommand DownloadAndRemoveCommand => new RelayCommand<KellerDevice>(DownloadDataAndRemove);

        public string ErrorMessage { get; set; }

        public string GatewayIp { get; set; }

        public List<KellerDevice> KellerDevices { get; set; }

        public ObservableCollection<MeasurementFileFormatHeader> LocalMeasurements { get; set; }

        public bool IsConnected { get; set; }

        private static readonly Logger Logger = LogManager.GetCurrentClassLogger();

        public void Disconnect()
        {
            KellerDevices.Clear();
            IsConnected = false;
            OnPropertyChanged(nameof(KellerDevices));
            OnPropertyChanged(nameof(IsConnected));
        }

        public void Connect()
        {
            LoadFromChirpNest();
        }

        public void LoadFromChirpNest()
        {
            ErrorMessage = string.Empty;
            try
            {
                var gateway = new Gateway(GatewayIp);
                KellerDevices = _deviceService.GetRegisteredDevices(gateway);
                IsConnected = true;
            }
            catch (Exception e)
            {
                Logger.Warn($"Connection to ChirpNest coult not be established with IP {GatewayIp}");
                ErrorMessage = e.Message;
            }
            OnPropertyChanged(nameof(ErrorMessage));
            OnPropertyChanged(nameof(KellerDevices));
            OnPropertyChanged(nameof(IsConnected));
        }

        public void LoadLocal()
        {
            LocalMeasurements.Clear();
            try
            {
                foreach (var header in _fileService.ReadTableOfContent())
                {
                    LocalMeasurements.Add(header);
                }
            }
            catch (Exception e)
            {
                Logger.Warn(e, "Failed to read local measurements");
            }
        }

        private void SelectData(KellerDevice kellerDevice)
        {
            SelectedRecord.Instance.ReadFromDeviceAndSelect(new Gateway(GatewayIp), kellerDevice);
            NavigateToVisualize();
        }
        private async void DownloadDataAndRemove(KellerDevice kellerDevice)
        {
            var existingFiles = _fileService.FindFilesFromDevice(kellerDevice.UniqueSerialNumber);
            if (existingFiles != null && existingFiles.Count > 0)
            {
                var view = new CombineFilesDialog
                {
                    DataContext = new CombineFilesDialogViewModel
                    {
                        FilesToCombine = existingFiles,
                        Device = kellerDevice
                    }
                };
                await DialogHost.Show(view, "ContentDialog", DownloadDialogClosingHandler);
            }
            else
            {
                Download(false, kellerDevice);
            }
        }

        private void DownloadDialogClosingHandler(object sender, DialogClosingEventArgs eventArgs)
        {
            var combineMeasurements = (bool) eventArgs.Parameter;
            var view = (CombineFilesDialog)eventArgs.Session.Content;
            var viewModel = (CombineFilesDialogViewModel)view.DataContext;
            Download(combineMeasurements, viewModel.Device, viewModel.FilesToCombine);
        }

        private void Download(bool combineMeasurements, KellerDevice device, List<MeasurementFileFormatHeader> files = null)
        {
            var gateway = new Gateway(GatewayIp);
            try
            {
                var measurement = _measurementService.GetMeasurements(gateway, device);
                if (combineMeasurements  && files != null)
                {
                    _fileService.WriteFileCombinedWith(measurement, files);
                }
                else
                {
                    _fileService.WriteFileFormat(measurement);
                }
                _measurementService.RemoveMeasurements(gateway, device);
                LoadLocal();
            }
            catch (Exception e)
            {
                Logger.Warn(e, "Failed to read measurements from Chirpnest");
            }
        }

        private void ShowLocalRecord(MeasurementFileFormatHeader header)
        {
            var record = _fileService.ReadFileFormat(header.RecordId);
            SelectedRecord.Instance.SelectLocalRecord(record);
            NavigateToVisualize();
        }

        private void NavigateToVisualize()
        {
            //This is a workaround to add the Devices page to the history
            if (_navigationService.CurrentPageKey != "Devices")
            {
                _navigationService.NavigateTo("Devices");
            }
            _navigationService.NavigateTo("Visualize");
        }

        public void NavigatedTo()
        {
            LoadLocal();
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
    }
}
