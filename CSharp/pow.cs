using System.Security.Cryptography;
using System.Dynamic;
using Newtonsoft.Json;

public class POW
{
    static string getHash(string input, int count)
    {
        dynamic json = new ExpandoObject();
        json = new dynamic[10];

        string padding = new string('0', count);

        for(int i = 0; i < 10; i++)
        {
            int postFix = 0;
            while (true)
            {
                postFix += 1;
                string encodedHash = BitConverter.ToString(new SHA256Managed().ComputeHash(Encoding.UTF8.GetBytes(input + postFix.ToString()))).Replace("-", "");
                if (encodedHash.IndexOf(padding) == 0)
                {
                    json[i] = new ExpandoObject();
                    json[i].postfix = postFix;
                    json[i].hash = encodedHash;
                    break;
                }
            }
        }
        return Convert.ToBase64String(Encoding.UTF8.GetBytes(JsonConvert.SerializeObject(json)));
    }
}
