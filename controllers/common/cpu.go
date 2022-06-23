package common

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
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

	//fmt.Println(cpuInfos)

	// CPU使用率
	/*
		for {
			percent, _ := cpu.Percent(time.Second, false)
			fmt.Printf("cpu percent:%v\n", percent)
		}
	*/
	h, _ := host.Info()
	fmt.Println(h)
	//内存
	r, _ := mem.VirtualMemory()
	fmt.Println(r)

}
