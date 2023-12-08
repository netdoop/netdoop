package igd

type FAPService struct {
	AccessMgmt   *AccessMgmt          `json:"AccessMgmt,omitempty" pn:"AccessMgmt"`
	Capabilities *Capabilities        `json:"Capabilities,omitempty" pn:"Capabilities"`
	CellConfig   *CellConfig          `json:"CellConfig,omitempty" pn:"CellConfig"`
	CellConfig1  map[int]*CellConfig1 `json:"CellConfig1,omitempty" pn:"CellConfig1"`

	FAPControl       *FAPControl       `json:"FAPControl,omitempty" pn:"FAPControl"`
	SelfDefine       *SelfDefine       `json:"SelfDefine,omitempty" pn:"SelfDefine"`
	SelfDefineConfig *SelfDefineConfig `json:"SelfDefineConfig,omitempty" pn:"SelfDefineConfig"`
	Transport        *Transport        `json:"Transport,omitempty" pn:"Transport"`

	REM *struct {
		LTE *struct {
			X_VENDOR_ECGIList   *string `json:"X_VENDOR_ECGIList,omitempty" pn:"X_VENDOR_ECGIList"`
			X_VENDOR_ECGISwitch *string `json:"X_VENDOR_ECGISwitch,omitempty" pn:"X_VENDOR_ECGISwitch"`
		} `json:"LTE,omitempty" pn:"LTE"`
	} `json:"REM,omitempty" pn:"REM"`
}

type AccessMgmt struct {
	HNBIdentifier *string `json:"HNBIdentifier,omitempty" pn:"HNBIdentifier"`
	LTE           *struct {
		MaxUEsServed            *string `json:"MaxUEsServed,omitempty" pn:"MaxUEsServed"`
		SupportActiveRRCNumbers *string `json:"SupportActiveRRCNumbers,omitempty" pn:"SupportActiveRRCNumbers"`
		UeInactiveTimer         *string `json:"UeInactiveTimer,omitempty" pn:"UeInactiveTimer"`
	} `json:"LTE,omitempty" pn:"LTE"`
}

type Capabilities struct {
	LTE *struct {
		NNSFSupported *string `json:"NNSFSupported,omitempty" pn:"NNSFSupported"`
	} `json:"LTE,omitempty" pn:"LTE"`
	MaxTxPower *string `json:"MaxTxPower,omitempty" pn:"MaxTxPower"`
}

type MME_Comm_Info struct {
	IPAddrMain             *string `json:"IPAddrMain,omitempty" pn:"IPAddrMain"`
	IPAddrSpare            *string `json:"IPAddrSpare,omitempty" pn:"IPAddrSpare"`
	NumIpAddr              *string `json:"NumIpAddr,omitempty" pn:"NumIpAddr"`
	RelOfMME               *string `json:"RelOfMME,omitempty" pn:"RelOfMME"`
	X_VENDOR_ASSIGN_ENB_IP *string `json:"X_VENDOR_ASSIGN_ENB_IP,omitempty" pn:"X_VENDOR_ASSIGN_ENB_IP"`
	X_VENDOR_MME_PLMNID    *string `json:"X_VENDOR_MME_PLMNID,omitempty" pn:"X_VENDOR_MME_PLMNID"`
}

type FAPControl_LTE_Gateway struct {
	AGPort1                      *string                `json:"AGPort1,omitempty" pn:"AGPort1"`
	AGPort2                      *string                `json:"AGPort2,omitempty" pn:"AGPort2"`
	AGPort3                      *string                `json:"AGPort3,omitempty" pn:"AGPort3"`
	AGServerEnable               *string                `json:"AGServerEnable,omitempty" pn:"AGServerEnable"`
	AGServerIp1                  *string                `json:"AGServerIp1,omitempty" pn:"AGServerIp1"`
	AGServerIp2                  *string                `json:"AGServerIp2,omitempty" pn:"AGServerIp2"`
	AGServerIp3                  *string                `json:"AGServerIp3,omitempty" pn:"AGServerIp3"`
	MMENumberOfEntries           *uint                  `json:"MMENumberOfEntries,omitempty" pn:"MMENumberOfEntries"`
	MME_Comm_Info                map[int]*MME_Comm_Info `json:"MME_Comm_Info,omitempty" pn:"MME_Comm_Info"`
	S1SigLinkPort                *string                `json:"S1SigLinkPort,omitempty" pn:"S1SigLinkPort"`
	SecGWServer1                 *string                `json:"SecGWServer1,omitempty" pn:"SecGWServer1"`
	SecGWServer2                 *string                `json:"SecGWServer2,omitempty" pn:"SecGWServer2"`
	SecGWServer3                 *string                `json:"SecGWServer3,omitempty" pn:"SecGWServer3"`
	X_VENDOR_RAN_SHARIING_CONFIG *string                `json:"X_VENDOR_RAN_SHARIING_CONFIG,omitempty" pn:"X_VENDOR_RAN_SHARIING_CONFIG"`
}

type SelfConfig struct {
	SONConfigParam *SONConfigParam `json:"SONConfigParam,omitempty" pn:"SONConfigParam"`
}

