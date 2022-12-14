package service

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

type SystemInfo struct {
	HostInfo     *HostInfo     `json:"host_info"`
	CPUInfo      *CPUInfo      `json:"cpu_info"`
	MemoryInfo   *MemoryInfo   `json:"memory_info"`
	NetworkInfo  *NetworkInfo  `json:"network_info"`
	DiskInfo     *DiskInfo     `json:"disk_info"`
	SystemStatus *SystemStatus `json:"system_status"`
}

type SystemStatus struct {
	CPUPercent   []float64     `json:"cpu_percent"`
	LoadAvg      *load.AvgStat `json:"load_avg"`
	MemoryStatus *MemoryStatus `json:"memory_status"`
	DiskInfo     *DiskInfo     `json:"disk_info"`
	NetworkInfo  *NetworkInfo  `json:"network_info"`
	Uptime       uint64        `json:"uptime"`
	Timestamp    int64         `json:"timestamp"`
}

func GetSystemInfo() *SystemInfo {
	return &SystemInfo{
		HostInfo:     GetHostInfo(),
		CPUInfo:      GetCPUInfo(),
		MemoryInfo:   GetMemoryInfo(),
		DiskInfo:     GetDiskInfo(),
		NetworkInfo:  GetNetworkInfo(),
		SystemStatus: GetSystemStatus(),
	}
}

func GetSystemStatus() *SystemStatus {
	cpuPercent, _ := cpu.Percent(0, true)
	loadAvg, _ := load.Avg()
	return &SystemStatus{
		CPUPercent:   cpuPercent,
		LoadAvg:      loadAvg,
		MemoryStatus: GetMemoryStatus(),
		DiskInfo:     GetDiskInfo(),
		NetworkInfo:  GetNetworkInfo(),
		Uptime:       GetUptime(),
		Timestamp:    time.Now().Unix(),
	}
}
