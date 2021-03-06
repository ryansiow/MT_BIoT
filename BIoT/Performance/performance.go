package main

import (
  "fmt"
  "time"

  "github.com/shirou/gopsutil/cpu"
  "github.com/shirou/gopsutil/load"
  "github.com/shirou/gopsutil/mem"
  "github.com/shirou/gopsutil/host"
  "github.com/shirou/gopsutil/disk"
  "github.com/shirou/gopsutil/net"
)

func main() {
fmt.Println("########################################################################")
getCpuInfo()
//getCpuLoad()
getMemInfo()
// getHostInfo()
getDiskInfo()
// getNetInfo()
}


// cpu info
func getCpuInfo() {
	// cpuInfos, err := cpu.Info()
	// if err != nil {
	// 	fmt.Printf("get cpu info failed, err:%v", err)
	// }
	// for _, ci := range cpuInfos {
	// 	fmt.Println(ci)
	// }
	//CPU utilization
	//for {
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("#CPU PERCENT:%v\n", percent)
	//}
}


func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}


// mem info
func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("#MEM INFO:%v\n", memInfo)
}


// host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}


// disk info
func getDiskInfo() {
	// parts, err := disk.Partitions(true)
	// if err != nil {
	// 	fmt.Printf("get Partitions failed, err:%v\n", err)
	// 	return
	// }
	// for _, part := range parts {
	// 	fmt.Printf("part:%v\n", part.String())
	// 	diskInfo, _ := disk.Usage(part.Mountpoint)
	// 	fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	// }

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("#IO %v:%v\n", k, v)
	}
}


func getNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}
