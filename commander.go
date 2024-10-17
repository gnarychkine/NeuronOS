package main

import (
	"net"
	"os"
	"time"

	// "github.com/go-ping/ping" // Doesn't work foe Go ver 1.23
	ping "github.com/prometheus-community/pro-bing"
)

// type Commander interface {
// 	Ping(host string) (PingResult, error)
// 	GetSystemInfo() (SystemInfo, error)
// }

// PingResult - struct keeps ping's result
type PingResult struct {
	Successful bool          `json:"success"`
	Time       time.Duration `json:"time"`
}

// SystemInfo - struct keeps sysinfo's result
type SystemInfo struct {
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ipaddress"`
}

// CommandResponse - struct is returned as data via http
type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

type Commander struct{}

func NewCommander() Commander {
	return Commander{}
}

// GetSystemInfo - gets host's system infofmation
func (c *Commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return SystemInfo{}, err
	}

	iPAddress := ""
	for _, address := range addrs {
		// Check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // Check if it's IPv4
				if len(iPAddress) > 0 {
					iPAddress += ", "
				}
				iPAddress += ipnet.IP.String()
				// fmt.Println(ipnet.IP.String())
			}
		}
	}
	return SystemInfo{
		Hostname:  hostname,
		IPAddress: iPAddress,
	}, nil
}

// Ping - runs ping to host
func (c *Commander) Ping(host string) (PingResult, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return PingResult{}, err
	}

	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	var pr PingResult
	pr.Time = stats.AvgRtt
	pr.Successful = false
	if stats.PacketsRecv > 0 {
		pr.Successful = true
	}

	return pr, nil
}
