package model

type ProcfsCPUInfo struct {
	Processor       uint
	VendorID        string
	CPUFamily       string
	Model           string
	ModelName       string
	Stepping        string
	Microcode       string
	CPUMHz          float64
	CacheSize       string
	PhysicalID      string
	Siblings        uint
	CoreID          string
	CPUCores        uint
	APICID          string
	InitialAPICID   string
	FPU             string
	FPUException    string
	CPUIDLevel      uint
	WP              string
	Flags           []string
	Bugs            []string
	BogoMips        float64
	CLFlushSize     uint
	CacheAlignment  uint
	AddressSizes    string
	PowerManagement string
}

func (cinfo ProcfsCPUInfo) PrintInfo() {
	println(cinfo.Processor)
}
