package igd

type CellConfig struct {
	LTE *struct {
		CAParam *struct {
			CASwitchDl *string `json:"CASwitchDl,omitempty" pn:"CASwitchDl"`
			CASwitchUl *string `json:"CASwitchUl,omitempty" pn:"CASwitchUl"`
		} `json:"CAParam,omitempty" pn:"CAParam"`
		EPC *struct {
			TAC                    *string `json:"TAC,omitempty" pn:"TAC"`
			EAID                   *string `json:"EAID,omitempty" pn:"EAID"`
			X_VENDOR_ESRVCC_SWITCH *string `json:"X_VENDOR_ESRVCC_SWITCH,omitempty" pn:"X_VENDOR_ESRVCC_SWITCH"`
			X_VENDOR_HO_EXT_CONFIG *string `json:"X_VENDOR_HO_EXT_CONFIG,omitempty" pn:"X_VENDOR_HO_EXT_CONFIG"`
			X_VENDOR_VOLTE_SWITCH  *string `json:"X_VENDOR_VOLTE_SWITCH,omitempty" pn:"X_VENDOR_VOLTE_SWITCH"`
			PLMNList               map[int]*struct {
				CellReservedForOperatorUse *string `json:"CellReservedForOperatorUse,omitempty" pn:"CellReservedForOperatorUse"`
				PLMNID                     *string `json:"PLMNID,omitempty" pn:"PLMNID"`
			} `json:"PLMNList,omitempty" pn:"PLMNList"`
			QoS map[int]*struct {
				PacketDelayBudget *int    `json:"PacketDelayBudget,omitempty" pn:"PacketDelayBudget"`
				Priority          *int    `json:"Priority,omitempty" pn:"Priority"`
				Type              *string `json:"Type,omitempty" pn:"Type"`
				X_VENDOR_BITRATE  *struct {
					MAX_DL_BITRATE *int `json:"MAX_DL_BITRATE,omitempty" pn:"MAX_DL_BITRATE"`
					MAX_UL_BITRATE *int `json:"MAX_UL_BITRATE,omitempty" pn:"MAX_UL_BITRATE"`
					MIN_DL_BITRATE *int `json:"MIN_DL_BITRATE,omitempty" pn:"MIN_DL_BITRATE"`
					MIN_UL_BITRATE *int `json:"MIN_UL_BITRATE,omitempty" pn:"MIN_UL_BITRATE"`
				} `json:"X_VENDOR_BITRATE,omitempty" pn:"X_VENDOR_BITRATE"`
				X_VENDOR_DISCARD_TIMER              *int `json:"X_VENDOR_DISCARD_TIMER,omitempty" pn:"X_VENDOR_DISCARD_TIMER"`
				X_VENDOR_DL_PARAM_T_REORDERING_AM   *int `json:"X_VENDOR_DL_PARAM_T_REORDERING_AM,omitempty" pn:"X_VENDOR_DL_PARAM_T_REORDERING_AM"`
				X_VENDOR_DL_PARAM_T_REORDERING_UM   *int `json:"X_VENDOR_DL_PARAM_T_REORDERING_UM,omitempty" pn:"X_VENDOR_DL_PARAM_T_REORDERING_UM"`
				X_VENDOR_LOGICAL_CHANNEL_GROUP      *int `json:"X_VENDOR_LOGICAL_CHANNEL_GROUP,omitempty" pn:"X_VENDOR_LOGICAL_CHANNEL_GROUP"`
				X_VENDOR_MAX_RETX_THRESHOLD         *int `json:"X_VENDOR_MAX_RETX_THRESHOLD,omitempty" pn:"X_VENDOR_MAX_RETX_THRESHOLD"`
				X_VENDOR_PACKETERRORLOSSRATE        *int `json:"X_VENDOR_PACKETERRORLOSSRATE,omitempty" pn:"X_VENDOR_PACKETERRORLOSSRATE"`
				X_VENDOR_RLC_MODE                   *int `json:"X_VENDOR_RLC_MODE,omitempty" pn:"X_VENDOR_RLC_MODE"`
				X_VENDOR_RLC_UM_PDCP_SN_SIZE        *int `json:"X_VENDOR_RLC_UM_PDCP_SN_SIZE,omitempty" pn:"X_VENDOR_RLC_UM_PDCP_SN_SIZE"`
				X_VENDOR_SN_LEN_DL_RLC              *int `json:"X_VENDOR_SN_LEN_DL_RLC,omitempty" pn:"X_VENDOR_SN_LEN_DL_RLC"`
				X_VENDOR_SN_LEN_UL_RLC              *int `json:"X_VENDOR_SN_LEN_UL_RLC,omitempty" pn:"X_VENDOR_SN_LEN_UL_RLC"`
				X_VENDOR_UE_INACTIVITY_TIMER_CONFIG *int `json:"X_VENDOR_UE_INACTIVITY_TIMER_CONFIG,omitempty" pn:"X_VENDOR_UE_INACTIVITY_TIMER_CONFIG"`
			} `json:"QoS,omitempty" pn:"QoS"`
		} `json:"EPC,omitempty" pn:"EPC"`
		RAN *struct {
			Mobility *struct {
				ConnMode *struct {
					EUTRA *ConnMode_EUTRA `json:"EUTRA,omitempty" pn:"EUTRA"`
					IRAT  *ConnMode_IRAT  `json:"IRAT,omitempty" pn:"IRAT"`
				} `json:"ConnMode,omitempty" pn:"ConnMode"`
				IdleMode *struct {
					Common    *IdleMode_Common    `json:"Common,omitempty" pn:"Common"`
					IRAT      *IdleMode_IRAT      `json:"IRAT,omitempty" pn:"IRAT"`
					InterFreq *IdleMode_InterFreq `json:"InterFreq,omitempty" pn:"InterFreq"`
					IntraFreq *IdleMode_IntraFreq `json:"IntraFreq,omitempty" pn:"IntraFreq"`
				} `json:"IdleMode,omitempty" pn:"IdleMode"`
			} `json:"Mobility,omitempty" pn:"Mobility"`
			MAC *struct {
				DRX *struct {
					DRXEnabled *int `json:"DRXEnabled,omitempty" pn:"DRXEnabled"`
				} `json:"DRX,omitempty" pn:"DRX"`
				DrxInitialParam map[int]*struct {
					DRXInactivityTimer     *string `json:"DRXInactivityTimer,omitempty" pn:"DRXInactivityTimer"`
					DRXRetransmissionTimer *string `json:"DRXRetransmissionTimer,omitempty" pn:"DRXRetransmissionTimer"`
					DRXShortCycleTimer     *string `json:"DRXShortCycleTimer,omitempty" pn:"DRXShortCycleTimer"`
					LongDRXCycle           *string `json:"LongDRXCycle,omitempty" pn:"LongDRXCycle"`
					ONDurationTimer        *string `json:"ONDurationTimer,omitempty" pn:"ONDurationTimer"`
					ShortDRXCycle          *string `json:"ShortDRXCycle,omitempty" pn:"ShortDRXCycle"`
				} `json:"DrxInitialParam,omitempty" pn:"DrxInitialParam"`
				RACH *struct {
					ContentionResolutionTimer          *string `json:"ContentionResolutionTimer,omitempty" pn:"ContentionResolutionTimer"`
					MaxHARQMsg3Tx                      *string `json:"MaxHARQMsg3Tx,omitempty" pn:"MaxHARQMsg3Tx"`
					MessagePowerOffsetGroupB           *string `json:"MessagePowerOffsetGroupB,omitempty" pn:"MessagePowerOffsetGroupB"`
					MessageSizeGroupA                  *string `json:"MessageSizeGroupA,omitempty" pn:"MessageSizeGroupA"`
					NumberOfRaPreambles                *string `json:"NumberOfRaPreambles,omitempty" pn:"NumberOfRaPreambles"`
					PowerRampingStep                   *string `json:"PowerRampingStep,omitempty" pn:"PowerRampingStep"`
					PreambleInitialReceivedTargetPower *string `json:"PreambleInitialReceivedTargetPower,omitempty" pn:"PreambleInitialReceivedTargetPower"`
					PreambleTransMax                   *string `json:"PreambleTransMax,omitempty" pn:"PreambleTransMax"`
					ResponseWindowSize                 *string `json:"ResponseWindowSize,omitempty" pn:"ResponseWindowSize"`
					SizeOfRaGroupA                     *string `json:"SizeOfRaGroupA,omitempty" pn:"SizeOfRaGroupA"`
				} `json:"RACH,omitempty" pn:"RACH"`
				ULSCH *struct {
					MaxHARQTx                       *string `json:"MaxHARQTx,omitempty" pn:"MaxHARQTx"`
					MaxUePerUlSf                    *string `json:"MaxUePerUlSf,omitempty" pn:"MaxUePerUlSf"`
					PeriodicBSRTimer                *string `json:"PeriodicBSRTimer,omitempty" pn:"PeriodicBSRTimer"`
					RetxBSRTimer                    *string `json:"RetxBSRTimer,omitempty" pn:"RetxBSRTimer"`
					TTIBundling                     *string `json:"TTIBundling,omitempty" pn:"TTIBundling"`
					X_VENDOR_CellExtraSchConfig     *string `json:"X_VENDOR_CellExtraSchConfig,omitempty" pn:"X_VENDOR_CellExtraSchConfig"`
					X_VENDOR_SR_final_periodicity   *string `json:"X_VENDOR_SR_final_periodicity,omitempty" pn:"X_VENDOR_SR_final_periodicity"`
					X_VENDOR_SR_initial_periodicity *string `json:"X_VENDOR_SR_initial_periodicity,omitempty" pn:"X_VENDOR_SR_initial_periodicity"`
					X_VENDOR_ulNumForcedGrants      *string `json:"X_VENDOR_ulNumForcedGrants,omitempty" pn:"X_VENDOR_ulNumForcedGrants"`
					X_VENDOR_ulSizeForcedGrant      *string `json:"X_VENDOR_ulSizeForcedGrant,omitempty" pn:"X_VENDOR_ulSizeForcedGrant"`
				} `json:"ULSCH,omitempty" pn:"ULSCH"`
				X_VENDOR_TCP_MSS *string `json:"X_VENDOR_TCP_MSS,omitempty" pn:"X_VENDOR_TCP_MSS"`
			} `json:"MAC,omitempty" pn:"MAC"`
			NeighborList *struct {
				InterRATCell *struct {
					GSM map[int]*struct {
						BCCHARFCN     *string `json:"BCCHARFCN,omitempty" pn:"BCCHARFCN"`
						BSIC          *string `json:"BSIC,omitempty" pn:"BSIC"`
						BandIndicator *string `json:"BandIndicator,omitempty" pn:"BandIndicator"`
						CI            *string `json:"CI,omitempty" pn:"CI"`
						LAC           *string `json:"LAC,omitempty" pn:"LAC"`
						PLMNID        *string `json:"PLMNID,omitempty" pn:"PLMNID"`
						RAC           *string `json:"RAC,omitempty" pn:"RAC"`
					} `json:"GSM,omitempty" pn:"GSM"`

					MaxNREntries *string `json:"MaxNREntries,omitempty" pn:"MaxNREntries"`

					NR map[int]*struct {
						Blacklisted              *string `json:"Blacklisted,omitempty" pn:"Blacklisted"`
						NbrGnbIdLength           *string `json:"NbrGnbIdLength,omitempty" pn:"NbrGnbIdLength"`
						PLMNID                   *string `json:"PLMNID,omitempty" pn:"PLMNID"`
						CarrierFreq_r15          *string `json:"CarrierFreq_r15,omitempty" pn:"CarrierFreq_r15"`
						Nr_cellid                *string `json:"Nr_cellid,omitempty" pn:"Nr_cellid"`
						Nr_pci                   *string `json:"Nr_pci,omitempty" pn:"Nr_pci"`
						SubcarrierSpacingSSB_r15 *string `json:"SubcarrierSpacingSSB_r15,omitempty" pn:"SubcarrierSpacingSSB_r15"`
						TAC                      *string `json:"TAC,omitempty" pn:"TAC"`
					} `json:"NR,omitempty" pn:"NR"`

					NRNumberOfEntries *uint `json:"NRNumberOfEntries,omitempty" pn:"NRNumberOfEntries"`

					UMTS map[int]*struct {
						CID                  *string `json:"CID,omitempty" pn:"CID"`
						EnbType              *string `json:"EnbType,omitempty" pn:"EnbType"`
						LAC                  *string `json:"LAC,omitempty" pn:"LAC"`
						PCPICHScramblingCode *string `json:"PCPICHScramblingCode,omitempty" pn:"PCPICHScramblingCode"`
						PCPICHTxPower        *string `json:"PCPICHTxPower,omitempty" pn:"PCPICHTxPower"`
						PLMNID               *string `json:"PLMNID,omitempty" pn:"PLMNID"`
						RAC                  *string `json:"RAC,omitempty" pn:"RAC"`
						RNCID                *string `json:"RNCID,omitempty" pn:"RNCID"`
						UARFCNDL             *string `json:"UARFCNDL,omitempty" pn:"UARFCNDL"`
						UARFCNUL             *string `json:"UARFCNUL,omitempty" pn:"UARFCNUL"`
						URA                  *string `json:"URA,omitempty" pn:"URA"`
					} `json:"UMTS,omitempty" pn:"UMTS"`
				} `json:"InterRATCell,omitempty" pn:"InterRATCell"`

				LTECell map[int]*struct {
					Blacklisted       *string `json:"Blacklisted,omitempty" pn:"Blacklisted"`
					CID               *string `json:"CID,omitempty" pn:"CID"`
					CIO               *string `json:"CIO,omitempty" pn:"CIO"`
					EUTRACarrierARFCN *string `json:"EUTRACarrierARFCN,omitempty" pn:"EUTRACarrierARFCN"`
					EnbType           *string `json:"EnbType,omitempty" pn:"EnbType"`
					PLMNID            *string `json:"PLMNID,omitempty" pn:"PLMNID"`
					PhyCellID         *string `json:"PhyCellID,omitempty" pn:"PhyCellID"`
					QOffset           *string `json:"QOffset,omitempty" pn:"QOffset"`
					RSTxPower         *string `json:"RSTxPower,omitempty" pn:"RSTxPower"`
					TAC               *string `json:"TAC,omitempty" pn:"TAC"`
					X_VENDOR_RSRP     *string `json:"X_VENDOR_RSRP,omitempty" pn:"X_VENDOR_RSRP"`
				} `json:"LTECell,omitempty" pn:"LTECell"`
			} `json:"NeighborList,omitempty" pn:"NeighborList"`
			PHY *struct {
				CarrierAggregationInfo *struct {
					CarrierNum *string `json:"CarrierNum,omitempty" pn:"CarrierNum"`
				} `json:"CarrierAggregationInfo,omitempty" pn:"CarrierAggregationInfo"`

				MBSFN *struct {
					NeighCellConfig *string `json:"NeighCellConfig,omitempty" pn:"NeighCellConfig"`

					SFConfigList map[int]*struct {
						RadioFrameAllocationPeriod *string `json:"RadioFrameAllocationPeriod,omitempty" pn:"RadioFrameAllocationPeriod"`
						RadioFrameAllocationSize   *string `json:"RadioFrameAllocationSize,omitempty" pn:"RadioFrameAllocationSize"`
						RadioframeAllocationOffset *string `json:"RadioframeAllocationOffset,omitempty" pn:"RadioframeAllocationOffset"`
						SubFrameAllocations        *string `json:"SubFrameAllocations,omitempty" pn:"SubFrameAllocations"`
						SyncStratumID              *string `json:"SyncStratumID,omitempty" pn:"SyncStratumID"`
					} `json:"SFConfigList,omitempty" pn:"SFConfigList"`
				} `json:"MBSFN,omitempty" pn:"MBSFN"`

				PDSCH *struct {
					Pa *string `json:"Pa,omitempty" pn:"Pa"`
					Pb *string `json:"Pb,omitempty" pn:"Pb"`
				} `json:"PDSCH,omitempty" pn:"PDSCH"`

				PRACH *struct {
					ConfigurationIndex        *string `json:"ConfigurationIndex,omitempty" pn:"ConfigurationIndex"`
					FreqOffset                *string `json:"FreqOffset,omitempty" pn:"FreqOffset"`
					HighSpeedFlag             *string `json:"HighSpeedFlag,omitempty" pn:"HighSpeedFlag"`
					RootSequenceIndex         *string `json:"RootSequenceIndex,omitempty" pn:"RootSequenceIndex"`
					ZeroCorrelationZoneConfig *string `json:"ZeroCorrelationZoneConfig,omitempty" pn:"ZeroCorrelationZoneConfig"`
				} `json:"PRACH,omitempty" pn:"PRACH"`

				PRS *struct {
					NumConsecutivePRSSubfames *string `json:"NumConsecutivePRSSubfames,omitempty" pn:"NumConsecutivePRSSubfames"`
					NumPRSResourceBlocks      *string `json:"NumPRSResourceBlocks,omitempty" pn:"NumPRSResourceBlocks"`
					PRSConfigurationIndex     *string `json:"PRSConfigurationIndex,omitempty" pn:"PRSConfigurationIndex"`
				} `json:"PRS,omitempty" pn:"PRS"`

				PUCCH *struct {
					CQIPUCCHResourceIndex *string `json:"CQIPUCCHResourceIndex,omitempty" pn:"CQIPUCCHResourceIndex"`
					DeltaPUCCHShift       *string `json:"DeltaPUCCHShift,omitempty" pn:"DeltaPUCCHShift"`
					K                     *string `json:"K,omitempty" pn:"K"`
					N1PUCCHAN             *string `json:"N1PUCCHAN,omitempty" pn:"N1PUCCHAN"`
					NCSAN                 *string `json:"NCSAN,omitempty" pn:"NCSAN"`
					NRBCQI                *string `json:"NRBCQI,omitempty" pn:"NRBCQI"`
				} `json:"PUCCH,omitempty" pn:"PUCCH"`

				PUSCH *struct {
					Enable64QAM   *string `json:"Enable64QAM,omitempty" pn:"Enable64QAM"`
					HoppingMode   *string `json:"HoppingMode,omitempty" pn:"HoppingMode"`
					HoppingOffset *string `json:"HoppingOffset,omitempty" pn:"HoppingOffset"`
					NSB           *string `json:"NSB,omitempty" pn:"NSB"`
				} `json:"PUSCH,omitempty" pn:"PUSCH"`

				PaParam *struct {
					PUCCHPowerCtrlSwitch *string `json:"PUCCHPowerCtrlSwitch,omitempty" pn:"PUCCHPowerCtrlSwitch"`
					PUSCHPowerCtrlSwitch *string `json:"PUSCHPowerCtrlSwitch,omitempty" pn:"PUSCHPowerCtrlSwitch"`
				} `json:"PaParam,omitempty" pn:"PaParam"`

				SRS *struct {
					AckNackSRSSimultaneousTransmission *string `json:"AckNackSRSSimultaneousTransmission,omitempty" pn:"AckNackSRSSimultaneousTransmission"`
					RootSequenceIndex                  *string `json:"RootSequenceIndex,omitempty" pn:"RootSequenceIndex"`
					SRSBandwidthConfig                 *string `json:"SRSBandwidthConfig,omitempty" pn:"SRSBandwidthConfig"`
					SRSEnabled                         *string `json:"SRSEnabled,omitempty" pn:"SRSEnabled"`
					SRSMaxUpPTS                        *string `json:"SRSMaxUpPTS,omitempty" pn:"SRSMaxUpPTS"`
				} `json:"SRS,omitempty" pn:"SRS"`

				TDDFrame *struct {
					SpecialSubframePatterns *string `json:"SpecialSubframePatterns,omitempty" pn:"SpecialSubframePatterns"`
					SubFrameAssignment      *string `json:"SubFrameAssignment,omitempty" pn:"SubFrameAssignment"`
				} `json:"TDDFrame,omitempty" pn:"TDDFrame"`

				ULPowerControl *struct {
					Alpha                    *string `json:"Alpha,omitempty" pn:"Alpha"`
					DeltaMCSEnabled          *string `json:"DeltaMCSEnabled,omitempty" pn:"DeltaMCSEnabled"`
					P0NominalPUCCH           *string `json:"P0NominalPUCCH,omitempty" pn:"P0NominalPUCCH"`
					P0NominalPUSCH           *string `json:"P0NominalPUSCH,omitempty" pn:"P0NominalPUSCH"`
					P0NominalPUSCHPersistent *string `json:"P0NominalPUSCHPersistent,omitempty" pn:"P0NominalPUSCHPersistent"`
				} `json:"ULPowerControl,omitempty" pn:"ULPowerControl"`
			} `json:"PHY,omitempty" pn:"PHY"`
			RF *struct {
				Ante1WithPssConfig *string `json:"Ante1WithPssConfig,omitempty" pn:"Ante1WithPssConfig"`
				DLBandwidth        *string `json:"DLBandwidth,omitempty" pn:"DLBandwidth"`
				EARFCNDL           *int    `json:"EARFCNDL,omitempty" pn:"EARFCNDL"`
				EARFCNUL           *int    `json:"EARFCNUL,omitempty" pn:"EARFCNUL"`
				FreqBandIndicator  *string `json:"FreqBandIndicator,omitempty" pn:"FreqBandIndicator"`
				PBCHPowerOffset    *string `json:"PBCHPowerOffset,omitempty" pn:"PBCHPowerOffset"`
				POWERAMPLIFIERList map[int]*struct {
					X_VENDOR_POWERAMPLIFIER_ADDR       *string `json:"X_VENDOR_POWERAMPLIFIER_ADDR,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_ADDR"`
					X_VENDOR_POWERAMPLIFIER_ATT        *string `json:"X_VENDOR_POWERAMPLIFIER_ATT,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_ATT"`
					X_VENDOR_POWERAMPLIFIER_FORPWRREF  *string `json:"X_VENDOR_POWERAMPLIFIER_FORPWRREF,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_FORPWRREF"`
					X_VENDOR_POWERAMPLIFIER_REVPWRREF  *string `json:"X_VENDOR_POWERAMPLIFIER_REVPWRREF,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_REVPWRREF"`
					X_VENDOR_POWERAMPLIFIER_RFSW       *string `json:"X_VENDOR_POWERAMPLIFIER_RFSW,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_RFSW"`
					X_VENDOR_POWERAMPLIFIER_TXPOWER    *string `json:"X_VENDOR_POWERAMPLIFIER_TXPOWER,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_TXPOWER"`
					X_VENDOR_POWERAMPLIFIER_VSWR       *string `json:"X_VENDOR_POWERAMPLIFIER_VSWR,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_VSWR"`
					X_VENDOR_POWERAMPLIFIER_VSWR_ALARM *string `json:"X_VENDOR_POWERAMPLIFIER_VSWR_ALARM,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_VSWR_ALARM"`
					X_VENDOR_POWERAMPLIFIER_VSWR_TEMP  *string `json:"X_VENDOR_POWERAMPLIFIER_VSWR_TEMP,omitempty" pn:"X_VENDOR_POWERAMPLIFIER_VSWR_TEMP"`
				} `json:"POWERAMPLIFIERList,omitempty" pn:"POWERAMPLIFIERList"`
				PSCHPowerOffset                *string `json:"PSCHPowerOffset,omitempty" pn:"PSCHPowerOffset"`
				PhyCellID                      *string `json:"PhyCellID,omitempty" pn:"PhyCellID"`
				ReferenceSignalPower           *string `json:"ReferenceSignalPower,omitempty" pn:"ReferenceSignalPower"`
				SSCHPowerOffset                *string `json:"SSCHPowerOffset,omitempty" pn:"SSCHPowerOffset"`
				ULBandwidth                    *string `json:"ULBandwidth,omitempty" pn:"ULBandwidth"`
				X_VENDOR_GPS_COMPENSATE_SWITCH *string `json:"X_VENDOR_GPS_COMPENSATE_SWITCH,omitempty" pn:"X_VENDOR_GPS_COMPENSATE_SWITCH"`
				X_VENDOR_GPS_COMPENSATE_VALUE  *string `json:"X_VENDOR_GPS_COMPENSATE_VALUE,omitempty" pn:"X_VENDOR_GPS_COMPENSATE_VALUE"`
				X_VENDOR_POWERAMPLIFIERLNUM    *string `json:"X_VENDOR_POWERAMPLIFIERLNUM,omitempty" pn:"X_VENDOR_POWERAMPLIFIERLNUM"`
				X_VENDOR_SYNC_METHOD           *string `json:"X_VENDOR_SYNC_METHOD,omitempty" pn:"X_VENDOR_SYNC_METHOD"`
				RB_OVER_DIMENSION_CONFIG       *string `json:"RB_OVER_DIMENSION_CONFIG,omitempty" pn:"RB_OVER_DIMENSION_CONFIG"`
			} `json:"RF,omitempty" pn:"RF"`
			RRCTimers *struct {
				N310      *string `json:"N310,omitempty" pn:"N310"`
				N311      *string `json:"N311,omitempty" pn:"N311"`
				T300      *string `json:"T300,omitempty" pn:"T300"`
				T301      *string `json:"T301,omitempty" pn:"T301"`
				T302      *string `json:"T302,omitempty" pn:"T302"`
				T304EUTRA *string `json:"T304EUTRA,omitempty" pn:"T304EUTRA"`
				T304IRAT  *string `json:"T304IRAT,omitempty" pn:"T304IRAT"`
				T310      *string `json:"T310,omitempty" pn:"T310"`
				T311      *string `json:"T311,omitempty" pn:"T311"`
				T320      *string `json:"T320,omitempty" pn:"T320"`
			} `json:"RRCTimers,omitempty" pn:"RRCTimers"`
			NeighborListInUse *struct {
				InterRATCell *struct {
					GSM map[int]*struct {
						BandIndicator *string `json:"BandIndicator,omitempty" pn:"BandIndicator"`
					} `json:"GSM,omitempty" pn:"GSM"`
				} `json:"InterRATCell,omitempty" pn:"InterRATCell"`
			} `json:"NeighborListInUse,omitempty" pn:"NeighborListInUse"`
			CellRestriction *struct {
				CellBarred *string `json:"CellBarred,omitempty" pn:"CellBarred"`
			} `json:"CellRestriction,omitempty" pn:"CellRestriction"`
			Common *struct {
				CellIdentity *string `json:"CellIdentity,omitempty" pn:"CellIdentity"`
				EnbType      *string `json:"EnbType,omitempty" pn:"EnbType"`
			} `json:"Common,omitempty" pn:"Common"`
			ANR *struct {
				X_VENDOR_ANR_ENABLE *string `json:"X_VENDOR_ANR_ENABLE,omitempty" pn:"X_VENDOR_ANR_ENABLE"`
			} `json:"ANR,omitempty" pn:"ANR"`
			X_VENDOR_ADDITIONAL_OPERATOR_INFO *struct {
				DEFAULT_PAGING_CYCLE *string `json:"DEFAULT_PAGING_CYCLE,omitempty" pn:"DEFAULT_PAGING_CYCLE"`
				SIB_SCHEDULING_INFO  map[int]*struct {
					NUM_SIB_TYPE   *string `json:"NUM_SIB_TYPE,omitempty" pn:"NUM_SIB_TYPE"`
					SIB_TYPE1      *string `json:"SIB_TYPE1,omitempty" pn:"SIB_TYPE1"`
					SIB_TYPE2      *string `json:"SIB_TYPE2,omitempty" pn:"SIB_TYPE2"`
					SIB_TYPE3      *string `json:"SIB_TYPE3,omitempty" pn:"SIB_TYPE3"`
					SIB_TYPE4      *string `json:"SIB_TYPE4,omitempty" pn:"SIB_TYPE4"`
					SI_PERIODICITY *string `json:"SI_PERIODICITY,omitempty" pn:"SI_PERIODICITY"`
				} `json:"SIB_SCHEDULING_INFO,omitempty" pn:"SIB_SCHEDULING_INFO"`
				SI_COUNT         *string `json:"SI_COUNT,omitempty" pn:"SI_COUNT"`
				SI_WINDOW_LENGTH *string `json:"SI_WINDOW_LENGTH,omitempty" pn:"SI_WINDOW_LENGTH"`
			} `json:"X_VENDOR_ADDITIONAL_OPERATOR_INFO,omitempty" pn:"X_VENDOR_ADDITIONAL_OPERATOR_INFO"`
			X_VENDOR_EMERGENCY_SERVICE *struct {
				IMS_EMERGENCY_SUPPORT_R9 *string `json:"IMS_EMERGENCY_SUPPORT_R9,omitempty" pn:"IMS_EMERGENCY_SUPPORT_R9"`
			} `json:"X_VENDOR_EMERGENCY_SERVICE,omitempty" pn:"X_VENDOR_EMERGENCY_SERVICE"`
		} `json:"RAN,omitempty" pn:"RAN"`
		MmePoolConfigParam map[int]*struct {
			MMECode    *string `json:"MMECode,omitempty" pn:"MMECode"`
			MMEGroupID *string `json:"MMEGroupID,omitempty" pn:"MMEGroupID"`
			MMEIp1     *string `json:"MMEIp1,omitempty" pn:"MMEIp1"`
			MMEIp2     *string `json:"MMEIp2,omitempty" pn:"MMEIp2"`
			PLMNID     *string `json:"PLMNID,omitempty" pn:"PLMNID"`
		} `json:"MmePoolConfigParam,omitempty" pn:"MmePoolConfigParam"`
		S1U map[int]*struct {
			FarIpSubnetworkList *string `json:"FarIpSubnetworkList,omitempty" pn:"FarIpSubnetworkList"`
			LocIpAddrList       *string `json:"LocIpAddrList,omitempty" pn:"LocIpAddrList"`
		} `json:"S1U,omitempty" pn:"S1U"`
		VoLTE *struct {
			PdcpInitParam map[int]*struct {
				RohcEn *string `json:"RohcEn,omitempty" pn:"RohcEn"`
			} `json:"PdcpInitParam,omitempty" pn:"PdcpInitParam"`
		} `json:"VoLTE,omitempty" pn:"VoLTE"`
		VoLTEParam *struct {
			SPSSwitchQCI1Ul *string `json:"SPSSwitchQCI1Ul,omitempty" pn:"SPSSwitchQCI1Ul"`
		} `json:"VoLTEParam,omitempty" pn:"VoLTEParam"`
		X_VENDOR_NR_FAST_RETURN     *string `json:"X_VENDOR_NR_FAST_RETURN,omitempty" pn:"X_VENDOR_NR_FAST_RETURN"`
		X_VENDOR_NR_SUPPORT_ENABLED *int    `json:"X_VENDOR_NR_SUPPORT_ENABLED,omitempty" pn:"X_VENDOR_NR_SUPPORT_ENABLED"`
	} `json:"LTE,omitempty" pn:"LTE"`
	SysInfoCtrlParam *struct {
		MultiBandInfoListSIB1 *string `json:"MultiBandInfoListSIB1,omitempty" pn:"MultiBandInfoListSIB1"`
		MultiBandInfoListSIB5 *string `json:"MultiBandInfoListSIB5,omitempty" pn:"MultiBandInfoListSIB5"`
	} `json:"SysInfoCtrlParam,omitempty" pn:"SysInfoCtrlParam"`
}
