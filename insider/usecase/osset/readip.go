package osset

import (
	"net"
)

func Readip() {
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		print(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err.
		if err != nil {
			print(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			print(ip)
		}
	}
}