type SONConfigParam struct {
	ANREnable                *string `json:"ANREnable,omitempty" pn:"ANREnable"`
	ANRGERANEnable           *string `json:"ANRGERANEnable,omitempty" pn:"ANRGERANEnable"`
	ANRInterFeqEnable        *string `json:"ANRInterFeqEnable,omitempty" pn:"ANRInterFeqEnable"`
	ANRUTRANEnable           *string `json:"ANRUTRANEnable,omitempty" pn:"ANRUTRANEnable"`
	ARFCNEnable              *string `json:"ARFCNEnable,omitempty" pn:"ARFCNEnable"`
	CandidateARFCNList       *string `json:"CandidateARFCNList,omitempty" pn:"CandidateARFCNList"`
	CandidatePCIList         *string `json:"CandidatePCIList,omitempty" pn:"CandidatePCIList"`
	GERANSnifferChannelList  *string `json:"GERANSnifferChannelList,omitempty" pn:"GERANSnifferChannelList"`
	GERANSnifferEnable       *string `json:"GERANSnifferEnable,omitempty" pn:"GERANSnifferEnable"`
	LTESnifferChannelList    *string `json:"LTESnifferChannelList,omitempty" pn:"LTESnifferChannelList"`
	LTESnifferFreqBandList   *string `json:"LTESnifferFreqBandList,omitempty" pn:"LTESnifferFreqBandList"`
	MROEnable                *string `json:"MROEnable,omitempty" pn:"MROEnable"`
	MaxGRANNeighbourCellNum  *string `json:"MaxGRANNeighbourCellNum,omitempty" pn:"MaxGRANNeighbourCellNum"`
	MaxLTENeighbourCellNum   *string `json:"MaxLTENeighbourCellNum,omitempty" pn:"MaxLTENeighbourCellNum"`
	MaxUTRANNeighbourCellNum *string `json:"MaxUTRANNeighbourCellNum,omitempty" pn:"MaxUTRANNeighbourCellNum"`
	PCIOptEnable             *string `json:"PCIOptEnable,omitempty" pn:"PCIOptEnable"`
	PCIReconfigWaitTime      *string `json:"PCIReconfigWaitTime,omitempty" pn:"PCIReconfigWaitTime"`
	PRACHConfigEnable        *string `json:"PRACHConfigEnable,omitempty" pn:"PRACHConfigEnable"`
	PowerEnable              *string `json:"PowerEnable,omitempty" pn:"PowerEnable"`
	ReSynCellEnable          *string `json:"ReSynCellEnable,omitempty" pn:"ReSynCellEnable"`
	RootSeqConfigEnable      *string `json:"RootSeqConfigEnable,omitempty" pn:"RootSeqConfigEnable"`
	SHEnable                 *string `json:"SHEnable,omitempty" pn:"SHEnable"`
	SONSysMode               *string `json:"SONSysMode,omitempty" pn:"SONSysMode"`
	SONWorkMode              *string `json:"SONWorkMode,omitempty" pn:"SONWorkMode"`
	SyncMode                 *string `json:"SyncMode,omitempty" pn:"SyncMode"`
	UTRANSnifferChannelList  *string `json:"UTRANSnifferChannelList,omitempty" pn:"UTRANSnifferChannelList"`
	UTRANSnifferEnable       *string `json:"UTRANSnifferEnable,omitempty" pn:"UTRANSnifferEnable"`
}

