package osset

import (
	_config "go/UI/config"
	_windows "go/UI/insider/usecase/osset/windows"
)

type Ossetinf interface {
	Setipv4(ip string, prefixLength string, gateway string)
	SetDHCP(DHCP bool)
	Setntp(address string)
	SetDNS(DNS string)
}

func NewWindowSetup(os _config.Os, app _config.App) Ossetinf {
	return &_windows.Windowsset{
		OS:  os,
		APP: app,
	}
}
