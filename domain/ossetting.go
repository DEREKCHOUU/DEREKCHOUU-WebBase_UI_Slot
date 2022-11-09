package domain

type Setting struct {
	Os      string `Json:"os"`
	Ip      string `json:"ip"`
	Subnet  string `json:"subnet"`
	Gateway string `json:"gateway"`
	DHCP    string `json:"DHCP"`
	Uiport  string `json:"uiport"`
	Sqltype string `json:"sqltype"`
	Sqlport string `json:"sqlport"`
	Sqlacc  string `json:"sqlacc"`
	Sqlpw   string `json:"sqlpw"`
	Dbexit  string `json:"dbexit"`
	UTC     string `json:"UTC"`
}

type SettingUse interface {
	Setipv4(ip string, prefixLength int, gateway string)
	SetDHCP()
	Setntp(address string)
	SetDNS(DNS string)
}
