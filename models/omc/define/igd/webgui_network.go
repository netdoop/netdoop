package igd

type Network struct {
	WANSettings   *WANSettings   `json:"WANSettings,omitempty" pn:"WANSettings"`
	LTE           *LTE           `json:"LTE,omitempty" pn:"LTE"`
	ScanMode      *ScanMode      `json:"ScanMode,omitempty" pn:"ScanMode"`
	APNManagement *APNManagement `json:"APNManagement,omitempty" pn:"APNManagement"`
	PINManagement *PINManagement `json:"PINManagement,omitempty" pn:"PINManagement"`
	SIMLock       *SIMLock       `json:"SIMLock,omitempty" pn:"SIMLock"`

	LANSettings *LANSettings `json:"LANSettings,omitempty" pn:"LANSettings"`
	DMZ         *DMZ         `json:"DMZ,omitempty" pn:"DMZ"`
	StaticRoute *StaticRoute `json:"StaticRoute,omitempty" pn:"StaticRoute"`
	NR_LTE      *NRLTE       `json:"NR_LTE,omitempty" pn:"NR-LTE"`
}

type WANSettings struct {
	NetworkMode *string `json:"NetworkMode,omitempty" pn:"NetworkMode"`
}

type LTE struct {
	ConnectStatus *string `json:"ConnectStatus,omitempty" pn:"ConnectStatus"`
	ConnectMethod *string `json:"ConnectMethod,omitempty" pn:"ConnectMethod"`
	Status        *struct {
		PCC *struct {
			DLMCS       *string `json:"DLMCS,omitempty" pn:"DLMCS"`
			ULMCS       *string `json:"ULMCS,omitempty" pn:"ULMCS"`
			DLFrequency *string `json:"DLFrequency,omitempty" pn:"DLFrequency"`
			ULFrequency *string `json:"ULFrequency,omitempty" pn:"ULFrequency"`
			Bandwidth   *string `json:"Bandwidth,omitempty" pn:"Bandwidth"`
			RSSI        *string `json:"RSSI,omitempty" pn:"RSSI"`
			RSRP0       *string `json:"RSRP0,omitempty" pn:"RSRP0"`
			RSRP1       *string `json:"RSRP1,omitempty" pn:"RSRP1"`
			RSRQ        *string `json:"RSRQ,omitempty" pn:"RSRQ"`
			SINR        *string `json:"SINR,omitempty" pn:"SINR"`
			TXPower     *string `json:"TXPower,omitempty" pn:"TXPower"`
			PCI         *string `json:"PCI,omitempty" pn:"PCI"`
			CINR0       *string `json:"CINR0,omitempty" pn:"CINR0"`
			CINR1       *string `json:"CINR1,omitempty" pn:"CINR1"`
			CellID      *string `json:"CellID,omitempty" pn:"CellID"`
			MCC         *string `json:"MCC,omitempty" pn:"MCC"`
			MNC         *string `json:"MNC,omitempty" pn:"MNC"`
		} `json:"PCC,omitempty" pn:"PCC"`
		SCC map[int]*struct {
			DLMCS       *string `json:"DLMCS,omitempty" pn:"DLMCS"`
			ULMCS       *string `json:"ULMCS,omitempty" pn:"ULMCS"`
			DLFrequency *string `json:"DLFrequency,omitempty" pn:"DLFrequency"`
			ULFrequency *string `json:"ULFrequency,omitempty" pn:"ULFrequency"`
			Bandwidth   *string `json:"Bandwidth,omitempty" pn:"Bandwidth"`
			RSSI        *string `json:"RSSI,omitempty" pn:"RSSI"`
			RSRP0       *string `json:"RSRP0,omitempty" pn:"RSRP0"`
			RSRP1       *string `json:"RSRP1,omitempty" pn:"RSRP1"`
			RSRQ        *string `json:"RSRQ,omitempty" pn:"RSRQ"`
			SINR        *string `json:"SINR,omitempty" pn:"SINR"`
			TXPower     *string `json:"TXPower,omitempty" pn:"TXPower"`
			PCI         *string `json:"PCI,omitempty" pn:"PCI"`
			CINR0       *string `json:"CINR0,omitempty" pn:"CINR0"`
			CINR1       *string `json:"CINR1,omitempty" pn:"CINR1"`
			CellID      *string `json:"CellID,omitempty" pn:"CellID"`
			MCC         *string `json:"MCC,omitempty" pn:"MCC"`
			MNC         *string `json:"MNC,omitempty" pn:"MNC"`
		} `json:"SCC,omitempty" pn:",SCC"`
	} `json:"Status,omitempty" pn:"Status"`
}
type NRLTE struct {
	RoamingSwitch *string `json:"RoamingSwitch,omitempty" pn:"RoamingSwitch"`
	ConnectMethod *string `json:"ConnectMethod,omitempty" pn:"ConnectMethod"`
	ConnectStatus *string `json:"ConnectStatus,omitempty" pn:"ConnectStatus"`
	MaxTxPower    *string `json:"MaxTxPower,omitempty" pn:"MaxTxPower"`

	Status *struct {
		LTE *struct {
			Band      *string `json:"Band,omitempty" pn:"Band"`
			TXPower   *string `json:"TXPower,omitempty" pn:"TXPower"`
			CQI       *string `json:"CQI,omitempty" pn:"CQI"`
			Bandwidth *string `json:"Bandwidth,omitempty" pn:"Bandwidth"`
			EARFCN    *string `json:"EARFCN,omitempty" pn:"EARFCN"`
			Frequency *string `json:"Frequency,omitempty" pn:"Frequency"`
			Rank      *string `json:"Rank,omitempty" pn:"Rank"`
			SINR      *string `json:"SINR,omitempty" pn:"SINR"`
			TM        *string `json:"TM,omitempty" pn:"TM"`
			RSRP      *string `json:"RSRP,omitempty" pn:"RSRP"`
			RSRP0     *string `json:"RSRP0,omitempty" pn:"RSRP0"`
			RSRP1     *string `json:"RSRP1,omitempty" pn:"RSRP1"`
			RSRQ      *string `json:"RSRQ,omitempty" pn:"RSRQ"`
			RSSI      *string `json:"RSSI,omitempty" pn:"RSSI"`
			ECGI      *string `json:"ECGI,omitempty" pn:"ECGI"`
			ENBID     *string `json:"ENBID,omitempty" pn:"eNBID"`
			CellID    *string `json:"CellID,omitempty" pn:"CellID"`
			TAC       *string `json:"TAC,omitempty" pn:"TAC"`
			PCI       *string `json:"PCI,omitempty" pn:"PCI"`
			MNC       *string `json:"MNC,omitempty" pn:"MNC"`
			MCC       *string `json:"MCC,omitempty" pn:"MCC"`
			PLMN      *string `json:"PLMN,omitempty" pn:"PLMN"`
			DLMCS     *string `json:"DLMCS,omitempty" pn:"DLMCS"`
			ULMCS     *string `json:"ULMCS,omitempty" pn:"ULMCS"`
		} `json:"LTE,omitempty" pn:"LTE"`
		NR *struct {
			SSB_BeamID   *string `json:"SSB_BeamID,omitempty" pn:"SSB_BeamID"`
			SSB_RSRP     *string `json:"SSB_RSRP,omitempty" pn:"SSB_RSRP"`
			SSB_RSRP0    *string `json:"SSB_RSRP0,omitempty" pn:"SSB_RSRP0"`
			SSB_RSRP1    *string `json:"SSB_RSRP1,omitempty" pn:"SSB_RSRP1"`
			SSB_RSRQ     *string `json:"SSB_RSRQ,omitempty" pn:"SSB_RSRQ"`
			SSB_SINR     *string `json:"SSB_SINR,omitempty" pn:"SSB_SINR"`
			SSB_RSSI     *string `json:"SSB_RSSI,omitempty" pn:"SSB_RSSI"`
			NR_ARFCN     *string `json:"NR_ARFCN,omitempty" pn:"NR-ARFCN"`
			NR_Bandwidth *string `json:"NR_Bandwidth,omitempty" pn:"NR_Bandwidth"`
			NR_Frequency *string `json:"NR_Frequency,omitempty" pn:"NR_Frequency"`
			NCGI         *string `json:"NCGI,omitempty" pn:"NCGI"`
			NR_PLMN      *string `json:"NR_PLMN,omitempty" pn:"NR_PLMN"`
			GNBID        *string `json:"GNBID,omitempty" pn:"gNBID"`
			NR_CellID    *string `json:"NR_CellID,omitempty" pn:"NR_CellID"`
			NR_TAC       *string `json:"NR_TAC,omitempty" pn:"NR_TAC"`
			NR_TXPower   *string `json:"NR_TXPower,omitempty" pn:"NR_TXPower"`
			NR_ULMCS     *string `json:"NR_ULMCS,omitempty" pn:"NR_ULMCS"`
			NR_DLMCS     *string `json:"NR_DLMCS,omitempty" pn:"NR_DLMCS"`
			NR_PCI       *string `json:"NR_PCI,omitempty" pn:"NR_PCI"`
			NR_CQI       *string `json:"NR_CQI,omitempty" pn:"NR_CQI"`
			NR_Rank      *string `json:"NR_Rank,omitempty" pn:"NR_Rank"`
			NR_Band      *string `json:"NR_Band,omitempty" pn:"NR_Band"`
			NR_MNC       *string `json:"NR_MNC,omitempty" pn:"NR_MNC"`
			NR_MCC       *string `json:"NR_MCC,omitempty" pn:"NR_MCC"`
			DLMCS        *string `json:"DLMCS,omitempty" pn:"DLMCS"`
			ULMCS        *string `json:"ULMCS,omitempty" pn:"ULMCS"`
		} `json:"NR,omitempty" pn:"NR"`
		PCC *struct {
			Band           *string `json:"Band,omitempty" pn:"Band"`
			TXPower        *string `json:"TXPower,omitempty" pn:"TXPower"`
			CQI            *string `json:"CQI,omitempty" pn:"CQI"`
			Bandwidth      *string `json:"Bandwidth,omitempty" pn:"Bandwidth"`
			Frequency      *string `json:"Frequency,omitempty" pn:"Frequency"`
			Rank           *string `json:"Rank,omitempty" pn:"Rank"`
			SINR           *string `json:"SINR,omitempty" pn:"SINR"`
			TM             *string `json:"TM,omitempty" pn:"TM"`
			TMMode         *string `json:"TMMode,omitempty" pn:"TMmode"`
			RSRP           *string `json:"RSRP,omitempty" pn:"RSRP"`
			RSRP0          *string `json:"RSRP0,omitempty" pn:"RSRP0"`
			RSRP1          *string `json:"RSRP1,omitempty" pn:"RSRP1"`
			RSRQ           *string `json:"RSRQ,omitempty" pn:"RSRQ"`
			RSSI           *string `json:"RSSI,omitempty" pn:"RSSI"`
			CINR0          *string `json:"CINR0,omitempty" pn:"CINR0"`
			CINR1          *string `json:"CINR1,omitempty" pn:"CINR1"`
			CellID         *string `json:"CellID,omitempty" pn:"CellID"`
			TAC            *string `json:"TAC,omitempty" pn:"TAC"`
			PCI            *string `json:"PCI,omitempty" pn:"PCI"`
			PLMN           *string `json:"PLMN,omitempty" pn:"PLMN"`
			MCC            *string `json:"MCC,omitempty" pn:"MCC"`
			MNC            *string `json:"MNC,omitempty" pn:"MNC"`
			DLMCS          *string `json:"DLMCS,omitempty" pn:"DLMCS"`
			ULMCS          *string `json:"ULMCS,omitempty" pn:"ULMCS"`
			ECGI_NCGI      *string `json:"ECGI_NCGI,omitempty" pn:"ECGI-NCGI"`
			SSB_BeamID     *string `json:"SSB_BeamID,omitempty" pn:"SSB_BeamID"`
			ENBID_GNBID    *string `json:"ENBID_GNBID,omitempty" pn:"eNBID-gNBID"`
			ENBID          *string `json:"ENBID,omitempty" pn:"eNBID"`
			EARFCN_NRARFCN *string `json:"EARFCN_NRARFCN,omitempty" pn:"EARFCN-NRARFCN"`
			DL_EARFCN      *string `json:"DL_EARFCN,omitempty" pn:"DLEarfcn"`
			ULFrequency    *string `json:"ULFrequency,omitempty" pn:"ULFrequency"`
			DLFrequency    *string `json:"DLFrequency,omitempty" pn:"DLFrequency"`
			DuplexMode     *string `json:"DuplexMode,omitempty" pn:"DuplexMode"`
		} `json:"PCC,omitempty" pn:"PCC"`
		NeighborCellList *struct {
			EARFCN_NBR *string `json:"EARFCN_NBR,omitempty" pn:"EARFCN_NBR"`
			PCI_NBR    *string `json:"PCI_NBR,omitempty" pn:"PCI_NBR"`
			RSRP_NBR   *string `json:"RSRP_NBR,omitempty" pn:"RSRP_NBR"`
		} `json:"NeighborCellList,omitempty" pn:"NeighborCellList"`
	} `json:"Status,omitempty" pn:"Status"`
}
type ScanMode struct {
	PreferMode              *string `json:"PreferMode,omitempty" pn:"PreferMode"`
	SuppBand                *string `json:"SuppBand,omitempty" pn:"SuppBand"`
	LockBand                *string `json:"LockBand,omitempty" pn:"LockBand"`
	LockBandList            *string `json:"LockBandList,omitempty" pn:"LockBandList"`
	LockEARFCN              *string `json:"LockEARFCN,omitempty" pn:"LockEARFCN"`
	LockEARFCNList          *string `json:"LockEARFCNList,omitempty" pn:"LockEARFCNList"`
	LockPCI                 *string `json:"LockPCI,omitempty" pn:"LockPCI"`
	LockPCIList             *string `json:"LockPCIList,omitempty" pn:"LockPCIList"`
	LockEARFCN_NRARFCN      *string `json:"LockEARFCN_NRARFCN,omitempty" pn:"LockEARFCN_NRARFCN"`
	LockEARFCN_NRARFCN_List *string `json:"LockEARFCN_NRARFCN_List,omitempty" pn:"LockEARFCN_NRARFCN_List"`
}

