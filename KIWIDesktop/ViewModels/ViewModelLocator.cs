/*
  In App.xaml:
  <Application.Resources>
      <vm:ViewModelLocator xmlns:vm="clr-namespace:KIWIDesktop"
                           x:Key="Locator" />
  </Application.Resources>
  
  In the View:
  DataContext="{Binding Source={StaticResource Locator}, Path=ViewModelName}"

  You can also use Blend to do all this with the tool's support.
  See http://www.galasoft.ch/mvvm
*/

using System.Collections.Generic;
using System.Windows.Navigation;
using Kiwi.Api;
using ChirpNestCommunication;
using ChirpNestCommunication.Mapping;
using ChirpNestCommunication.Models;
using CommonServiceLocator;
using GalaSoft.MvvmLight.Ioc;
using KellerAg.Shared.Entities.FileFormat;
using KIWIDesktop.Navigation;
using KIWIDesktop.Services;
using DeviceService = ChirpNestCommunication.DeviceService;
using MeasurementService = ChirpNestCommunication.MeasurementService;

namespace KIWIDesktop.ViewModels
{
    /// <summary>
    /// This class contains static references to all the view models in the
    /// application and provides an entry point for the bindings.
    /// </summary>
    public class ViewModelLocator
    {
        /// <summary>
        /// Initializes a new instance of the ViewModelLocator class.
        /// </summary>
        public ViewModelLocator()
        {
            ServiceLocator.SetLocatorProvider(() => SimpleIoc.Default);

            RegisterMapping();
            RegisterServices();
            RegisterViewModels();
        }

        private void RegisterViewModels()
        {
            SimpleIoc.Default.Register<MainViewModel>();
            SimpleIoc.Default.Register<VisualizeViewModel>();
            SimpleIoc.Default.Register<DevicesViewModel>();
        }

        private void RegisterMapping()
        {
            SimpleIoc.Default.Register<IMapping<GetMeasurementsResponse, MeasurementFileFormat>, GetMeasurementsResponseMapping> ();
            SimpleIoc.Default.Register<IMapping<ListDeviceResponse, List<KellerDevice>>, ListDeviceResponseMapping> ();
        }

        private void RegisterServices()
        {
            if (!SimpleIoc.Default.IsRegistered<IViewNavigationService>())
            {
                var navigationService = new ViewNavigationService();
                SimpleIoc.Default.Register<IViewNavigationService>(() => navigationService);
            }
            SimpleIoc.Default.Register<IDeviceService, DeviceService>();
            SimpleIoc.Default.Register<IMeasurementService, MeasurementService>();
            SimpleIoc.Default.Register<IKellerFileService, KellerFileService>();
            SimpleIoc.Default.Register<IExportService, ExportService>();
        }


        public MainViewModel Main => ServiceLocator.Current.GetInstance<MainViewModel>();

        public VisualizeViewModel Visualize => ServiceLocator.Current.GetInstance<VisualizeViewModel>();

        public DevicesViewModel Devices => ServiceLocator.Current.GetInstance<DevicesViewModel>();

        public static void Cleanup()
        {
        }
    }
}