type FAPControl struct {
	LTE *struct {
		AdminState          *string                 `json:"AdminState,omitempty" pn:"AdminState"`
		Gateway             *FAPControl_LTE_Gateway `json:"Gateway,omitempty" pn:"Gateway"`
		OpState             *string                 `json:"OpState,omitempty" pn:"OpState"`
		RFTxStatus          *string                 `json:"RFTxStatus,omitempty" pn:"RFTxStatus"`
		RUHOControlEnable   *string                 `json:"RUHOControlEnable,omitempty" pn:"RUHOControlEnable"`
		SelfConfig          *SelfConfig             `json:"SelfConfig,omitempty" pn:"SelfConfig"`
		SourcedSignalEnable *string                 `json:"SourcedSignalEnable,omitempty" pn:"SourcedSignalEnable"`
	} `json:"LTE,omitempty" pn:"LTE"`
	LocalBreakout *struct {
		Enable   *string `json:"Enable,omitempty" pn:"Enable"`
		DnsRules *struct {
			Enable   *string `json:"Enable,omitempty" pn:"Enable"`
			RuleList map[int]*struct {
				DomainName   *string `json:"DomainName,omitempty" pn:"DomainName"`
				DomainIpList map[int]*struct {
					DomainIp *string `json:"DomainIp,omitempty" pn:"DomainIp"`
				} `json:"DomainIpList,omitempty" pn:"DomainIpList"`
			} `json:"RuleList,omitempty" pn:"RuleList"`
		} `json:"DnsRules,omitempty" pn:"DnsRules"`
		IpRules *struct {
			Enable   *string `json:"Enable,omitempty" pn:"Enable"`
			RuleList map[int]*struct {
				IPversion          *string `json:"IPversion,omitempty" pn:"IPversion"`
				ProtocolType       *string `json:"ProtocolType,omitempty" pn:"ProtocolType"`
				SourceIP           *string `json:"SourceIP,omitempty" pn:"SourceIP"`
				SourceMask         *string `json:"SourceMask,omitempty" pn:"SourceMask"`
				SourcePort         *string `json:"SourcePort,omitempty" pn:"SourcePort"`
				SourceIPv6         *string `json:"SourceIPv6,omitempty" pn:"SourceIPv6"`
				SourcePrefixLength *string `json:"SourcePrefixLength,omitempty" pn:"SourcePrefixLength"`
				TargetIP           *string `json:"TargetIP,omitempty" pn:"TargetIP"`
				TargetMask         *string `json:"TargetMask,omitempty" pn:"TargetMask"`
				TargetPort         *string `json:"TargetPort,omitempty" pn:"TargetPort"`
				TargetIPv6         *string `json:"TargetIPv6,omitempty" pn:"TargetIPv6"`
				TargetPrefixLength *string `json:"TargetPrefixLength,omitempty" pn:"TargetPrefixLength"`
			} `json:"RuleList,omitempty" pn:"RuleList"`
		} `json:"IpRules,omitempty" pn:"IpRules"`
		PlmnRules *struct {
			Enable   *string `json:"Enable,omitempty" pn:"Enable"`
			RuleList map[int]*struct {
				PLMNID *string `json:"PLMNID,omitempty" pn:"PLMNID"`
			} `json:"RuleList,omitempty" pn:"RuleList"`
		} `json:"PlmnRules,omitempty" pn:"PlmnRules"`
		NssaiRules *struct {
			Enable   *string `json:"Enable,omitempty" pn:"Enable"`
			RuleList map[int]*struct {
				SNSSAI *string `json:"SNSSAI,omitempty" pn:"SNSSAI"`
			} `json:"RuleList,omitempty" pn:"RuleList"`
		} `json:"NssaiRules,omitempty" pn:"NssaiRules"`

		TrafficRules *struct {
			RuleList map[int]*struct {
				CycleTo      *string `json:"CycleTo,omitempty" pn:"CycleTo"`
				ServerIP     *string `json:"ServerIP,omitempty" pn:"ServerIP"`
				CycleFrom    *string `json:"CycleFrom,omitempty" pn:"CycleFrom"`
				DLRateLimit  *string `json:"DLRateLimit,omitempty" pn:"DLRateLimit"`
				ULRateLimit  *string `json:"ULRateLimit,omitempty" pn:"ULRateLimit"`
				TrafficLimit *string `json:"TrafficLimit,omitempty" pn:"TrafficLimit"`
			} `json:"RuleList,omitempty" pn:"RuleList"`
		} `json:"TrafficRules,omitempty" pn:"TrafficRules"`

		TrafficStatistics *struct {
			Plmn map[int]*struct {
				PLMNID     *string `json:"PLMNID,omitempty" pn:"PLMNID"`
				Throughput *string `json:"Throughput,omitempty" pn:"Throughput"`
			} `json:"Plmn,omitempty" pn:"Plmn"`
			IP map[int]*struct {
				IPAddress  *string `json:"IPAddress,omitempty" pn:"IPAddress"`
				Throughput *string `json:"Throughput,omitempty" pn:"Throughput"`
			} `json:"IP,omitempty" pn:"IP"`
			Dns map[int]*struct {
				DomainName   *string `json:"DomainName,omitempty" pn:"DomainName"`
				Throughput   *string `json:"Throughput,omitempty" pn:"Throughput"`
				DomainIpList map[int]*struct {
					DomainIp *string `json:"DomainIp,omitempty" pn:"DomainIp"`
				} `json:"DomainIpList,omitempty" pn:"DomainIpList"`
			} `json:"Dns,omitempty" pn:"Dns"`
			Nssai map[int]*struct {
				SNSSAI     *string `json:"SNSSAI,omitempty" pn:"SNSSAI"`
				Throughput *string `json:"Throughput,omitempty" pn:"Throughput"`
			} `json:"Nssai,omitempty" pn:"Nssai"`
		} `json:"TrafficStatistics,omitempty" pn:"TrafficStatistics"`
	} `json:"LocalBreakout,omitempty" pn:"LocalBreakout"`
	StartUp *struct {
		Stage        *string `json:"Stage,omitempty" pn:"Stage"`
		Status       *string `json:"Status,omitempty" pn:"Status"`
		FailureCause *string `json:"FailureCause,omitempty" pn:"FailureCause"`
	} `json:"StartUp,omitempty" pn:"StartUp"`
}

type ConnMode_EUTRA struct {
	A1MeasureCtrl map[int]*struct {
		A1ThresholdRSRP *string `json:"A1ThresholdRSRP,omitempty" pn:"A1ThresholdRSRP"`
		A1ThresholdRSRQ *string `json:"A1ThresholdRSRQ,omitempty" pn:"A1ThresholdRSRQ"`
		Enable          *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis      *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells  *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose  *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount    *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval  *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		ReportQuantity  *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
		TimeToTrigger   *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
		TriggerQuantity *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
	} `json:"A1MeasureCtrl,omitempty" pn:"A1MeasureCtrl"`
	A2MeasureCtrl map[int]*struct {
		A2ThresholdRSRP *string `json:"A2ThresholdRSRP,omitempty" pn:"A2ThresholdRSRP"`
		A2ThresholdRSRQ *string `json:"A2ThresholdRSRQ,omitempty" pn:"A2ThresholdRSRQ"`
		Enable          *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis      *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells  *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose  *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount    *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval  *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		ReportQuantity  *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
		TimeToTrigger   *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
		TriggerQuantity *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
	} `json:"A2MeasureCtrl,omitempty" pn:"A2MeasureCtrl"`
	A3MeasureCtrl map[int]*struct {
		A3Offset        *string `json:"A3Offset,omitempty" pn:"A3Offset"`
		Enable          *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis      *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells  *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose  *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount    *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval  *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		ReportOnLeave   *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
		ReportQuantity  *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
		TimeToTrigger   *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
		TriggerQuantity *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
	} `json:"A3MeasureCtrl,omitempty" pn:"A3MeasureCtrl"`
	A4MeasureCtrl map[int]*struct {
		A4ThresholdRSRP *string `json:"A4ThresholdRSRP,omitempty" pn:"A4ThresholdRSRP"`
		A4ThresholdRSRQ *string `json:"A4ThresholdRSRQ,omitempty" pn:"A4ThresholdRSRQ"`
		Enable          *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis      *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells  *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose  *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount    *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval  *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		ReportQuantity  *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
		TimeToTrigger   *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
		TriggerQuantity *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
	} `json:"A4MeasureCtrl,omitempty" pn:"A4MeasureCtrl"`
	A5MeasureCtrl map[int]*struct {
		A5Threshold1RSRP *string `json:"A5Threshold1RSRP,omitempty" pn:"A5Threshold1RSRP"`
		A5Threshold1RSRQ *string `json:"A5Threshold1RSRQ,omitempty" pn:"A5Threshold1RSRQ"`
		A5Threshold2RSRP *string `json:"A5Threshold2RSRP,omitempty" pn:"A5Threshold2RSRP"`
		A5Threshold2RSRQ *string `json:"A5Threshold2RSRQ,omitempty" pn:"A5Threshold2RSRQ"`
		Enable           *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis       *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells   *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose   *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount     *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval   *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		ReportQuantity   *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
		TimeToTrigger    *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
		TriggerQuantity  *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
	} `json:"A5MeasureCtrl,omitempty" pn:"A5MeasureCtrl"`
	MeasureCtrl *struct {
		Smeasure *string `json:"Smeasure,omitempty" pn:"Smeasure"`
	} `json:"MeasureCtrl,omitempty" pn:"MeasureCtrl"`
	PeriodMeasCtrl map[int]*struct {
		MaxReportCells *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount   *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
	} `json:"PeriodMeasCtrl,omitempty" pn:"PeriodMeasCtrl"`
}

