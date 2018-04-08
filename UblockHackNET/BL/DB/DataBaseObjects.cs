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
        public string EthPrvKey { get; set; } = "";

        public string EthAddress { get; set; } = "";

        public string EncryptedSeed { get; set; }
    }

    public class Vote
    {
        public string description { get; set; } = "";

        public List<string> approvedAddresses { get; set; } = new List<string>();

        public DateTime startTime { get; set; } = DateTime.MinValue;
        public DateTime endTime { get; set; } = DateTime.MinValue;
        public bool end { get; set; } = false;
        public string _id { get; set; } = "";
        public string valisatorsAddress { get; set; } = "";
        public int num { get; set; } = 0;
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
