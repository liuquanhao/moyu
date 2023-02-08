package service

import (
	"regexp"

	"github.com/shirou/gopsutil/v3/net"
)

type Ifce struct {
	Name     string `json:"name"`
	ByteSend uint64 `json:"send_byte"`
	ByteRecv uint64 `json:"recv_byte"`
}

type NetworkInfo struct {
	Ifces []*Ifce `json:"ifces"`
}

func (networkInfo *NetworkInfo) AppendIfce(ifce *Ifce) {
	networkInfo.Ifces = append(networkInfo.Ifces, ifce)
}

func GetNetworkInfo() *NetworkInfo {
	networkInfo := new(NetworkInfo)
	ioStats, _ := net.IOCounters(true)
	r, _ := regexp.Compile("^(eth|enp).*")
	for _, ioStat := range ioStats {
		if !r.MatchString(ioStat.Name) {
			continue
		}
		ifce := &Ifce{
			Name:     ioStat.Name,
			ByteSend: ioStat.BytesSent,
			ByteRecv: ioStat.BytesRecv,
		}
		networkInfo.AppendIfce(ifce)
	}
	return networkInfo
}
