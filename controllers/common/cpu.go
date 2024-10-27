package common

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func Getcpu() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci.Cores)
		fmt.Println(ci.ModelName)
	}

	//fmt.Println("cpu:", cpuInfos)
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	// h, _ := host.Info()
	// fmt.Println("host:", h)
	//内存
	r, _ := mem.VirtualMemory()
	fmt.Println("memory:", r.Total, r.Used)
	info, _ := load.Avg()
	fmt.Printf("cpuavg:%v\n", info)
	//CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu 使用率:%v\n", percent)
	}
}

func Getcpuload() {
	info, _ := load.Avg()
	fmt.Printf("cpuavg:%v\n", info)
}
