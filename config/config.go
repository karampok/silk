package config

import (
	"net"

	"github.com/containernetworking/cni/pkg/ns"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
)

//go:generate counterfeiter -o fakes/netNS.go --fake-name NetNS . netNS
type netNS interface {
	ns.NetNS
}

type DualAddress struct {
	Hardware net.HardwareAddr
	IP       net.IP
}

type Config struct {
	Container struct {
		DeviceName          string
		TemporaryDeviceName string
		Namespace           netNS
		Address             DualAddress
		MTU                 int
	}
	Host struct {
		DeviceName string
		Namespace  netNS
		Address    DualAddress
	}
}

func (c *Config) AsCNIResult() *current.Result {
	return &current.Result{
		Interfaces: []*current.Interface{
			&current.Interface{
				Name:    c.Host.DeviceName,
				Mac:     c.Host.Address.Hardware.String(),
				Sandbox: "",
			},
			&current.Interface{
				Name:    c.Container.DeviceName,
				Mac:     c.Container.Address.Hardware.String(),
				Sandbox: c.Container.Namespace.Path(),
			},
		},
		IPs: []*current.IPConfig{
			&current.IPConfig{
				Version:   "4",
				Interface: 1,
				Address: net.IPNet{
					IP:   c.Container.Address.IP,
					Mask: []byte{255, 255, 255, 255},
				},
				Gateway: c.Host.Address.IP,
			},
		},
		Routes: []*types.Route{
			&types.Route{
				Dst: net.IPNet{
					IP:   net.IPv4zero,
					Mask: net.IPv4Mask(0, 0, 0, 0),
				},
				GW: c.Host.Address.IP,
			},
		},
		DNS: types.DNS{},
	}
}