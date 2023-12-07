package igd

import (
	"github.com/netdoop/netdoop/utils"
)

type InternetGatewayDevice struct {
	WebGui *WebGui `json:"WebGui,omitempty" pn:"WEB_GUI"`

	DeviceSummary        *string `json:"DeviceSummary,omitempty" pn:"DeviceSummary"`
	RootDataModelVersion *string `json:"RootDataModelVersion,omitempty" pn:"RootDataModelVersion"`

	DeviceInfo       *DeviceInfo       `json:"DeviceInfo,omitempty" pn:"DeviceInfo"`
	ManagementServer *ManagementServer `json:"ManagementServer,omitempty" pn:"ManagementServer"`
	Time             *Time             `json:"Time,omitempty" pn:"Time"`
	IPsec            *IPsec            `json:"IPsec,omitempty" pn:"IPsec"`
	LogMgmt          *LogMgmt          `json:"LogMgmt,omitempty" pn:"LogMgmt"`
	UDPEchoConfig    *UDPEchoConfig    `json:"UDPEchoConfig,omitempty" pn:"UDPEchoConfig"`
	SoftwareCtrl     *SoftwareCtrl     `json:"SoftwareCtrl,omitempty" pn:"SoftwareCtrl"`
	SyncMode         *SyncMode         `json:"SyncMode,omitempty" pn:"SyncMode"`
	Ethernet         *Ethernet         `json:"Ethernet,omitempty" pn:"Ethernet"`

	FaultMgmt *FaultMgmt `json:"FaultMgmt,omitempty" pn:"FaultMgmt"`

	IPPingDiagnostics     *IPPingDiagnostics     `json:"IPPingDiagnostics,omitempty" pn:"IPPingDiagnostics"`
	TraceRouteDiagnostics *TraceRouteDiagnostics `json:"TraceRouteDiagnostics,omitempty" pn:"TraceRouteDiagnostics"`
	LinkDiagnostics       *LinkDiagnostics       `json:"LinkDiagnostics,omitempty" pn:"LinkDiagnostics"`
	DownloadDiagnostics   *DownloadDiagnostics   `json:"DownloadDiagnostics,omitempty" pn:"DownloadDiagnostics"`
	UploadDiagnostics     *UploadDiagnostics     `json:"UploadDiagnostics,omitempty" pn:"UploadDiagnostics"`

	InterfaceStackNumberOfEntries *uint              `json:"InterfaceStackNumberOfEntries,omitempty" pn:"InterfaceStackNumberOfEntries"`
	LANDeviceNumberOfEntries      *uint              `json:"LANDeviceNumberOfEntries,omitempty" pn:"LANDeviceNumberOfEntries"`
	LANDevice                     map[int]*LANDevice `json:"LANDevice,omitempty" pn:"LANDevice"`
	WANDeviceNumberOfEntries      *uint              `json:"WANDeviceNumberOfEntries,omitempty" pn:"WANDeviceNumberOfEntries"`
	WANDevice                     map[int]*WANDevice `json:"WANDevice,omitempty" pn:"WANDevice"`

	FAP      *FAP `json:"FAP,omitempty" pn:"FAP"`
	Services *struct {
		FAPService map[int]*FAPService `json:"FAPService,omitempty" pn:"FAPService"`
	} `json:"Services,omitempty" pn:"Services"`

	CustomConfig *struct {
		Custom_TA *string `json:"Custom_TA,omitempty" pn:"Custom_TA"`
	}

	WAN       *WAN  `json:"WAN,omitempty" pn:"WAN"`
	AU_TYPE   *uint `json:"AU_TYPE,omitempty" pn:"au_type"`
	EURU_TYPE *uint `json:"EURU_TYPE,omitempty" pn:"euru_update"`
	RU_TYPE   *uint `json:"RU_TYPE,omitempty" pn:"ru_type"`

	SelfCureMgmt *struct {
		EuSystemTimeSyncEnable         *string `json:"EuSystemTimeSyncEnable,omitempty" pn:"EuSystemTimeSyncEnable"`
		EuRebootRestartInterval        *string `json:"EuRebootRestartInterval,omitempty" pn:"EuRebootRestartInterval"`
		NoHeartbeatSelfCureSwitch      *string `json:"NoHeartbeatSelfCureSwitch,omitempty" pn:"NoHeartbeatSelfCureSwitch"`
		NoHeartbeatSelfCureDelayPeriod *string `json:"NoHeartbeatSelfCureDelayPeriod,omitempty" pn:"NoHeartbeatSelfCureDelayPeriod"`
	} `json:"SelfCureMgmt,omitempty" pn:"SelfCureMgmt"`

	Auth *struct {
		SerialNumberSec *string `json:"SerialNumberSec,omitempty" pn:"SerialNumberSec"`
	} `json:"Auth,omitempty" pn:"Auth"`
}

func (m *InternetGatewayDevice) ReadValues(values map[string]string) {
	utils.ReadParamList(values, m)
}

