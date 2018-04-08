using System;
using System.Windows.Input;
using UnblockHackNET.UI;
using UnblockHackNET.BL.DB;

namespace UnblockHackNET.BL.ViewModels.Auth
{
    public class AuthViewModel : BaseViewModel
    {
        public ICommand AuthCommand => MakeCommand(async () =>
        {
            //var tmp = PrivateKeyAccount.GenerateSeed();
            //WavesCS.AddressEncoding.
            //var enc = Crypt.Encrypt(tmp, "123");
            var tmp = await DataBaseService.CreateSeed("1234");
            NavigationService.Instance.SetMainMasterDetailPage(Pages.Menu, Pages.Main);
        });
    }
}