type ConnMode_IRAT struct {
	B1MeasureCtrl map[int]*struct {
		B1ThresholdCDMA2000 *string `json:"B1ThresholdCDMA2000,omitempty" pn:"B1ThresholdCDMA2000"`
		B1ThresholdGERAN    *string `json:"B1ThresholdGERAN,omitempty" pn:"B1ThresholdGERAN"`
		B1ThresholdUTRAEcN0 *string `json:"B1ThresholdUTRAEcN0,omitempty" pn:"B1ThresholdUTRAEcN0"`
		B1ThresholdUTRARSCP *string `json:"B1ThresholdUTRARSCP,omitempty" pn:"B1ThresholdUTRARSCP"`
		Enable              *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis          *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells      *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose      *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount        *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval      *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		TimeToTrigger       *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
	} `json:"B1MeasureCtrl,omitempty" pn:"B1MeasureCtrl"`
	B2MeasureCtrl map[int]*struct {
		B2Threshold1EutraRSRP *string `json:"B2Threshold1EutraRSRP,omitempty" pn:"B2Threshold1EutraRSRP"`
		B2Threshold1EutraRSRQ *string `json:"B2Threshold1EutraRSRQ,omitempty" pn:"B2Threshold1EutraRSRQ"`
		B2Threshold2CDMA2000  *string `json:"B2Threshold2CDMA2000,omitempty" pn:"B2Threshold2CDMA2000"`
		B2Threshold2GERAN     *string `json:"B2Threshold2GERAN,omitempty" pn:"B2Threshold2GERAN"`
		B2Threshold2UTRAEcN0  *string `json:"B2Threshold2UTRAEcN0,omitempty" pn:"B2Threshold2UTRAEcN0"`
		B2Threshold2UTRARSCP  *string `json:"B2Threshold2UTRARSCP,omitempty" pn:"B2Threshold2UTRARSCP"`
		Enable                *string `json:"Enable,omitempty" pn:"Enable"`
		Hysteresis            *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
		MaxReportCells        *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
		MeasurePurpose        *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
		ReportAmount          *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
		ReportInterval        *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
		TimeToTrigger         *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
	} `json:"B2MeasureCtrl,omitempty" pn:"B2MeasureCtrl"`
	B2hreshold_R1 *struct {
		B2Threshold1R15      *string `json:"B2Threshold1R15,omitempty" pn:"B2Threshold1R15"`
		B2Threshold1R15RSRP  *string `json:"B2Threshold1R15RSRP,omitempty" pn:"B2Threshold1R15RSRP"`
		B2Threshold1R15RSRQ  *string `json:"B2Threshold1R15RSRQ,omitempty" pn:"B2Threshold1R15RSRQ"`
		B2hreshold2NRR15     *string `json:"B2hreshold2NRR15,omitempty" pn:"B2hreshold2NRR15"`
		B2hreshold2NRR15RSRP *string `json:"B2hreshold2NRR15RSRP,omitempty" pn:"B2hreshold2NRR15RSRP"`
		B2hreshold2NRR15RSRQ *string `json:"B2hreshold2NRR15RSRQ,omitempty" pn:"B2hreshold2NRR15RSRQ"`
		B2hreshold2NRR15SINR *string `json:"B2hreshold2NRR15SINR,omitempty" pn:"B2hreshold2NRR15SINR"`
		B2hreshold_ENABLE    *string `json:"B2hreshold_ENABLE,omitempty" pn:"B2hreshold_ENABLE"`
		MeasurePurpose       *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
	} `json:"B2hreshold_R1,omitempty" pn:"B2hreshold_R1"`
	B2hreshold_R2 *struct {
		B2Threshold1R15      *string `json:"B2Threshold1R15,omitempty" pn:"B2Threshold1R15"`
		B2Threshold1R15RSRP  *string `json:"B2Threshold1R15RSRP,omitempty" pn:"B2Threshold1R15RSRP"`
		B2Threshold1R15RSRQ  *string `json:"B2Threshold1R15RSRQ,omitempty" pn:"B2Threshold1R15RSRQ"`
		B2hreshold2NRR15     *string `json:"B2hreshold2NRR15,omitempty" pn:"B2hreshold2NRR15"`
		B2hreshold2NRR15RSRP *string `json:"B2hreshold2NRR15RSRP,omitempty" pn:"B2hreshold2NRR15RSRP"`
		B2hreshold2NRR15RSRQ *string `json:"B2hreshold2NRR15RSRQ,omitempty" pn:"B2hreshold2NRR15RSRQ"`
		B2hreshold2NRR15SINR *string `json:"B2hreshold2NRR15SINR,omitempty" pn:"B2hreshold2NRR15SINR"`
		B2hreshold_ENABLE    *string `json:"B2hreshold_ENABLE,omitempty" pn:"B2hreshold_ENABLE"`
		MeasurePurpose       *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
	} `json:"B2hreshold_R2,omitempty" pn:"B2hreshold_R2"`
	MaxReportRSIndexR15       *string `json:"MaxReportRSIndexR15,omitempty" pn:"MaxReportRSIndexR15"`
	MeasQuantityGERAN         *string `json:"MeasQuantityGERAN,omitempty" pn:"MeasQuantityGERAN"`
	MeasQuantityUTRAFDD       *string `json:"MeasQuantityUTRAFDD,omitempty" pn:"MeasQuantityUTRAFDD"`
	QoffsetGERAN              *string `json:"QoffsetGERAN,omitempty" pn:"QoffsetGERAN"`
	QoffsetUTRA               *string `json:"QoffsetUTRA,omitempty" pn:"QoffsetUTRA"`
	QuantityConfigNRListcount *string `json:"QuantityConfigNRListcount,omitempty" pn:"QuantityConfigNRListcount"`
	QuantityConfigNRLIST      *struct {
		MaxReportRSIndexR15        *string `json:"MaxReportRSIndexR15,omitempty" pn:"MaxReportRSIndexR15"`
		QuantityConfigNRListcount  *string `json:"QuantityConfigNRListcount,omitempty" pn:"QuantityConfigNRListcount"`
		QuantityConfigNRListcount2 *string `json:"QuantityConfigNRListcount2,omitempty" pn:"QuantityConfigNRListcount2"`
		QuantityConfigNR           map[int]*struct {
			MeasQuantityCellNRr15_rsrp    *string `json:"MeasQuantityCellNRr15_rsrp,omitempty" pn:"MeasQuantityCellNRr15_rsrp"`
			MeasQuantityCellNRr15_rsrq    *string `json:"MeasQuantityCellNRr15_rsrq,omitempty" pn:"MeasQuantityCellNRr15_rsrq"`
			MeasQuantityCellNRr15_sinr    *string `json:"MeasQuantityCellNRr15_sinr,omitempty" pn:"MeasQuantityCellNRr15_sinr"`
			MeasQuantityRSIndexNR_enable  *string `json:"MeasQuantityRSIndexNR_enable,omitempty" pn:"MeasQuantityRSIndexNR_enable"`
			MeasQuantityRSIndexNRr15_rsrp *string `json:"MeasQuantityRSIndexNRr15_rsrp,omitempty" pn:"MeasQuantityRSIndexNRr15_rsrp"`
			MeasQuantityRSIndexNRr15_rsrq *string `json:"MeasQuantityRSIndexNRr15_rsrq,omitempty" pn:"MeasQuantityRSIndexNRr15_rsrq"`
			MeasQuantityRSIndexNRr15_sinr *string `json:"MeasQuantityRSIndexNRr15_sinr,omitempty" pn:"MeasQuantityRSIndexNRr15_sinr"`
		} `json:"QuantityConfigNR,omitempty" pn:"QuantityConfigNR"`
	} `json:"QuantityConfigNRLIST,omitempty" pn:"QuantityConfigNRLIST"`
	ReportOnLeaveR15           *string `json:"ReportOnLeaveR15,omitempty" pn:"ReportOnLeaveR15"`
	ReportQuantityCellNRR15    *string `json:"ReportQuantityCellNRR15,omitempty" pn:"ReportQuantityCellNRR15"`
	ReportQuantityRSIndexNRR15 *string `json:"ReportQuantityRSIndexNRR15,omitempty" pn:"ReportQuantityRSIndexNRR15"`
}

