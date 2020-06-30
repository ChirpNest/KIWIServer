using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Media;
using KIWIDesktop.Annotations;
using KIWIDesktop.ViewModels;

namespace KIWIDesktop.Navigation
{
    public class ViewNavigationService : IViewNavigationService
    {
        private readonly Dictionary<string, UserControl> _pagesByKey;
        private readonly List<string> _historic;
        private string _currentPageKey;
        private bool _canNavigate;

        public string CurrentPageKey
        {
            get
            {
                // ReSharper disable once ArrangeAccessorOwnerBody
                return _currentPageKey;
            }

            private set
            {
                if (_currentPageKey != null && _currentPageKey == value)
                {
                    return;
                }

                _currentPageKey = value;
            }
        }

        public bool CanNavigate
        {
            get
            {
                // ReSharper disable once ArrangeAccessorOwnerBody
                return _canNavigate;
            }
            set
            {
                _canNavigate = value;
                OnPropertyChanged(nameof(CanNavigate));
            }
        }

        public object Parameter { get; private set; }

        public ViewNavigationService()
        {
            _pagesByKey = new Dictionary<string, UserControl>();
            _historic = new List<string>();
            CanNavigate = true;
        }

        public void GoBack()
        {
            if (_historic.Count > 1)
            {
                _historic.RemoveAt(_historic.Count - 1);
                NavigateTo(_historic.Last(), null);
                if (CurrentPageKey != _historic.Last())
                {
                    _historic.Add(CurrentPageKey);
                }
            }
        }

        public void NavigateTo(string pageKey)
        {
            NavigateTo(pageKey, null);
        }

        public virtual void NavigateTo(string pageKey, object parameter)
        {
            lock (_pagesByKey)
            {
                if (!_pagesByKey.ContainsKey(pageKey))
                {
                    throw new ArgumentException("Page not found", nameof(pageKey));
                }

                var control = GetDescendantFromName(Application.Current.MainWindow, "MainArea") as UserControl;

                Parameter = parameter;

                PerformNavigateTo(pageKey, control);
            }
        }

        private void PerformNavigateTo(string pageKey, UserControl control)
        {
            if (control != null)
            {
                ((IViewModel)((UserControl)control.Content).DataContext).NavigatedFrom();
                control.Content = _pagesByKey[pageKey];
            }

            _historic.Add(pageKey);
            CurrentPageKey = pageKey;

            ((IViewModel)((UserControl)control?.Content)?.DataContext)?.NavigatedTo();
        }

        public void Configure(string key, UserControl pageType)
        {
            lock (_pagesByKey)
            {
                if (_pagesByKey.ContainsKey(key))
                {
                    _pagesByKey[key] = pageType;
                }
                else
                {
                    _pagesByKey.Add(key, pageType);
                }
            }
        }

        private static FrameworkElement GetDescendantFromName(DependencyObject parent, string name)
        {
            var count = VisualTreeHelper.GetChildrenCount(parent);

            if (count < 1)
            {
                return null;
            }

            for (var i = 0; i < count; i++)
            {
                if (!(VisualTreeHelper.GetChild(parent, i) is FrameworkElement frameworkElement))
                {
                    continue;
                }
                if (frameworkElement.Name == name)
                {
                    return frameworkElement;
                }

                frameworkElement = GetDescendantFromName(frameworkElement, name);
                if (frameworkElement != null)
                {
                    return frameworkElement;
                }
            }
            return null;
        }

        public event PropertyChangedEventHandler PropertyChanged;

        [NotifyPropertyChangedInvocator]
        protected virtual void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }
}
