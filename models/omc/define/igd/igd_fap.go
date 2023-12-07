package igd

type FAP struct {
	GPS           *FAP_GPS       `json:"GPS,omitempty" pn:"GPS"`
	CDRMgmt       *CDRMgmt       `json:"CDRMgmt,omitempty" pn:"CDRMgmt"`
	MRMgmt        *MRMgmt        `json:"MRMgmt,omitempty" pn:"MRMgmt"`
	NRMMgmt       *NRMMgmt       `json:"NRMMgmt,omitempty" pn:"NRMMgmt"`
	PerfMgmt      *PerfMgmt      `json:"PerfMgmt,omitempty" pn:"PerfMgmt"`
	Tunnel        *Tunnel        `json:"Tunnel,omitempty" pn:"Tunnel"`
	CellTraceMgmt *CellTraceMgmt `json:"CellTraceMgmt,omitempty" pn:"CellTraceMgmt"`
	SCMgmt        *string        `json:"SCMgmt,omitempty" pn:"SCMgmt"`
}

type CDRMgmt struct {
	Enable                 *string `json:"Enable,omitempty" pn:"Enable"`
	Password               *string `json:"Password,omitempty" pn:"Password"`
	PeriodicUploadInterval *string `json:"PeriodicUploadInterval,omitempty" pn:"PeriodicUploadInterval"`
	PeriodicUploadTime     *string `json:"PeriodicUploadTime,omitempty" pn:"PeriodicUploadTime"`
	URL                    *string `json:"URL,omitempty" pn:"URL"`
	Username               *string `json:"Username,omitempty" pn:"Username"`
}

type FAP_GPS struct {
	ContinuousGPSStatus *struct {
		GPSReadHardWare    *string `json:"GPSReadHardWare,omitempty" pn:"GPSReadHardWare"`
		X_VENDOR_Latitude  *string `json:"X_VENDOR_Latitude,omitempty" pn:"X_VENDOR_Latitude"`
		X_VENDOR_Longitude *string `json:"X_VENDOR_Longitude,omitempty" pn:"X_VENDOR_Longitude"`
	} `json:"ContinuousGPSStatus,omitempty" pn:"ContinuousGPSStatus"`
	LockedLatitude     *string `json:"LockedLatitude,omitempty" pn:"LockedLatitude"`
	LockedLongitude    *string `json:"LockedLongitude,omitempty" pn:"LockedLongitude"`
	NumberOfSatellites *string `json:"NumberOfSatellites,omitempty" pn:"NumberOfSatellites"`
}

type MRMgmt struct {
	Config map[int]*MRMgmtConfig `json:"Config,omitempty" pn:"Config"`
}

type MRMgmtConfig struct {
	MeasureType     *string `json:"MeasureType,omitempty" pn:"MeasureType"`
	MrEnable        *string `json:"MrEnable,omitempty" pn:"MrEnable"`
	MrPassword      *string `json:"MrPassword,omitempty" pn:"MrPassword"`
	MrUrl           *string `json:"MrUrl,omitempty" pn:"MrUrl"`
	MrUsername      *string `json:"MrUsername,omitempty" pn:"MrUsername"`
	OmcName         *string `json:"OmcName,omitempty" pn:"OmcName"`
	PrbNum          *string `json:"PrbNum,omitempty" pn:"PrbNum"`
	SampleBeginTime *string `json:"SampleBeginTime,omitempty" pn:"SampleBeginTime"`
	SampleEndTime   *string `json:"SampleEndTime,omitempty" pn:"SampleEndTime"`
	SamplePeriod    *string `json:"SamplePeriod,omitempty" pn:"SamplePeriod"`
	SubFrameNum     *string `json:"SubFrameNum,omitempty" pn:"SubFrameNum"`
	UploadPeriod    *string `json:"UploadPeriod,omitempty" pn:"UploadPeriod"`
	MRNCGIList      *string `json:"MRNCGIList,omitempty" pn:"MRNCGIList"`
	UserLabel       *string `json:"UserLabel,omitempty" pn:"UserLabel"`
	MeasureItems    *string `json:"MeasureItems,omitempty" pn:"MeasureItems"`
}
type NRMMgmt struct {
	Enable                 *string `json:"Enable,omitempty" pn:"Enable"`
	PeriodicUploadInterval *string `json:"PeriodicUploadInterval,omitempty" pn:"PeriodicUploadInterval"`
	URL                    *string `json:"URL,omitempty" pn:"URL"`
}

type PerfMgmt struct {
	ALL_ADORN *string                 `json:"ALL_ADORN,omitempty" pn:"ALL_ADORN"`
	Config    map[int]*PerfMgmtConfig `json:"Config,omitempty" pn:"Config"`
}