type IdleMode_Common struct {
	IntraFreqReselection *string `json:"IntraFreqReselection,omitempty" pn:"IntraFreqReselection"`
	NCellChangeHigh      *string `json:"NCellChangeHigh,omitempty" pn:"NCellChangeHigh"`
	NCellChangeMedium    *string `json:"NCellChangeMedium,omitempty" pn:"NCellChangeMedium"`
	QHystSFHigh          *string `json:"QHystSFHigh,omitempty" pn:"QHystSFHigh"`
	QHystSFMedium        *string `json:"QHystSFMedium,omitempty" pn:"QHystSFMedium"`
	Qhyst                *string `json:"Qhyst,omitempty" pn:"Qhyst"`
	TEvaluation          *string `json:"TEvaluation,omitempty" pn:"TEvaluation"`
	THystNormal          *string `json:"THystNormal,omitempty" pn:"THystNormal"`
}

type IdleMode_IRAT struct {
	GERAN *struct {
		TReselectionGERAN *string `json:"TReselectionGERAN,omitempty" pn:"TReselectionGERAN"`
		GERANFreqGroup    map[int]*struct {
			BCCHARFCN                    *string `json:"BCCHARFCN,omitempty" pn:"BCCHARFCN"`
			CellReselectionPriority      *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
			PMaxGERAN                    *string `json:"PMaxGERAN,omitempty" pn:"PMaxGERAN"`
			QRxLevMin                    *string `json:"QRxLevMin,omitempty" pn:"QRxLevMin"`
			ThreshXHigh                  *string `json:"ThreshXHigh,omitempty" pn:"ThreshXHigh"`
			ThreshXLow                   *string `json:"ThreshXLow,omitempty" pn:"ThreshXLow"`
			X_VENDOR_EXPLICIT_ARFCN_LIST *string `json:"X_VENDOR_EXPLICIT_ARFCN_LIST,omitempty" pn:"X_VENDOR_EXPLICIT_ARFCN_LIST"`
		} `json:"GERANFreqGroup,omitempty" pn:"GERANFreqGroup"`
	} `json:"GERAN,omitempty" pn:"GERAN"`

	NR_SIB_LIST *struct {
		CARRIERFREQNR map[int]*struct {
			Measurement_length               *string `json:"Measurement_length,omitempty" pn:"Measurement_length"`
			NR_NUMMULTIBANDINFOSUL_LENGTH    *string `json:"NR_NUMMULTIBANDINFOSUL_LENGTH,omitempty" pn:"NR_NUMMULTIBANDINFOSUL_LENGTH"`
			NR_NUMMULTIBANDINFO_LENGTH       *string `json:"NR_NUMMULTIBANDINFO_LENGTH,omitempty" pn:"NR_NUMMULTIBANDINFO_LENGTH"`
			SF                               *string `json:"SF,omitempty" pn:"SF"`
			SFBIT                            *string `json:"SFBIT,omitempty" pn:"SFBIT"`
			AdditionalPmaxNR_r15             *string `json:"AdditionalPmaxNR_r15,omitempty" pn:"AdditionalPmaxNR_r15"`
			AdditionalSpectrumEmissionNR_r15 *string `json:"AdditionalSpectrumEmissionNR_r15,omitempty" pn:"AdditionalSpectrumEmissionNR_r15"`
			CarrierFreq_r15                  *string `json:"CarrierFreq_r15,omitempty" pn:"CarrierFreq_r15"`
			CellReselectionPriority_r15      *string `json:"CellReselectionPriority_r15,omitempty" pn:"CellReselectionPriority_r15"`
			CellReselectionSubPriority_r15   *string `json:"CellReselectionSubPriority_r15,omitempty" pn:"CellReselectionSubPriority_r15"`
			DeriveSSB_IndexFromCell_r15      *string `json:"DeriveSSB_IndexFromCell_r15,omitempty" pn:"DeriveSSB_IndexFromCell_r15"`
			EndSymbol_r15                    *string `json:"EndSymbol_r15,omitempty" pn:"EndSymbol_r15"`
			MaxRS_IndexCellQual_r15          *string `json:"MaxRS_IndexCellQual_r15,omitempty" pn:"MaxRS_IndexCellQual_r15"`
			MeasurementSlots                 *string `json:"MeasurementSlots,omitempty" pn:"MeasurementSlots"`
			NR_RSRP_r15                      *string `json:"NR_RSRP_r15,omitempty" pn:"NR_RSRP_r15"`
			NR_RSRQ_r15                      *string `json:"NR_RSRQ_r15,omitempty" pn:"NR_RSRQ_r15"`
			NR_SINR_r15                      *string `json:"NR_SINR_r15,omitempty" pn:"NR_SINR_r15"`
			NR_multiBandInfo                 *string `json:"NR_multiBandInfo,omitempty" pn:"NR_multiBandInfo"`
			NR_multiBandInfosul              *string `json:"NR_multiBandInfosul,omitempty" pn:"NR_multiBandInfosul"`
			P_MaxNR_r15                      *string `json:"P_MaxNR_r15,omitempty" pn:"P_MaxNR_r15"`
			Q_QualMin_r15                    *string `json:"Q_QualMin_r15,omitempty" pn:"Q_QualMin_r15"`
			Q_RxLevMinSUL_r15                *string `json:"Q_RxLevMinSUL_r15,omitempty" pn:"Q_RxLevMinSUL_r15"`
			Q_RxLevMin_r15                   *string `json:"Q_RxLevMin_r15,omitempty" pn:"Q_RxLevMin_r15"`
			Size_of_NS_PmaxValueNR           *string `json:"Size_of_NS_PmaxValueNR,omitempty" pn:"Size_of_NS_PmaxValueNR"`
			Ssb_Duration_r15                 *string `json:"Ssb_Duration_r15,omitempty" pn:"Ssb_Duration_r15"`
			SubcarrierSpacingSSB_r15         *string `json:"SubcarrierSpacingSSB_r15,omitempty" pn:"SubcarrierSpacingSSB_r15"`
			ThreshRS_Index                   *string `json:"ThreshRS_Index,omitempty" pn:"ThreshRS_Index"`
			ThreshX_HighQ_r15                *string `json:"ThreshX_HighQ_r15,omitempty" pn:"ThreshX_HighQ_r15"`
			ThreshX_High_r15                 *string `json:"ThreshX_High_r15,omitempty" pn:"ThreshX_High_r15"`
			ThreshX_LowQ_r15                 *string `json:"ThreshX_LowQ_r15,omitempty" pn:"ThreshX_LowQ_r15"`
			ThreshX_Low_r15                  *string `json:"ThreshX_Low_r15,omitempty" pn:"ThreshX_Low_r15"`
		} `json:"CARRIERFREQNR,omitempty" pn:"CARRIERFREQNR"`
	} `json:"NR_SIB_LIST,omitempty" pn:"NR_SIB_LIST"`

	UTRA *struct {
		IRAT_EUTRAN_TO_UTRAN_TDD_CARRIERS map[int]*struct {
			CELL_RESELECTION_PRIORITY *string `json:"CELL_RESELECTION_PRIORITY,omitempty" pn:"CELL_RESELECTION_PRIORITY"`
			P_MAX_UTRA                *string `json:"P_MAX_UTRA,omitempty" pn:"P_MAX_UTRA"`
			Q_RX_LEV_MIN              *string `json:"Q_RX_LEV_MIN,omitempty" pn:"Q_RX_LEV_MIN"`
			THRESH_X_HIGH             *string `json:"THRESH_X_HIGH,omitempty" pn:"THRESH_X_HIGH"`
			THRESH_X_LOW              *string `json:"THRESH_X_LOW,omitempty" pn:"THRESH_X_LOW"`
			UTRA_CARRIER_ARFCN        *string `json:"UTRA_CARRIER_ARFCN,omitempty" pn:"UTRA_CARRIER_ARFCN"`
		} `json:"IRAT_EUTRAN_TO_UTRAN_TDD_CARRIERS,omitempty" pn:"IRAT_EUTRAN_TO_UTRAN_TDD_CARRIERS"`

		NUM_IRAT_EUTRAN_TO_UTRAN_TDD_CARRIERS *string `json:"NUM_IRAT_EUTRAN_TO_UTRAN_TDD_CARRIERS,omitempty" pn:"NUM_IRAT_EUTRAN_TO_UTRAN_TDD_CARRIERS"`
		TReselectionUTRA                      *string `json:"TReselectionUTRA,omitempty" pn:"TReselectionUTRA"`
		UTRANFDDFreqNumberOfEntries           *string `json:"UTRANFDDFreqNumberOfEntries,omitempty" pn:"UTRANFDDFreqNumberOfEntries"`
	} `json:"UTRA,omitempty" pn:"UTRA"`
}
type IdleMode_IntraFreq struct {
	AllowedMeasBandwidth      *string `json:"AllowedMeasBandwidth,omitempty" pn:"AllowedMeasBandwidth"`
	CellReselectionPriority   *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
	PMax                      *string `json:"PMax,omitempty" pn:"PMax"`
	QQualMinOffsetR9          *string `json:"QQualMinOffsetR9,omitempty" pn:"QQualMinOffsetR9"`
	QQualMinR9Reselection     *string `json:"QQualMinR9Reselection,omitempty" pn:"QQualMinR9Reselection"`
	QQualMinR9Selection       *string `json:"QQualMinR9Selection,omitempty" pn:"QQualMinR9Selection"`
	QRxLevMinOffset           *string `json:"QRxLevMinOffset,omitempty" pn:"QRxLevMinOffset"`
	QRxLevMinSIB1             *string `json:"QRxLevMinSIB1,omitempty" pn:"QRxLevMinSIB1"`
	QRxLevMinSIB3             *string `json:"QRxLevMinSIB3,omitempty" pn:"QRxLevMinSIB3"`
	SIntraSearch              *string `json:"SIntraSearch,omitempty" pn:"SIntraSearch"`
	SIntraSearchPR9           *string `json:"SIntraSearchPR9,omitempty" pn:"SIntraSearchPR9"`
	SIntraSearchQR9           *string `json:"SIntraSearchQR9,omitempty" pn:"SIntraSearchQR9"`
	SNonIntraSearch           *string `json:"SNonIntraSearch,omitempty" pn:"SNonIntraSearch"`
	SNonIntraSearchPR9        *string `json:"SNonIntraSearchPR9,omitempty" pn:"SNonIntraSearchPR9"`
	SNonIntraSearchQR9        *string `json:"SNonIntraSearchQR9,omitempty" pn:"SNonIntraSearchQR9"`
	TReselectionEUTRA         *string `json:"TReselectionEUTRA,omitempty" pn:"TReselectionEUTRA"`
	TReselectionEUTRASFHigh   *string `json:"TReselectionEUTRASFHigh,omitempty" pn:"TReselectionEUTRASFHigh"`
	TReselectionEUTRASFMedium *string `json:"TReselectionEUTRASFMedium,omitempty" pn:"TReselectionEUTRASFMedium"`
	ThreshServingLow          *string `json:"ThreshServingLow,omitempty" pn:"ThreshServingLow"`
	ThreshServingLowQR9       *string `json:"ThreshServingLowQR9,omitempty" pn:"ThreshServingLowQR9"`
}
type IdleMode_InterFreq struct {
	Carrier map[int]*struct {
		CellReselectionPriority            *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
		EUTRACarrierARFCN                  *string `json:"EUTRACarrierARFCN,omitempty" pn:"EUTRACarrierARFCN"`
		PMax                               *string `json:"PMax,omitempty" pn:"PMax"`
		QOffsetFreq                        *string `json:"QOffsetFreq,omitempty" pn:"QOffsetFreq"`
		QQualMinR9Reselection              *string `json:"QQualMinR9Reselection,omitempty" pn:"QQualMinR9Reselection"`
		QRxLevMinSIB5                      *string `json:"QRxLevMinSIB5,omitempty" pn:"QRxLevMinSIB5"`
		TReselectionEUTRA                  *string `json:"TReselectionEUTRA,omitempty" pn:"TReselectionEUTRA"`
		TReselectionEUTRASFHigh            *string `json:"TReselectionEUTRASFHigh,omitempty" pn:"TReselectionEUTRASFHigh"`
		TReselectionEUTRASFMedium          *string `json:"TReselectionEUTRASFMedium,omitempty" pn:"TReselectionEUTRASFMedium"`
		ThreshXHigh                        *string `json:"ThreshXHigh,omitempty" pn:"ThreshXHigh"`
		ThreshXHighQR9                     *string `json:"ThreshXHighQR9,omitempty" pn:"ThreshXHighQR9"`
		ThreshXLow                         *string `json:"ThreshXLow,omitempty" pn:"ThreshXLow"`
		ThreshXLowQR9                      *string `json:"ThreshXLowQR9,omitempty" pn:"ThreshXLowQR9"`
		X_VENDOR_MEAS_BANDWIDTH_FOR_EARFCN *string `json:"X_VENDOR_MEAS_BANDWIDTH_FOR_EARFCN,omitempty" pn:"X_VENDOR_MEAS_BANDWIDTH_FOR_EARFCN"`
	} `json:"Carrier,omitempty" pn:"Carrier"`
}

