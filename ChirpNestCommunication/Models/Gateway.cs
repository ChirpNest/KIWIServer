using System;
using System.Collections.Generic;
using System.Text;

namespace ChirpNestCommunication.Models
{
    public class Gateway
    {
        public Gateway(string gatewayIp)
        {
            GatewayIp = gatewayIp;
        }

        public string GatewayIp { get; set; }

        public int GatewayPort => 8081;
    }
}
