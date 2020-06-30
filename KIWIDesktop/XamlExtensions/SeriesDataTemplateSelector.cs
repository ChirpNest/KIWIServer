using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;

namespace KIWIDesktop.XamlExtensions
{
    public class SeriesDataTemplateSelector : DataTemplateSelector
    {
        public DataTemplate LineSeriesTemplate { get; set; }

        public override DataTemplate SelectTemplate(object item, DependencyObject container)
        {

            DataTemplate selectedDataTemplate;

            if (item == null)
            {
                return null;
            }

            selectedDataTemplate = LineSeriesTemplate;



            return selectedDataTemplate;
        }
    }
}