type APNManagement struct {
	DefaultGateway *int32           `json:"DefaultGateway,omitempty" pn:"DefaultGateway"`
	APNList        map[int]*APNList `json:"APNList,omitempty" pn:"APNList"`
}

type APNList struct {
	Index    *uint   `json:"Index,omitempty" pn:"index"`
	Enable   *string `json:"Enable,omitempty" pn:"Enable"`
	Label    *string `json:"Label,omitempty" pn:"Lable"`
	APNName  *string `json:"APNName,omitempty" pn:"APNName"`
	ApplyTo  *string `json:"ApplyTo,omitempty" pn:"ApplyTo"`
	AuthType *uint32 `json:"AuthType,omitempty" pn:"AuthType"`
	Username *string `json:"Username,omitempty" pn:"Username"`
	Password *string `json:"Password,omitempty" pn:"Password"`
	PDNType  *string `json:"PDNType,omitempty" pn:"PDNType"`
}

type PINManagement struct {
	VerifyEnable         *string `json:"VerifyEnable,omitempty" pn:"VerifyEnable"`
	PUK                  *string `json:"PUK,omitempty" pn:"PUK"`
	ChangePIN            *string `json:"ChangePIN,omitempty" pn:"ChangePIN"`
	PUKRemainingAttempts *string `json:"PUKRemainingAttempts,omitempty" pn:"PUKRemainingAttempts"`
	PINRemainingAttempts *string `json:"PINRemainingAttempts,omitempty" pn:"PINRemainingAttempts"`
	PINStatus            *string `json:"PINStatus,omitempty" pn:"PINStatus"`
}

type SIMLock struct {
	PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
}

type LANSettings struct {
	LANHost *LANHost `json:"LANHost,omitempty" pn:"LANHost"`
	DHCP    *DHCP    `json:"DHCP,omitempty"  pn:"DHCP"`
}

type LANHost struct {
	IPAddress  *string `json:"IPAddress,omitempty" pn:"IPAddress"`
	SubnetMask *string `json:"SubnetMask,omitempty" pn:"SubnetMask"`
}

type DHCP struct {
	ServerEnable *string `json:"ServerEnable,omitempty" pn:"ServerEnable"`
	StartIP      *string `json:"StartIP,omitempty" pn:"StartIP"`
	EndIP        *string `json:"EndIP,omitempty" pn:"EndIP"`
	LeaseTime    *int    `json:"LeaseTime,omitempty" pn:"LeaseTime"`
}

type DMZ struct {
	Enable             *string `json:"Enable,omitempty" pn:"Enable"`
	ICMPRedirectEnable *string `json:"ICMPRedirectEnable,omitempty" pn:"ICMPRedirectEnable"`
	HostAddress        *string `json:"HostAddress,omitempty" pn:"HostAddress"`
}
