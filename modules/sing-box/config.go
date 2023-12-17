package singbox

import (
	"net/netip"

	"github.com/dickymuliafiqri/BenchBox/modules/helper"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
)

func GenerateConfig(outbound *option.Outbound) (option.Options, uint) {
	listenPort := helper.GetFreePort()
	options := option.Options{
		Log: &option.LogOptions{
			Disabled: true,
		},
		DNS: &option.DNSOptions{
			Servers: []option.DNSServerOptions{
				{
					Tag:     "dns-direct",
					Address: "tls://1.1.1.1",
					Detour:  "direct",
				},
				{
					Tag:     "dns-proxy",
					Address: "udp://1.1.1.1",
					Detour:  outbound.Tag,
				},
			},
			Rules: []option.DNSRule{
				{
					DefaultOptions: option.DefaultDNSRule{
						Domain: []string{"twilio.com"},
						Server: "dns-proxy",
					},
				},
			},
			Final: "dns-direct",
		},
		Inbounds: []option.Inbound{
			{
				Type: C.TypeMixed,
				MixedOptions: option.HTTPMixedInboundOptions{
					ListenOptions: option.ListenOptions{
						Listen:     option.NewListenAddress(netip.IPv4Unspecified()),
						ListenPort: uint16(listenPort),
					},
				},
			},
		},
		Outbounds: []option.Outbound{
			{
				Type: C.TypeDirect,
				Tag:  C.TypeDirect,
			},
			*outbound,
		},
	}

	return options, listenPort
}
