package util

import (
	"github.com/jaypipes/ghw"
)

func CPUInfo() string {
	if cpu, err := ghw.CPU(); err == nil {
		if len(cpu.Processors) > 0 {
			return cpu.Processors[0].Model
		}
	}

	return "Unknown"
}

func GPUInfo() string {
	if gpu, err := ghw.GPU(); err == nil {
		if len(gpu.GraphicsCards) > 0 {
			return gpu.GraphicsCards[0].DeviceInfo.Product.Name
		}
	}

	return "Unknown"
}

func MemoryInfo() string {
	if mem, err := ghw.Memory(); err == nil {
		var str string
		if len(mem.Modules) > 0 {
			str = mem.Modules[0].Vendor + " "
		}
		str += mem.String()[7:]
		return str
	}

	return "Unknown"
}
