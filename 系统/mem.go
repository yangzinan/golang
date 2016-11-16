package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	//v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	//fmt.Printf("Total:%v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	//fmt.Println(v)
	t, _ := cpu.Times(false)
	fmt.Println(t[0].User)
}