type PerfMgmtConfig struct {
	Alias                  *string `json:"Alias,omitempty" pn:"Alias"`
	Enable                 *string `json:"Enable,omitempty" pn:"Enable"`
	Password               *string `json:"Password,omitempty" pn:"Password"`
	PeriodicUploadInterval *string `json:"PeriodicUploadInterval,omitempty" pn:"PeriodicUploadInterval"`
	PeriodicUploadTime     *string `json:"PeriodicUploadTime,omitempty" pn:"PeriodicUploadTime"`
	URL                    *string `json:"URL,omitempty" pn:"URL"`
	Username               *string `json:"Username,omitempty" pn:"Username"`
	ReplenishEndTime       *string `json:"ReplenishEndTime,omitempty" pn:"ReplenishEndTime"`
	ReplenishEnable        *string `json:"ReplenishEnable,omitempty" pn:"ReplenishEnable"`
}

type Tunnel struct {
	CryptoProfile map[int]*struct {
		AkaAuthType        *string `json:"Aka_Auth_Type,omitempty" pn:"Aka_Auth_Type"`
		AkaC1              *string `json:"Aka_C1,omitempty" pn:"Aka_C1"`
		AkaC2              *string `json:"Aka_C2,omitempty" pn:"Aka_C2"`
		AkaC3              *string `json:"Aka_C3,omitempty" pn:"Aka_C3"`
		AkaC4              *string `json:"Aka_C4,omitempty" pn:"Aka_C4"`
		AkaC5              *string `json:"Aka_C5,omitempty" pn:"Aka_C5"`
		AkaKey             *string `json:"Aka_Key,omitempty" pn:"Aka_Key"`
		AkaOpc             *string `json:"Aka_Opc,omitempty" pn:"Aka_Opc"`
		AkaR1              *string `json:"Aka_R1,omitempty" pn:"Aka_R1"`
		AkaR2              *string `json:"Aka_R2,omitempty" pn:"Aka_R2"`
		AkaR3              *string `json:"Aka_R3,omitempty" pn:"Aka_R3"`
		AkaR4              *string `json:"Aka_R4,omitempty" pn:"Aka_R4"`
		AkaR5              *string `json:"Aka_R5,omitempty" pn:"Aka_R5"`
		AuthMethod         *string `json:"AuthMethod,omitempty" pn:"AuthMethod"`
		CertificateKeyFile *string `json:"CertificateKeyFile,omitempty" pn:"CertificateKeyFile"`
		CertificateName    *string `json:"CertificateName,omitempty" pn:"CertificateName"`
		CertificateTheme   *string `json:"CertificateTheme,omitempty" pn:"CertificateTheme"`
		ESPEncrypt         *string `json:"ESPEncrypt,omitempty" pn:"ESPEncrypt"`
		IKEEncrypt         *string `json:"IKEEncrypt,omitempty" pn:"IKEEncrypt"`
		LeftId             *string `json:"LeftId,omitempty" pn:"LeftId"`
		PskKey             *string `json:"PskKey,omitempty" pn:"PskKey"`
		RightCert          *string `json:"RightCert,omitempty" pn:"RightCert"`
		RightId            *string `json:"RightId,omitempty" pn:"RightId"`
		RightSubnet        *string `json:"RightSubnet,omitempty" pn:"RightSubnet"`
		SecureGateWayIp    *string `json:"SecureGateWayIp,omitempty" pn:"SecureGateWayIp"`
	} `json:"CryptoProfile,omitempty" pn:"CryptoProfile"`
	CryptoProfileNumberOfEntries *uint `json:"CryptoProfileNumberOfEntries,omitempty" pn:"CryptoProfileNumberOfEntries"`
}

type CellTraceMgmt struct {
	Config map[int]*struct {
		CellId            *string `json:"CellId,omitempty" pn:"CellId,omitempty"`
		Enable            *string `json:"Enable,omitempty" pn:"Enable,omitempty"`
		InterfacesToTrace *string `json:"InterfacesToTrace,omitempty" pn:"InterfacesToTrace,omitempty"`
		TraceReference    *string `json:"TraceReference,omitempty" pn:"TraceReference,omitempty"`
		TraceDepth        *string `json:"TraceDepth,omitempty" pn:"TraceDepth,omitempty"`
		CallTraceUrl      *string `json:"CallTraceUrl,omitempty" pn:"CallTraceUrl,omitempty"`
		CallTraceUsername *string `json:"CallTraceUsername,omitempty" pn:"CallTraceUsername,omitempty"`
		CallTracePassword *string `json:"CallTracePassword,omitempty" pn:"CallTracePassword,omitempty"`
		SamplePeriod      *string `json:"SamplePeriod,omitempty" pn:"SamplePeriod,omitempty"`
		UploadPeriod      *string `json:"UploadPeriod,omitempty" pn:"UploadPeriod,omitempty"`
	} `json:"Config,omitempty" pn:"Config,omitempty"`
}
