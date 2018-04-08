using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace UnblockHackNET.BL.DB
{
    public static class DataBaseService
    {
        private static string _apiUri = "http://rosum.westeurope.cloudapp.azure.com:3001/";

        public static async Task<List<string>> CreateSeed(string log)
        {
            var result = new List<string>();
            using (var client = new HttpClient())
            {
                var request = new HttpRequestMessage
                {
                    RequestUri = new Uri(_apiUri + "createSeed"),
                    Method = HttpMethod.Post
                };

                //request.Headers.Add("Content-Type", "application/x-www-form-urlencoded");
                List<KeyValuePair<string, string>> tmp = new List<KeyValuePair<string, string>>();
                tmp.Add(new KeyValuePair<string, string>("userID", log));

                request.Content = new FormUrlEncodedContent(tmp);
                var response = await client.SendAsync(request);
                if (response.StatusCode == HttpStatusCode.OK)
                {
                    var json = await response.Content.ReadAsStringAsync();
                    var myObj = json.Split();
                    if (!Equals(myObj, null))
                        result.AddRange(myObj);
                }
            }


            return result;
        }

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
