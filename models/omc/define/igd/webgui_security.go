package igd

type Security struct {
	Firewall    *Firewall    `json:"Firewall,omitempty" pn:"Firewall"`
	MACFilter   *MACFilter   `json:"MACFilter,omitempty" pn:"MACFilter"`
	IPFilter    *IPFilter    `json:"IPFilter,omitempty" pn:"IPFilter"`
	URLFilter   *URLFilter   `json:"URLFilter,omitempty" pn:"URLFilter"`
	PortForward *PortForward `json:"PortForward,omitempty" pn:"PortForward"`
	PortTrigger *PortTrigger `json:"PortTrigger,omitempty" pn:"PortTrigger"`

	AccessRestriction *AccessRestriction `json:"AccessRestriction,omitempty" pn:"AccessRestriction"`
	UPNP              *UPNP              `json:"UPNP,omitempty" pn:"UPNP"`
	DoS               *DoS               `json:"DoS,omitempty" pn:"DoS"`
}

type Firewall struct {
	Enable *string `json:"Enable,omitempty" pn:"Enable"`
}

type MACFilter struct {
	Enable        *string `json:"Enable,omitempty" pn:"Enable"`
	AccessNetwork *string `json:"AccessNetwork,omitempty" pn:"AccessNetwork"`
	List          map[int]*struct {
		MACAddress *string `json:"MACAddress,omitempty" pn:"MACAddress"`
	} `json:"List,omitempty" pn:"List"`
}

type IPFilter struct {
	Enable        *string `json:"Enable,omitempty" pn:"Enable"`
	AccessNetwork *string `json:"AccessNetwork,omitempty" pn:"AccessNetwork"`
	List          map[int]*struct {
		Protocol   *string `json:"Protocol,omitempty" pn:"Protocol"`
		SourceIP   *string `json:"SourceIP,omitempty" pn:"SourceIP"`
		SourcePort *string `json:"SourcePort,omitempty" pn:"SourcePort"`
		DestIP     *string `json:"DestIP,omitempty" pn:"DestIP"`
		DestPort   *string `json:"DestPort,omitempty" pn:"DestPort"`
		Status     *string `json:"Status,omitempty" pn:"Status"`
	} `json:"List,omitempty" pn:"List"`
}

type URLFilter struct {
	Enable *string `json:"Enable,omitempty" pn:"Enable"`
	List   map[int]*struct {
		URL *string `json:"List,omitempty" pn:"URL"`
	} `json:"List,omitempty" pn:"List"`
}

type PortForward struct {
	List map[int]*struct {
		Protocol   *string `json:"Protocol,omitempty" pn:"Protocol"`
		RemotePort *string `json:"RemotePort,omitempty" pn:"RemotePort"`
		LocalHost  *string `json:"LocalHost,omitempty" pn:"LocalHost"`
		LocalPort  *string `json:"LocalPort,omitempty" pn:"LocalPort"`
	} `json:"List,omitempty" pn:"List"`
}

type PortTrigger struct {
	Enable *string `json:"Enable,omitempty" pn:"Enable"`
}

type AccessRestriction struct {
	List map[int]*struct {
		Enable    *string `json:"Enable,omitempty" pn:"Enable"`
		Name      *string `json:"Name,omitempty" pn:"Name"`
		Device    *string `json:"Device,omitempty" pn:"Device"`
		Weekdays  *string `json:"Weekdays,omitempty" pn:"Weekdays"`
		TimeStart *string `json:"TimeStart,omitempty" pn:"TimeStart"`
		TimeStop  *string `json:"TimeStop,omitempty" pn:"TimeStop"`
	} `json:"List,omitempty" pn:"List"`
}

type UPNP struct {
	Enable *string `json:"Enable,omitempty" pn:"Enable"`
	List   map[int]*struct {
		InternalPort *string `json:"Internalport,omitempty" pn:"Internalport"`
		ExternalPort *string `json:"ExternalPort,omitempty" pn:"ExternalPort"`
		IPAddress    *string `json:"IPAddress,omitempty" pn:"IPAddress"`
		Protocol     *string `json:"Protocol,omitempty" pn:"Protocol"`
		Description  *string `json:"Description,omitempty" pn:"Description"`
	} `json:"List,omitempty" pn:"List"`
}

type DoS struct {
	DoSEnable  *string `json:"DoSEnable,omitempty" pn:"DoSEnable"`
	SyncEnable *string `json:"SyncEnable,omitempty" pn:"SyncEnable"`
	PingEnable *string `json:"PingEnable,omitempty" pn:"PingEnable"`
	TCPEnable  *string `json:"TCPEnable,omitempty" pn:"TCPEnable"`
	UDPEnable  *string `json:"UDPEnable,omitempty" pn:"UDPEnable"`
}
