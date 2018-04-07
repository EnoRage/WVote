using System;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace BlockChainNET.BL.DB
{
    public static class DataBaseService
    {
        private static string _apiUri = "http://itts-worker.westeurope.cloudapp.azure.com:6000/api/";

        public static async Task<bool> UserAuthCheck(string log)
        {
            var result = false;
            using (var client = new HttpClient())
            {
                var request = new HttpRequestMessage
                {
                    RequestUri = new Uri(_apiUri + $"users/{log}"),
                    Method = HttpMethod.Get
                };

                request.Headers.Add("Accept", "application/json");
                var response = await client.SendAsync(request);
                if (response.StatusCode == HttpStatusCode.OK)
                {
                    var json = await response.Content.ReadAsStringAsync();
                    var myObj = JsonConvert.DeserializeObject<bool>(json);
                    if (!Equals(myObj, null))
                        result = myObj;
                }
            }

            return result;
        }

        public static async Task<User> GetUser(string log)
        {
            var result = new User();
            using (var client = new HttpClient())
            {
                var request = new HttpRequestMessage
                {
                    RequestUri = new Uri(_apiUri + $"users/{log}/getuser"),
                    Method = HttpMethod.Get
                };

                request.Headers.Add("Accept", "application/json");
                var response = await client.SendAsync(request);
                if (response.StatusCode == HttpStatusCode.OK)
                {
                    var json = await response.Content.ReadAsStringAsync();
                    var myObj = JsonConvert.DeserializeObject<User>(json);
                    if (!Equals(myObj, null))
                        result = myObj;
                }
            }

            return result;
        }
    }
}
