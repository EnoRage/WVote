using System;
using System.Collections.Generic;

namespace UnblockHackNET.BL.DB
{
    public class FoundationOptions
    {
        public string Id { get; set; }

        public string Name { get; set; } = "";
        //gggg
        public int FoundedDate { get; set; } = 0;

        public float Capital { get; set; } = 0.0F;

        public string Country { get; set; } = "";

        public string Mission { get; set; } = "";
    }

    public class User
    {
        public string Id { get; set; }

        public string Name { get; set; } = "";

        public string EthPrvKey { get; set; } = "";

        public string EthAddress { get; set; } = "";

        public List<Tuple<string, string, string, string>> Foundations { get; set; }
            = new List<Tuple<string, string, string, string>>();
    }

    public class TransactionHistory
    {
        public string Id { get; set; }

        public string UserId { get; set; }

        public string OrgId { get; set; }

        //gggg.mm.dd*hh:mm:ss
        public string DateTime { get; set; }

        public float Summ { get; set; }
    }
}
