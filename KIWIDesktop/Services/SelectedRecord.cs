using System;
using System.ComponentModel;
using System.Runtime.CompilerServices;
using System.Threading.Tasks;
using ChirpNestCommunication;
using ChirpNestCommunication.Models;
using GalaSoft.MvvmLight.Ioc;
using KellerAg.Shared.Entities.FileFormat;
using KIWIDesktop.Annotations;

namespace KIWIDesktop.Services
{
    public class SelectedRecord :INotifyPropertyChanged
    {
        private readonly IMeasurementService _measurementService;
        public SelectedRecord()
        {
            //Is this a clean way to get the service?
            _measurementService = SimpleIoc.Default.GetInstance<IMeasurementService>();
        }

        private static SelectedRecord _instance;

        private static readonly object SingletonLock = new object();

        public MeasurementFileFormat SelectedFile { get; private set; }

        public bool IsSelectedFileLocal { get; private set; }

        public Gateway SelectedGateway { get; private set; }

        public KellerDevice SelectedKellerDevice { get; private set; }

        public event EventHandler SelectedRecordChanged;

        public static SelectedRecord Instance
        {
            get
            {
                if (_instance == null)
                {
                    lock (SingletonLock)
                    {
                        if (_instance == null)
                            _instance = new SelectedRecord();
                    }
                }

                return _instance;
            }
        }

        public async void ReadFromDeviceAndSelect(Gateway gateway, KellerDevice device)
        {
            var file = await Task.Run(() => _measurementService.GetMeasurements(gateway, device));
            SelectedGateway = gateway;
            SelectedKellerDevice = device;
            IsSelectedFileLocal = false;
            SelectRecord(file);

        }

        public void SelectLocalRecord(MeasurementFileFormat file)
        {
            IsSelectedFileLocal = true; 
            SelectedGateway = null;
            SelectedKellerDevice = null;
            SelectRecord(file);
        }

        private void SelectRecord(MeasurementFileFormat file)
        {
            SelectedFile = file;
            SelectedRecordChanged?.Invoke(this, null);
        }


        public event PropertyChangedEventHandler PropertyChanged;

        [NotifyPropertyChangedInvocator]
        protected virtual void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }
}