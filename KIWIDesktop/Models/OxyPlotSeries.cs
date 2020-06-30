using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using System.Web.UI.WebControls;
using KIWIDesktop.Annotations;
using OxyPlot;
using UnitType = KellerAg.Shared.Entities.Units.UnitType;

namespace KIWIDesktop.Models
{
    public class OxyPlotSeries : INotifyPropertyChanged
    {
        public OxyPlotSeries(string title, UnitType unitType, IList<DataPoint> points)
        {
            Title = title;
            Points = points;
            UnitType = unitType;
        }

        public string Title { get; set; }

        public IList<DataPoint> Points { get; private set; }

        public UnitType UnitType { get; set; }

        public event PropertyChangedEventHandler PropertyChanged;

        [NotifyPropertyChangedInvocator]
        protected virtual void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }
}