type SelfDefine struct {
	AdaptionHarq        *string `json:"AdaptionHarq,omitempty" pn:"AdaptionHarq"`
	Bsd                 *string `json:"Bsd,omitempty" pn:"Bsd"`
	DRBNum              *string `json:"DRBNum,omitempty" pn:"DRBNum"`
	DlPathlossChange    *string `json:"DlPathlossChange,omitempty" pn:"DlPathlossChange"`
	DlSchedulerStrategy *string `json:"DlSchedulerStrategy,omitempty" pn:"DlSchedulerStrategy"`
	DynamicPdcch        *string `json:"DynamicPdcch,omitempty" pn:"DynamicPdcch"`
	ERABmanage          *string `json:"ERABmanage,omitempty" pn:"ERABmanage"`
	EutranGapOffsetType *string `json:"EutranGapOffsetType,omitempty" pn:"EutranGapOffsetType"`
	FreqSelSwitchDL     *string `json:"FreqSelSwitchDL,omitempty" pn:"FreqSelSwitchDL"`
	FreqSelSwitchUL     *string `json:"FreqSelSwitchUL,omitempty" pn:"FreqSelSwitchUL"`
	Lcp                 *string `json:"Lcp,omitempty" pn:"Lcp"`
	LoadBlanceConfig    *string `json:"LoadBlanceConfig,omitempty" pn:"LoadBlanceConfig"`
	Msg3PowerCtrlUl     *string `json:"Msg3PowerCtrlUl,omitempty" pn:"Msg3PowerCtrlUl"`
	PCIconflict         *string `json:"PCIconflict,omitempty" pn:"PCIconflict"`
	PUCCHperiod         *string `json:"PUCCHperiod,omitempty" pn:"PUCCHperiod"`
	PUCCHtype           *string `json:"PUCCHtype,omitempty" pn:"PUCCHtype"`
	PUSCHPC             *string `json:"PUSCHPC,omitempty" pn:"PUSCHPC"`
	Pbr                 *string `json:"Pbr,omitempty" pn:"Pbr"`
	PeriodicPHRTimer    *string `json:"PeriodicPHRTimer,omitempty" pn:"PeriodicPHRTimer"`
	SIB                 *string `json:"SIB,omitempty" pn:"SIB"`
	TmMode              *string `json:"TmMode,omitempty" pn:"TmMode"`
	UEcontextmanage     *string `json:"UEcontextmanage,omitempty" pn:"UEcontextmanage"`
	ULMeanNL            *string `json:"ULMeanNL,omitempty" pn:"ULMeanNL"`
	Uecapacity          *string `json:"Uecapacity,omitempty" pn:"Uecapacity"`
	UlSchedulerStrategy *string `json:"UlSchedulerStrategy,omitempty" pn:"UlSchedulerStrategy"`
	LinkAdaption        *string `json:"LinkAdaption,omitempty" pn:"LinkAdaption"`
	Modulation          *string `json:"Modulation,omitempty" pn:"Modulation"`
}

