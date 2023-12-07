package igd

type Overview struct {
	SystemInfo  *SystemInfo  `json:"SystemInfo,omitempty" pn:"SystemInfo"`
	VersionInfo *VersionInfo `json:"VersionInfo,omitempty" pn:"VersionInfo"`
	ModuleInfo  *ModuleInfo  `json:"ModuleInfo,omitempty" pn:"ModuleInfo"`

	InternetStatus *InternetStatus     `json:"InternetStatus,omitempty" pn:"InternetStatus"`
	LANStatus      *LANStatus          `json:"LANStatus,omitempty" pn:"LANStatus"`
	WiFiStatus     map[int]*WiFiStatus `json:"WiFiStatus,omitempty" pn:"WiFiStatus"`
	Status4G       *Status4G           `json:"4GStatus,omitempty" pn:"4GStatus"`
	WANStatus      *WANStatus          `json:"WANStatus,omitempty" pn:"WANStatus"`
	GPS            *GPS                `json:"GPS,omitempty" pn:"GPS"`

	CPUUsage                 *CPUUsage                         `json:"CPUUsage,omitempty" pn:"CPUUsage"`
	MemoryUsage              *MemoryUsage                      `json:"MemoryUsage,omitempty" pn:"MmoryUsage"`
	ThroughputStatisticsList map[int]*ThroughputStatisticsList `json:"ThroughputStatisticsList,omitempty" pn:"ThroughputStatisticsList"`
	DeviceList               map[int]*DeviceList               `json:"DeviceList,omitempty" pn:"DeviceList"`
	WANStatusList            map[int]*WANStatusList            `json:"WANStatusList,omitempty" pn:"WANStatusList"`
}

type GPS struct {
	Longitude *string `json:"Longitude,omitempty" pn:"Longitude"`
	Latitude  *string `json:"Latitude,omitempty" pn:"Latitude"`
	Altitude  *string `json:"Altitude,omitempty" pn:"Altitude"`
}

type SystemInfo struct {
	OnlineTime          *string `json:"OnlineTime,omitempty" pn:"OnlineTime"`
	ConnectionTotalTime *string `json:"ConnectionTotalTime,omitempty" pn:"ConnectionTotalTime"`
	RunTime             *string `json:"RunTime,omitempty" pn:"RunTime"`
}

type ModuleInfo struct {
	IMEI    *string `json:"IMEI,omitempty" pn:"IMEI"`
	IMSI    *string `json:"IMSI,omitempty" pn:"IMSI"`
	Version *string `json:"Version,omitempty" pn:"Version"`
	Model   *string `json:"Model,omitempty" pn:"Model"`
}

type VersionInfo struct {
	ProductModel          *string `json:"ProductModel,omitempty" pn:"ProductModel"`
	HardVersion           *string `json:"HardVersion,omitempty" pn:"HardVersion"`
	SoftwareVersion       *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
	UBOOTVersion          *string `json:"UBOOTVersion,omitempty" pn:"UBOOTVersion"`
	SerialNumber          *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
	BoardSN               *string `json:"BoardSN,omitempty" pn:"BoardSN"`
	ProductName           *string `json:"ProductName,omitempty" pn:"ProductName"`
	BackupSoftwareVersion *string `json:"BackupSoftwareVersion,omitempty" pn:"BackupSoftwareVersion"`
}

type LANStatus struct {
	MACAddress *string `json:"MACAddress,omitempty" pn:"MACAddress"`
	IPAddress  *string `json:"IPAddress,omitempty" pn:"IPAddress"`
	SubnetMask *string `json:"SubnetMask,omitempty" pn:"SubnetMask"`
}

type InternetStatus struct {
	Mode     *string `json:"Mode,omitempty" pn:"Mode"`
	RSRP     *string `json:"RSRP,omitempty" pn:"RSRP"`
	RSRQ     *string `json:"RSRQ,omitempty" pn:"RSRQ"`
	CellID   *string `json:"CellID,omitempty" pn:"CellID"`
	Status   *string `json:"Status,omitempty" pn:"Status"`
	Operator *string `json:"Operator,omitempty" pn:"Operator"`
}

