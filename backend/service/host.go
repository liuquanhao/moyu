package service

import (
	"github.com/shirou/gopsutil/v3/host"
)

type HostInfo struct {
	Hostname        string `json:"hostname"`
	Distribution    string `json:"distribution"`
	Arch            string `json:"arch"`
	Kernel          string `json:"kernel"`
	VirtualPlatform string `json:"virtual_platform"`
	Uptime          uint64 `json:"uptime"`
}

func GetHostInfo() *HostInfo {
	hostInfo, _ := host.Info()
	return &HostInfo{
		Hostname:        hostInfo.Hostname,
		Distribution:    hostInfo.Platform + " " + hostInfo.PlatformVersion,
		Arch:            hostInfo.KernelArch,
		Kernel:          hostInfo.KernelVersion,
		VirtualPlatform: hostInfo.VirtualizationSystem,
		Uptime:          hostInfo.Uptime,
	}
}
