package igd

type Statistics struct {
	TotalUpload          *string `json:"TotalUpload,omitempty" pn:"TotalUpload"`
	TotalDownload        *string `json:"TotalDownload,omitempty" pn:"TotalDownload"`
	ThroughputStatistics *struct {
		Reset *string `json:"Reset,omitempty" pn:"Reset"`
	} `json:"ThroughputStatistics,omitempty" pn:"ThroughputStatistics"`
	ThroughputStatisticsList map[int]*ThroughputStatisticsList `json:"ThroughputStatisticsList,omitempty" pn:"ThroughputStatisticsList"`
	DeviceList               map[int]*struct {
		Type       *string `json:"Type,omitempty" pn:"Type"`
		DeviceName *string `json:"DeviceName,omitempty" pn:"DeviceName"`
		MACAddress *string `json:"MACAddress,omitempty" pn:"MACAddress"`
		IPAddress  *string `json:"IPAddress,omitempty" pn:"IPAddress"`
		LeaseTime  *uint   `json:"LeaseTime,omitempty" pn:"LeaseTime"`
	} `json:"DeviceList,omitempty" pn:"DeviceList"`
}
