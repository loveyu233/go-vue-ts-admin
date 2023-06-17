package test

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestSys(t *testing.T) {
	// 获取CPU使用率
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("CPU使用率:%.3f%%\n", cpuPercent[0])

	// 获取内存使用率
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("内存使用率:%.3f%%\n", memInfo.UsedPercent)

	// 获取环境变量
	//envVars := os.Environ()
	//for _, envVar := range envVars {
	//	fmt.Println(envVar)
	//}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("主机名:", hostname)
	fmt.Println("系统类型:", runtime.GOOS)

	info, err := host.Info()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("系统运行时间:%d小时\n", info.Uptime/60/60)
	fmt.Println("========================")

	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usage, err := disk.Usage(partitions[0].Mountpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("路径:", usage.Path)
	fmt.Println("共:", usage.Total/(1024*1024*1024))
	fmt.Println("使用:", usage.Used/(1024*1024*1024))
	fmt.Println("剩余:", usage.Free/(1024*1024*1024))
	fmt.Printf("百分比:%.3f%%\n", usage.UsedPercent)
}
