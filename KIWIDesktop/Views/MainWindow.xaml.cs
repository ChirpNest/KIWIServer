using System;
using System.Windows;
using KIWIDesktop.Navigation;

namespace KIWIDesktop.Views
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
            var vn = ViewsNavigation;
            if (vn == null)
            {
                Console.Out.Write("navigation not loaded");
            }
        }
        private ViewsNavigation _viewsNavigation;
        public ViewsNavigation ViewsNavigation => _viewsNavigation ?? (_viewsNavigation = new ViewsNavigation());
    }
}
