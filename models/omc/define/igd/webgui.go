package igd

type WebGui struct {
	Overview       *Overview       `json:"Overview,omitempty"  pn:"Overview"`
	ScheduleReboot *ScheduleReboot `json:"ScheduleReboot"  pn:"ScheduleReboot"`

	System           *System           `json:"System,omitempty"  pn:"System"`
	Network          *Network          `json:"Network,omitempty" pn:"Network"`
	WiFi             *WiFi             `json:"WiFi,omitempty" pn:"WiFi,omitempty"`
	ParentalControls *ParentalControls `json:"ParentalControl,omitempty" pn:"ParentalControls"`
	Security         *Security         `json:"Security,omitempty" pn:"Security"`
	VPN              *VPN              `json:"VPN,omitempty" pn:"VPN"`
	IPv6             *IPv6             `json:"IPv6,omitempty" pn:"IPv6"`
	VoIP             *VoIP             `json:"VoIP,omitempty" pn:"VoIP"`

	Statistics           *Statistics           `json:"Statistics,omitempty" pn:"Statistics"`
	ThroughputStatistics *ThroughputStatistics `json:"ThroughputStatistics,omitempty" pn:"ThroughputStatistics"`
}

type ParentalControls struct {
	Enable *string                       `json:"Enable,omitempty" pn:"Enable"`
	List   map[int]*ParentalControlsList `json:"List,omitempty" pn:"List"`
}
type ParentalControlsList struct {
	Name      *string `json:"Name,omitempty" pn:"Name"`
	Device    *string `json:"Device,omitempty" pn:"Device"`
	Weekdays  *string `json:"Weekdays,omitempty" pn:"Weekdays"`
	TimeStart *string `json:"TimeStart,omitempty" pn:"TimeStart"`
	TimeStop  *string `json:"TimeStop,omitempty" pn:"TimeStop"`
	Status    *string `json:"Status,omitempty" pn:"Status"`
}
type StaticRoute struct {
	List map[int]*StaticRouteList `json:"List,omitempty" pn:"List"`
}
type StaticRouteList struct {
	Index         uint    `json:"index,omitempty" pn:"index"`
	DestIPAddress *string `json:"DestIPAddress,omitempty" pn:"DestIPAddress"`
	SubnetMask    *string `json:"SubnetMask,omitempty" pn:"SubnetMask"`
	Interface     *string `json:"Interface,omitempty" pn:"Interface"`
	Gateway       *string `json:"Gateway,omitempty" pn:"Gateway"`
	Status        *string `json:"Status,omitempty" pn:"Status"`
}
type VPN struct {
	Enable         *string `json:"Enable,omitempty" pn:"Enable"`
	Server         *string `json:"Server,omitempty" pn:"Server"`
	Username       *string `json:"Username,omitempty" pn:"Username"`
	Password       *string `json:"Password,omitempty" pn:"Password"`
	Protocol       *string `json:"Protocol,omitempty" pn:"Protocol"`
	BearDevice     *string `json:"BearDevice,omitempty" pn:"BearDevice"`
	DefaultGateway *string `json:"DefaultGateway,omitempty" pn:"DefaultGateway"`
	IPsecEnable    *string `json:"IPsecEnable,omitempty" pn:"IPsecEnable"`
	IPsecPassword  *string `json:"IPsecPassword,omitempty" pn:"IPsecPassword"`

	DefaultRouteEnable *string    `json:"DefaultRouteEnable,omitempty" pn:"DefaultRouteEnable"`
	Status             *VPNStatus `json:"Status,omitempty" pn:"Status"`

	GRE    *GRE    `json:"GRE,omitempty" pn:"GRE"`
	L2TPv3 *L2TPv3 `json:"L2TPv3,omitempty" pn:"L2TPv3"`
}
type VPNStatus struct {
	List map[int]*VPNStatusList `json:"List,omitempty" pn:"List"`
}
type VPNStatusList struct {
	Index         *uint   `json:"Index,omitempty" pn:"index"`
	Username      *string `json:"Username,omitempty" pn:"Username"`
	LocalAddress  *string `json:"LocalAddress,omitempty" pn:"LocalAddress"`
	RemoteAddress *string `json:"RemoteAddress,omitempty" pn:"RemoteAddress"`
	OnlineTime    *string `json:"OnlineTime,omitempty" pn:"OnlineTime"`
}

