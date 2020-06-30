using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using OxyPlot.Wpf;

namespace KIWIDesktop.XamlExtensions
{
    public class PlotExtension : Plot
    {

        public static readonly DependencyProperty SourceProperty =
            DependencyProperty.Register("Source", typeof(object), typeof(PlotExtension), new PropertyMetadata(null, OnPropertyChanged));

        public static readonly DependencyProperty SeriesTemplateProperty =
            DependencyProperty.Register("SeriesTemplate", typeof(DataTemplate), typeof(PlotExtension), new PropertyMetadata(null, OnPropertyChanged));

        public static readonly DependencyProperty SeriesDataTemplateSelectorProperty =
            DependencyProperty.Register("SeriesDataTemplateSelector", typeof(SeriesDataTemplateSelector), typeof(PlotExtension), new PropertyMetadata(null, OnPropertyChanged));


        //Gets or sets the ItemsSource of collection of collections.

        public object Source
        {
            get => GetValue(SourceProperty);
            set => SetValue(SourceProperty, value);
        }

        //Gets or sets the template for the series to be generated.

        public DataTemplate SeriesTemplate
        {
            get => (DataTemplate)GetValue(SeriesTemplateProperty);
            set => SetValue(SeriesTemplateProperty, value);
        }

        //Get or sets the DataTemplateSelector for the multiple series generation.
        public SeriesDataTemplateSelector SeriesDataTemplateSelector
        {
            get => (SeriesDataTemplateSelector)GetValue(SeriesDataTemplateSelectorProperty);
            set => SetValue(SeriesDataTemplateSelectorProperty, value);
        }

        private static void OnPropertyChanged(DependencyObject d, DependencyPropertyChangedEventArgs e)
        {
            (d as PlotExtension)?.GenerateSeries();
        }

        //Generate the series per the counts in the itemssource.
        private void GenerateSeries()
        {
            Series.Clear();
            if (Source == null || (SeriesDataTemplateSelector == null && SeriesTemplate == null))
                return;
            var commonItemsSource = (Source as IEnumerable)?.GetEnumerator();

            while (commonItemsSource != null && commonItemsSource.MoveNext())
            {
                DataPointSeries series = null;

                //The conditions checked for setting the SeriesTemplate or SeriesDataTemplateSelector.
                if (SeriesTemplate != null)
                {
                    series = SeriesTemplate.LoadContent() as DataPointSeries;
                }
                else
                {
                    var selectedseriesTemplate =
                        SeriesDataTemplateSelector?.SelectTemplate(commonItemsSource.Current,
                            null);
                    if (selectedseriesTemplate != null) series = selectedseriesTemplate.LoadContent() as DataPointSeries;
                }

                if (series != null)
                {
                    series.DataContext = commonItemsSource.Current;
                    Series.Add(series);
                }
            }
        }
    }
}