type DeviceInfo struct {
	Manufacturer                    *string  `json:"Manufacturer,omitempty" pn:"Manufacturer"`
	ManufacturerOUI                 *string  `json:"ManufacturerOUI,omitempty" pn:"ManufacturerOUI"`
	ModelName                       *string  `json:"ModelName,omitempty" pn:"ModelName"`
	Description                     *string  `json:"Description,omitempty" pn:"Description"`
	ProductClass                    *string  `json:"ProductClass,omitempty" pn:"ProductClass"`
	SerialNumber                    *string  `json:"SerialNumber,omitempty" pn:"SerialNumber"`
	HardwareVersion                 *string  `json:"HardwareVersion,omitempty" pn:"HardwareVersion"`
	SoftwareVersion                 *string  `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
	ModuleVersion                   *string  `json:"ModuleVersion,omitempty" pn:"ModuleVersion"`
	ModemFirmwareVersion            *string  `json:"ModemFirmwareVersion,omitempty" pn:"ModemFirmwareVersion"`
	SpecVersion                     *string  `json:"SpecVersion,omitempty" pn:"SpecVersion"`
	ProvisioningCode                *string  `json:"ProvisioningCode,omitempty" pn:"ProvisioningCode"`
	UpTime                          *float64 `json:"UpTime,omitempty" pn:"UpTime"`
	FirstUseDate                    *string  `json:"FirstUseDate,omitempty" pn:"FirstUseDate"`
	DeviceLog                       *string  `json:"DeviceLog,omitempty" pn:"DeviceLog"`
	VendorConfigFileNumberOfEntries *uint    `json:"VendorConfigFileNumberOfEntries,omitempty" pn:"VendorConfigFileNumberOfEntries"`
	ThreeGPPSpecVersion             *string  `json:"3GPPSpecVersion,omitempty" pn:"3GPPSpecVersion"`
	AddEndPoint                     *string  `json:"AddEndPoint,omitempty" pn:"AddEndPoint"`
	AdditionalHardwareVersion       *string  `json:"AdditionalHardwareVersion,omitempty" pn:"AdditionalHardwareVersion"`
	AdditionalSoftwareVersion       *string  `json:"AdditionalSoftwareVersion,omitempty" pn:"AdditionalSoftwareVersion"`
	BSMode                          *string  `json:"BSMode,omitempty" pn:"BSMode"`
	DelEndPonit                     *string  `json:"DelEndPonit,omitempty" pn:"DelEndPonit"`
	HardwarePlatform                *string  `json:"HardwarePlatform,omitempty" pn:"HardwarePlatform"`
	VendorLogFileNumberOfEntries    *uint    `json:"VendorLogFileNumberOfEntries,omitempty" pn:"VendorLogFileNumberOfEntries"`

	OpenAccountLocal *string `json:"OpenAccountLocal,omitempty" pn:"OpenAccountLocal"`
	License          *struct {
		ImportTime *string `json:"ImportTime,omitempty" pn:"ImportTime"`
		MinLeft    *string `json:"MinLeft,omitempty" pn:"MinLeft"`
		Status     *string `json:"Status,omitempty" pn:"Status"`
	} `json:"License,omitempty" pn:"License"`
	PHY *struct {
		NMode *string `json:"nMode,omitempty" pn:"nMode"`
	}
	EURUinfo *struct {
		Eunum *string `json:"Eunum,omitempty" pn:"eunum"`
		Runum *string `json:"Runum,omitempty" pn:"runum"`
	} `json:"EURUinfo,omitempty" pn:"EURUinfo"`
	EuRuAlarmEnable        *string `json:"EuRuAlarmEnable,omitempty" pn:"EuRuAlarmEnable"`
	X_VENDOR_AUT_IPADDR    *string `json:"X_VENDOR_AUT_IPADDR,omitempty" pn:"X_VENDOR_AUT_IPADDR"`
	X_VENDOR_DEVICE_TYPE   *string `json:"X_VENDOR_DEVICE_TYPE,omitempty" pn:"X_VENDOR_DEVICE_TYPE"`
	X_VENDOR_ENODEB_STATUS *struct {
		X_VENDOR_CELL_STATUS *string `json:"X_VENDOR_CELL_STATUS,omitempty" pn:"X_VENDOR_CELL_STATUS"`
		X_VENDOR_MMEACT_NUM  *string `json:"X_VENDOR_MMEACT_NUM,omitempty" pn:"X_VENDOR_MMEACT_NUM"`
		X_VENDOR_UEACT_NUM   *string `json:"X_VENDOR_UEACT_NUM,omitempty" pn:"X_VENDOR_UEACT_NUM"`
	} `json:"X_VENDOR_ENODEB_STATUS,omitempty" pn:"X_VENDOR_ENODEB_STATUS"`
	X_VENDOR_FEATURE_CODE      *string `json:"X_VENDOR_FEATURE_CODE,omitempty" pn:"X_VENDOR_FEATURE_CODE"`
	X_VENDOR_IPADDR_AUT_SWITCH *string `json:"X_VENDOR_IPADDR_AUT_SWITCH,omitempty" pn:"X_VENDOR_IPADDR_AUT_SWITCH"`

	MU map[int]*MU `json:"MU,omitempty" pn:"MU"`
}

type ManageableDevice struct {
	ManufacturerOUI *string `json:"ManufacturerOUI,omitempty" pn:"ManufacturerOUI"`
	SerialNumber    *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
	ProductClass    *string `json:"ProductClass,omitempty" pn:"ProductClass"`
}

type ManagementServer struct {
	EnableCWMP                                   *int                      `json:"EnableCWMP,omitempty" pn:"EnableCWMP"`
	URL                                          *string                   `json:"URL,omitempty" pn:"URL"`
	Username                                     *string                   `json:"Username,omitempty" pn:"Username"`
	Password                                     *string                   `json:"Password,omitempty" pn:"Password"`
	PeriodicInformEnable                         *string                   `json:"PeriodicInformEnable,omitempty" pn:"PeriodicInformEnable"`
	PeriodicInformInterval                       *uint                     `json:"PeriodicInformInterval,omitempty" pn:"PeriodicInformInterval"`
	PeriodicInformTime                           *string                   `json:"PeriodicInformTime,omitempty" pn:"PeriodicInformTime"`
	CWMPRetryIntervalMultiplier                  *uint                     `json:"CWMPRetryIntervalMultiplier,omitempty" pn:"CWMPRetryIntervalMultiplier"`
	CWMPRetryMinimumWaitInterval                 *uint                     `json:"CWMPRetryMinimumWaitInterval,omitempty" pn:"CWMPRetryMinimumWaitInterval"`
	ParameterKey                                 *string                   `json:"ParameterKey,omitempty" pn:"ParameterKey"`
	ConnectionRequestURL                         *string                   `json:"ConnectionRequestURL,omitempty" pn:"ConnectionRequestURL"`
	ConnectionRequestUsername                    *string                   `json:"ConnectionRequestUsername,omitempty" pn:"ConnectionRequestUsername"`
	ConnectionRequestPassword                    *string                   `json:"ConnectionRequestPassword,omitempty" pn:"ConnectionRequestPassword"`
	UDPConnectionRequestAddress                  *string                   `json:"UDPConnectionRequestAddress,omitempty" pn:"UDPConnectionRequestAddress"`
	UDPConnectionRequestAddressNotificationLimit *uint                     `json:"UDPConnectionRequestAddressNotificationLimit,omitempty" pn:"UDPConnectionRequestAddressNotificationLimit"`
	STUNEnable                                   *string                   `json:"STUNEnable,omitempty" pn:"STUNEnable"`
	STUNServerAddress                            *string                   `json:"STUNServerAddress,omitempty" pn:"STUNServerAddress"`
	STUNServerPort                               *uint                     `json:"STUNServerPort,omitempty" pn:"STUNServerPort"`
	STUNUsername                                 *string                   `json:"STUNUsername,omitempty" pn:"STUNUsername"`
	STUNPassword                                 *string                   `json:"STUNPassword,omitempty" pn:"STUNPassword"`
	STUNMinimumKeepAlivePeriod                   *uint                     `json:"STUNMinimumKeepAlivePeriod,omitempty" pn:"STUNMinimumKeepAlivePeriod"`
	STUNMaximumKeepAlivePeriod                   *uint                     `json:"STUNMaximumKeepAlivePeriod,omitempty" pn:"STUNMaximumKeepAlivePeriod"`
	NATDetected                                  *int                      `json:"NATDetected,omitempty" pn:"NATDetected"`
	NATLocalPort                                 *string                   `json:"NATLocalPort,omitempty" pn:"NATLocalPort"`
	ManageableDevice                             map[int]*ManageableDevice `json:"ManageableDevice,omitempty" pn:"ManageableDevice"`
	Parameterkey                                 *string                   `json:"Parameterkey,omitempty" pn:"Parameterkey"`
	TcpMaxSeg                                    *string                   `json:"TcpMaxSeg,omitempty" pn:"TcpMaxSeg"`
	InitServerURL                                *string                   `json:"InitServerURL,omitempty" pn:"InitServerURL"`
	X_VENDOR_IHEMS_URL                           *string                   `json:"X_VENDOR_IHEMS_URL,omitempty" pn:"X_VENDOR_IHEMS_URL"`
}

type UDPEchoConfig struct {
	Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
	UDPPort                 *uint   `json:"UDPPort,omitempty" pn:"UDPPort"`
	Interface               *string `json:"Interface,omitempty" pn:"Interface"`
	SourceIPAddress         *string `json:"SourceIPAddress,omitempty" pn:"SourceIPAddress"`
	BytesReceived           *uint   `json:"BytesReceived,omitempty" pn:"BytesReceived"`
	BytesResponded          *uint   `json:"BytesResponded,omitempty" pn:"BytesResponded"`
	PacketsReceived         *uint   `json:"PacketsReceived,omitempty" pn:"PacketsReceived"`
	PacketsResponded        *uint   `json:"PacketsResponded,omitempty" pn:"PacketsResponded"`
	EchoPlusEnabled         *int    `json:"EchoPlusEnabled,omitempty" pn:"EchoPlusEnabled"`
	EchoPlusSupported       *int    `json:"EchoPlusSupported,omitempty" pn:"EchoPlusSupported"`
	TimeFirstPacketReceived *string `json:"TimeFirstPacketReceived,omitempty" pn:"TimeFirstPacketReceived"`
	TimeLastPacketReceived  *string `json:"TimeLastPacketReceived,omitempty" pn:"TimeLastPacketReceived"`
}

type Time struct {
	CurrentLocalTime          *string `json:"CurrentLocalTime,omitempty" pn:"CurrentLocalTime"`
	Enable                    *string `json:"Enable,omitempty" pn:"Enable"`
	LocalTimeZone             *string `json:"LocalTimeZone,omitempty" pn:"LocalTimeZone"`
	NTPServer1                *string `json:"NTPServer1,omitempty" pn:"NTPServer1"`
	NTPServer2                *string `json:"NTPServer2,omitempty" pn:"NTPServer2"`
	NTPServer3                *string `json:"NTPServer3,omitempty" pn:"NTPServer3"`
	NTPServer4                *string `json:"NTPServer4,omitempty" pn:"NTPServer4"`
	NTPServer5                *string `json:"NTPServer5,omitempty" pn:"NTPServer5"`
	X_VENDOR_OVERSEASTIMEZOME *string `json:"X_VENDOR_OVERSEASTIMEZOME,omitempty" pn:"X_VENDOR_OVERSEASTIMEZOME"`
	X_VENDOR_SYNC_MODE        *string `json:"X_VENDOR_SYNC_MODE,omitempty" pn:"X_VENDOR_SYNC_MODE"`
	SystemTime                *string `json:"SystemTime,omitempty" pn:"SystemTime"`
	PTPServer                 *string `json:"PTPServer,omitempty" pn:"PTPServer"`
}
type IPsec struct {
	AHSupported                           *string `json:"AHSupported,omitempty" pn:"AHSupported"`
	ESPSupportedEncryptionAlgorithms      *string `json:"ESPSupportedEncryptionAlgorithms,omitempty" pn:"ESPSupportedEncryptionAlgorithms"`
	Enable                                *string `json:"Enable,omitempty" pn:"Enable"`
	IKEv2SupportedEncryptionAlgorithms    *string `json:"IKEv2SupportedEncryptionAlgorithms,omitempty" pn:"IKEv2SupportedEncryptionAlgorithms"`
	IKEv2SupportedPseudoRandomFunctions   *string `json:"IKEv2SupportedPseudoRandomFunctions,omitempty" pn:"IKEv2SupportedPseudoRandomFunctions"`
	MyKeyMode                             *string `json:"MyKeyMode,omitempty" pn:"MyKeyMode"`
	Status                                *string `json:"Status,omitempty" pn:"Status"`
	SupportedDiffieHellmanGroupTransforms *string `json:"SupportedDiffieHellmanGroupTransforms,omitempty" pn:"SupportedDiffieHellmanGroupTransforms"`
	SupportedIntegrityAlgorithms          *string `json:"SupportedIntegrityAlgorithms,omitempty" pn:"SupportedIntegrityAlgorithms"`
	Profile                               map[int]*struct {
		IKEv2AuthenticationMethod                *string `json:"IKEv2AuthenticationMethod,omitempty" pn:"IKEv2AuthenticationMethod,omitempty"`
		AHAllowedIntegrityAlgorithms             *string `json:"AHAllowedIntegrityAlgorithms,omitempty" pn:"AHAllowedIntegrityAlgorithms,omitempty"`
		ESPAllowedIntegrityAlgorithms            *string `json:"ESPAllowedIntegrityAlgorithms,omitempty" pn:"ESPAllowedIntegrityAlgorithms,omitempty"`
		ESPAllowedEncryptionAlgorithms           *string `json:"ESPAllowedEncryptionAlgorithms,omitempty" pn:"ESPAllowedEncryptionAlgorithms,omitempty"`
		IKEv2AllowedIntegrityAlgorithms          *string `json:"IKEv2AllowedIntegrityAlgorithms,omitempty" pn:"IKEv2AllowedIntegrityAlgorithms,omitempty"`
		IKEv2AllowedEncryptionAlgorithms         *string `json:"IKEv2AllowedEncryptionAlgorithms,omitempty" pn:"IKEv2AllowedEncryptionAlgorithms,omitempty"`
		IKEv2AllowedPseudoRandomFunctions        *string `json:"IKEv2AllowedPseudoRandomFunctions,omitempty" pn:"IKEv2AllowedPseudoRandomFunctions,omitempty"`
		IKEv2AllowedDiffieHellmanGroupTransforms *string `json:"IKEv2AllowedDiffieHellmanGroupTransforms,omitempty" pn:"IKEv2AllowedDiffieHellmanGroupTransforms,omitempty"`
	} `json:"Profile,omitempty" pn:"Profile"`
	Gateway *struct {
		SecGWServer1 *string `json:"SecGWServer1,omitempty" pn:"SecGWServer1"`
		SecGWServer2 *string `json:"SecGWServer2,omitempty" pn:"SecGWServer2"`
		SecGWServer3 *string `json:"SecGWServer3,omitempty" pn:"SecGWServer3"`
	} `json:"Gateway,omitempty" pn:"Gateway"`
}
type LogMgmt struct {
	LogLevel               *uint   `json:"LogLevel,omitempty" pn:"LogLevel"`
	Password               *string `json:"Password,omitempty" pn:"Password"`
	PeriodicUploadEnable   *string `json:"PeriodicUploadEnable,omitempty" pn:"PeriodicUploadEnable"`
	PeriodicUploadInterval *uint   `json:"PeriodicUploadInterval,omitempty" pn:"PeriodicUploadInterval"`
	URL                    *string `json:"URL,omitempty" pn:"URL"`
	Username               *string `json:"Username,omitempty" pn:"Username"`
}

type IPInterface struct {
	IPInterfaceIPAddress  *string `json:"IPInterfaceIPAddress,omitempty" pn:"IPInterfaceIPAddress"`
	IPInterfaceSubnetMask *string `json:"IPInterfaceSubnetMask,omitempty" pn:"IPInterfaceSubnetMask"`
}

type LANEthernetInterfaceConfig struct {
	MACAddress *string `json:"MACAddress,omitempty" pn:"MACAddress"`
}

type LANHostConfigManagement struct {
	IPInterface map[int]*IPInterface `json:"IPInterface,omitempty" pn:"IPInterface"`
}

type LANDevice struct {
	LANHostConfigManagement    *LANHostConfigManagement            `json:"LANHostConfigManagement,omitempty" pn:"LANHostConfigManagement"`
	LANEthernetInterfaceConfig map[int]*LANEthernetInterfaceConfig `json:"LANEthernetInterfaceConfig,omitempty" pn:"LANEthernetInterfaceConfig"`
}

type WANDevice struct {
	WANConnectionDevice          map[int]*WANConnectionDevice `json:"WANConnectionDevice,omitempty" pn:"WANConnectionDevice"`
	WANConnectionNumberOfEntries *uint                        `json:"fWANConnectionNumberOfEntries,omitempty" pn:"WANConnectionNumberOfEntries"`
	WANEthernetInterfaceConfig   *struct {
		DuplexMode     *string `json:"DuplexMode,omitempty" pn:"DuplexMode"`
		Enable         *string `json:"Enable,omitempty" pn:"Enable"`
		MACAddress     *string `json:"MACAddress,omitempty" pn:"MACAddress"`
		MaxBitRate     *string `json:"MaxBitRate,omitempty" pn:"MaxBitRate"`
		SignTransMedia *string `json:"SignTransMedia,omitempty" pn:"SignTransMedia"`
		Status         *string `json:"Status,omitempty" pn:"Status"`
	} `json:"WANEthernetInterfaceConfig,omitempty" pn:"WANEthernetInterfaceConfig"`
	WANCommonInterfaceConfig *struct {
		WANAccessType *string `json:"WANAccessType,omitempty" pn:"WANAccessType"`
	} `json:"WANCommonInterfaceConfig,omitempty" pn:"WANCommonInterfaceConfig"`
}

type WANConnectionDevice struct {
	WANIPConnection                 map[int]*WANIPConnection `json:"WANIPConnection,omitempty" pn:"WANIPConnection"`
	WANIPConnectionNumberOfEntries  *uint                    `json:"WANIPConnectionNumberOfEntries,omitempty" pn:"WANIPConnectionNumberOfEntries"`
	WANPPPConnectionNumberOfEntries *uint                    `json:"WANPPPConnectionNumberOfEntries,omitempty" pn:"WANPPPConnectionNumberOfEntries"`
}
type WANIPConnection struct {
	AddressingType    *string `json:"AddressingType,omitempty" pn:"AddressingType"`
	DNSServers        *string `json:"DNSServers,omitempty" pn:"DNSServers"`
	DefaultGateway    *string `json:"DefaultGateway,omitempty" pn:"DefaultGateway"`
	Enable            *string `json:"Enable,omitempty" pn:"Enable"`
	ExternalIPAddress *string `json:"ExternalIPAddress,omitempty" pn:"ExternalIPAddress"`
	MACAddress        *string `json:"MACAddress,omitempty" pn:"MACAddress"`
	PortType          *string `json:"PortType,omitempty" pn:"PortType"`
	Stats             *struct {
		EthernetBytesReceived   *string `json:"EthernetBytesReceived,omitempty" pn:"EthernetBytesReceived"`
		EthernetBytesSent       *string `json:"EthernetBytesSent,omitempty" pn:"EthernetBytesSent"`
		EthernetPacketsReceived *string `json:"EthernetPacketsReceived,omitempty" pn:"EthernetPacketsReceived"`
		EthernetPacketsSent     *string `json:"EthernetPacketsSent,omitempty" pn:"EthernetPacketsSent"`
	} `json:"Stats,omitempty" pn:"Stats"`
	SubnetMask *string `json:"SubnetMask,omitempty" pn:"SubnetMask"`
}
type SoftwareCtrl struct {
	ActivateEnable       *string `json:"ActivateEnable,omitempty" pn:"ActivateEnable"`
	ActivateTime         *string `json:"ActivateTime,omitempty" pn:"ActivateTime"`
	AutoActivateEnable   *string `json:"AutoActivateEnable,omitempty" pn:"AutoActivateEnable"`
	SystemBackupVersion  *string `json:"SystemBackupVersion,omitempty" pn:"SystemBackupVersion"`
	SystemCurrentVersion *string `json:"SystemCurrentVersion,omitempty" pn:"SystemCurrentVersion"`
}

type SyncMode struct {
	Acr1588 struct {
		DURATION_FIELD_Acr1588               *string `json:"DURATION_FIELD_Acr1588,omitempty" pn:"DURATION_FIELD_Acr1588"`
		FREQ_SYNC_ENABLE_Acr1588             *string `json:"FREQ_SYNC_ENABLE_Acr1588,omitempty" pn:"FREQ_SYNC_ENABLE_Acr1588"`
		LOG_INTER_MESSAGE_PERION_Acr1588     *string `json:"LOG_INTER_MESSAGE_PERION_Acr1588,omitempty" pn:"LOG_INTER_MESSAGE_PERION_Acr1588"`
		NETWORK_INTERFACE_Acr1588            *string `json:"NETWORK_INTERFACE_Acr1588,omitempty" pn:"NETWORK_INTERFACE_Acr1588"`
		REQUST_SIGNALING_INTERVAL_T1_Acr1588 *string `json:"REQUST_SIGNALING_INTERVAL_T1_Acr1588,omitempty" pn:"REQUST_SIGNALING_INTERVAL_T1_Acr1588"`
		REQUST_SIGNALING_INTERVAL_T2_Acr1588 *string `json:"REQUST_SIGNALING_INTERVAL_T2_Acr1588,omitempty" pn:"REQUST_SIGNALING_INTERVAL_T2_Acr1588"`
		REQUST_SIGNALING_INTERVAL_T3_Acr1588 *string `json:"REQUST_SIGNALING_INTERVAL_T3_Acr1588,omitempty" pn:"REQUST_SIGNALING_INTERVAL_T3_Acr1588"`
		TIME_SYNC_ENABLE_Acr1588             *string `json:"TIME_SYNC_ENABLE_Acr1588,omitempty" pn:"TIME_SYNC_ENABLE_Acr1588"`
		UNICASE_ADDRRESS_Acr1588             *string `json:"UNICASE_ADDRRESS_Acr1588,omitempty" pn:"UNICASE_ADDRRESS_Acr1588"`
		FREQACC_Acr1588                      *string `json:"FREQACC_Acr1588,omitempty" pn:"freqACC_Acr1588"`
		SYNCStatus_Acr1588                   *string `json:"SYNCStatus_Acr1588,omitempty" pn:"syncStatus_Acr1588"`
	} `json:"Acr1588,omitempty"`
	I1588v2 struct {
		AnnoDuration1588v2              *string `json:"AnnoDuration1588v2,omitempty" pn:"AnnoDuration1588v2"`
		AnnoInter1588v2                 *string `json:"AnnoInter1588v2,omitempty" pn:"AnnoInter1588v2"`
		AnnoRequest1588v2               *string `json:"AnnoRequest1588v2,omitempty" pn:"AnnoRequest1588v2"`
		ClockMode1588v2                 *string `json:"ClockMode1588v2,omitempty" pn:"ClockMode1588v2"`
		DelayAsymmetryNanoSeconds1588v2 *string `json:"DelayAsymmetryNanoSeconds1588v2,omitempty" pn:"DelayAsymmetryNanoSeconds1588v2"`
		DelayAsymmetrySeconds1588v2     *string `json:"DelayAsymmetrySeconds1588v2,omitempty" pn:"DelayAsymmetrySeconds1588v2"`
		DelayDuration1588v2             *string `json:"DelayDuration1588v2,omitempty" pn:"DelayDuration1588v2"`
		DelayInter1588v2                *string `json:"DelayInter1588v2,omitempty" pn:"DelayInter1588v2"`
		DelayMechanism1588v2            *string `json:"DelayMechanism1588v2,omitempty" pn:"DelayMechanism1588v2"`
		DelayRequest1588v2              *string `json:"DelayRequest1588v2,omitempty" pn:"DelayRequest1588v2"`
		DomainNumber1588v2              *string `json:"DomainNumber1588v2,omitempty" pn:"DomainNumber1588v2"`
		Interface1588v2                 *string `json:"Interface1588v2,omitempty" pn:"Interface1588v2"`
		MeanPathDelay                   *string `json:"MeanPathDelay,omitempty"`
		NatEnable1588v2                 *string `json:"NatEnable1588v2,omitempty" pn:"NatEnable1588v2"`
		OffsetFromMaster                *string `json:"OffsetFromMaster,omitempty"`
		PackageFormat1588v2             *string `json:"PackageFormat1588v2,omitempty" pn:"PackageFormat1588v2"`
		PortMode1588v2                  *string `json:"PortMode1588v2,omitempty" pn:"PortMode1588v2"`
		SyncDuration1588v2              *string `json:"SyncDuration1588v2,omitempty" pn:"SyncDuration1588v2"`
		SyncInter1588v2                 *string `json:"SyncInter1588v2,omitempty" pn:"SyncInter1588v2"`
		SyncRequest1588v2               *string `json:"SyncRequest1588v2,omitempty" pn:"SyncRequest1588v2"`
		UnicastAddr1588v2               *string `json:"UnicastAddr1588v2,omitempty" pn:"UnicastAddr1588v2"`
	} `json:"I1588v2,omitempty"`
}

type Ethernet struct {
	ReserveVlan  *string `json:"ReserveVlan,omitempty" pn:"ReserveVlan"`
	ConfigStatus *string `json:"ConfigStatus,omitempty" pn:"ConfigStatus"`
	IpRoute      map[int]*struct {
		IpVer            *string `json:"IpVer,omitempty" pn:"IpVer"`
		DstIpNetwork     *string `json:"DstIpNetwork,omitempty" pn:"DstIpNetwork"`
		PrefixLength     *string `json:"PrefixLength,omitempty" pn:"PrefixLength"`
		InterfaceName    *string `json:"InterfaceName,omitempty" pn:"InterfaceName"`
		GatewayIpAddress *string `json:"GatewayIpAddress,omitempty" pn:"GatewayIpAddress"`
	} `json:"IpRoute,omitempty" pn:"IpRoute"`
	Interface map[int]*struct {
		UserLabel      *string `json:"UserLabel,omitempty" pn:"UserLabel"`
		Name           *string `json:"Name,omitempty" pn:"Name"`
		Enable         *string `json:"Enable,omitempty" pn:"Enable"`
		Status         *string `json:"Status,omitempty" pn:"Status"`
		DuplexMode     *string `json:"DuplexMode,omitempty" pn:"DuplexMode"`
		MACAddress     *string `json:"MACAddress,omitempty" pn:"MACAddress"`
		MaxBitRate     *string `json:"MaxBitRate,omitempty" pn:"MaxBitRate"`
		PortLocation   *string `json:"PortLocation,omitempty" pn:"PortLocation"`
		SignTransMedia *string `json:"SignTransMedia,omitempty" pn:"SignTransMedia"`
		IPv4Address    map[int]*struct {
			PortType       *string `json:"PortType,omitempty" pn:"PortType"`
			IPAddress      *string `json:"IPAddress,omitempty" pn:"IPAddress"`
			SubnetMask     *string `json:"SubnetMask,omitempty" pn:"SubnetMask"`
			AddressingType *string `json:"AddressingType,omitempty" pn:"AddressingType"`
			DefaultGateway *string `json:"DefaultGateway,omitempty" pn:"DefaultGateway"`
		} `json:"IPv4Address,omitempty" pn:"IPv4Address"`
	} `json:"Interface,omitempty" pn:"Interface"`
}

type EXTRA_IP_LIST struct {
	X_VENDOR_EXTRA_IP_CONFIG          *string `json:"X_VENDOR_EXTRA_IP_CONFIG,omitempty" pn:"X_VENDOR_EXTRA_IP_CONFIG"`
	X_VENDOR_EXTRA_IP_DEFAULT_GW      *string `json:"X_VENDOR_EXTRA_IP_DEFAULT_GW,omitempty" pn:"X_VENDOR_EXTRA_IP_DEFAULT_GW"`
	X_VENDOR_EXTRA_IP_IPADDR_MASK     *string `json:"X_VENDOR_EXTRA_IP_IPADDR_MASK,omitempty" pn:"X_VENDOR_EXTRA_IP_IPADDR_MASK"`
	X_VENDOR_EXTRA_IP_IPADDR_METHOD   *string `json:"X_VENDOR_EXTRA_IP_IPADDR_METHOD,omitempty" pn:"X_VENDOR_EXTRA_IP_IPADDR_METHOD"`
	X_VENDOR_EXTRA_IP_IPADDR_STATIC   *string `json:"X_VENDOR_EXTRA_IP_IPADDR_STATIC,omitempty" pn:"X_VENDOR_EXTRA_IP_IPADDR_STATIC"`
	X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP *string `json:"X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP,omitempty" pn:"X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP"`
	X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG  *string `json:"X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG,omitempty" pn:"X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG"`
	X_VENDOR_EXTRA_IP_VLAN1_VID       *string `json:"X_VENDOR_EXTRA_IP_VLAN1_VID,omitempty" pn:"X_VENDOR_EXTRA_IP_VLAN1_VID"`
}
type WAN struct {
	EXTRA_IP_LIST map[int]*EXTRA_IP_LIST `json:"EXTRA_IP_LIST,omitempty" pn:"EXTRA_IP_LIST"`

	X_VENDOR_DEFAULT_GW             *string `json:"X_VENDOR_DEFAULT_GW,omitempty" pn:"X_VENDOR_DEFAULT_GW"`
	X_VENDOR_IPADDR_MASK            *string `json:"X_VENDOR_IPADDR_MASK,omitempty" pn:"X_VENDOR_IPADDR_MASK"`
	X_VENDOR_IPADDR_METHOD          *string `json:"X_VENDOR_IPADDR_METHOD,omitempty" pn:"X_VENDOR_IPADDR_METHOD"`
	X_VENDOR_IPADDR_STATIC          *string `json:"X_VENDOR_IPADDR_STATIC,omitempty" pn:"X_VENDOR_IPADDR_STATIC"`
	X_VENDOR_PPPOE_PWD              *string `json:"X_VENDOR_PPPOE_PWD,omitempty" pn:"X_VENDOR_PPPOE_PWD"`
	X_VENDOR_PPPOE_USER             *string `json:"X_VENDOR_PPPOE_USER,omitempty" pn:"X_VENDOR_PPPOE_USER"`
	X_VENDOR_STATIC_ROUTE_2         *string `json:"X_VENDOR_STATIC_ROUTE_2,omitempty" pn:"X_VENDOR_STATIC_ROUTE_2"`
	X_VENDOR_STATIC_ROUTE_NEETHOP_2 *string `json:"X_VENDOR_STATIC_ROUTE_NEETHOP_2,omitempty" pn:"X_VENDOR_STATIC_ROUTE_NEETHOP_2"`
	X_VENDOR_STATIC_ROUTE_NETMASK_2 *string `json:"X_VENDOR_STATIC_ROUTE_NETMASK_2,omitempty" pn:"X_VENDOR_STATIC_ROUTE_NETMASK_2"`
	X_VENDOR_VLAN1_DEFAULT_GW       *string `json:"X_VENDOR_VLAN1_DEFAULT_GW,omitempty" pn:"X_VENDOR_VLAN1_DEFAULT_GW"`
	X_VENDOR_VLAN1_IPADDR_MASK      *string `json:"X_VENDOR_VLAN1_IPADDR_MASK,omitempty" pn:"X_VENDOR_VLAN1_IPADDR_MASK"`
	X_VENDOR_VLAN1_IPADDR_METHOD    *string `json:"X_VENDOR_VLAN1_IPADDR_METHOD,omitempty" pn:"X_VENDOR_VLAN1_IPADDR_METHOD"`
	X_VENDOR_VLAN1_IPADDR_STATIC    *string `json:"X_VENDOR_VLAN1_IPADDR_STATIC,omitempty" pn:"X_VENDOR_VLAN1_IPADDR_STATIC"`
	X_VENDOR_VLAN1_TAG_FLAG         *string `json:"X_VENDOR_VLAN1_TAG_FLAG,omitempty" pn:"X_VENDOR_VLAN1_TAG_FLAG"`
	X_VENDOR_VLAN1_VID              *string `json:"X_VENDOR_VLAN1_VID,omitempty" pn:"X_VENDOR_VLAN1_VID"`
	X_VENDOR_VLAN2_DEFAULT_GW       *string `json:"X_VENDOR_VLAN2_DEFAULT_GW,omitempty" pn:"X_VENDOR_VLAN2_DEFAULT_GW"`
	X_VENDOR_VLAN2_IPADDR_MASK      *string `json:"X_VENDOR_VLAN2_IPADDR_MASK,omitempty" pn:"X_VENDOR_VLAN2_IPADDR_MASK"`
	X_VENDOR_VLAN2_IPADDR_METHOD    *string `json:"X_VENDOR_VLAN2_IPADDR_METHOD,omitempty" pn:"X_VENDOR_VLAN2_IPADDR_METHOD"`
	X_VENDOR_VLAN2_IPADDR_STATIC    *string `json:"X_VENDOR_VLAN2_IPADDR_STATIC,omitempty" pn:"X_VENDOR_VLAN2_IPADDR_STATIC"`
	X_VENDOR_VLAN2_TAG_FLAG         *string `json:"X_VENDOR_VLAN2_TAG_FLAG,omitempty" pn:"X_VENDOR_VLAN2_TAG_FLAG"`
	X_VENDOR_VLAN2_VID              *string `json:"X_VENDOR_VLAN2_VID,omitempty" pn:"X_VENDOR_VLAN2_VID"`
	X_VENDOR_WAN_MODE_CONFIG        *string `json:"X_VENDOR_WAN_MODE_CONFIG,omitempty" pn:"X_VENDOR_WAN_MODE_CONFIG"`
	X_VENDOR_WAN_MTU_CONFIG         *string `json:"X_VENDOR_WAN_MTU_CONFIG,omitempty" pn:"X_VENDOR_WAN_MTU_CONFIG"`
	X_VENDOR_WAN_VLAN_TYPE          *string `json:"X_VENDOR_WAN_VLAN_TYPE,omitempty" pn:"X_VENDOR_WAN_VLAN_TYPE"`
}