type WANStatus struct {
	DLRateCurrent   *string `json:"DLRateCurrent,omitempty" pn:"DLRateCurrent"`
	DLRateMin       *string `json:"DLRateMin,omitempty" pn:"DLRateMin"`
	DLRateMax       *string `json:"DLRateMax,omitempty" pn:"DLRateMax"`
	ULRateCurrent   *string `json:"ULRateCurrent,omitempty" pn:"ULRateCurrent"`
	ULRateMin       *string `json:"ULRateMin,omitempty" pn:"ULRateMin"`
	ULRateMax       *string `json:"ULRateMax,omitempty" pn:"ULRateMax"`
	MaxDLThroughput *string `json:"MaxDLThroughput,omitempty" pn:"MaxDLThroughput"`
	MaxULThroughput *string `json:"MaxULThroughput,omitempty" pn:"MaxULThroughput"`
}

type WiFiStatus struct {
	Enable  *string `json:"Enable,omitempty" pn:"Enable"`
	Mode    *string `json:"Mode,omitempty" pn:"Mode"`
	Band    *string `json:"Band,omitempty" pn:"Band"`
	Channel *string `json:"Channel,omitempty" pn:"Channel"`
	SSID    *string `json:"SSID,omitempty" pn:"SSID"`
}

type Status4G struct {
	Status   *string `json:"Status,omitempty" pn:"Status"`
	Operator *string `json:"Operator,omitempty" pn:"Operator"`
	Mode     *string `json:"Mode,omitempty" pn:"Mode"`
	CellID   *string `json:"CellID,omitempty" pn:"CellID"`
	RSRP0    *string `json:"RSRP0,omitempty" pn:"RSRP0"`
	RSRP1    *string `json:"RSRP1,omitempty" pn:"RSRP1"`
	RSRQ     *string `json:"RSRQ,omitempty" pn:"RSRQ"`
}

type CPUUsage struct {
	Current *string `json:"Current,omitempty" pn:"Current"`
	Max     *string `json:"Max,omitempty" pn:"Max"`
	Min     *string `json:"Min,omitempty" pn:"Min"`
}

type MemoryUsage struct {
	TotalMemory *string `json:"TotalMemory,omitempty" pn:"TotalMemory"`
	Current     *string `json:"Current,omitempty" pn:"Current"`
	Max         *string `json:"Max,omitempty" pn:"Max"`
	Min         *string `json:"Min,omitempty" pn:"Min"`
}

type ThroughputStatistics struct {
	Reset *int `json:"Reset,omitempty" pn:"Reset"`
}

type ThroughputStatisticsList struct {
	Index           *string `json:"index,omitempty" pn:"index"`
	Port            *string `json:"Port,omitempty" pn:"Port"`
	ReceivedTotal   *string `json:"ReceivedTotal,omitempty" pn:"ReceivedTotal"`
	ReceivedPackets *string `json:"ReceivedPackets,omitempty" pn:"ReceivedPackets"`
	ReceivedErrors  *string `json:"ReceivedErrors,omitempty" pn:"ReceivedErrors"`
	ReceivedDropped *string `json:"ReceivedDropped,omitempty" pn:"ReceivedDropped"`
	SentTotal       *string `json:"SentTotal,omitempty" pn:"SentTotal"`
	SentPackets     *string `json:"SentPackets,omitempty" pn:"SentPackets"`
	SentErrors      *string `json:"SentErrors,omitempty" pn:"SentErrors"`
	SentDropped     *string `json:"SentDropped,omitempty" pn:"SentDropped"`
}

type DeviceList struct {
	Index      *string `json:"index,omitempty" pn:"index"`
	DeviceName *string `json:"DeviceName,omitempty" pn:"DeviceName"`
	MACAddress *string `json:"MACAddress,omitempty" pn:"MACAddress"`
	IPAddress  *string `json:"IPAddress,omitempty" pn:"IPAddress"`
	LeaseTime  *string `json:"LeaseTime,omitempty" pn:"LeaseTime"`
	Type       *string `json:"Type,omitempty" pn:"Type"`
}

type WANStatusList struct {
	Index   *string `json:"index,omitempty" pn:"index"`
	APNName *string `json:"APNName,omitempty" pn:"APNName"`
	Status  *string `json:"Status,omitempty" pn:"Status"`
	IP      *string `json:"IP,omitempty" pn:"IP"`
	Mask    *string `json:"Mask,omitempty" pn:"Mask"`
}
