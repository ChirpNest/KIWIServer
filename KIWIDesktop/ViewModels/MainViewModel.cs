using System;
using System.Collections.Generic;
using System.Windows.Input;
using ChirpNestCommunication;
using ChirpNestCommunication.Models;
using GalaSoft.MvvmLight;
using GalaSoft.MvvmLight.CommandWpf;
using KIWIDesktop.Navigation;
using KIWIDesktop.Services;

namespace KIWIDesktop.ViewModels
{
    /// <summary>
    /// This class contains properties that the main View can data bind to.
    /// <para>
    /// Use the <strong>mvvminpc</strong> snippet to add bindable properties to this ViewModel.
    /// </para>
    /// <para>
    /// You can also use Blend to data bind with the tool's support.
    /// </para>
    /// <para>
    /// See http://www.galasoft.ch/mvvm
    /// </para>
    /// </summary>
    public class MainViewModel : ViewModelBase
    {
        private readonly IViewNavigationService _navigationService;

        /// <summary>
        /// Initializes a new instance of the MainViewModel class.
        /// </summary>
        public MainViewModel(IViewNavigationService navigationService)
        {
            _navigationService = navigationService;
        }

        public ICommand NavigateBackCommand => new RelayCommand(NavigateBack);

        private void NavigateBack()
        {
            _navigationService.GoBack();
        }
    }
}