type GRE struct {
	VPNLayer                     *string `json:"VPNLayer,omitempty" pn:"VPNLayer"`
	HostIPAddress                *string `json:"HostIPAddress,omitempty" pn:"HostIPAddress"`
	RemoteIPAddress              *string `json:"RemoteIPAddress,omitempty" pn:"RemoteIPAddress"`
	RemotePrivateIPAddress       *string `json:"RemotePrivateIPAddress,omitempty" pn:"RemotePrivateIPAddress"`
	RemotePrivateIPAddressPrefix *string `json:"RemotePrivateIPAddressPrefix,omitempty" pn:"RemotePrivateIPAddressPrefix"`
	GREDestinationAddress        *string `json:"GREDestinationAddress,omitempty" pn:"GREDestinationAddress"`
	BearDevice                   *string `json:"BearDevice,omitempty" pn:"BearDevice"`
}
type L2TPv3 struct {
	VLANIDEnable         *string `json:"VLANIDEnable,omitempty" pn:"VLANIDEnable"`
	VLANID               *string `json:"VLANID,omitempty" pn:"VLANID"`
	LocalCookie          *string `json:"LocalCookie,omitempty" pn:"LocalCookie"`
	LocalTunnelID        *string `json:"LocalTunnelID,omitempty" pn:"LocalTunnelID"`
	LocalTunnelIPAddress *string `json:"LocalTunnelIPAddress,omitempty" pn:"LocalTunnelIPAddress"`
	LocalSessionID       *string `json:"LocalSessionID,omitempty" pn:"LocalSessionID"`

	RemoteCookie          *string `json:"RemoteCookie,omitempty" pn:"RemoteCookie"`
	RemoteTunnelID        *string `json:"RemoteTunnelID,omitempty" pn:"RemoteTunnelID"`
	RemoteTunnelIPAddress *string `json:"RemoteTunnelIPAddress,omitempty" pn:"RemoteTunnelIPAddress"`
	RemoteIPAddress       *string `json:"RemoteIPAddress,omitempty" pn:"RemoteIPAddress"`
	RemoteSessionID       *string `json:"RemoteSessionID,omitempty" pn:"RemoteSessionID"`

	EncapsulationType  *string `json:"EncapsulationType,omitempty" pn:"EncapsulationType"`
	UDPSourcePort      *uint   `json:"UDPSourcePort,omitempty" pn:"UDPSourcePort"`
	UDPDestinationPort *uint   `json:"UDPDestinationPort,omitempty" pn:"UDPDestinationPort"`
}

type IPv6 struct {
	WAN    *IPv6WAN    `json:"WAN,omitempty" pn:"WAN"`
	LAN    *IPv6LAN    `json:"LAN,omitempty" pn:"LAN"`
	Status *IPv6Status `json:"Status,omitempty" pn:"Status"`
}

type IPv6WAN struct {
	Enable     *string `json:"Enable,omitempty" pn:"Enable"`
	Type       *string `json:"Type,omitempty" pn:"Type"`
	GlobalAddr *string `json:"GlobalAddress,omitempty" pn:"GlobalAddress"`
	DNSForm    *string `json:"DNSForm,omitempty" pn:"DNSForm"`
	PDEnable   *string `json:"PDEnable,omitempty" pn:"PDEnable"`
	IPv6Pre    *string `json:"IPv6Pre,omitempty" pn:"IPv6Pre"`
	Gateway    *string `json:"Gateway,omitempty" pn:"Gateway"`
	DNS1       *string `json:"DNS1,omitempty" pn:"DNS1"`
	DNS2       *string `json:"DNS2,omitempty" pn:"DNS2"`
}

