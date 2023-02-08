package service

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryInfo struct {
	Memory uint64 `json:"memory"`
	Swap   uint64 `json:"swap"`
}

type MemoryStatus struct {
	MemoryPercent float64 `json:"memory_percent"`
	SwapPercent   float64 `json:"swap_percent"`
}

func GetMemoryInfo() *MemoryInfo {
	vmStat, _ := mem.VirtualMemory()
	swapStat, _ := mem.SwapMemory()
	return &MemoryInfo{
		Memory: vmStat.Total,
		Swap:   swapStat.Total,
	}
}

func GetMemoryStatus() *MemoryStatus {
	memInfo, _ := mem.VirtualMemory()
	swapInfo, _ := mem.SwapMemory()
	return &MemoryStatus{
		MemoryPercent: memInfo.UsedPercent,
		SwapPercent:   swapInfo.UsedPercent,
	}
}
