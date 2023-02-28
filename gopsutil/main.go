package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"time"
)

func main() {

	h, _ := host.Info()
	fmt.Printf("Host: %v\n", h)

	cpuInfo, _ := cpu.Info()
	count, _ := cpu.Counts(false)
	fmt.Printf("CpuInfo : %v\n", cpuInfo[0].ModelName)
	fmt.Printf("CpuCores : %v\n", count)
	cpuPercent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("CpuUsedPercent:%v%%\n", float32(cpuPercent[0]))

	memory, _ := mem.VirtualMemory()
	fmt.Printf("RamSize:%vGB\n", float32(memory.Total)/1024/1024/1024)
	fmt.Printf("RamUsed:%vGB\n", float32(memory.Used)/1024/1024/1024)

	parts, _ := disk.Partitions(false)
	var diskSize uint64
	var diskUsed uint64
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		diskSize += diskInfo.Total
		diskUsed += diskInfo.Used
	}
	fmt.Printf("DiskSzie:%vGB\n", float32(diskSize)/1024/1024/1024)
	fmt.Printf("DiskUsed:%vGB\n", float32(diskUsed)/1024/1024/1024)

	on, _ := net.Connections("tcp")
	fmt.Printf("OS Proc Net Connect Nums: %#v\n", len(on))

	proc, _ := process.NewProcess(int32(os.Getpid()))
	pc, _ := proc.CPUPercent()
	fmt.Printf("Proc %d uses %v%% CPU\n", proc.Pid, float32(pc))

	pm, _ := proc.MemoryPercent()
	fmt.Printf("Proc %d uses %v%% MEM\n", proc.Pid, pm)

}
