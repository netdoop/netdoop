package igd

type IPPingDiagnostics struct {
	Interface           *string `json:"Interface,omitempty" pn:"Interface"`
	DiagnosticsState    *string `json:"DiagnosticsState,omitempty" pn:"DiagnosticsState"`
	Host                *string `json:"Host,omitempty" pn:"Host"`
	NumberOfRepetitions *uint   `json:"NumberOfRepetitions,omitempty" pn:"NumberOfRepetitions"`
	Timeout             *uint   `json:"Timeout,omitempty" pn:"Timeout"`
	DataBlockSize       *uint   `json:"DataBlockSize,omitempty" pn:"DataBlockSize"`
	SuccessCount        *uint   `json:"SuccessCount,omitempty" pn:"SuccessCount"`
	FailureCount        *uint   `json:"FailureCount,omitempty" pn:"FailureCount"`
	AverageResponseTime *uint   `json:"AverageResponseTime,omitempty" pn:"AverageResponseTime"`
	MinimumResponseTime *uint   `json:"MinimumResponseTime,omitempty" pn:"MinimumResponseTime"`
	MaximumResponseTime *uint   `json:"MaximumResponseTime,omitempty" pn:"MaximumResponseTime"`
	DSCP                *uint   `json:"DSCP,omitempty" pn:"DSCP"`
}

type TraceRouteDiagnostics struct {
	DiagnosticsState         *string `json:"DiagnosticsState,omitempty" pn:"DiagnosticsState"`
	Interface                *string `json:"Interface,omitempty" pn:"Interface"`
	Host                     *string `json:"Host,omitempty" pn:"Host"`
	NumberOfTries            *uint   `json:"NumberOfTries,omitempty" pn:"NumberOfTries"`
	Timeout                  *uint   `json:"Timeout,omitempty" pn:"Timeout"`
	MaxHopCount              *uint   `json:"MaxHopCount,omitempty" pn:"MaxHopCount"`
	DataBlockSize            *uint   `json:"DataBlockSize,omitempty" pn:"DataBlockSize"`
	ResponseTime             *uint   `json:"ResponseTime,omitempty" pn:"ResponseTime"`
	RouteHopsNumberOfEntries *uint   `json:"RouteHopsNumberOfEntries,omitempty" pn:"RouteHopsNumberOfEntries"`
	RouteHops                map[int]*struct {
		HopHost        *string `json:"HopHost,omitempty" pn:"HopHost"`
		HopHostAddress *string `json:"HopHostAddress,omitempty" pn:"HopHostAddress"`
		HopErrorCode   *uint   `json:"HopErrorCode,omitempty" pn:"HopErrorCode"`
		HopRTTimes     *string `json:"HopRTTimes,omitempty" pn:"HopRTTimes"`
	} `json:"RouteHops,omitempty" pn:"RouteHops"`
	DSCP *uint `json:"DSCP,omitempty" pn:"DSCP"`
}

type LinkDiagnostics struct {
	DiagnosticsState *string `json:"DiagnosticsState,omitempty" pn:"DiagnosticsState"`
}

type DownloadDiagnostics struct {
	DiagnosticsState    *string `json:"DiagnosticsState,omitempty" pn:"DiagnosticsState"`
	Interface           *string `json:"Interface,omitempty" pn:"Interface"`
	DownloadURL         *string `json:"DownloadURL,omitempty" pn:"DownloadURL"`
	DSCP                *uint   `json:"DSCP,omitempty" pn:"DSCP"`
	EthernetPriority    *uint   `json:"EthernetPriority,omitempty" pn:"EthernetPriority"`
	ROMTime             *string `json:"ROMTime,omitempty" pn:"ROMTime"`
	BOMTime             *string `json:"BOMTime,omitempty" pn:"BOMTime"`
	EOMTime             *string `json:"EOMTime,omitempty" pn:"EOMTime"`
	TestBytesReceived   *uint   `json:"TestBytesReceived,omitempty" pn:"TestBytesReceived"`
	TotalBytesReceived  *uint   `json:"TotalBytesReceived,omitempty" pn:"TotalBytesReceived"`
	TCPOpenRequestTime  *string `json:"TCPOpenRequestTime,omitempty" pn:"TCPOpenRequestTime"`
	TCPOpenResponseTime *string `json:"TCPOpenResponseTime,omitempty" pn:"TCPOpenResponseTime"`
}

type UploadDiagnostics struct {
	DiagnosticsState    *string `json:"DiagnosticsState,omitempty" pn:"DiagnosticsState"`
	Interface           *string `json:"Interface,omitempty" pn:"Interface"`
	UploadURL           *string `json:"UploadURL,omitempty" pn:"UploadURL"`
	DSCP                *uint   `json:"DSCP,omitempty" pn:"DSCP"`
	EthernetPriority    *uint   `json:"EthernetPriority,omitempty" pn:"EthernetPriority"`
	TestFileLength      *uint   `json:"TestFileLength,omitempty" pn:"TestFileLength"`
	ROMTime             *string `json:"ROMTime,omitempty" pn:"ROMTime"`
	BOMTime             *string `json:"BOMTime,omitempty" pn:"BOMTime"`
	EOMTime             *string `json:"EOMTime,omitempty" pn:"EOMTime"`
	TotalBytesSent      *uint   `json:"TotalBytesSent,omitempty" pn:"TotalBytesSent"`
	TCPOpenRequestTime  *string `json:"TCPOpenRequestTime,omitempty" pn:"TCPOpenRequestTime"`
	TCPOpenResponseTime *string `json:"TCPOpenResponseTime,omitempty" pn:"TCPOpenResponseTime"`
}