type SelfDefineConfig struct {
	CellSelfDetectConfig *string `json:"CellSelfDetectConfig,omitempty" pn:"CellSelfDetectConfig"`
	L2LogLevel           *string `json:"L2LogLevel,omitempty" pn:"L2LogLevel"`
	L3LogLevel           *string `json:"L3LogLevel,omitempty" pn:"L3LogLevel"`
	OAMLogLevel          *string `json:"OAMLogLevel,omitempty" pn:"OAMLogLevel"`
	RRMLogLevel          *string `json:"RRMLogLevel,omitempty" pn:"RRMLogLevel"`
	RebootIntervalDays   *string `json:"RebootIntervalDays,omitempty" pn:"RebootIntervalDays"`
	RebootSwitch         *string `json:"RebootSwitch,omitempty" pn:"RebootSwitch"`
	RebootTime           *string `json:"RebootTime,omitempty" pn:"RebootTime"`
	SONLogLevel          *string `json:"SONLogLevel,omitempty" pn:"SONLogLevel"`
	SetLogLevel          *string `json:"SetLogLevel,omitempty" pn:"SetLogLevel"`
	DhcpLogLevel         *string `json:"DhcpLogLevel,omitempty" pn:"DhcpLogLevel"`
	IpsecLogLevel        *string `json:"IpsecLogLevel,omitempty" pn:"IpsecLogLevel"`
}

