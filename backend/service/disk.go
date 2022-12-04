package service

import "github.com/shirou/gopsutil/v3/disk"

type Partition struct {
	Device      string  `json:"device"`
	MountPoint  string  `json:"mount_point"`
	Size        uint64  `json:"size"`
	UsedPercent float64 `json:"used_percent"`
}
type DiskInfo struct {
	Partitions []*Partition `json:"partitions"`
}

func (diskInfo *DiskInfo) AppendPartition(partition *Partition) {
	diskInfo.Partitions = append(diskInfo.Partitions, partition)
}

func GetDiskInfo() *DiskInfo {
	diskInfo := new(DiskInfo)
	partitionStats, _ := disk.Partitions(false)
	for _, partitionStat := range partitionStats {
		usageStat, _ := disk.Usage(partitionStat.Mountpoint)
		partition := &Partition{
			Device:      partitionStat.Device,
			MountPoint:  partitionStat.Mountpoint,
			Size:        usageStat.Total,
			UsedPercent: usageStat.UsedPercent,
		}
		diskInfo.AppendPartition(partition)
	}
	return diskInfo
}
