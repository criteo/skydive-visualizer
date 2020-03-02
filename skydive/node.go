package skydive

type NodeResponse struct {
	Nodes []Node `json:"Nodes"`
}

type Node struct {
	ID        string   `json:"ID"`
	Metadata  Metadata `json:"Metadata,omitempty"`
	Host      string   `json:"Host"`
	Origin    string   `json:"Origin"`
	CreatedAt int64    `json:"CreatedAt"`
	UpdatedAt int64    `json:"UpdatedAt"`
	DeletedAt int64    `json:"DeletedAt"`
	Revision  int      `json:"Revision"`
}

type Metadata struct {
	Name                 string        `json:"Name"`
	KernelVersion        string        `json:"KernelVersion"`
	Platform             string        `json:"Platform"`
	PlatformVersion      string        `json:"PlatformVersion"`
	Hostname             string        `json:"Hostname"`
	Sockets              []Socket      `json:"Sockets"`
	Type                 string        `json:"Type"`
	PlatformFamily       string        `json:"PlatformFamily"`
	TID                  string        `json:"TID"`
	CPU                  []CPU         `json:"CPU"`
	KernelCmdLine        KernelCMDLine `json:"KernelCmdLine"`
	OS                   string        `json:"OS"`
	VirtualizationRole   string        `json:"VirtualizationRole"`
	VirtualizationSystem string        `json:"VirtualizationSystem"`
}

type Socket struct {
	LocalAddress  string `json:"LocalAddress"`
	LocalPort     int    `json:"LocalPort"`
	Name          string `json:"Name"`
	PID           int    `json:"Pid"`
	Process       string `json:"Process"`
	Protocol      string `json:"Protocol"`
	RemoteAddress string `json:"RemoteAddress"`
	RemotePort    int    `json:"RemotePort"`
	State         string `json:"State"`
}

type CPU struct {
	CPU        int    `json:"CPU"`
	CacheSize  int    `json:"CacheSize"`
	CoreID     string `json:"CoreID"`
	Cores      int    `json:"Cores"`
	Family     string `json:"Family"`
	Mhz        int    `json:"Mhz"`
	Microcode  string `json:"Microcode"`
	Model      string `json:"Model"`
	ModelName  string `json:"ModelName"`
	PhysicalID string `json:"PhysicalID"`
	Stepping   int    `json:"Stepping"`
	VendorID   string `json:"VendorID"`
}

type KernelCMDLine struct {
	BOOTIMAGE string `json:"BOOT_IMAGE"`
	Console   string `json:"console"`
	Ro        bool   `json:"ro"`
	Root      string `json:"root"`
}

type Metric struct {
	Multicast int   `json:"Multicast"`
	RxBytes   int   `json:"RxBytes"`
	RxPackets int   `json:"RxPackets"`
	TxBytes   int   `json:"TxBytes"`
	TxPackets int   `json:"TxPackets"`
	Start     int64 `json:"Start"`
	Last      int64 `json:"Last"`
}