type IPv6LAN struct {
	Type         *string `json:"Type,omitempty" pn:"Type"`
	LocalAddress *string `json:"LocalAddress,omitempty" pn:"LocalAddress"`
}

type IPv6Status struct {
	LANAddress *struct {
		GlobalAddress *string `json:"GlobalAddress,omitempty" pn:"GlobalAddress"`
		LocalAddress  *string `json:"LocalAddress,omitempty" pn:"LocalAddress"`
		Type          *string `json:"Type,omitempty" pn:"Type"`
	} `json:"LANAddress,omitempty" pn:"LANAddress"`
	Information *struct {
		Status        *string `json:"Status,omitempty" pn:"Status"`
		GlobalAddress *string `json:"GlobalAddress,omitempty" pn:"GlobalAddress"`
		Type          *string `json:"Type,omitempty" pn:"Type"`
	} `json:"Information,omitempty" pn:"Information"`
}

type VoIP struct {
	Info       *VoIPInfo       `json:"Info,omitempty" pn:"Info"`
	SIPServer  *VoIPSIPServer  `json:"SIPServer,omitempty" pn:"SIPServer"`
	SIPAccount *VoIPSIPAccount `json:"SIPAccount,omitempty" pn:"SIPAccount"`
	Advanced   *VoIPAdvanced   `json:"Advanced,omitempty" pn:"VoIPAdvanced"`
}

type VoIPInfo struct {
	SIPAccount         *string `json:"SIPAccount,omitempty" pn:"SIPAccount"`
	LineStatus         *string `json:"LineStatus,omitempty" pn:"LineStatus"`
	RegistrationStatus *string `json:"RegistrationStatus,omitempty" pn:"RegistrationStatus"`
}

type VoIPSIPServer struct {
	UserAgentPort      uint32 `json:"UserAgentPort,omitempty" pn:"UserAgentPort"`
	RegistrationServer *struct {
		SIPServerDomainName       *string `json:"SIPServerDomainName,omitempty" pn:"SIPServerDomainName"`
		ProxyServerAddress        *string `json:"ProxyServerAddress,omitempty" pn:"ProxyServerAddress"`
		ProxyServerPort           uint32  `json:"ProxyServerPort,omitempty" pn:"ProxyServerPort"`
		RegistrationServerAddress *string `json:"RegistrationServerAddress,omitempty" pn:"RegistrationServerAddress"`
		RegistrationServerPort    uint32  `json:"RegistrationServerPort,omitempty" pn:"RegistrationServerPort"`
		RegistrationExpiryTime    uint32  `json:"RegistrationExpiryTime,omitempty" pn:"RegistrationExpiryTime"`
	} `json:"RegistrationServer,omitempty" pn:"RegistrationServer"`
}

type VoIPSIPAccount struct {
	Enable                 *string `json:"Enable,omitempty" pn:"Enable"`
	Username               *string `json:"Username,omitempty" pn:"Username"`
	Password               *string `json:"Password,omitempty" pn:"Password"`
	PhoneNumber            *string `json:"PhoneNumber,omitempty" pn:"PhoneNumber"`
	DisplayName            *string `json:"DisplayName,omitempty" pn:"DisplayName"`
	CodecPriority          *string `json:"CodecPriority,omitempty" pn:"CodecPriority"`
	SupportedCodecPriority *string `json:"SupportedCodecPriority,omitempty" pn:"SupportedCodecPriority"`
}

type VoIPAdvanced struct {
	DigitMap                   *string `json:"DigitMap,omitempty" pn:"DigitMap"`
	ThreewayConversationEnable *string `json:"ThreewayConversationEnable,omitempty" pn:"ThreewayConversationEnable"`
}
