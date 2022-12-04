package service

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfo struct {
	Hostname        string       `json:"hostanme"`
	Distribution    string       `json:"distribution"`
	Arch            string       `json:"arch"`
	Kernel          string       `json:"kernel"`
	VirtualPlatform string       `json:"virtual_platform"`
	Uptime          uint64       `json:"uptime"`
	CPUInfo         *CPUInfo     `json:"cpu_info"`
	Memory          uint64       `json:"memory"`
	Swap            uint64       `json:"swap"`
	DiskInfo        *DiskInfo    `json:"disk_info"`
	NetworkInfo     *NetworkInfo `json:"network_info"`
	Timestamp       int64        `json:"timestamp"`
}

type SystemStatus struct {
	CPUPercent    []float64     `json:"cpu_percent"`
	LoadAvg       *load.AvgStat `json:"load_avg"`
	MemoryPercent float64       `json:"memory_percent"`
	SwapPercent   float64       `json:"swap_percent"`
	DiskInfo      *DiskInfo     `json:"disk_info"`
	NetworkInfo   *NetworkInfo  `json:"network_info"`
}

func GetSystemInfo() *SystemInfo {
	hostInfo, _ := host.Info()
	vmStat, _ := mem.VirtualMemory()
	swapStat, _ := mem.SwapMemory()
	return &SystemInfo{
		Hostname:        hostInfo.Hostname,
		Distribution:    hostInfo.Platform + " " + hostInfo.PlatformVersion,
		Arch:            hostInfo.KernelArch,
		Kernel:          hostInfo.KernelVersion,
		VirtualPlatform: hostInfo.VirtualizationSystem,
		Uptime:          hostInfo.Uptime,
		CPUInfo:         GetCPUInfo(),
		Memory:          vmStat.Total,
		Swap:            swapStat.Total,
		DiskInfo:        GetDiskInfo(),
		NetworkInfo:     GetNetworkInfo(),
		Timestamp:       time.Now().Unix(),
	}
}

func GetSystemStatus() *SystemStatus {
	cpuPercent, _ := cpu.Percent(0, true)
	loadAvg, _ := load.Avg()
	memInfo, _ := mem.VirtualMemory()
	swapInfo, _ := mem.SwapMemory()
	return &SystemStatus{
		CPUPercent:    cpuPercent,
		LoadAvg:       loadAvg,
		MemoryPercent: memInfo.UsedPercent,
		SwapPercent:   swapInfo.UsedPercent,
		DiskInfo:      GetDiskInfo(),
		NetworkInfo:   GetNetworkInfo(),
	}
}