type Transport struct {
	SCTP *SCTP `json:"SCTP,omitempty" pn:"SCTP"`
}

type SCTP struct {
	Assoc map[int]*struct {
		LocalPort          *string `json:"LocalPort,omitempty" pn:"LocalPort"`
		PrimaryPeerAddress *string `json:"PrimaryPeerAddress,omitempty" pn:"PrimaryPeerAddress"`
		RemotePort         *string `json:"RemotePort,omitempty" pn:"RemotePort"`
		SCTPAssocLocalAddr *string `json:"SCTPAssocLocalAddr,omitempty" pn:"SCTPAssocLocalAddr"`
	} `json:",omitempty" pn:""`
	Enable                    *string `json:"Enable,omitempty" pn:"Enable"`
	HBInterval                *string `json:"HBInterval,omitempty" pn:"HBInterval"`
	MaxAssociationRetransmits *string `json:"MaxAssociationRetransmits,omitempty" pn:"MaxAssociationRetransmits"`
	MaxInitRetransmits        *string `json:"MaxInitRetransmits,omitempty" pn:"MaxInitRetransmits"`
	MaxPathRetransmits        *string `json:"MaxPathRetransmits,omitempty" pn:"MaxPathRetransmits"`
	RTOInitial                *string `json:"RTOInitial,omitempty" pn:"RTOInitial"`
	RTOMax                    *string `json:"RTOMax,omitempty" pn:"RTOMax"`
	RTOMin                    *string `json:"RTOMin,omitempty" pn:"RTOMin"`
	ValCookieLife             *string `json:"ValCookieLife,omitempty" pn:"ValCookieLife"`
}
