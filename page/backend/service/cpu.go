package service

import "github.com/shirou/gopsutil/v3/cpu"

type Core struct {
	CPU       int32    `json:"cpu"`
	CoreID    string   `json:"core_id"`
	ModelName string   `json:"model"`
	Mhz       float64  `json:"mhz"`
	Flags     []string `json:"flags"`
}
type CPUInfo struct {
	Count int     `json:"count"`
	Cores []*Core `json:"cores"`
}

func (cpuInfo *CPUInfo) AppendCore(core *Core) {
	cpuInfo.Cores = append(cpuInfo.Cores, core)
}

func GetCPUInfo() *CPUInfo {
	cpuInfoStats, _ := cpu.Info()
	cpuInfo := new(CPUInfo)
	cpuInfo.Count, _ = cpu.Counts(true)
	for _, infoStat := range cpuInfoStats {
		core := &Core{
			CPU:       infoStat.CPU,
			CoreID:    infoStat.CoreID,
			ModelName: infoStat.ModelName,
			Mhz:       infoStat.Mhz,
			Flags:     infoStat.Flags,
		}
		cpuInfo.AppendCore(core)
	}
	return cpuInfo
}
