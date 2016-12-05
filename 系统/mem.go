package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	//v, _ := mem.VirtualMemory()

	//almost every return value is a struct
	//fmt.Printf("Total:%v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	//convert to JSON. String() is also implemented
	//fmt.Println(v)
	t, _ := cpu.Times(false)
	fmt.Println(t)
	x, _ := disk.Partitions(true)
	//fmt.Println(x)
	fmt.Println(x)
	a, _ := disk.Usage("C:")
	fmt.Println(a)

	//fmt.Println(pro.)

}
