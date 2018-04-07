using System;
using System.Windows.Input;
using UnblockHackNET.UI;

namespace UnblockHackNET.BL.ViewModels.Auth
{
    public class AuthViewModel : BaseViewModel
    {
        public ICommand AuthCommand => MakeCommand(()=>{
            /*
             *
             * Some Logic
             *
             */
            NavigationService.Instance.SetMainMasterDetailPage(Pages.Menu, Pages.Main);
        });
    }
}
