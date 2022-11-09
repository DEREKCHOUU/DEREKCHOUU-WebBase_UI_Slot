package windows

import (
	"bytes"
	"fmt"
	_config "go/UI/config"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

//usecase :os setting

type Windowsset struct {
	OS  _config.Os
	APP _config.App
}

var Windowsseting Windowsset

func (m *Windowsset) Setipv4(ip string, prefixLength string, gateway string) {

	//check if ip avialable
	forCheck := strings.Split(ip, ".")
	for i := range forCheck {
		e, err := strconv.Atoi(forCheck[i])
		if err != nil || e < 255 || e > 0 {
			fmt.Print("ip input error. ", err)
			return
		}
	}

	subnet, err := strconv.Atoi(prefixLength)
	if err != nil {
		fmt.Print("subnet input error. ", err)
		return
	}

	forCheck = strings.Split(gateway, ".")
	for i := range forCheck {
		e, err := strconv.Atoi(forCheck[i])
		if err != nil || e < 255 || e > 0 {
			fmt.Print("gateway input error. ", err)
			return
		}
	}

	m.OS.Ip = ip
	m.OS.Subnet = subnet
	m.OS.Gateway = gateway

	err = _config.UpdateConfig(m)
	if err != nil {
		print("IP config update fail")
		return
	}

	indoxNo := getNetworkIndex()
	shellCommand("Set-NetIPInterface -InterfaceIndex " + indoxNo +
		" -Dhcp Disabled")
	shellCommand("Set-NetIPAddress  -InterfaceIndex " + indoxNo +
		" -IPAddress " + ip + " -AddressFamily IPv4 -PrefixLength " +
		prefixLength + " -DefaultGateway " + gateway)
}

func (m *Windowsset) SetDHCP(DHCP bool) {
	DHCPcmd := ""
	if !DHCP {
		DHCPcmd = "diable"
	} else {
		DHCPcmd = "enable"
	}

	m.OS.DHCP = DHCP
	err := _config.UpdateConfig(m)
	if err != nil {
		print("DHCP config update fail")
		return
	}

	indoxNo := getNetworkIndex()
	shellCommand("Set-NetIPInterface -InterfaceIndex " + indoxNo +
		" -Dhcp " + DHCPcmd)
}

func (m *Windowsset) Setntp(address string) {
	shellCommand(`"Set-ItemProperty -Path 
	"HKLM:\SYSTEM\CurrentControlSet\Services\w32time\Parameters" -Name 
	"NtpServer" -Value ` + address + `,0x9`)
	shellCommand("Restart-Service w32Time")
}

func (m *Windowsset) SetDNS(DNS string) {
	shellCommand(`Set-DnsClientServerAddress -InterfaceIndex 4 
	-ServerAddresses ` + DNS)
}

// `shellOutput` returns the output of powershell command, and any errors.
func shellOutput(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("powershell.exe", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

// `shellCommand` executes the powershell command.
func shellCommand(command string) {
	out, errout, err := shellOutput(command)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

func getNetworkIndex() string {
	indoxNo, errout, err := shellOutput(`Get-NetAdapter | Where-Object -Property Status 
	-EQ Up | Select-Object -Property ifIndex`)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
	s := strings.Split(indoxNo, "\n")
	n := strings.ReplaceAll(s[3], " ", "")
	return n
}

// func shellOutputWithDir(command, dir string) (string, string, error) {
// 	var stdout bytes.Buffer
// 	var stderr bytes.Buffer

// 	cmd := exec.Command("powershell.exe", command)
// 	cmd.Dir = dir
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()

// 	return stdout.String(), stderr.String(), err
// }
