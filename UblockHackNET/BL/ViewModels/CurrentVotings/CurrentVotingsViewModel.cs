using System;
using System.Collections.ObjectModel;
using System.Windows.Input;
using UnblockHackNET.BL.DB;

namespace UnblockHackNET.BL.ViewModels.CurrentVotings
{
    public class CurrentVotingsViewModel : BaseViewModel
    {
        public bool IsRefreshing
        {
            get => Get(false);
            set => Set(value);
        }

        public ICommand RefreshCommand => MakeCommand(()=>
        {
            RefreshChartItemSource = new ObservableCollection<Vote>();
            LoadVotes();

        });

        public ObservableCollection<Vote> RefreshChartItemSource
        {
            get => Get(new ObservableCollection<Vote>());
            set => Set(value);
        }

        public CurrentVotingsViewModel()
        {
            LoadVotes();
        }

        private async void LoadVotes()
        {
            IsRefreshing = true;
            RefreshChartItemSource = await DataBaseService.GetAllVotes();
            OnPropertyChanged(nameof(RefreshChartItemSource));
            IsRefreshing = false;
        }
    }
}
