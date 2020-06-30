using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace KIWIDesktop.ViewModels
{
    public interface IViewModel : INotifyPropertyChanged
    {
        string PageTitle { get; }

        void NavigatedTo();

        void NavigatedFrom();
    }
}
