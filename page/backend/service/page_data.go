package service

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type PageData struct {
	Uptime        uint64  `json:"uptime"`
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskPercent   float64 `json:"disk_percent"`
	NetSendTotal  uint64  `json:"net_send"`
	NetRecvTotal  uint64  `json:"net_recv"`
	Timestamp     int64   `json:"timestamp"`
}

func GetPageData() *PageData {
	hostInfo, _ := host.Info()
	cpuPercent, _ := cpu.Percent(time.Duration(1)*time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")
	ioStats, _ := net.IOCounters(false)
	return &PageData{
		Uptime:        hostInfo.Uptime,
		CPUPercent:    cpuPercent[0],
		MemoryPercent: memInfo.UsedPercent,
		DiskPercent:   diskUsage.UsedPercent,
		NetSendTotal:  ioStats[0].BytesSent,
		NetRecvTotal:  ioStats[0].BytesRecv,
		Timestamp:     time.Now().Unix(),
	}
}
