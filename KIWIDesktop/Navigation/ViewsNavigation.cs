using CommonServiceLocator;
using GalaSoft.MvvmLight.CommandWpf;
using System.Windows.Controls;
using KIWIDesktop.Views;

namespace KIWIDesktop.Navigation
{
    public class ViewsNavigation
    {
        public UserControl VisualizeControl { get; }
        public UserControl DevicesControl { get; }

        public ViewNavigationService NavigationService { get; set; }

        private RelayCommand<string> _switchCommand;

        public RelayCommand<string> SwitchCommand
        {
            get
            {
                return _switchCommand ??
                       (_switchCommand = new RelayCommand<string>
                           (
                               newControl =>
                               {
                                   NavigationService.NavigateTo(newControl);
                               })
                       );
            }
        }

        public bool CanNavigate => NavigationService.CanNavigate;

        public ViewsNavigation()
        {
            VisualizeControl = new Visualize();
            DevicesControl = new Devices();

            ConfigureNavigation();

        }

        private void ConfigureNavigation()
        {
            NavigationService = ServiceLocator.Current.GetInstance<IViewNavigationService>() as ViewNavigationService;

            if (NavigationService == null)
            {
                return;
            }

            NavigationService.Configure("Visualize", VisualizeControl);
            NavigationService.Configure("Devices", DevicesControl);
        }
    }

}
