package domain

type DeviceData struct {
	Id          int    `json:"id"`
	DeviceId    int    `json:"deviceid"`
	Description string `json:"accessdata"`
	Actiontime  string `json:"time"`
}

type OsSetting struct {
	Ipv4 int    `json:"ip"`
	Port int    `json:"port"`
	Os   string `json:"os"`
}
