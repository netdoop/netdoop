package igd

type CellConfig1 struct {
	NR *struct {
		NNSFSupported *string `json:"NNSFSupported,omitempty" pn:"NNSFSupported"`
		CN            *struct {
			EmergencyAreaID          *string `json:"EmergencyAreaID,omitempty" pn:"EmergencyAreaID"`
			CellReservedForOtherUse  *string `json:"CellReservedForOtherUse,omitempty" pn:"CellReservedForOtherUse"`
			EnableSliceAccessControl *string `json:"EnableSliceAccessControl,omitempty" pn:"EnableSliceAccessControl"`
			TA                       map[int]*struct {
				TAC                        *string `json:"TAC,omitempty" pn:"TAC"`
				GnbId                      *string `json:"GnbId,omitempty" pn:"GnbId"`
				GnbIdLength                *string `json:"GnbIdLength,omitempty" pn:"GnbIdLength"`
				Ranac                      *string `json:"Ranac,omitempty" pn:"Ranac"`
				NrcellIdentity             *string `json:"NrcellIdentity,omitempty" pn:"NrcellIdentity"`
				CellReservedForOperatorUse *string `json:"CellReservedForOperatorUse,omitempty" pn:"CellReservedForOperatorUse"`
				PLMNList                   map[int]*struct {
					PLMNID    *string `json:"PLMNID,omitempty" pn:"PLMNID"`
					SliceList map[int]*struct {
						SNSSAI *string `json:"SNSSAI,omitempty" pn:"SNSSAI"`
					} `json:"SliceList,omitempty" pn:"SliceList"`
				} `json:"PLMNList,omitempty" pn:"PLMNList"`
			} `json:"TA,omitempty" pn:"TA"`
		} `json:"CN,omitempty" pn:"CN"`
		RAN *struct {
			HoMode         *string `json:"HoMode,omitempty" pn:"HoMode"`
			RuList         *string `json:"RuList,omitempty" pn:"RuList"`
			OpState        *string `json:"OpState,omitempty" pn:"OpState"`
			CellMode       *string `json:"CellMode,omitempty" pn:"CellMode"`
			CellState      *string `json:"CellState,omitempty" pn:"CellState"`
			UserLabel      *string `json:"UserLabel,omitempty" pn:"UserLabel"`
			MaxTxPower     *string `json:"MaxTxPower,omitempty" pn:"MaxTxPower"`
			RouteIndexList *string `json:"RouteIndexList,omitempty" pn:"RouteIndexList"`
			Common         *struct {
				CellType        *string `json:"CellType,omitempty" pn:"CellType"`
				CellIdWithinGnb *string `json:"CellIdWithinGnb,omitempty" pn:"CellIdWithinGnb"`
				Security        *struct {
					Cipher map[int]*struct {
						Algorithm *string `json:"Algorithm,omitempty" pn:"Algorithm"`
						Priority  *string `json:"Priority,omitempty" pn:"Priority"`
					} `json:"Cipher,omitempty" pn:"Cipher"`
					Integrity map[int]*struct {
						Algorithm *string `json:"Algorithm,omitempty" pn:"Algorithm"`
						Priority  *string `json:"Priority,omitempty" pn:"Priority"`
					} `json:"Integrity,omitempty" pn:"Integrity"`
				} `json:"Security,omitempty" pn:"Security"`
			} `json:"Common,omitempty" pn:"Common"`
			CellEnable *struct {
				AdminState *string `json:"AdminState,omitempty" pn:"AdminState"`
			} `json:"CellEnable,omitempty" pn:"CellEnable"`
			CellRestriction *struct {
				CellBarred *string `json:"CellBarred,omitempty" pn:"CellBarred"`
			} `json:"CellRestriction,omitempty" pn:"CellRestriction"`
			CipheringAlgorithm *struct {
				Nea0Enable *string `json:"Nea0Enable,omitempty" pn:"Nea0Enable"`
				Nea1Enable *string `json:"Nea1Enable,omitempty" pn:"Nea1Enable"`
				Nea2Enable *string `json:"Nea2Enable,omitempty" pn:"Nea2Enable"`
				Nea3Enable *string `json:"Nea3Enable,omitempty" pn:"Nea3Enable"`
			} `json:"CipheringAlgorithm,omitempty" pn:"CipheringAlgorithm"`
			Drx map[int]*struct {
				QCI                      *string `json:"QCI,omitempty" pn:"QCI"`
				DrxAlgoSwitch            *string `json:"DrxAlgoSwitch,omitempty" pn:"DrxAlgoSwitch"`
				DrxShortCycle            *string `json:"DrxShortCycle,omitempty" pn:"DrxShortCycle"`
				DrxShortCycleEnabled     *string `json:"DrxShortCycleEnabled,omitempty" pn:"DrxShortCycleEnabled"`
				DrxShortCycleTimer       *string `json:"DrxShortCycleTimer,omitempty" pn:"DrxShortCycleTimer"`
				DrxLongCycle             *string `json:"DrxLongCycle,omitempty" pn:"DrxLongCycle"`
				DrxonDurationTimer       *string `json:"DrxonDurationTimer,omitempty" pn:"DrxonDurationTimer"`
				DrxInactivityTimer       *string `json:"DrxInactivityTimer,omitempty" pn:"DrxInactivityTimer"`
				DrxHARQRTTTimerDL        *string `json:"DrxHARQRTTTimerDL,omitempty" pn:"DrxHARQRTTTimerDL"`
				DrxHARQRTTTimerUL        *string `json:"DrxHARQRTTTimerUL,omitempty" pn:"DrxHARQRTTTimerUL"`
				DrxRetransmissionTimerDL *string `json:"DrxRetransmissionTimerDL,omitempty" pn:"DrxRetransmissionTimerDL"`
				DrxRetransmissionTimerUL *string `json:"DrxRetransmissionTimerUL,omitempty" pn:"DrxRetransmissionTimerUL"`
				PLMNList                 map[int]*struct {
					PLMNID *string `json:"PLMNID,omitempty" pn:"PLMNID"`
				} `json:"PLMNList,omitempty" pn:"PLMNList"`
			} `json:"Drx,omitempty" pn:"Drx"`
			EnergySaveParam *struct {
				AutoXnCellActivationEnable *string `json:"AutoXnCellActivationEnable,omitempty" pn:"AutoXnCellActivationEnable"`
				EnergySavingDelayTime      *string `json:"EnergySavingDelayTime,omitempty" pn:"EnergySavingDelayTime"`
				EnergySavingStage          *string `json:"EnergySavingStage,omitempty" pn:"EnergySavingStage"`
				EnergySavingTime           *string `json:"EnergySavingTime,omitempty" pn:"EnergySavingTime"`
				EnergySavingType           *string `json:"EnergySavingType,omitempty" pn:"EnergySavingType"`
				PrbLowThreshold            *string `json:"PrbLowThreshold,omitempty" pn:"PrbLowThreshold"`
				RRCLowThreshold            *string `json:"RRCLowThreshold,omitempty" pn:"RRCLowThreshold"`
			} `json:"EnergySaveParam,omitempty" pn:"EnergySaveParam"`
			Mobility *struct {
				ConnMode *struct {
					EUTRA *struct {
						MeasureCtrl *struct {
							Smeasure *string `json:"Smeasure,omitempty" pn:"Smeasure"`
						} `json:"MeasureCtrl,omitempty" pn:"MeasureCtrl"`
					} `json:"EUTRA,omitempty" pn:"EUTRA"`
					IRAT *struct {
						B1MeasureCtrl map[int]*struct {
							B1ThresholdEUTRARsrp *string `json:"B1ThresholdEUTRARsrp,omitempty" pn:"B1ThresholdEUTRARsrp"`
							B1ThresholdEUTRARsrq *string `json:"B1ThresholdEUTRARsrq,omitempty" pn:"B1ThresholdEUTRARsrq"`
							B1ThresholdEUTRASinr *string `json:"B1ThresholdEUTRASinr,omitempty" pn:"B1ThresholdEUTRASinr"`
							Enable               *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis           *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							MaxReportCells       *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose       *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAmount         *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval       *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave        *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity       *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							TimeToTrigger        *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity      *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							PlmnId               *struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"B1MeasureCtrl,omitempty" pn:"B1MeasureCtrl"`
						B2MeasureCtrl map[int]*struct {
							B2Threshold1RSRP      *string `json:"B2Threshold1RSRP,omitempty" pn:"B2Threshold1RSRP"`
							B2Threshold1RSRQ      *string `json:"B2Threshold1RSRQ,omitempty" pn:"B2Threshold1RSRQ"`
							B2Threshold1SINR      *string `json:"B2Threshold1SINR,omitempty" pn:"B2Threshold1SINR"`
							B2Threshold2EUTRARSRP *string `json:"B2Threshold2EUTRARSRP,omitempty" pn:"B2Threshold2EUTRARSRP"`
							B2Threshold2EUTRARSRQ *string `json:"B2Threshold2EUTRARSRQ,omitempty" pn:"B2Threshold2EUTRARSRQ"`
							B2Threshold2EUTRASINR *string `json:"B2Threshold2EUTRASINR,omitempty" pn:"B2Threshold2EUTRASINR"`
							Enable                *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis            *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							MaxReportCells        *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose        *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAmount          *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval        *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave         *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity        *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							TimeToTrigger         *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity       *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							PlmnId                *struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"B2MeasureCtrl,omitempty" pn:"B2MeasureCtrl"`
					} `json:"IRAT,omitempty" pn:"IRAT"`
					NR *struct {
						MeasureCtrl *struct {
							IntraCellHoEnabled *string `json:"IntraCellHoEnabled,omitempty" pn:"IntraCellHoEnabled"`
							Smeasure           *string `json:"Smeasure,omitempty" pn:"Smeasure"`
							InterFreq          *struct {
								SSBCarrier map[int]*struct {
									CsirsQOffsetRsrp            *string `json:"CsirsQOffsetRsrp,omitempty" pn:"CsirsQOffsetRsrp"`
									CsirsQOffsetRsrq            *string `json:"CsirsQOffsetRsrq,omitempty" pn:"CsirsQOffsetRsrq"`
									CsirsQOffsetSinr            *string `json:"CsirsQOffsetSinr,omitempty" pn:"CsirsQOffsetSinr"`
									DeriveSSBIndexFromCell      *string `json:"DeriveSSBIndexFromCell,omitempty" pn:"DeriveSSBIndexFromCell"`
									NrofCSIRsResourcesToAverage *string `json:"NrofCSIRsResourcesToAverage,omitempty" pn:"NrofCSIRsResourcesToAverage"`
									NrofSSBlocksToAverage       *string `json:"NrofSSBlocksToAverage,omitempty" pn:"NrofSSBlocksToAverage"`
									QuantityConfigIndex         *string `json:"QuantityConfigIndex,omitempty" pn:"QuantityConfigIndex"`
									RefFreqCSIRs                *string `json:"RefFreqCSIRs,omitempty" pn:"RefFreqCSIRs"`
									SSBSubcarrierSpacing        *string `json:"SSBSubcarrierSpacing,omitempty" pn:"SSBSubcarrierSpacing"`
									SsbFreq                     *string `json:"SsbFreq,omitempty" pn:"SsbFreq"`
									SsbQOffsetRsrp              *string `json:"SsbQOffsetRsrp,omitempty" pn:"SsbQOffsetRsrp"`
									SsbQOffsetRsrq              *string `json:"SsbQOffsetRsrq,omitempty" pn:"SsbQOffsetRsrq"`
									SsbQOffsetSinr              *string `json:"SsbQOffsetSinr,omitempty" pn:"SsbQOffsetSinr"`
									SsbToMeasureBitLength       *string `json:"SsbToMeasureBitLength,omitempty" pn:"SsbToMeasureBitLength"`
									SsbToMeasureSsbPosition     *string `json:"SsbToMeasureSsbPosition,omitempty" pn:"SsbToMeasureSsbPosition"`
									PlmnId                      *struct {
										PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
									} `json:"PlmnId,omitempty" pn:"PlmnId"`
									AbsThreshCSIRsConsolidation *struct {
										ThresholdRSRP *string `json:"ThresholdRSRP,omitempty" pn:"ThresholdRSRP"`
										ThresholdRSRQ *string `json:"ThresholdRSRQ,omitempty" pn:"ThresholdRSRQ"`
										ThresholdSINR *string `json:"ThresholdSINR,omitempty" pn:"ThresholdSINR"`
									} `json:"AbsThreshCSIRsConsolidation,omitempty" pn:"AbsThreshCSIRsConsolidation"`
									AbsThreshSSBlocksConsolidation *struct {
										ThresholdRSRP *string `json:"ThresholdRSRP,omitempty" pn:"ThresholdRSRP"`
										ThresholdRSRQ *string `json:"ThresholdRSRQ,omitempty" pn:"ThresholdRSRQ"`
										ThresholdSINR *string `json:"ThresholdSINR,omitempty" pn:"ThresholdSINR"`
									} `json:"AbsThreshSSBlocksConsolidation,omitempty" pn:"AbsThreshSSBlocksConsolidation"`
								} `json:"SSBCarrier,omitempty" pn:"SSBCarrier"`
							} `json:"InterFreq,omitempty" pn:"InterFreq"`
							InterRat *struct {
								EutraMeasObj map[int]*struct {
									AllowedMeasBandwidth *string `json:"AllowedMeasBandwidth,omitempty" pn:"AllowedMeasBandwidth"`
									EUTRACarrierARFCN    *string `json:"EUTRACarrierARFCN,omitempty" pn:"EUTRACarrierARFCN"`
									EutraQOffsetRange    *string `json:"EutraQOffsetRange,omitempty" pn:"EutraQOffsetRange"`
									PresenceAntennaPort1 *string `json:"PresenceAntennaPort1,omitempty" pn:"PresenceAntennaPort1"`
									PlmnId               *struct {
										PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
									} `json:"PlmnId,omitempty" pn:"PlmnId"`
								} `json:"EutraMeasObj,omitempty" pn:"EutraMeasObj"`
							} `json:"InterRat,omitempty" pn:"InterRat"`
							IntraFreq *struct {
								CsirsQOffsetRsrp            *string `json:"CsirsQOffsetRsrp,omitempty" pn:"CsirsQOffsetRsrp"`
								CsirsQOffsetRsrq            *string `json:"CsirsQOffsetRsrq,omitempty" pn:"CsirsQOffsetRsrq"`
								CsirsQOffsetSinr            *string `json:"CsirsQOffsetSinr,omitempty" pn:"CsirsQOffsetSinr"`
								DeriveSSBIndexFromCell      *string `json:"DeriveSSBIndexFromCell,omitempty" pn:"DeriveSSBIndexFromCell"`
								NrofCSIRsResourcesToAverage *string `json:"NrofCSIRsResourcesToAverage,omitempty" pn:"NrofCSIRsResourcesToAverage"`
								NrofSSBlocksToAverage       *string `json:"NrofSSBlocksToAverage,omitempty" pn:"NrofSSBlocksToAverage"`
								QuantityConfigIndex         *string `json:"QuantityConfigIndex,omitempty" pn:"QuantityConfigIndex"`
								SsbQOffsetRsrp              *string `json:"SsbQOffsetRsrp,omitempty" pn:"SsbQOffsetRsrp"`
								SsbQOffsetRsrq              *string `json:"SsbQOffsetRsrq,omitempty" pn:"SsbQOffsetRsrq"`
								SsbQOffsetSinr              *string `json:"SsbQOffsetSinr,omitempty" pn:"SsbQOffsetSinr"`
								SsbToMeasureBitLength       *string `json:"SsbToMeasureBitLength,omitempty" pn:"SsbToMeasureBitLength"`
								SsbToMeasureSsbPosition     *string `json:"SsbToMeasureSsbPosition,omitempty" pn:"SsbToMeasureSsbPosition"`
								AbsThreshCSIRsConsolidation *struct {
									ThresholdRSRP *string `json:"ThresholdRSRP,omitempty" pn:"ThresholdRSRP"`
									ThresholdRSRQ *string `json:"ThresholdRSRQ,omitempty" pn:"ThresholdRSRQ"`
									ThresholdSINR *string `json:"ThresholdSINR,omitempty" pn:"ThresholdSINR"`
								} `json:"AbsThreshCSIRsConsolidation,omitempty" pn:"AbsThreshCSIRsConsolidation"`
								AbsThreshSSBlocksConsolidation *struct {
									ThresholdRSRP *string `json:"ThresholdRSRP,omitempty" pn:"ThresholdRSRP"`
									ThresholdRSRQ *string `json:"ThresholdRSRQ,omitempty" pn:"ThresholdRSRQ"`
									ThresholdSINR *string `json:"ThresholdSINR,omitempty" pn:"ThresholdSINR"`
								} `json:"AbsThreshSSBlocksConsolidation,omitempty" pn:"AbsThreshSSBlocksConsolidation"`
							} `json:"IntraFreq,omitempty" pn:"IntraFreq"`
							QuantityConfig *struct {
								QuantityConfigEutra *struct {
									FilterRSRP *string `json:"FilterRSRP,omitempty" pn:"FilterRSRP"`
									FilterRSRQ *string `json:"FilterRSRQ,omitempty" pn:"FilterRSRQ"`
									FilterSINR *string `json:"FilterSINR,omitempty" pn:"FilterSINR"`
								} `json:"QuantityConfigEutra,omitempty" pn:"QuantityConfigEutra"`
								QuantityConfigNR map[int]*struct {
									QuantityConfigCell *struct {
										CsirsFilterRSRP *string `json:"CsirsFilterRSRP,omitempty" pn:"CsirsFilterRSRP"`
										CsirsFilterRSRQ *string `json:"CsirsFilterRSRQ,omitempty" pn:"CsirsFilterRSRQ"`
										CsirsFilterSINR *string `json:"CsirsFilterSINR,omitempty" pn:"CsirsFilterSINR"`
										SsbFilterRSRP   *string `json:"SsbFilterRSRP,omitempty" pn:"SsbFilterRSRP"`
										SsbFilterRSRQ   *string `json:"SsbFilterRSRQ,omitempty" pn:"SsbFilterRSRQ"`
										SsbFilterSINR   *string `json:"SsbFilterSINR,omitempty" pn:"SsbFilterSINR"`
									} `json:"QuantityConfigCell,omitempty" pn:"QuantityConfigCell"`
									QuantityConfigRsIndex *struct {
										CsirsFilterRSRP *string `json:"CsirsFilterRSRP,omitempty" pn:"CsirsFilterRSRP"`
										CsirsFilterRSRQ *string `json:"CsirsFilterRSRQ,omitempty" pn:"CsirsFilterRSRQ"`
										CsirsFilterSINR *string `json:"CsirsFilterSINR,omitempty" pn:"CsirsFilterSINR"`
										SsbFilterRSRP   *string `json:"SsbFilterRSRP,omitempty" pn:"SsbFilterRSRP"`
										SsbFilterRSRQ   *string `json:"SsbFilterRSRQ,omitempty" pn:"SsbFilterRSRQ"`
										SsbFilterSINR   *string `json:"SsbFilterSINR,omitempty" pn:"SsbFilterSINR"`
									} `json:"QuantityConfigRsIndex,omitempty" pn:"QuantityConfigRsIndex"`
								} `json:"QuantityConfigNR,omitempty" pn:"QuantityConfigNR"`
							} `json:"QuantityConfig,omitempty" pn:"QuantityConfig"`
						} `json:"MeasureCtrl,omitempty" pn:"MeasureCtrl"`
						PeriodMeasCtrl map[int]*struct {
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							UseWhiteCellList        *string `json:"UseWhiteCellList,omitempty" pn:"UseWhiteCellList"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"PeriodMeasCtrl,omitempty" pn:"PeriodMeasCtrl"`
						ReportCGI map[int]*struct {
							PhysCellId *string `json:"PhysCellId,omitempty" pn:"PhysCellId"`
						} `json:"ReportCGI,omitempty" pn:"ReportCGI"`
						A1MeasureCtrl map[int]*struct {
							A1ThresholdRSRP         *string `json:"A1ThresholdRSRP,omitempty" pn:"A1ThresholdRSRP"`
							A1ThresholdRSRQ         *string `json:"A1ThresholdRSRQ,omitempty" pn:"A1ThresholdRSRQ"`
							A1ThresholdSINR         *string `json:"A1ThresholdSINR,omitempty" pn:"A1ThresholdSINR"`
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis              *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAddNeighMeas      *string `json:"ReportAddNeighMeas,omitempty" pn:"ReportAddNeighMeas"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave           *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							TimeToTrigger           *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity         *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"A1MeasureCtrl,omitempty" pn:"A1MeasureCtrl"`
						A2MeasureCtrl map[int]*struct {
							A2ThresholdRSRP         *string `json:"A2ThresholdRSRP,omitempty" pn:"A2ThresholdRSRP"`
							A2ThresholdRSRQ         *string `json:"A2ThresholdRSRQ,omitempty" pn:"A2ThresholdRSRQ"`
							A2ThresholdSINR         *string `json:"A2ThresholdSINR,omitempty" pn:"A2ThresholdSINR"`
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis              *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAddNeighMeas      *string `json:"ReportAddNeighMeas,omitempty" pn:"ReportAddNeighMeas"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave           *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							TimeToTrigger           *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity         *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"A2MeasureCtrl,omitempty" pn:"A2MeasureCtrl"`
						A3MeasureCtrl map[int]*struct {
							A3OffsetRSRP            *string `json:"A3OffsetRSRP,omitempty" pn:"A3OffsetRSRP"`
							A3OffsetRSRQ            *string `json:"A3OffsetRSRQ,omitempty" pn:"A3OffsetRSRQ"`
							A3OffsetSINR            *string `json:"A3OffsetSINR,omitempty" pn:"A3OffsetSINR"`
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis              *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAddNeighMeas      *string `json:"ReportAddNeighMeas,omitempty" pn:"ReportAddNeighMeas"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave           *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							TimeToTrigger           *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity         *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							UseWhiteCellList        *string `json:"UseWhiteCellList,omitempty" pn:"UseWhiteCellList"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"A3MeasureCtrl,omitempty" pn:"A3MeasureCtrl"`
						A4MeasureCtrl map[int]*struct {
							A4ThresholdRSRP         *string `json:"A4ThresholdRSRP,omitempty" pn:"A4ThresholdRSRP"`
							A4ThresholdRSRQ         *string `json:"A4ThresholdRSRQ,omitempty" pn:"A4ThresholdRSRQ"`
							A4ThresholdSINR         *string `json:"A4ThresholdSINR,omitempty" pn:"A4ThresholdSINR"`
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis              *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAddNeighMeas      *string `json:"ReportAddNeighMeas,omitempty" pn:"ReportAddNeighMeas"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave           *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							TimeToTrigger           *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity         *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							UseWhiteCellList        *string `json:"UseWhiteCellList,omitempty" pn:"UseWhiteCellList"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"A4MeasureCtrl,omitempty" pn:"A4MeasureCtrl"`
						A5MeasureCtrl map[int]*struct {
							A5Threshold1RSRP        *string `json:"A5Threshold1RSRP,omitempty" pn:"A5Threshold1RSRP"`
							A5Threshold1RSRQ        *string `json:"A5Threshold1RSRQ,omitempty" pn:"A5Threshold1RSRQ"`
							A5Threshold1SINR        *string `json:"A5Threshold1SINR,omitempty" pn:"A5Threshold1SINR"`
							A5Threshold2RSRP        *string `json:"A5Threshold2RSRP,omitempty" pn:"A5Threshold2RSRP"`
							A5Threshold2RSRQ        *string `json:"A5Threshold2RSRQ,omitempty" pn:"A5Threshold2RSRQ"`
							A5Threshold2SINR        *string `json:"A5Threshold2SINR,omitempty" pn:"A5Threshold2SINR"`
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis              *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAddNeighMeas      *string `json:"ReportAddNeighMeas,omitempty" pn:"ReportAddNeighMeas"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave           *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							TimeToTrigger           *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity         *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							UseWhiteCellList        *string `json:"UseWhiteCellList,omitempty" pn:"UseWhiteCellList"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"A5MeasureCtrl,omitempty" pn:"A5MeasureCtrl"`
						A6MeasureCtrl map[int]*struct {
							A6OffsetRSRP            *string `json:"A6OffsetRSRP,omitempty" pn:"A6OffsetRSRP"`
							A6OffsetRSRQ            *string `json:"A6OffsetRSRQ,omitempty" pn:"A6OffsetRSRQ"`
							A6OffsetSINR            *string `json:"A6OffsetSINR,omitempty" pn:"A6OffsetSINR"`
							Enable                  *string `json:"Enable,omitempty" pn:"Enable"`
							Hysteresis              *string `json:"Hysteresis,omitempty" pn:"Hysteresis"`
							IncludeBeamMeasurements *string `json:"IncludeBeamMeasurements,omitempty" pn:"IncludeBeamMeasurements"`
							MaxNrofRSIndexToReport  *string `json:"MaxNrofRSIndexToReport,omitempty" pn:"MaxNrofRSIndexToReport"`
							MaxReportCells          *string `json:"MaxReportCells,omitempty" pn:"MaxReportCells"`
							MeasurePurpose          *string `json:"MeasurePurpose,omitempty" pn:"MeasurePurpose"`
							ReportAddNeighMeas      *string `json:"ReportAddNeighMeas,omitempty" pn:"ReportAddNeighMeas"`
							ReportAmount            *string `json:"ReportAmount,omitempty" pn:"ReportAmount"`
							ReportInterval          *string `json:"ReportInterval,omitempty" pn:"ReportInterval"`
							ReportOnLeave           *string `json:"ReportOnLeave,omitempty" pn:"ReportOnLeave"`
							ReportQuantity          *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
							RptQuantityRsIndex      *string `json:"RptQuantityRsIndex,omitempty" pn:"RptQuantityRsIndex"`
							RsType                  *string `json:"RsType,omitempty" pn:"RsType"`
							TimeToTrigger           *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
							TriggerQuantity         *string `json:"TriggerQuantity,omitempty" pn:"TriggerQuantity"`
							UseWhiteCellList        *string `json:"UseWhiteCellList,omitempty" pn:"UseWhiteCellList"`
							PlmnId                  map[int]*struct {
								PLMN *string `json:"PLMN,omitempty" pn:"PLMN"`
							} `json:"PlmnId,omitempty" pn:"PlmnId"`
						} `json:"A6MeasureCtrl,omitempty" pn:"A6MeasureCtrl"`
						Timer *struct {
							NgRelocOverAllTimer         *string `json:"NgRelocOverAllTimer,omitempty" pn:"NgRelocOverAllTimer"`
							NgRelocPrepTimer            *string `json:"NgRelocPrepTimer,omitempty" pn:"NgRelocPrepTimer"`
							PathSwitchAckTimer          *string `json:"PathSwitchAckTimer,omitempty" pn:"PathSwitchAckTimer"`
							PsHoEutranRelocOverAllTimer *string `json:"PsHoEutranRelocOverAllTimer,omitempty" pn:"PsHoEutranRelocOverAllTimer"`
							PsHoEutranRelocPrepTimer    *string `json:"PsHoEutranRelocPrepTimer,omitempty" pn:"PsHoEutranRelocPrepTimer"`
							UeHoInProgressTimer         *string `json:"UeHoInProgressTimer,omitempty" pn:"UeHoInProgressTimer"`
						} `json:"Timer,omitempty" pn:"Timer"`
					} `json:"NR,omitempty" pn:"NR"`
				} `json:"ConnMode,omitempty" pn:"ConnMode"`
				IdleMode *struct {
					Common *struct {
						IntraFreqReselection      *string `json:"IntraFreqReselection,omitempty" pn:"IntraFreqReselection"`
						NrofSSBlocksToAverage     *string `json:"NrofSSBlocksToAverage,omitempty" pn:"NrofSSBlocksToAverage"`
						Qhyst                     *string `json:"Qhyst,omitempty" pn:"Qhyst"`
						RangeToBestCell           *string `json:"RangeToBestCell,omitempty" pn:"RangeToBestCell"`
						TEvaluation               *string `json:"TEvaluation,omitempty" pn:"TEvaluation"`
						THystNormal               *string `json:"THystNormal,omitempty" pn:"THystNormal"`
						ThresholdRSRP             *string `json:"ThresholdRSRP,omitempty" pn:"ThresholdRSRP"`
						ThresholdRSRQ             *string `json:"ThresholdRSRQ,omitempty" pn:"ThresholdRSRQ"`
						ThresholdSINR             *string `json:"ThresholdSINR,omitempty" pn:"ThresholdSINR"`
						SpeedStateReselectionPars *struct {
							Enable   *string `json:"Enable,omitempty" pn:"Enable"`
							MobState *struct {
								NCellChangeHigh   *string `json:"NCellChangeHigh,omitempty" pn:"NCellChangeHigh"`
								NCellChangeMedium *string `json:"NCellChangeMedium,omitempty" pn:"NCellChangeMedium"`
								TEvaluation       *string `json:"TEvaluation,omitempty" pn:"TEvaluation"`
								THystNormal       *string `json:"THystNormal,omitempty" pn:"THystNormal"`
							} `json:"MobState,omitempty" pn:"MobState"`
							QHystSF *struct {
								SFHigh   *string `json:"SFHigh,omitempty" pn:"SFHigh"`
								SFMedium *string `json:"SFMedium,omitempty" pn:"SFMedium"`
							} `json:"QHystSF,omitempty" pn:"QHystSF"`
						} `json:"SpeedStateReselectionPars,omitempty" pn:"SpeedStateReselectionPars"`
					} `json:"Common,omitempty" pn:"Common"`
					EUTRA *struct {
						TReselectionEUTRA *string `json:"TReselectionEUTRA,omitempty" pn:"TReselectionEUTRA"`
						TReselectionSF    *struct {
							Enable   *string `json:"Enable,omitempty" pn:"Enable"`
							SFHigh   *string `json:"SFHigh,omitempty" pn:"SFHigh"`
							SFMedium *string `json:"SFMedium,omitempty" pn:"SFMedium"`
						} `json:"TReselectionSF,omitempty" pn:"TReselectionSF"`
						Carrier map[int]*struct {
							AllowedMeasBandwidth           *string `json:"AllowedMeasBandwidth,omitempty" pn:"AllowedMeasBandwidth"`
							BlackPhysCellIdRange           *string `json:"BlackPhysCellIdRange,omitempty" pn:"BlackPhysCellIdRange"`
							BlackPhysCellIdStart           *string `json:"BlackPhysCellIdStart,omitempty" pn:"BlackPhysCellIdStart"`
							CellReselectionPriority        *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
							CellReselectionSubPriority     *string `json:"CellReselectionSubPriority,omitempty" pn:"CellReselectionSubPriority"`
							EUTRACarrierARFCN              *string `json:"EUTRACarrierARFCN,omitempty" pn:"EUTRACarrierARFCN"`
							PMaxEUTRA                      *string `json:"PMaxEUTRA,omitempty" pn:"PMaxEUTRA"`
							PresenceAntennaPort1           *string `json:"PresenceAntennaPort1,omitempty" pn:"PresenceAntennaPort1"`
							QQualMin                       *string `json:"QQualMin,omitempty" pn:"QQualMin"`
							QRxLevMin                      *string `json:"QRxLevMin,omitempty" pn:"QRxLevMin"`
							TReselectionEUTRA              *string `json:"TReselectionEUTRA,omitempty" pn:"TReselectionEUTRA"`
							ThreshXHigh                    *string `json:"ThreshXHigh,omitempty" pn:"ThreshXHigh"`
							ThreshXHighQ                   *string `json:"ThreshXHighQ,omitempty" pn:"ThreshXHighQ"`
							ThreshXLow                     *string `json:"ThreshXLow,omitempty" pn:"ThreshXLow"`
							ThreshXLowQ                    *string `json:"ThreshXLowQ,omitempty" pn:"ThreshXLowQ"`
							CellReselectionPrioritySpecial map[int]*struct {
								CellReselectionPriority    *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
								CellReselectionSubPriority *string `json:"CellReselectionSubPriority,omitempty" pn:"CellReselectionSubPriority"`
								Plmn                       *string `json:"Plmn,omitempty" pn:"Plmn"`
							} `json:"CellReselectionPrioritySpecial,omitempty" pn:"CellReselectionPrioritySpecial"`
						} `json:"Carrier,omitempty" pn:"Carrier"`
					} `json:"EUTRA,omitempty" pn:"EUTRA"`
					InterFreq *struct {
						Carrier map[int]*struct {
							AbsThreshSSBlocksConsolidation *string `json:"AbsThreshSSBlocksConsolidation,omitempty" pn:"AbsThreshSSBlocksConsolidation"`
							BlackPhysCellIdRange           *string `json:"BlackPhysCellIdRange,omitempty" pn:"BlackPhysCellIdRange"`
							BlackPhysCellIdStart           *string `json:"BlackPhysCellIdStart,omitempty" pn:"BlackPhysCellIdStart"`
							CellReselectionPriority        *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
							CellReselectionSubPriority     *string `json:"CellReselectionSubPriority,omitempty" pn:"CellReselectionSubPriority"`
							DeriveSSBIndexFromCell         *string `json:"DeriveSSBIndexFromCell,omitempty" pn:"DeriveSSBIndexFromCell"`
							DlCarrierFreq                  *string `json:"DlCarrierFreq,omitempty" pn:"DlCarrierFreq"`
							NrofSSBlocksToAverage          *string `json:"NrofSSBlocksToAverage,omitempty" pn:"NrofSSBlocksToAverage"`
							Pmax                           *string `json:"Pmax,omitempty" pn:"Pmax"`
							QOffsetFreq                    *string `json:"QOffsetFreq,omitempty" pn:"QOffsetFreq"`
							QQualMin                       *string `json:"QQualMin,omitempty" pn:"QQualMin"`
							QRxLevMin                      *string `json:"QRxLevMin,omitempty" pn:"QRxLevMin"`
							SsbSubcarrierSpacing           *string `json:"SsbSubcarrierSpacing,omitempty" pn:"SsbSubcarrierSpacing"`
							SsbToMeasureBitLength          *string `json:"SsbToMeasureBitLength,omitempty" pn:"SsbToMeasureBitLength"`
							SsbToMeasureSsbPosition        *string `json:"SsbToMeasureSsbPosition,omitempty" pn:"SsbToMeasureSsbPosition"`
							TReselectionNR                 *string `json:"TReselectionNR,omitempty" pn:"TReselectionNR"`
							ThreshXHighP                   *string `json:"ThreshXHighP,omitempty" pn:"ThreshXHighP"`
							ThreshXHighQ                   *string `json:"ThreshXHighQ,omitempty" pn:"ThreshXHighQ"`
							ThreshXLowP                    *string `json:"ThreshXLowP,omitempty" pn:"ThreshXLowP"`
							ThreshXLowQ                    *string `json:"ThreshXLowQ,omitempty" pn:"ThreshXLowQ"`
							ThresholdRSRQ                  *string `json:"ThresholdRSRQ,omitempty" pn:"ThresholdRSRQ"`
							ThresholdSINR                  *string `json:"ThresholdSINR,omitempty" pn:"ThresholdSINR"`
							CellReselectionPrioritySpecial map[int]*struct {
								CellReselectionPriority    *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
								CellReselectionSubPriority *string `json:"CellReselectionSubPriority,omitempty" pn:"CellReselectionSubPriority"`
								Plmn                       *string `json:"Plmn,omitempty" pn:"Plmn"`
								Snssai                     *string `json:"Snssai,omitempty" pn:"Snssai"`
							} `json:"CellReselectionPrioritySpecial,omitempty" pn:"CellReselectionPrioritySpecial"`
							MultiFrequencyBandListNRSib map[int]*struct {
								FreqBandIndicatorNR *string `json:"FreqBandIndicatorNR,omitempty" pn:"FreqBandIndicatorNR"`
								NrNsPmax            map[int]*struct {
									AdditionalPmax             *string `json:"AdditionalPmax,omitempty" pn:"AdditionalPmax"`
									AdditionalSpectrumEmission *string `json:"AdditionalSpectrumEmission,omitempty" pn:"AdditionalSpectrumEmission"`
								} `json:"NrNsPmax,omitempty" pn:"NrNsPmax"`
							} `json:"MultiFrequencyBandListNRSib,omitempty" pn:"MultiFrequencyBandListNRSib"`
							TReselectionNRSF *struct {
								Enable   *string `json:"Enable,omitempty" pn:"Enable"`
								SFHigh   *string `json:"SFHigh,omitempty" pn:"SFHigh"`
								SFMedium *string `json:"SFMedium,omitempty" pn:"SFMedium"`
							} `json:"TReselectionNRSF,omitempty" pn:"TReselectionNRSF"`
						} `json:"Carrier,omitempty" pn:"Carrier"`
					} `json:"InterFreq,omitempty" pn:"InterFreq"`
					IntraFreq *struct {
						BlackPhysCellIdRange           *string `json:"BlackPhysCellIdRange,omitempty" pn:"BlackPhysCellIdRange"`
						BlackPhysCellIdStart           *string `json:"BlackPhysCellIdStart,omitempty" pn:"BlackPhysCellIdStart"`
						CellReselectionPriority        *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
						CellReselectionSubPriority     *string `json:"CellReselectionSubPriority,omitempty" pn:"CellReselectionSubPriority"`
						DeriveSSBIndexFromCell         *string `json:"DeriveSSBIndexFromCell,omitempty" pn:"DeriveSSBIndexFromCell"`
						Pmax                           *string `json:"Pmax,omitempty" pn:"Pmax"`
						QQualMinSIB1                   *string `json:"QQualMinSIB1,omitempty" pn:"QQualMinSIB1"`
						QQualMinSIB2                   *string `json:"QQualMinSIB2,omitempty" pn:"QQualMinSIB2"`
						QRxLevMinOffset                *string `json:"QRxLevMinOffset,omitempty" pn:"QRxLevMinOffset"`
						QRxLevMinSIB1                  *string `json:"QRxLevMinSIB1,omitempty" pn:"QRxLevMinSIB1"`
						QRxLevMinSIB2                  *string `json:"QRxLevMinSIB2,omitempty" pn:"QRxLevMinSIB2"`
						Qoffsettemp                    *string `json:"Qoffsettemp,omitempty" pn:"Qoffsettemp"`
						Qqualminoffset                 *string `json:"Qqualminoffset,omitempty" pn:"Qqualminoffset"`
						SIntraSearchP                  *string `json:"SIntraSearchP,omitempty" pn:"SIntraSearchP"`
						SIntraSearchQ                  *string `json:"SIntraSearchQ,omitempty" pn:"SIntraSearchQ"`
						SNonIntraSearchP               *string `json:"SNonIntraSearchP,omitempty" pn:"SNonIntraSearchP"`
						SNonIntraSearchQ               *string `json:"SNonIntraSearchQ,omitempty" pn:"SNonIntraSearchQ"`
						SsbToMeasureBitLength          *string `json:"SsbToMeasureBitLength,omitempty" pn:"SsbToMeasureBitLength"`
						SsbToMeasureSsbPosition        *string `json:"SsbToMeasureSsbPosition,omitempty" pn:"SsbToMeasureSsbPosition"`
						TReselectionNR                 *string `json:"TReselectionNR,omitempty" pn:"TReselectionNR"`
						ThreshServingLowP              *string `json:"ThreshServingLowP,omitempty" pn:"ThreshServingLowP"`
						ThreshServingLowQ              *string `json:"ThreshServingLowQ,omitempty" pn:"ThreshServingLowQ"`
						CellReselectionPrioritySpecial map[int]*struct {
							CellReselectionPriority    *string `json:"CellReselectionPriority,omitempty" pn:"CellReselectionPriority"`
							CellReselectionSubPriority *string `json:"CellReselectionSubPriority,omitempty" pn:"CellReselectionSubPriority"`
							Plmn                       *string `json:"Plmn,omitempty" pn:"Plmn"`
							Snssai                     *string `json:"Snssai,omitempty" pn:"Snssai"`
						} `json:"CellReselectionPrioritySpecial,omitempty" pn:"CellReselectionPrioritySpecial"`
						MultiFrequencyBandListNRSib map[int]*struct {
							FreqBandIndicatorNR *string `json:"FreqBandIndicatorNR,omitempty" pn:"FreqBandIndicatorNR"`
							NrNsPmax            map[int]*struct {
								AdditionalPmax             *string `json:"AdditionalPmax,omitempty" pn:"AdditionalPmax"`
								AdditionalSpectrumEmission *string `json:"AdditionalSpectrumEmission,omitempty" pn:"AdditionalSpectrumEmission"`
							} `json:"NrNsPmax,omitempty" pn:"NrNsPmax"`
						} `json:"MultiFrequencyBandListNRSib,omitempty" pn:"MultiFrequencyBandListNRSib"`
						TReselectionNRSF map[int]*struct {
							Enable   *string `json:"Enable,omitempty" pn:"Enable"`
							SFHigh   *string `json:"SFHigh,omitempty" pn:"SFHigh"`
							SFMedium *string `json:"SFMedium,omitempty" pn:"SFMedium"`
						} `json:"TReselectionNRSF,omitempty" pn:"TReselectionNRSF"`
					} `json:"IntraFreq,omitempty" pn:"IntraFreq"`
					InactiveMode *struct {
						RANPagingTimer *string `json:"RANPagingTimer,omitempty" pn:"RANPagingTimer"`
					} `json:"InactiveMode,omitempty" pn:"InactiveMode"`
				} `json:"IdleMode,omitempty" pn:"IdleMode"`
			} `json:"Mobility,omitempty" pn:"Mobility"`
			NeighborList *struct {
				LTECell map[int]*struct {
					Blacklisted         *string `json:"Blacklisted,omitempty" pn:"Blacklisted"`
					CID                 *string `json:"CID,omitempty" pn:"CID"`
					CIO                 *string `json:"CIO,omitempty" pn:"CIO"`
					EUTRACarrierARFCN   *string `json:"EUTRACarrierARFCN,omitempty" pn:"EUTRACarrierARFCN"`
					EnbType             *string `json:"EnbType,omitempty" pn:"EnbType"`
					NoHOEnable          *string `json:"NoHOEnable,omitempty" pn:"NoHOEnable"`
					NoRemoveEnable      *string `json:"NoRemoveEnable,omitempty" pn:"NoRemoveEnable"`
					NoX2Enable          *string `json:"NoX2Enable,omitempty" pn:"NoX2Enable"`
					PLMNID              *string `json:"PLMNID,omitempty" pn:"PLMNID"`
					PhyCellID           *string `json:"PhyCellID,omitempty" pn:"PhyCellID"`
					QQualMinOffsetCell  *string `json:"QQualMinOffsetCell,omitempty" pn:"QQualMinOffsetCell"`
					QRxLevMinOffsetCell *string `json:"QRxLevMinOffsetCell,omitempty" pn:"QRxLevMinOffsetCell"`
					Qoffset             *string `json:"Qoffset,omitempty" pn:"Qoffset"`
					Status              *string `json:"Status,omitempty" pn:"Status"`
					TAC                 *string `json:"TAC,omitempty" pn:"TAC"`
					TargetPLMNID        *string `json:"TargetPLMNID,omitempty" pn:"TargetPLMNID"`
					TcRAT               *string `json:"TcRAT,omitempty" pn:"TcRAT"`
				} `json:"LTECell,omitempty" pn:"LTECell"`
				NRCell map[int]*struct {
					Blacklisted          *string `json:"Blacklisted,omitempty" pn:"Blacklisted"`
					CID                  *string `json:"CID,omitempty" pn:"CID"`
					CIO                  *string `json:"CIO,omitempty" pn:"CIO"`
					CIOCSIRSRP           *string `json:"CIOCSIRSRP,omitempty" pn:"CIOCSIRSRP"`
					CIOCSIRSRQ           *string `json:"CIOCSIRSRQ,omitempty" pn:"CIOCSIRSRQ"`
					CIOCSISINR           *string `json:"CIOCSISINR,omitempty" pn:"CIOCSISINR"`
					CIOSSBRSRP           *string `json:"CIOSSBRSRP,omitempty" pn:"CIOSSBRSRP"`
					CIOSSBRSRQ           *string `json:"CIOSSBRSRQ,omitempty" pn:"CIOSSBRSRQ"`
					CIOSSBSINR           *string `json:"CIOSSBSINR,omitempty" pn:"CIOSSBSINR"`
					NRCarrierARFCN       *string `json:"NRCarrierARFCN,omitempty" pn:"NRCarrierARFCN"`
					NbrGnbIdLength       *string `json:"NbrGnbIdLength,omitempty" pn:"NbrGnbIdLength"`
					NoHOEnable           *string `json:"NoHOEnable,omitempty" pn:"NoHOEnable"`
					NoRemoveEnable       *string `json:"NoRemoveEnable,omitempty" pn:"NoRemoveEnable"`
					NoXnEnable           *string `json:"NoXnEnable,omitempty" pn:"NoXnEnable"`
					PLMNID               *string `json:"PLMNID,omitempty" pn:"PLMNID"`
					PhyCellID            *string `json:"PhyCellID,omitempty" pn:"PhyCellID"`
					QQualMinOffsetCell   *string `json:"QQualMinOffsetCell,omitempty" pn:"QQualMinOffsetCell"`
					QRxLevMinOffsetCell  *string `json:"QRxLevMinOffsetCell,omitempty" pn:"QRxLevMinOffsetCell"`
					Qoffset              *string `json:"Qoffset,omitempty" pn:"Qoffset"`
					RsType               *string `json:"RsType,omitempty" pn:"RsType"`
					Status               *string `json:"Status,omitempty" pn:"Status"`
					TAC                  *string `json:"TAC,omitempty" pn:"TAC"`
					TargetPLMNID         *string `json:"TargetPLMNID,omitempty" pn:"TargetPLMNID"`
					SsbFrequency         *string `json:"SsbFrequency,omitempty" pn:"SsbFrequency"`
					SsbSubcarrierSpacing *string `json:"SsbSubcarrierSpacing,omitempty" pn:"SsbSubcarrierSpacing"`
				} `json:"NRCell,omitempty" pn:"NRCell"`
			} `json:"NeighborList,omitempty" pn:"NeighborList"`
			NoiseStatus *struct {
				RRUOffset     *string `json:"RRUOffset,omitempty" pn:"RRUOffset"`
				MinRachRssi   *string `json:"MinRachRssi,omitempty" pn:"MinRachRssi"`
				HoMinRachRssi *string `json:"HoMinRachRssi,omitempty" pn:"HoMinRachRssi"`
			} `json:"NoiseStatus,omitempty" pn:"NoiseStatus"`
			Ocng *struct {
				OcngDlCceOccupiedPercent *string `json:"OcngDlCceOccupiedPercent,omitempty" pn:"OcngDlCceOccupiedPercent"`
				OcngDlPrbOccupiedPercent *string `json:"OcngDlPrbOccupiedPercent,omitempty" pn:"OcngDlPrbOccupiedPercent"`
				OcngEnabled              *string `json:"OcngEnabled,omitempty" pn:"OcngEnabled"`
			} `json:"Ocng,omitempty" pn:"Ocng"`
			PCCHConfig *struct {
				DefaultPagingCycle               *string `json:"DefaultPagingCycle,omitempty" pn:"DefaultPagingCycle"`
				N                                *string `json:"N,omitempty" pn:"N"`
				Ns                               *string `json:"Ns,omitempty" pn:"Ns"`
				PagingFrameOffset                *string `json:"PagingFrameOffset,omitempty" pn:"PagingFrameOffset"`
				FirstPdcchMonitoringOccasionOfPo *string `json:"FirstPdcchMonitoringOccasionOfPo,omitempty" pn:"FirstPdcchMonitoringOccasionOfPo"`
			} `json:"PCCHConfig,omitempty" pn:"PCCHConfig"`
			PDCP map[int]*struct {
				QCI          *string `json:"QCI,omitempty" pn:"QCI"`
				DiscardTimer *string `json:"DiscardTimer,omitempty" pn:"DiscardTimer"`
				PdcpSnSizeDL *string `json:"PdcpSnSizeDL,omitempty" pn:"PdcpSnSizeDL"`
				PdcpSnSizeUL *string `json:"PdcpSnSizeUL,omitempty" pn:"PdcpSnSizeUL"`
			} `json:"PDCP,omitempty" pn:"PDCP"`
			PHY *struct {
				DmrsTypeAPosition        *string `json:"DmrsTypeAPosition,omitempty" pn:"DmrsTypeAPosition"`
				SubCarrierSpacingCommon  *string `json:"SubCarrierSpacingCommon,omitempty" pn:"SubCarrierSpacingCommon"`
				TimeAlignmentTimerCommon *string `json:"TimeAlignmentTimerCommon,omitempty" pn:"TimeAlignmentTimerCommon"`
				AdmissionControl         *struct {
					MildCongestionThreshold *string `json:"MildCongestionThreshold,omitempty" pn:"MildCongestionThreshold"`
					NgbrDetectionTime       *string `json:"NgbrDetectionTime,omitempty" pn:"NgbrDetectionTime"`
					PrbAverageDuration      *string `json:"PrbAverageDuration,omitempty" pn:"PrbAverageDuration"`
					PrbUsageGbrThreshold    *string `json:"PrbUsageGbrThreshold,omitempty" pn:"PrbUsageGbrThreshold"`
				} `json:"AdmissionControl,omitempty" pn:"AdmissionControl"`
				Antenna *struct {
					NumOfRxAntenna *string `json:"NumOfRxAntenna,omitempty" pn:"NumOfRxAntenna"`
					NumOfTxAntenna *string `json:"NumOfTxAntenna,omitempty" pn:"NumOfTxAntenna"`
				} `json:"Antenna,omitempty" pn:"Antenna"`
				BWP *struct {
					BWPUL map[int]*struct {
						Bandwidth                       *string `json:"Bandwidth,omitempty" pn:"Bandwidth"`
						ConfiguredGrantConfig           *string `json:"ConfiguredGrantConfig,omitempty" pn:"ConfiguredGrantConfig"`
						CyclicPrefix                    *string `json:"CyclicPrefix,omitempty" pn:"CyclicPrefix"`
						SrsAlpha                        *string `json:"SrsAlpha,omitempty" pn:"SrsAlpha"`
						SrsP0                           *string `json:"SrsP0,omitempty" pn:"SrsP0"`
						SrsPowerControlAdjustmentStates *string `json:"SrsPowerControlAdjustmentStates,omitempty" pn:"SrsPowerControlAdjustmentStates"`
						StartPrbPosition                *string `json:"StartPrbPosition,omitempty" pn:"StartPrbPosition"`
						SubcarrierSpacing               *string `json:"SubcarrierSpacing,omitempty" pn:"SubcarrierSpacing"`
						UlBWPId                         *string `json:"UlBWPId,omitempty" pn:"UlBWPId"`
						PRACH                           *struct {
							MaxHarqMsg3ReTx            *string `json:"MaxHarqMsg3ReTx,omitempty" pn:"MaxHarqMsg3ReTx"`
							NumberOfDedicatedPreambles *string `json:"NumberOfDedicatedPreambles,omitempty" pn:"NumberOfDedicatedPreambles"`
							NumberOfRaPreambles        *string `json:"NumberOfRaPreambles,omitempty" pn:"NumberOfRaPreambles"`
							RaBackoffIndex             *string `json:"RaBackoffIndex,omitempty" pn:"RaBackoffIndex"`
						} `json:"PRACH,omitempty" pn:"PRACH"`
						PUCCH *struct {
							PucchRBNum *string `json:"PucchRBNum,omitempty" pn:"PucchRBNum"`
						} `json:"PUCCH,omitempty" pn:"PUCCH"`
						PUCCHConfig *struct {
							IntraSlotFrequencyHoppingPeriodEnable *string `json:"IntraSlotFrequencyHoppingPeriodEnable,omitempty" pn:"IntraSlotFrequencyHoppingPeriodEnable"`
							PucchTpcEnabled                       *string `json:"PucchTpcEnabled,omitempty" pn:"PucchTpcEnabled"`
						} `json:"PUCCHConfig,omitempty" pn:"PUCCHConfig"`
						PUCCHConfigCommon *struct {
							HoppingId           *string `json:"HoppingId,omitempty" pn:"HoppingId"`
							P0nominal           *string `json:"P0nominal,omitempty" pn:"P0nominal"`
							PucchGroupHopping   *string `json:"PucchGroupHopping,omitempty" pn:"PucchGroupHopping"`
							PucchResourceCommon *string `json:"PucchResourceCommon,omitempty" pn:"PucchResourceCommon"`
						} `json:"PUCCHConfigCommon,omitempty" pn:"PUCCHConfigCommon"`
						PUCCHPowerControl *struct {
							DeltaFPUCCHf0 *string `json:"DeltaFPUCCHf0,omitempty" pn:"DeltaFPUCCHf0"`
							DeltaFPUCCHf1 *string `json:"DeltaFPUCCHf1,omitempty" pn:"DeltaFPUCCHf1"`
							DeltaFPUCCHf2 *string `json:"DeltaFPUCCHf2,omitempty" pn:"DeltaFPUCCHf2"`
							DeltaFPUCCHf3 *string `json:"DeltaFPUCCHf3,omitempty" pn:"DeltaFPUCCHf3"`
							DeltaFPUCCHf4 *string `json:"DeltaFPUCCHf4,omitempty" pn:"DeltaFPUCCHf4"`
							PucchPcSwitch *string `json:"PucchPcSwitch,omitempty" pn:"PucchPcSwitch"`
						} `json:"PUCCHPowerControl,omitempty" pn:"PUCCHPowerControl"`
						PUSCH *struct {
							PuschTpcAccumulation *string `json:"PuschTpcAccumulation,omitempty" pn:"PuschTpcAccumulation"`
							UlAdditionalDmrsPos  *string `json:"UlAdditionalDmrsPos,omitempty" pn:"UlAdditionalDmrsPos"`
							UlDmrsMaxLength      *string `json:"UlDmrsMaxLength,omitempty" pn:"UlDmrsMaxLength"`
							UlTimeAlignmentTimer *string `json:"UlTimeAlignmentTimer,omitempty" pn:"UlTimeAlignmentTimer"`
						} `json:"PUSCH,omitempty" pn:"PUSCH"`
						PUSCHConfig *struct {
							FrequencyHopping                *string `json:"FrequencyHopping,omitempty" pn:"FrequencyHopping"`
							MinUlRbAllowed                  *string `json:"MinUlRbAllowed,omitempty" pn:"MinUlRbAllowed"`
							PuschMaxValidRssi               *string `json:"PuschMaxValidRssi,omitempty" pn:"PuschMaxValidRssi"`
							PuschMaxValidRssiDev            *string `json:"PuschMaxValidRssiDev,omitempty" pn:"PuschMaxValidRssiDev"`
							PuschTpcEnabled                 *string `json:"PuschTpcEnabled,omitempty" pn:"PuschTpcEnabled"`
							ResourceAllocationAdaptiveValue *string `json:"ResourceAllocationAdaptiveValue,omitempty" pn:"ResourceAllocationAdaptiveValue"`
							ResourceAllocationType          *string `json:"ResourceAllocationType,omitempty" pn:"ResourceAllocationType"`
							TpPi2BpskEnable                 *string `json:"TpPi2BpskEnable,omitempty" pn:"TpPi2BpskEnable"`
							TransformPrecoder               *string `json:"TransformPrecoder,omitempty" pn:"TransformPrecoder"`
							TransformPrecoderEnable         *string `json:"TransformPrecoderEnable,omitempty" pn:"TransformPrecoderEnable"`
							TransformPrecoderSwitch         *string `json:"TransformPrecoderSwitch,omitempty" pn:"TransformPrecoderSwitch"`
							TxConfig                        *string `json:"TxConfig,omitempty" pn:"TxConfig"`
							UlFreqSelectScheduling          *string `json:"UlFreqSelectScheduling,omitempty" pn:"UlFreqSelectScheduling"`
							DmrsUplinkForPUSCHMappingTypeA  *struct {
								DmrsAdditionalPosition *string `json:"DmrsAdditionalPosition,omitempty" pn:"DmrsAdditionalPosition"`
								DmrsType               *string `json:"DmrsType,omitempty" pn:"DmrsType"`
							} `json:"DmrsUplinkForPUSCHMappingTypeA,omitempty" pn:"DmrsUplinkForPUSCHMappingTypeA"`
							DmrsUplinkForPUSCHMappingTypeB *struct {
								DmrsAdditionalPosition *string `json:"DmrsAdditionalPosition,omitempty" pn:"DmrsAdditionalPosition"`
								DmrsType               *string `json:"DmrsType,omitempty" pn:"DmrsType"`
							} `json:"DmrsUplinkForPUSCHMappingTypeB,omitempty" pn:"DmrsUplinkForPUSCHMappingTypeB"`
							FrequencyHoppingOffsetLists map[int]*struct {
								FrequencyHoppingOffset *string `json:"FrequencyHoppingOffset,omitempty" pn:"FrequencyHoppingOffset"`
							} `json:"FrequencyHoppingOffsetLists,omitempty" pn:"FrequencyHoppingOffsetLists"`
							UciOnPUSCH *struct {
								BetaOffsetACKIndex1      *string `json:"BetaOffsetACKIndex1,omitempty" pn:"BetaOffsetACKIndex1"`
								BetaOffsetACKIndex2      *string `json:"BetaOffsetACKIndex2,omitempty" pn:"BetaOffsetACKIndex2"`
								BetaOffsetACKIndex3      *string `json:"BetaOffsetACKIndex3,omitempty" pn:"BetaOffsetACKIndex3"`
								BetaOffsetCsiPart1Index1 *string `json:"BetaOffsetCsiPart1Index1,omitempty" pn:"BetaOffsetCsiPart1Index1"`
								BetaOffsetCsiPart1Index2 *string `json:"BetaOffsetCsiPart1Index2,omitempty" pn:"BetaOffsetCsiPart1Index2"`
								BetaOffsetCsiPart2Index1 *string `json:"BetaOffsetCsiPart2Index1,omitempty" pn:"BetaOffsetCsiPart2Index1"`
								BetaOffsetCsiPart2Index2 *string `json:"BetaOffsetCsiPart2Index2,omitempty" pn:"BetaOffsetCsiPart2Index2"`
								BetaOffsetType           *string `json:"BetaOffsetType,omitempty" pn:"BetaOffsetType"`
							} `json:"UciOnPUSCH,omitempty" pn:"UciOnPUSCH"`
							PUSCHPowerControl *struct {
								DeltaMCS    *string `json:"DeltaMCS,omitempty" pn:"DeltaMCS"`
								P0AlphaSets map[int]*struct {
									Alpha *string `json:"Alpha,omitempty" pn:"Alpha"`
								} `json:"P0AlphaSets,omitempty" pn:"P0AlphaSets"`
							} `json:"PUSCHPowerControl,omitempty" pn:"PUSCHPowerControl"`
						} `json:"PUSCHConfig,omitempty" pn:"PUSCHConfig"`
						PUSCHConfigCommon *struct {
							GroupHoppingEnabledTransformPrecoding *string `json:"GroupHoppingEnabledTransformPrecoding,omitempty" pn:"GroupHoppingEnabledTransformPrecoding"`
							Msg3DeltaPreamble                     *string `json:"Msg3DeltaPreamble,omitempty" pn:"Msg3DeltaPreamble"`
							P0NominalWithGrant                    *string `json:"P0NominalWithGrant,omitempty" pn:"P0NominalWithGrant"`
						} `json:"PUSCHConfigCommon,omitempty" pn:"PUSCHConfigCommon"`
						PUSCHTimeDomainAllocationList map[int]*struct {
							K2                   *string `json:"K2,omitempty" pn:"K2"`
							MappingType          *string `json:"MappingType,omitempty" pn:"MappingType"`
							StartSymbolAndLength *string `json:"StartSymbolAndLength,omitempty" pn:"StartSymbolAndLength"`
						} `json:"PUSCHTimeDomainAllocationList,omitempty" pn:"PUSCHTimeDomainAllocationList"`
						RACHConfigCommon *struct {
							CBPreamblesPerSSB           *string `json:"CBPreamblesPerSSB,omitempty" pn:"CBPreamblesPerSSB"`
							MessagePowerOffsetGroupB    *string `json:"MessagePowerOffsetGroupB,omitempty" pn:"MessagePowerOffsetGroupB"`
							Msg1SubcarrierSpacing       *string `json:"Msg1SubcarrierSpacing,omitempty" pn:"Msg1SubcarrierSpacing"`
							Msg3TransformPrecoder       *string `json:"Msg3TransformPrecoder,omitempty" pn:"Msg3TransformPrecoder"`
							NumberOfRAPreamblesGroupA   *string `json:"NumberOfRAPreamblesGroupA,omitempty" pn:"NumberOfRAPreamblesGroupA"`
							RaContentionResolutionTimer *string `json:"RaContentionResolutionTimer,omitempty" pn:"RaContentionResolutionTimer"`
							RaMsg3SizeGroupA            *string `json:"RaMsg3SizeGroupA,omitempty" pn:"RaMsg3SizeGroupA"`
							RestrictedSetConfig         *string `json:"RestrictedSetConfig,omitempty" pn:"RestrictedSetConfig"`
							RootSequenceIndex           *string `json:"RootSequenceIndex,omitempty" pn:"RootSequenceIndex"`
							RootType                    *string `json:"RootType,omitempty" pn:"RootType"`
							RsrpThresholdSSB            *string `json:"RsrpThresholdSSB,omitempty" pn:"RsrpThresholdSSB"`
							SSBPerRACHOccasionChoice    *string `json:"SSBPerRACHOccasionChoice,omitempty" pn:"SSBPerRACHOccasionChoice"`
							TotalNumberOfRAPreambles    *string `json:"TotalNumberOfRAPreambles,omitempty" pn:"TotalNumberOfRAPreambles"`
						} `json:"RACHConfigCommon,omitempty" pn:"RACHConfigCommon"`
						RACHConfigGeneric *struct {
							Msg1FDM                     *string `json:"Msg1FDM,omitempty" pn:"Msg1FDM"`
							Msg1FrequencyStart          *string `json:"Msg1FrequencyStart,omitempty" pn:"Msg1FrequencyStart"`
							PowerRampingStep            *string `json:"PowerRampingStep,omitempty" pn:"PowerRampingStep"`
							PrachConfigurationIndex     *string `json:"PrachConfigurationIndex,omitempty" pn:"PrachConfigurationIndex"`
							PreambleReceivedTargetPower *string `json:"PreambleReceivedTargetPower,omitempty" pn:"PreambleReceivedTargetPower"`
							PreambleTransMax            *string `json:"PreambleTransMax,omitempty" pn:"PreambleTransMax"`
							RaResponseWindow            *string `json:"RaResponseWindow,omitempty" pn:"RaResponseWindow"`
							ZeroCorrelationZoneConfig   *string `json:"ZeroCorrelationZoneConfig,omitempty" pn:"ZeroCorrelationZoneConfig"`
						} `json:"RACHConfigGeneric,omitempty" pn:"RACHConfigGeneric"`
						SrsConfig *struct {
							SrsP0  *string `json:"SrsP0,omitempty" pn:"SrsP0"`
							Config map[int]*struct {
								B_SRS                  *string `json:"B_SRS,omitempty" pn:"B_SRS"`
								B_hop                  *string `json:"B_hop,omitempty" pn:"B_hop"`
								C_SRS                  *string `json:"C_SRS,omitempty" pn:"C_SRS"`
								CycleShift             *string `json:"CycleShift,omitempty" pn:"CycleShift"`
								GroupOrSequenceHopping *string `json:"GroupOrSequenceHopping,omitempty" pn:"GroupOrSequenceHopping"`
								NrofSymbols            *string `json:"NrofSymbols,omitempty" pn:"NrofSymbols"`
								Period                 *string `json:"Period,omitempty" pn:"Period"`
								RepetitionFactor       *string `json:"RepetitionFactor,omitempty" pn:"RepetitionFactor"`
								ResourceType           *string `json:"ResourceType,omitempty" pn:"ResourceType"`
								TransComb              *string `json:"TransComb,omitempty" pn:"TransComb"`
							} `json:"Config,omitempty" pn:"Config"`
						} `json:"SrsConfig,omitempty" pn:"SrsConfig"`
					} `json:"BWPUL,omitempty" pn:"BWPUL"`
					BWPDL map[int]*struct {
						Bandwidth    *string `json:"Bandwidth,omitempty" pn:"Bandwidth"`
						CyclicPrefix *string `json:"CyclicPrefix,omitempty" pn:"CyclicPrefix"`
						DlBWPId      *string `json:"DlBWPId,omitempty" pn:"DlBWPId"`
						PDCCH        *struct {
							DlAdditionalDmrsPos      *string `json:"DlAdditionalDmrsPos,omitempty" pn:"DlAdditionalDmrsPos"`
							DlDmrsMaxLength          *string `json:"DlDmrsMaxLength,omitempty" pn:"DlDmrsMaxLength"`
							CommonControlResourceSet *struct {
								ControlResourceSetId     *string `json:"ControlResourceSetId,omitempty" pn:"ControlResourceSetId"`
								Duration                 *string `json:"Duration,omitempty" pn:"Duration"`
								FrequencyDomainResources *string `json:"FrequencyDomainResources,omitempty" pn:"FrequencyDomainResources"`
								PdcchDMRSScramblingID    *string `json:"PdcchDMRSScramblingID,omitempty" pn:"PdcchDMRSScramblingID"`
								PrecoderGranularity      *string `json:"PrecoderGranularity,omitempty" pn:"PrecoderGranularity"`
								CceRegMappingType        *struct {
									InterleaverSize *string `json:"InterleaverSize,omitempty" pn:"InterleaverSize"`
									RegBundleSize   *string `json:"RegBundleSize,omitempty" pn:"RegBundleSize"`
									RegType         *string `json:"RegType,omitempty" pn:"RegType"`
									ShiftIndex      *string `json:"ShiftIndex,omitempty" pn:"ShiftIndex"`
								} `json:"CceRegMappingType,omitempty" pn:"CceRegMappingType"`
							} `json:"CommonControlResourceSet,omitempty" pn:"CommonControlResourceSet"`
							CommonSearchSpaceList map[int]*struct {
								ControlResourceSetId        *string `json:"ControlResourceSetId,omitempty" pn:"ControlResourceSetId"`
								DciFormat00And10En          *string `json:"DciFormat00And10En,omitempty" pn:"DciFormat00And10En"`
								Duration                    *string `json:"Duration,omitempty" pn:"Duration"`
								MonitoringSlotOffset        *string `json:"MonitoringSlotOffset,omitempty" pn:"MonitoringSlotOffset"`
								MonitoringSlotPeriodicity   *string `json:"MonitoringSlotPeriodicity,omitempty" pn:"MonitoringSlotPeriodicity"`
								MonitoringSymbolsWithinSlot *string `json:"MonitoringSymbolsWithinSlot,omitempty" pn:"MonitoringSymbolsWithinSlot"`
								SearchSpaceId               *string `json:"SearchSpaceId,omitempty" pn:"SearchSpaceId"`
								SearchSpaceType             *string `json:"SearchSpaceType,omitempty" pn:"SearchSpaceType"`
								UeSpecificDciFormats        *string `json:"UeSpecificDciFormats,omitempty" pn:"UeSpecificDciFormats"`
								NrofCandidates              *struct {
									AggregationLevel1  *string `json:"AggregationLevel1,omitempty" pn:"AggregationLevel1"`
									AggregationLevel2  *string `json:"AggregationLevel2,omitempty" pn:"AggregationLevel2"`
									AggregationLevel4  *string `json:"AggregationLevel4,omitempty" pn:"AggregationLevel4"`
									AggregationLevel8  *string `json:"AggregationLevel8,omitempty" pn:"AggregationLevel8"`
									AggregationLevel16 *string `json:"AggregationLevel16,omitempty" pn:"AggregationLevel16"`
								} `json:"NrofCandidates,omitempty" pn:"nrofCandidates"`
							} `json:"CommonSearchSpaceList,omitempty" pn:"CommonSearchSpaceList"`
							ControlResourceSetToAddModList map[int]*struct {
								DciType             *string `json:"DciType,omitempty" pn:"DciType"`
								Duration            *string `json:"Duration,omitempty" pn:"Duration"`
								PrecoderGranularity *string `json:"PrecoderGranularity,omitempty" pn:"PrecoderGranularity"`
								RbNumber            *string `json:"RbNumber,omitempty" pn:"RbNumber"`
								RbStart             *string `json:"RbStart,omitempty" pn:"RbStart"`
							} `json:"ControlResourceSetToAddModList,omitempty" pn:"ControlResourceSetToAddModList"`
							PDCCHConfigCommon *struct {
								ControlResourceSetZero            *string `json:"ControlResourceSetZero,omitempty" pn:"ControlResourceSetZero"`
								PagingSearchSpace                 *string `json:"PagingSearchSpace,omitempty" pn:"PagingSearchSpace"`
								RaSearchSpace                     *string `json:"RaSearchSpace,omitempty" pn:"RaSearchSpace"`
								SearchSpaceOtherSystemInformation *string `json:"SearchSpaceOtherSystemInformation,omitempty" pn:"SearchSpaceOtherSystemInformation"`
								SearchSpaceSIB1                   *string `json:"SearchSpaceSIB1,omitempty" pn:"SearchSpaceSIB1"`
								SearchSpaceZero                   *string `json:"SearchSpaceZero,omitempty" pn:"SearchSpaceZero"`
							} `json:"PDCCHConfigCommon,omitempty" pn:"PDCCHConfigCommon"`
						} `json:"PDCCH,omitempty" pn:"PDCCH"`
						PDSCH *struct {
							ResourceAllocation              *string `json:"ResourceAllocation,omitempty" pn:"ResourceAllocation"`
							ResourceAllocationAdaptiveValue *string `json:"ResourceAllocationAdaptiveValue,omitempty" pn:"ResourceAllocationAdaptiveValue"`
							DownlinkForPDSCHMappingTypeA    *struct {
								DmrsAdditionalPosition *string `json:"DmrsAdditionalPosition,omitempty" pn:"DmrsAdditionalPosition"`
								DmrsType               *string `json:"DmrsType,omitempty" pn:"DmrsType"`
							} `json:"DownlinkForPDSCHMappingTypeA,omitempty" pn:"DownlinkForPDSCHMappingTypeA"`
							DownlinkForPDSCHMappingTypeB *struct {
								DmrsAdditionalPosition *string `json:"DmrsAdditionalPosition,omitempty" pn:"DmrsAdditionalPosition"`
								DmrsType               *string `json:"DmrsType,omitempty" pn:"DmrsType"`
							} `json:"DownlinkForPDSCHMappingTypeB,omitempty" pn:"DownlinkForPDSCHMappingTypeB"`
							PDSCHTimeDomainResourceAllocationList map[int]*struct {
								K0                   *string `json:"K0,omitempty" pn:"K0"`
								MappingType          *string `json:"MappingType,omitempty" pn:"MappingType"`
								StartSymbolAndLength *string `json:"StartSymbolAndLength,omitempty" pn:"StartSymbolAndLength"`
							} `json:"PDSCHTimeDomainResourceAllocationList,omitempty" pn:"PDSCHTimeDomainResourceAllocationList"`
							PrbBundlingType *struct {
								BundlingType    *string `json:"BundlingType,omitempty" pn:"BundlingType"`
								DynamicBundling *struct {
									BundleSizeSet1 *string `json:"BundleSizeSet1,omitempty" pn:"BundleSizeSet1"`
									BundleSizeSet2 *string `json:"BundleSizeSet2,omitempty" pn:"BundleSizeSet2"`
								} `json:"PDSCH,omitempty" pn:"PDSCH"`
								StaticBundling *struct {
									BundleSize *string `json:"BundleSize,omitempty" pn:"BundleSize"`
								} `json:"StaticBundling,omitempty" pn:"StaticBundling"`
							} `json:"PrbBundlingType,omitempty" pn:"PrbBundlingType"`
						} `json:"PDSCH,omitempty" pn:"PDSCH"`
					} `json:"BWPDL,omitempty" pn:"BWPDL"`
				} `json:"BWP,omitempty" pn:"BWP"`
				BwpConfig *struct {
					BwpSwitchType      *string `json:"BwpSwitchType,omitempty" pn:"BwpSwitchType"`
					DLSwitchThreshold  *string `json:"DLSwitchThreshold,omitempty" pn:"DLSwitchThreshold"`
					DefaultActiveBwpId *string `json:"DefaultActiveBwpId,omitempty" pn:"DefaultActiveBwpId"`
					NumOfBwpConfig     *string `json:"NumOfBwpConfig,omitempty" pn:"NumOfBwpConfig"`
					ULSwitchThreshold  *string `json:"ULSwitchThreshold,omitempty" pn:"ULSwitchThreshold"`
				} `json:"BwpConfig,omitempty" pn:"BwpConfig"`
				CSI *struct {
					CSI_RS *struct {
						IsTrsEnable                    *string `json:"IsTrsEnable,omitempty" pn:"IsTrsEnable"`
						NumberOfNzpCsiRsResources      *string `json:"NumberOfNzpCsiRsResources,omitempty" pn:"NumberOfNzpCsiRsResources"`
						ResourcesForChannelMeasurement *struct {
							FirstOFDMSymbolInTimeDomain *string `json:"FirstOFDMSymbolInTimeDomain,omitempty" pn:"FirstOFDMSymbolInTimeDomain"`
							FrequencyDomainAllocation   *string `json:"FrequencyDomainAllocation,omitempty" pn:"FrequencyDomainAllocation"`
							FreqBand                    *struct {
								NrofRBs    *string `json:"NrofRBs,omitempty" pn:"NrofRBs"`
								StartingRB *string `json:"StartingRB,omitempty" pn:"StartingRB"`
							} `json:"FreqBand,omitempty" pn:"FreqBand"`
						} `json:"ResourcesForChannelMeasurement,omitempty" pn:"ResourcesForChannelMeasurement"`
					} `json:"CSI_RS,omitempty" pn:"CSI_RS"`
					CSI_Report *struct {
						InvalidThr                 *string `json:"InvalidThr,omitempty" pn:"InvalidThr"`
						ReportConfigType           *string `json:"ReportConfigType,omitempty" pn:"ReportConfigType"`
						ReportFrequencyGranularity *string `json:"ReportFrequencyGranularity,omitempty" pn:"ReportFrequencyGranularity"`
						ReportPeriodic             *string `json:"ReportPeriodic,omitempty" pn:"ReportPeriodic"`
						ReportQuantity             *string `json:"ReportQuantity,omitempty" pn:"ReportQuantity"`
						SubbandSize                *string `json:"SubbandSize,omitempty" pn:"SubbandSize"`
					} `json:"CSI_Report,omitempty" pn:"CSI_Report"`
					PUCCH_Resource *struct {
						SimultaneousHARQAckCsi *string `json:"SimultaneousHARQAckCsi,omitempty" pn:"SimultaneousHARQAckCsi"`
					} `json:"PUCCH_Resource,omitempty" pn:"PUCCH_Resource"`
				} `json:"CSI,omitempty" pn:"CSI"`
				FrequencyInfoDLSIB *struct {
					OffsetToPointA              *string `json:"OffsetToPointA,omitempty" pn:"OffsetToPointA"`
					AbsoluteFrequencyPointA     *string `json:"AbsoluteFrequencyPointA,omitempty" pn:"AbsoluteFrequencyPointA"`
					Pmax                        *string `json:"Pmax,omitempty" pn:"Pmax"`
					MultiFrequencyBandListNRSIB map[int]*struct {
						FreqBandIndicatorNR *string `json:"FreqBandIndicatorNR,omitempty" pn:"FreqBandIndicatorNR"`
					} `json:"MultiFrequencyBandListNRSIB,omitempty" pn:"MultiFrequencyBandListNRSIB"`
					ScsSpecificCarrierList map[int]*struct {
						SCSSpecificCarrier *struct {
							CarrierBandwidth  *string `json:"CarrierBandwidth,omitempty" pn:"CarrierBandwidth"`
							OffsetToCarrier   *string `json:"OffsetToCarrier,omitempty" pn:"OffsetToCarrier"`
							SubcarrierSpacing *string `json:"SubcarrierSpacing,omitempty" pn:"SubcarrierSpacing"`
						} `json:"SCSSpecificCarrier,omitempty" pn:"SCSSpecificCarrier"`
					} `json:"ScsSpecificCarrierList,omitempty" pn:"ScsSpecificCarrierList"`
				} `json:"FrequencyInfoDLSIB,omitempty" pn:"FrequencyInfoDLSIB"`
				FrequencyInfoULSIB *struct {
					ScsSpecificCarrierList map[int]*struct {
						SCSSpecificCarrier *struct {
							CarrierBandwidth  *string `json:"CarrierBandwidth,omitempty" pn:"CarrierBandwidth"`
							OffsetToCarrier   *string `json:"OffsetToCarrier,omitempty" pn:"OffsetToCarrier"`
							SubcarrierSpacing *string `json:"SubcarrierSpacing,omitempty" pn:"SubcarrierSpacing"`
						} `json:"SCSSpecificCarrier,omitempty" pn:"SCSSpecificCarrier"`
					} `json:"ScsSpecificCarrierList,omitempty" pn:"ScsSpecificCarrierList"`
				} `json:"FrequencyInfoULSIB,omitempty" pn:"FrequencyInfoULSIB"`
				MAC *struct {
					RachUlGrantSizeGroupB *string `json:"RachUlGrantSizeGroupB,omitempty" pn:"RachUlGrantSizeGroupB"`
					Dl                    *struct {
						DlLayers                 *string `json:"DlLayers,omitempty" pn:"DlLayers"`
						DlAntennas               *string `json:"DlAntennas,omitempty" pn:"DlAntennas"`
						DlMaxLayers              *string `json:"DlMaxLayers,omitempty" pn:"DlMaxLayers"`
						DlFreqSelectScheduling   *string `json:"DlFreqSelectScheduling,omitempty" pn:"DlFreqSelectScheduling"`
						MaxNumHarqReTxCcchDl     *string `json:"MaxNumHarqReTxCcchDl,omitempty" pn:"MaxNumHarqReTxCcchDl"`
						MaxNumHarqReTxDcchDtchDl *string `json:"MaxNumHarqReTxDcchDtchDl,omitempty" pn:"MaxNumHarqReTxDcchDtchDl"`
						PdschMcsTable            *string `json:"PdschMcsTable,omitempty" pn:"PdschMcsTable"`
						TargetCceOccupancy       *string `json:"TargetCceOccupancy,omitempty" pn:"TargetCceOccupancy"`
						TargetVrbOccupancy       *string `json:"TargetVrbOccupancy,omitempty" pn:"TargetVrbOccupancy"`
						TrafficGeneratorEnabled  *string `json:"TrafficGeneratorEnabled,omitempty" pn:"TrafficGeneratorEnabled"`
						TtiDuty                  *string `json:"TtiDuty,omitempty" pn:"TtiDuty"`
						PDSCHServingCellConfig   *struct {
							CodeBlockGroupTransmission *struct {
								CodeBlockGroupFlushIndicator        *string `json:"CodeBlockGroupFlushIndicator,omitempty" pn:"CodeBlockGroupFlushIndicator"`
								MaxCodeBlockGroupsPerTransportBlock *string `json:"MaxCodeBlockGroupsPerTransportBlock,omitempty" pn:"MaxCodeBlockGroupsPerTransportBlock"`
							} `json:"CodeBlockGroupTransmission,omitempty" pn:"CodeBlockGroupTransmission"`
						} `json:"PDSCHServingCellConfig,omitempty" pn:"PDSCHServingCellConfig"`
						PowerAllocation *struct {
							DmrsEpreRatioToSsb                 *string `json:"DmrsEpreRatioToSsb,omitempty" pn:"DmrsEpreRatioToSsb"`
							Msg2PdcchEpreRatioToSsb            *string `json:"Msg2PdcchEpreRatioToSsb,omitempty" pn:"Msg2PdcchEpreRatioToSsb"`
							Msg2PdschEpreRatioToSsb            *string `json:"Msg2PdschEpreRatioToSsb,omitempty" pn:"Msg2PdschEpreRatioToSsb"`
							PagingPdcchEpreRatioToSsb          *string `json:"PagingPdcchEpreRatioToSsb,omitempty" pn:"PagingPdcchEpreRatioToSsb"`
							PagingPdschEpreRatioToSsb          *string `json:"PagingPdschEpreRatioToSsb,omitempty" pn:"PagingPdschEpreRatioToSsb"`
							PdcchDynamicPowerAllocationEnabled *string `json:"PdcchDynamicPowerAllocationEnabled,omitempty" pn:"PdcchDynamicPowerAllocationEnabled"`
							PdcchEpreRatioToSsb                *string `json:"PdcchEpreRatioToSsb,omitempty" pn:"PdcchEpreRatioToSsb"`
							PdcchEpreRatioToSsbMaxValue        *string `json:"PdcchEpreRatioToSsbMaxValue,omitempty" pn:"PdcchEpreRatioToSsbMaxValue"`
							PdcchEpreRatioToSsbMinValue        *string `json:"PdcchEpreRatioToSsbMinValue,omitempty" pn:"PdcchEpreRatioToSsbMinValue"`
							PdschDynamicPowerAllocationEnabled *string `json:"PdschDynamicPowerAllocationEnabled,omitempty" pn:"PdschDynamicPowerAllocationEnabled"`
							PdschEpreRatioToSsb                *string `json:"PdschEpreRatioToSsb,omitempty" pn:"PdschEpreRatioToSsb"`
							PdschEpreRatioToSsbMaxValue        *string `json:"PdschEpreRatioToSsbMaxValue,omitempty" pn:"PdschEpreRatioToSsbMaxValue"`
							PdschEpreRatioToSsbMinValue        *string `json:"PdschEpreRatioToSsbMinValue,omitempty" pn:"PdschEpreRatioToSsbMinValue"`
							PowerControlOffsetSS               *string `json:"PowerControlOffsetSS,omitempty" pn:"PowerControlOffsetSS"`
							SibPdcchEpreRatioToSsb             *string `json:"SibPdcchEpreRatioToSsb,omitempty" pn:"SibPdcchEpreRatioToSsb"`
							SsbPowerOffset                     *string `json:"SsbPowerOffset,omitempty" pn:"SsbPowerOffset"`
						} `json:"PowerAllocation,omitempty" pn:"PowerAllocation"`
					} `json:"Dl,omitempty" pn:"Dl"`
					Ul *struct {
						ListNumOfSR           *string `json:"ListNumOfSR,omitempty" pn:"ListNumOfSR"`
						MaxHarqReTx           *string `json:"MaxHarqReTx,omitempty" pn:"MaxHarqReTx"`
						PeriodicBsrTimer      *string `json:"PeriodicBsrTimer,omitempty" pn:"PeriodicBsrTimer"`
						PucchResource0Fmt     *string `json:"PucchResource0Fmt,omitempty" pn:"PucchResource0Fmt"`
						PucchResource1Fmt     *string `json:"PucchResource1Fmt,omitempty" pn:"PucchResource1Fmt"`
						PuschManualFss        *string `json:"PuschManualFss,omitempty" pn:"PuschManualFss"`
						PuschManualFssEndRb   *string `json:"PuschManualFssEndRb,omitempty" pn:"PuschManualFssEndRb"`
						PuschManualFssSlot    *string `json:"PuschManualFssSlot,omitempty" pn:"PuschManualFssSlot"`
						PuschManualFssStartRb *string `json:"PuschManualFssStartRb,omitempty" pn:"PuschManualFssStartRb"`
						PuschMcsTable         *string `json:"PuschMcsTable,omitempty" pn:"PuschMcsTable"`
						RetxBsrTimer          *string `json:"RetxBsrTimer,omitempty" pn:"RetxBsrTimer"`
						TagListNum            *string `json:"TagListNum,omitempty" pn:"TagListNum"`
						UlLayers              *string `json:"UlLayers,omitempty" pn:"UlLayers"`
						UlAntennas            *string `json:"UlAntennas,omitempty" pn:"UlAntennas"`
						UlBoReportMsgLen      *string `json:"UlBoReportMsgLen,omitempty" pn:"UlBoReportMsgLen"`
						TagTimeAlignmentList  map[int]*struct {
							Timer *string `json:"Timer,omitempty" pn:"Timer"`
						} `json:"TagTimeAlignmentList,omitempty" pn:"TagTimeAlignmentList"`
					} `json:"Ul,omitempty" pn:"Ul"`
					Reserved *struct {
						CrntiTimerRelease    *string `json:"CrntiTimerRelease,omitempty" pn:"CrntiTimerRelease"`
						DlGrantNumPerTti     *string `json:"DlGrantNumPerTti,omitempty" pn:"DlGrantNumPerTti"`
						IslandDeltaSinrRange *string `json:"IslandDeltaSinrRange,omitempty" pn:"IslandDeltaSinrRange"`
						NoiseReportSwitch    *string `json:"NoiseReportSwitch,omitempty" pn:"NoiseReportSwitch"`
						PacketSchedMode      *string `json:"PacketSchedMode,omitempty" pn:"PacketSchedMode"`
						RachTargetSinr       *string `json:"RachTargetSinr,omitempty" pn:"RachTargetSinr"`
						UlGrantNumPerTti     *string `json:"UlGrantNumPerTti,omitempty" pn:"UlGrantNumPerTti"`
						UlPacketSchedMode    *string `json:"UlPacketSchedMode,omitempty" pn:"UlPacketSchedMode"`
						Dl                   *struct {
							McsOffsetTableBe *struct {
								Cqi_0  *string `json:"Cqi_0,omitempty" pn:"Cqi_0"`
								Cqi_1  *string `json:"Cqi_1,omitempty" pn:"Cqi_1"`
								Cqi_2  *string `json:"Cqi_2,omitempty" pn:"Cqi_2"`
								Cqi_3  *string `json:"Cqi_3,omitempty" pn:"Cqi_3"`
								Cqi_4  *string `json:"Cqi_4,omitempty" pn:"Cqi_4"`
								Cqi_5  *string `json:"Cqi_5,omitempty" pn:"Cqi_5"`
								Cqi_6  *string `json:"Cqi_6,omitempty" pn:"Cqi_6"`
								Cqi_7  *string `json:"Cqi_7,omitempty" pn:"Cqi_7"`
								Cqi_8  *string `json:"Cqi_8,omitempty" pn:"Cqi_8"`
								Cqi_9  *string `json:"Cqi_9,omitempty" pn:"Cqi_9"`
								Cqi_10 *string `json:"Cqi_10,omitempty" pn:"Cqi_10"`
								Cqi_11 *string `json:"Cqi_11,omitempty" pn:"Cqi_11"`
								Cqi_12 *string `json:"Cqi_12,omitempty" pn:"Cqi_12"`
								Cqi_13 *string `json:"Cqi_13,omitempty" pn:"Cqi_13"`
								Cqi_14 *string `json:"Cqi_14,omitempty" pn:"Cqi_14"`
								Cqi_15 *string `json:"Cqi_15,omitempty" pn:"Cqi_15"`
							} `json:"McsOffsetTableBe,omitempty" pn:"McsOffsetTableBe"`
							McsOffsetTableVoip *struct {
								Cqi_0  *string `json:"Cqi_0,omitempty" pn:"Cqi_0"`
								Cqi_1  *string `json:"Cqi_1,omitempty" pn:"Cqi_1"`
								Cqi_2  *string `json:"Cqi_2,omitempty" pn:"Cqi_2"`
								Cqi_3  *string `json:"Cqi_3,omitempty" pn:"Cqi_3"`
								Cqi_4  *string `json:"Cqi_4,omitempty" pn:"Cqi_4"`
								Cqi_5  *string `json:"Cqi_5,omitempty" pn:"Cqi_5"`
								Cqi_6  *string `json:"Cqi_6,omitempty" pn:"Cqi_6"`
								Cqi_7  *string `json:"Cqi_7,omitempty" pn:"Cqi_7"`
								Cqi_8  *string `json:"Cqi_8,omitempty" pn:"Cqi_8"`
								Cqi_9  *string `json:"Cqi_9,omitempty" pn:"Cqi_9"`
								Cqi_10 *string `json:"Cqi_10,omitempty" pn:"Cqi_10"`
								Cqi_11 *string `json:"Cqi_11,omitempty" pn:"Cqi_11"`
								Cqi_12 *string `json:"Cqi_12,omitempty" pn:"Cqi_12"`
								Cqi_13 *string `json:"Cqi_13,omitempty" pn:"Cqi_13"`
								Cqi_14 *string `json:"Cqi_14,omitempty" pn:"Cqi_14"`
								Cqi_15 *string `json:"Cqi_15,omitempty" pn:"Cqi_15"`
							} `json:"McsOffsetTableVoip,omitempty" pn:"McsOffsetTableVoip"`
						} `json:"Dl,omitempty" pn:"Dl"`
						Ul *struct {
							SinrToMcsTableQos map[int]*struct {
								Mcs_0  *string `json:"Mcs_0,omitempty" pn:"Mcs_0"`
								Mcs_1  *string `json:"Mcs_1,omitempty" pn:"Mcs_1"`
								Mcs_2  *string `json:"Mcs_2,omitempty" pn:"Mcs_2"`
								Mcs_3  *string `json:"Mcs_3,omitempty" pn:"Mcs_3"`
								Mcs_4  *string `json:"Mcs_4,omitempty" pn:"Mcs_4"`
								Mcs_5  *string `json:"Mcs_5,omitempty" pn:"Mcs_5"`
								Mcs_6  *string `json:"Mcs_6,omitempty" pn:"Mcs_6"`
								Mcs_7  *string `json:"Mcs_7,omitempty" pn:"Mcs_7"`
								Mcs_8  *string `json:"Mcs_8,omitempty" pn:"Mcs_8"`
								Mcs_9  *string `json:"Mcs_9,omitempty" pn:"Mcs_9"`
								Mcs_10 *string `json:"Mcs_10,omitempty" pn:"Mcs_10"`
								Mcs_11 *string `json:"Mcs_11,omitempty" pn:"Mcs_11"`
								Mcs_12 *string `json:"Mcs_12,omitempty" pn:"Mcs_12"`
								Mcs_13 *string `json:"Mcs_13,omitempty" pn:"Mcs_13"`
								Mcs_14 *string `json:"Mcs_14,omitempty" pn:"Mcs_14"`
								Mcs_15 *string `json:"Mcs_15,omitempty" pn:"Mcs_15"`
								Mcs_16 *string `json:"Mcs_16,omitempty" pn:"Mcs_16"`
								Mcs_17 *string `json:"Mcs_17,omitempty" pn:"Mcs_17"`
								Mcs_18 *string `json:"Mcs_18,omitempty" pn:"Mcs_18"`
								Mcs_19 *string `json:"Mcs_19,omitempty" pn:"Mcs_19"`
								Mcs_20 *string `json:"Mcs_20,omitempty" pn:"Mcs_20"`
								Mcs_21 *string `json:"Mcs_21,omitempty" pn:"Mcs_21"`
								Mcs_22 *string `json:"Mcs_22,omitempty" pn:"Mcs_22"`
								Mcs_23 *string `json:"Mcs_23,omitempty" pn:"Mcs_23"`
								Mcs_24 *string `json:"Mcs_24,omitempty" pn:"Mcs_24"`
								Mcs_25 *string `json:"Mcs_25,omitempty" pn:"Mcs_25"`
								Mcs_26 *string `json:"Mcs_26,omitempty" pn:"Mcs_26"`
								Mcs_27 *string `json:"Mcs_27,omitempty" pn:"Mcs_27"`
								Mcs_28 *string `json:"Mcs_28,omitempty" pn:"Mcs_28"`
							} `json:"SinrToMcsTableQos,omitempty" pn:"SinrToMcsTableQos"`
						} `json:"Ul,omitempty" pn:"Ul"`
					} `json:"Reserved,omitempty" pn:"Reserved"`
					LinkAdaptation *struct {
						Dl *struct {
							ContinuousLayerRptThr1 *string `json:"ContinuousLayerRptThr1,omitempty" pn:"ContinuousLayerRptThr1"`
							ContinuousLayerRptThr2 *string `json:"ContinuousLayerRptThr2,omitempty" pn:"ContinuousLayerRptThr2"`
							CqiBpCycleKFds         *string `json:"CqiBpCycleKFds,omitempty" pn:"CqiBpCycleKFds"`
							CqiReportPeriodNFds    *string `json:"CqiReportPeriodNFds,omitempty" pn:"CqiReportPeriodNFds"`
							DeltaMcsForRI          *string `json:"DeltaMcsForRI,omitempty" pn:"DeltaMcsForRI"`
							DlBlerWindow           *string `json:"DlBlerWindow,omitempty" pn:"DlBlerWindow"`
							EnterOneLayerBler      *string `json:"EnterOneLayerBler,omitempty" pn:"EnterOneLayerBler"`
							EnterOneLayerMcs       *string `json:"EnterOneLayerMcs,omitempty" pn:"EnterOneLayerMcs"`
							EnterTwoLayerBler      *string `json:"EnterTwoLayerBler,omitempty" pn:"EnterTwoLayerBler"`
							EnterTwoLayerMcs       *string `json:"EnterTwoLayerMcs,omitempty" pn:"EnterTwoLayerMcs"`
							FeedbackErrorThreshold *string `json:"FeedbackErrorThreshold,omitempty" pn:"FeedbackErrorThreshold"`
							InnerAlpha             *string `json:"InnerAlpha,omitempty" pn:"InnerAlpha"`
							MaxMcsIdx              *string `json:"MaxMcsIdx,omitempty" pn:"MaxMcsIdx"`
							McsInitIndex           *string `json:"McsInitIndex,omitempty" pn:"McsInitIndex"`
							MinMcsIdx              *string `json:"MinMcsIdx,omitempty" pn:"MinMcsIdx"`
							NomPdschRsEpreOffset   *string `json:"NomPdschRsEpreOffset,omitempty" pn:"NomPdschRsEpreOffset"`
							RiReportPeriodMFds     *string `json:"RiReportPeriodMFds,omitempty" pn:"RiReportPeriodMFds"`
							TargetBler             *string `json:"TargetBler,omitempty" pn:"TargetBler"`
							ThresholdDeltafBf      *string `json:"ThresholdDeltafBf,omitempty" pn:"ThresholdDeltafBf"`
							ThresholdDeltafCl      *string `json:"ThresholdDeltafCl,omitempty" pn:"ThresholdDeltafCl"`
							ThresholdDeltafOl      *string `json:"ThresholdDeltafOl,omitempty" pn:"ThresholdDeltafOl"`
							TxReassessTimeBf       *string `json:"TxReassessTimeBf,omitempty" pn:"TxReassessTimeBf"`
							TxReassessTimeCl       *string `json:"TxReassessTimeCl,omitempty" pn:"TxReassessTimeCl"`
							TxReassessTimeOl       *string `json:"TxReassessTimeOl,omitempty" pn:"TxReassessTimeOl"`
						} `json:"Dl,omitempty" pn:"Dl"`
						Ul *struct {
							ATBRbNum                   *string `json:"ATBRbNum,omitempty" pn:"ATBRbNum"`
							AlfaPusch                  *string `json:"AlfaPusch,omitempty" pn:"AlfaPusch"`
							ContinuousLayerRptThr1     *string `json:"ContinuousLayerRptThr1,omitempty" pn:"ContinuousLayerRptThr1"`
							ContinuousLayerRptThr2     *string `json:"ContinuousLayerRptThr2,omitempty" pn:"ContinuousLayerRptThr2"`
							EnterATBMcs                *string `json:"EnterATBMcs,omitempty" pn:"EnterATBMcs"`
							EnterOneLayerMcs           *string `json:"EnterOneLayerMcs,omitempty" pn:"EnterOneLayerMcs"`
							EnterTwoLayerMcs           *string `json:"EnterTwoLayerMcs,omitempty" pn:"EnterTwoLayerMcs"`
							LeaveATBMcs                *string `json:"LeaveATBMcs,omitempty" pn:"LeaveATBMcs"`
							MaxMcsIdx                  *string `json:"MaxMcsIdx,omitempty" pn:"MaxMcsIdx"`
							MinMcsIdx                  *string `json:"MinMcsIdx,omitempty" pn:"MinMcsIdx"`
							Msg3McsGroupA              *string `json:"Msg3McsGroupA,omitempty" pn:"Msg3McsGroupA"`
							Msg3McsGroupB              *string `json:"Msg3McsGroupB,omitempty" pn:"Msg3McsGroupB"`
							Msg3McsHO                  *string `json:"Msg3McsHO,omitempty" pn:"Msg3McsHO"`
							OlrcDeltaBad               *string `json:"OlrcDeltaBad,omitempty" pn:"OlrcDeltaBad"`
							OlrcDeltaBadMax            *string `json:"OlrcDeltaBadMax,omitempty" pn:"OlrcDeltaBadMax"`
							OlrcDeltaGood              *string `json:"OlrcDeltaGood,omitempty" pn:"OlrcDeltaGood"`
							OlrcDeltaGoodMax           *string `json:"OlrcDeltaGoodMax,omitempty" pn:"OlrcDeltaGoodMax"`
							OlrcDeltaUeConnect         *string `json:"OlrcDeltaUeConnect,omitempty" pn:"OlrcDeltaUeConnect"`
							OlrcEnabled                *string `json:"OlrcEnabled,omitempty" pn:"OlrcEnabled"`
							OlrcMaxBadInActivityWindow *string `json:"OlrcMaxBadInActivityWindow,omitempty" pn:"OlrcMaxBadInActivityWindow"`
							SrPeriod                   *string `json:"SrPeriod,omitempty" pn:"SrPeriod"`
							TargetBler                 *string `json:"TargetBler,omitempty" pn:"TargetBler"`
							DlReflectUlLayerSwitch     *string `json:"DlReflectUlLayerSwitch,omitempty" pn:"DlReflectUlLayerSwitch"`
						} `json:"Ul,omitempty" pn:"Ul"`
						Reserved *struct {
							CsiInvalidRlfTimes *string `json:"CsiInvalidRlfTimes,omitempty" pn:"CsiInvalidRlfTimes"`
							Dci0LostRlfTimes   *string `json:"Dci0LostRlfTimes,omitempty" pn:"Dci0LostRlfTimes"`
							DtxRlfTimes        *string `json:"DtxRlfTimes,omitempty" pn:"DtxRlfTimes"`
							McsForDbch         *string `json:"McsForDbch,omitempty" pn:"McsForDbch"`
							McsForPaging       *string `json:"McsForPaging,omitempty" pn:"McsForPaging"`
							McsForRar          *string `json:"McsForRar,omitempty" pn:"McsForRar"`
						} `json:"Reserved,omitempty" pn:"Reserved"`
					} `json:"LinkAdaptation,omitempty" pn:"LinkAdaptation"`
				} `json:"MAC,omitempty" pn:"MAC"`
				PDCCH *struct {
					TA_ALFA               *string `json:"TA_ALFA,omitempty" pn:"TA_ALFA"`
					TA_ENABLE             *string `json:"TA_ENABLE,omitempty" pn:"TA_ENABLE"`
					TA_FILTER             *string `json:"TA_FILTER,omitempty" pn:"TA_FILTER"`
					TA_UPDATE             *string `json:"TA_UPDATE,omitempty" pn:"TA_UPDATE"`
					TA_NUM_AVG            *string `json:"TA_NUM_AVG,omitempty" pn:"TA_NUM_AVG"`
					TA_OUT_SYNC_RANGE     *string `json:"TA_OUT_SYNC_RANGE,omitempty" pn:"TA_OUT_SYNC_RANGE"`
					TA_OUT_SYNC_TIMER     *string `json:"TA_OUT_SYNC_TIMER,omitempty" pn:"TA_OUT_SYNC_TIMER"`
					CommonSearchSpaceList map[int]*struct {
						ControlResourceSetId *string `json:"ControlResourceSetId,omitempty" pn:"ControlResourceSetId"`
						SearchSpaceId        *string `json:"SearchSpaceId,omitempty" pn:"SearchSpaceId"`
						NrofCandidates       *struct {
							AggregationLevel1  *string `json:"AggregationLevel1,omitempty" pn:"AggregationLevel1"`
							AggregationLevel2  *string `json:"AggregationLevel2,omitempty" pn:"AggregationLevel2"`
							AggregationLevel4  *string `json:"AggregationLevel4,omitempty" pn:"AggregationLevel4"`
							AggregationLevel8  *string `json:"AggregationLevel8,omitempty" pn:"AggregationLevel8"`
							AggregationLevel16 *string `json:"AggregationLevel16,omitempty" pn:"AggregationLevel16"`
						} `json:"NrofCandidates,omitempty" pn:"nrofCandidates"`
					} `json:"CommonSearchSpaceList,omitempty" pn:"CommonSearchSpaceList"`
				} `json:"PDCCH,omitempty" pn:"PDCCH"`
				PDCCHConfig *struct {
					CssNoAdapAggregationLevel           *string `json:"CssNoAdapAggregationLevel,omitempty" pn:"CssNoAdapAggregationLevel"`
					PdcchAggregationLevelAdaptionEnable *string `json:"PdcchAggregationLevelAdaptionEnable,omitempty" pn:"PdcchAggregationLevelAdaptionEnable"`
					UssNoAdapAggregationLevel           *string `json:"UssNoAdapAggregationLevel,omitempty" pn:"UssNoAdapAggregationLevel"`
					ControlResourceSetToAddModList      map[int]*struct {
						LimitedUEsInDefaultDedicatedSearchSpace *string `json:"LimitedUEsInDefaultDedicatedSearchSpace,omitempty" pn:"LimitedUEsInDefaultDedicatedSearchSpace"`
					} `json:"ControlResourceSetToAddModList,omitempty" pn:"ControlResourceSetToAddModList"`
					SearchSpacesToAddModList map[int]*struct {
						CQI_0                       *string `json:"CQI_0,omitempty" pn:"CQI_0"`
						CQI_1                       *string `json:"CQI_1,omitempty" pn:"CQI_1"`
						CQI_2                       *string `json:"CQI_2,omitempty" pn:"CQI_2"`
						CQI_3                       *string `json:"CQI_3,omitempty" pn:"CQI_3"`
						CQI_4                       *string `json:"CQI_4,omitempty" pn:"CQI_4"`
						CQI_5                       *string `json:"CQI_5,omitempty" pn:"CQI_5"`
						CQI_6                       *string `json:"CQI_6,omitempty" pn:"CQI_6"`
						CQI_7                       *string `json:"CQI_7,omitempty" pn:"CQI_7"`
						CQI_8                       *string `json:"CQI_8,omitempty" pn:"CQI_8"`
						CQI_9                       *string `json:"CQI_9,omitempty" pn:"CQI_9"`
						CQI_10                      *string `json:"CQI_10,omitempty" pn:"CQI_10"`
						CQI_11                      *string `json:"CQI_11,omitempty" pn:"CQI_11"`
						CQI_12                      *string `json:"CQI_12,omitempty" pn:"CQI_12"`
						CQI_13                      *string `json:"CQI_13,omitempty" pn:"CQI_13"`
						Cqi_14                      *string `json:"Cqi_14,omitempty" pn:"Cqi_14"`
						CQI_15                      *string `json:"CQI_15,omitempty" pn:"CQI_15"`
						MonitoringSymbolsWithinSlot *string `json:"MonitoringSymbolsWithinSlot,omitempty" pn:"MonitoringSymbolsWithinSlot"`
						NrofCandidates              *struct {
							AggregationLevel1  *string `json:"AggregationLevel1,omitempty" pn:"AggregationLevel1"`
							AggregationLevel2  *string `json:"AggregationLevel2,omitempty" pn:"AggregationLevel2"`
							AggregationLevel4  *string `json:"AggregationLevel4,omitempty" pn:"AggregationLevel4"`
							AggregationLevel8  *string `json:"AggregationLevel8,omitempty" pn:"AggregationLevel8"`
							AggregationLevel16 *string `json:"AggregationLevel16,omitempty" pn:"AggregationLevel16"`
						} `json:"NrofCandidates,omitempty" pn:"nrofCandidates"`
					} `json:"SearchSpacesToAddModList,omitempty" pn:"SearchSpacesToAddModList"`
				} `json:"PDCCHConfig,omitempty" pn:"PDCCHConfig"`
				PDSCH *struct {
					DmrsDlLaEnable *string `json:"DmrsDlLaEnable,omitempty" pn:"DmrsDlLaEnable"`
					MinDlRbAllowed *string `json:"MinDlRbAllowed,omitempty" pn:"MinDlRbAllowed"`
				} `json:"PDSCH,omitempty" pn:"PDSCH"`
				PUCCHConfig *struct {
					PucchMaxRssiProtection  *string `json:"PucchMaxRssiProtection,omitempty" pn:"PucchMaxRssiProtection"`
					PucchMaxValidRssi       *string `json:"PucchMaxValidRssi,omitempty" pn:"PucchMaxValidRssi"`
					PucchTpcTargetSinrHigh  *string `json:"PucchTpcTargetSinrHigh,omitempty" pn:"PucchTpcTargetSinrHigh"`
					PucchTpcTargetSinrLow   *string `json:"PucchTpcTargetSinrLow,omitempty" pn:"PucchTpcTargetSinrLow"`
					PucchValidRssiDeviation *string `json:"PucchValidRssiDeviation,omitempty" pn:"PucchValidRssiDeviation"`
					PucchValidSinrCSI       *string `json:"PucchValidSinrCSI,omitempty" pn:"PucchValidSinrCSI"`
					PucchValidSinrHARQandSR *string `json:"PucchValidSinrHARQandSR,omitempty" pn:"PucchValidSinrHARQandSR"`
				} `json:"PUCCHConfig,omitempty" pn:"PUCCHConfig"`
				PUSCHConfig *struct {
					DmrsUlLaEnable      *string `json:"DmrsUlLaEnable,omitempty" pn:"DmrsUlLaEnable"`
					MaxPhrValue         *string `json:"MaxPhrValue,omitempty" pn:"MaxPhrValue"`
					MaxPuschTpcAccVal   *string `json:"MaxPuschTpcAccVal,omitempty" pn:"MaxPuschTpcAccVal"`
					MinPhrValue         *string `json:"MinPhrValue,omitempty" pn:"MinPhrValue"`
					MinPuschTpcAccVal   *string `json:"MinPuschTpcAccVal,omitempty" pn:"MinPuschTpcAccVal"`
					PuschTargetSinr     *string `json:"PuschTargetSinr,omitempty" pn:"PuschTargetSinr"`
					PuschTargetSinrDev  *string `json:"PuschTargetSinrDev,omitempty" pn:"PuschTargetSinrDev"`
					PuschTpcForbitTimer *string `json:"PuschTpcForbitTimer,omitempty" pn:"PuschTpcForbitTimer"`
				} `json:"PUSCHConfig,omitempty" pn:"PUSCHConfig"`
				PUSCHConfigCommon *struct {
					K1K2Slot1Enable *string `json:"K1K2Slot1Enable,omitempty" pn:"K1K2Slot1Enable"`
				} `json:"PUSCHConfigCommon,omitempty" pn:"PUSCHConfigCommon"`
				PhysicalCellGroupConfig *struct {
					PdschHarqAckCodebook *string `json:"PdschHarqAckCodebook,omitempty" pn:"PdschHarqAckCodebook"`
					PNrFr1               *string `json:"PNrFr1,omitempty" pn:"PNrFr1"`
				} `json:"PhysicalCellGroupConfig,omitempty" pn:"PhysicalCellGroupConfig"`
				PowerControl *struct {
					PeriodicPhrTimer *string `json:"PeriodicPhrTimer,omitempty" pn:"PeriodicPhrTimer"`
					PhrEnable        *string `json:"PhrEnable,omitempty" pn:"PhrEnable"`
					PhrModeOtherCg   *string `json:"PhrModeOtherCg,omitempty" pn:"PhrModeOtherCg"`
					PhrTxPowerFactor *string `json:"PhrTxPowerFactor,omitempty" pn:"PhrTxPowerFactor"`
					ProhibitPhrTimer *string `json:"ProhibitPhrTimer,omitempty" pn:"ProhibitPhrTimer"`
				} `json:"PowerControl,omitempty" pn:"PowerControl"`
				RadioBearerType *struct {
					Srb1 *struct {
						MacConfig *struct {
							MaxNumHarqReTxSRB12 *string `json:"MaxNumHarqReTxSRB12,omitempty" pn:"MaxNumHarqReTxSRB12"`
						} `json:"MacConfig,omitempty" pn:"MacConfig"`
					} `json:"Srb1,omitempty" pn:"Srb1"`
				} `json:"RadioBearerType,omitempty" pn:"RadioBearerType"`
				SIB1 *struct {
					ECallOverIMSSupport *string `json:"ECallOverIMSSupport,omitempty" pn:"ECallOverIMSSupport"`
					ImsEmergencySupport *string `json:"ImsEmergencySupport,omitempty" pn:"ImsEmergencySupport"`
					UseFullResumeID     *string `json:"UseFullResumeID,omitempty" pn:"UseFullResumeID"`
					CellSelectionInfo   *struct {
						QRxLevMinSUL *string `json:"QRxLevMinSUL,omitempty" pn:"QRxLevMinSUL"`
					} `json:"CellSelectionInfo,omitempty" pn:"CellSelectionInfo"`
					ConnEstFailureControl *struct {
						ConnEstFailCount          *string `json:"ConnEstFailCount,omitempty" pn:"ConnEstFailCount"`
						ConnEstFailOffset         *string `json:"ConnEstFailOffset,omitempty" pn:"ConnEstFailOffset"`
						ConnEstFailOffsetValidity *string `json:"ConnEstFailOffsetValidity,omitempty" pn:"ConnEstFailOffsetValidity"`
					} `json:"ConnEstFailureControl,omitempty" pn:"ConnEstFailureControl"`
					ServingCellConfigCommon *struct {
						NTimingAdvanceOffset      *string `json:"NTimingAdvanceOffset,omitempty" pn:"NTimingAdvanceOffset"`
						SsPBCHBlockPower          *string `json:"SsPBCHBlockPower,omitempty" pn:"SsPBCHBlockPower"`
						SsbPeriodicityServingCell *string `json:"SsbPeriodicityServingCell,omitempty" pn:"SsbPeriodicityServingCell"`
						SsbPositionsInBurst       *struct {
							GroupPresence *string `json:"GroupPresence,omitempty" pn:"GroupPresence"`
							InOneGroup    *string `json:"InOneGroup,omitempty" pn:"InOneGroup"`
						} `json:"SsbPositionsInBurst,omitempty" pn:"SsbPositionsInBurst"`
					} `json:"ServingCellConfigCommon,omitempty" pn:"ServingCellConfigCommon"`
					SiSchedulingInfo *struct {
						SiRequestPeriod         *string `json:"SiRequestPeriod,omitempty" pn:"SiRequestPeriod"`
						SystemInformationAreaID *string `json:"SystemInformationAreaID,omitempty" pn:"SystemInformationAreaID"`
						RachOccasionsSI         *struct {
							SsbperRACHOccasion *string `json:"SsbperRACHOccasion,omitempty" pn:"SsbperRACHOccasion"`
						} `json:"RachOccasionsSI,omitempty" pn:"RachOccasionsSI"`
						SiRequestConfig *struct {
							SiRequestPeriod *string `json:"SiRequestPeriod,omitempty" pn:"SiRequestPeriod"`
							RachOccasionsSI *struct {
								SsbperRACHOccasion *string `json:"SsbperRACHOccasion,omitempty" pn:"SsbperRACHOccasion"`
							} `json:"RachOccasionsSI,omitempty" pn:"RachOccasionsSI"`
						} `json:"SiRequestConfig,omitempty" pn:"SiRequestConfig"`
						SiRequestConfigSUL *struct {
							SiRequestPeriod *string `json:"SiRequestPeriod,omitempty" pn:"SiRequestPeriod"`
							RachOccasionsSI *struct {
								SsbperRACHOccasion *string `json:"SsbperRACHOccasion,omitempty" pn:"SsbperRACHOccasion"`
							} `json:"RachOccasionsSI,omitempty" pn:"RachOccasionsSI"`
						} `json:"SiRequestConfigSUL,omitempty" pn:"SiRequestConfigSUL"`
					} `json:"SiSchedulingInfo,omitempty" pn:"SiSchedulingInfo"`
					UacBarringInfo *struct {
						UacAccessCategory1SelectionAssistanceInfo *struct {
							PlmnCommon *string `json:"PlmnCommon,omitempty" pn:"PlmnCommon"`
						} `json:"UacAccessCategory1SelectionAssistanceInfo,omitempty" pn:"UacAccessCategory1SelectionAssistanceInfo"`
					} `json:"UacBarringInfo,omitempty" pn:"UacBarringInfo"`
				} `json:"SIB1,omitempty" pn:"SIB1"`
				SSB *struct {
					SsPBCHBlockPower          *string `json:"SsPBCHBlockPower,omitempty" pn:"SsPBCHBlockPower"`
					SsbPeriodicityServingCell *string `json:"SsbPeriodicityServingCell,omitempty" pn:"SsbPeriodicityServingCell"`
					SsbSubcarrierOffset       *string `json:"SsbSubcarrierOffset,omitempty" pn:"SsbSubcarrierOffset"`
					SsbPositionsInBurst       *struct {
						InOneGroup *string `json:"InOneGroup,omitempty" pn:"InOneGroup"`
					} `json:"SsbPositionsInBurst,omitempty" pn:"SsbPositionsInBurst"`
				} `json:"SSB,omitempty" pn:"SSB"`
				SiConfig *struct {
					Enable                       *string `json:"Enable,omitempty" pn:"Enable"`
					SystemInformationAreaId      *string `json:"SystemInformationAreaId,omitempty" pn:"SystemInformationAreaId"`
					WindowLength                 *string `json:"WindowLength,omitempty" pn:"WindowLength"`
					SchedulingInfoAllocationList map[int]*struct {
						BroadcastStatus *string `json:"BroadcastStatus,omitempty" pn:"BroadcastStatus"`
						OnDemandTimes   *string `json:"OnDemandTimes,omitempty" pn:"OnDemandTimes"`
						Periodicity     *string `json:"Periodicity,omitempty" pn:"Periodicity"`
						SibMap          *string `json:"SibMap,omitempty" pn:"SibMap"`
					} `json:"SchedulingInfoAllocationList,omitempty" pn:"SchedulingInfoAllocationList"`
				} `json:"SiConfig,omitempty" pn:"SiConfig"`
				SrsConfig *struct {
					SrsAutoConfig        *string `json:"SrsAutoConfig,omitempty" pn:"SrsAutoConfig"`
					SrsEnable            *string `json:"SrsEnable,omitempty" pn:"SrsEnable"`
					UeNumUsingSrsConfig1 *string `json:"UeNumUsingSrsConfig1,omitempty" pn:"UeNumUsingSrsConfig1"`
				} `json:"SrsConfig,omitempty" pn:"SrsConfig"`
				SsbConfig *struct {
					PdcchConfigSIB1       *string `json:"PdcchConfigSIB1,omitempty" pn:"PdcchConfigSIB1"`
					SsbPositionsInBurstMB *string `json:"SsbPositionsInBurstMB,omitempty" pn:"SsbPositionsInBurstMB"`
					SsbPower              *string `json:"SsbPower,omitempty" pn:"SsbPower"`
				} `json:"SsbConfig,omitempty" pn:"SsbConfig"`
				SubframeConfig *struct {
					SubframeShutdownEnable *string `json:"SubframeShutdownEnable,omitempty" pn:"SubframeShutdownEnable"`
				} `json:"SubframeConfig,omitempty" pn:"SubframeConfig"`
				TddULDLConfigurationCommon *struct {
					IsPattern2Present          *string `json:"IsPattern2Present,omitempty" pn:"IsPattern2Present"`
					ReferenceSubcarrierSpacing *string `json:"ReferenceSubcarrierSpacing,omitempty" pn:"ReferenceSubcarrierSpacing"`
					Pattern1                   *struct {
						DlULTransmissionPeriodicity *string `json:"DlULTransmissionPeriodicity,omitempty" pn:"DlULTransmissionPeriodicity"`
						NrofDownlinkSlots           *string `json:"NrofDownlinkSlots,omitempty" pn:"NrofDownlinkSlots"`
						NrofDownlinkSymbols         *string `json:"NrofDownlinkSymbols,omitempty" pn:"NrofDownlinkSymbols"`
						NrofUplinkSlots             *string `json:"NrofUplinkSlots,omitempty" pn:"NrofUplinkSlots"`
						NrofUplinkSymbols           *string `json:"NrofUplinkSymbols,omitempty" pn:"NrofUplinkSymbols"`
					} `json:"Pattern1,omitempty" pn:"pattern1"`
					Pattern2 *struct {
						DlULTransmissionPeriodicity *string `json:"DlULTransmissionPeriodicity,omitempty" pn:"DlULTransmissionPeriodicity"`
						NrofDownlinkSlots           *string `json:"NrofDownlinkSlots,omitempty" pn:"NrofDownlinkSlots"`
						NrofDownlinkSymbols         *string `json:"NrofDownlinkSymbols,omitempty" pn:"NrofDownlinkSymbols"`
						NrofUplinkSlots             *string `json:"NrofUplinkSlots,omitempty" pn:"NrofUplinkSlots"`
						NrofUplinkSymbols           *string `json:"NrofUplinkSymbols,omitempty" pn:"NrofUplinkSymbols"`
					} `json:"Pattern2,omitempty" pn:"pattern2"`
				} `json:"TddULDLConfigurationCommon,omitempty" pn:"TddULDLConfigurationCommon"`
			} `json:"PHY,omitempty" pn:"PHY"`
			RF *struct {
				DLBandwidth         *string `json:"DLBandwidth,omitempty" pn:"DLBandwidth"`
				ForgetFactorRbNoise *string `json:"ForgetFactorRbNoise,omitempty" pn:"ForgetFactorRbNoise"`
				FreqBandIndicator   *string `json:"FreqBandIndicator,omitempty" pn:"FreqBandIndicator"`
				NRARFCNDL           *string `json:"NRARFCNDL,omitempty" pn:"NRARFCNDL"`
				NRARFCNUL           *string `json:"NRARFCNUL,omitempty" pn:"NRARFCNUL"`
				PSCHPowerOffset     *string `json:"PSCHPowerOffset,omitempty" pn:"PSCHPowerOffset"`
				PhyCellID           *string `json:"PhyCellID,omitempty" pn:"PhyCellID"`
				SsPBCHBlockPower    *string `json:"SsPBCHBlockPower,omitempty" pn:"SsPBCHBlockPower"`
				SsbFrequency        *string `json:"SsbFrequency,omitempty" pn:"SsbFrequency"`
				ULBandwidth         *string `json:"ULBandwidth,omitempty" pn:"ULBandwidth"`
				GNBDUId             *string `json:"GNBDUId,omitempty" pn:"GNBDUId"`
			} `json:"RF,omitempty" pn:"RF"`
			RLC map[int]*struct {
				RlcMode               *string `json:"RlcMode,omitempty" pn:"RlcMode"`
				DlRlcSnSize           *string `json:"DlRlcSnSize,omitempty" pn:"DlRlcSnSize"`
				UlRlcSnSize           *string `json:"UlRlcSnSize,omitempty" pn:"UlRlcSnSize"`
				GNBRlcReassemblyTimer *string `json:"GNBRlcReassemblyTimer,omitempty" pn:"GNBRlcReassemblyTimer"`
				QCI                   *string `json:"QCI,omitempty" pn:"QCI"`
			} `json:"RLC,omitempty" pn:"RLC"`
			RRCTimers *struct {
				N310               *string `json:"N310,omitempty" pn:"N310"`
				N311               *string `json:"N311,omitempty" pn:"N311"`
				T300               *string `json:"T300,omitempty" pn:"T300"`
				T301               *string `json:"T301,omitempty" pn:"T301"`
				T302               *string `json:"T302,omitempty" pn:"T302"`
				T304               *string `json:"T304,omitempty" pn:"T304"`
				T310               *string `json:"T310,omitempty" pn:"T310"`
				T311               *string `json:"T311,omitempty" pn:"T311"`
				T319               *string `json:"T319,omitempty" pn:"T319"`
				T320               *string `json:"T320,omitempty" pn:"T320"`
				T380               *string `json:"T380,omitempty" pn:"T380"`
				RrcMsgRetries      *string `json:"RrcMsgRetries,omitempty" pn:"RrcMsgRetries"`
				RrcRetransmitTimer *string `json:"RrcRetransmitTimer,omitempty" pn:"RrcRetransmitTimer"`
			} `json:"RRCTimers,omitempty" pn:"RRCTimers"`
			RadioBearParam map[int]*struct {
				FiveQI          *string `json:"FiveQI,omitempty" pn:"FiveQI"`
				MappingDrbIndex *string `json:"MappingDrbIndex,omitempty" pn:"MappingDrbIndex"`
			} `json:"RadioBearParam,omitempty" pn:"RadioBearParam"`
			IntegrityProtAlgorithm *struct {
				Nia0Enable *string `json:"Nia0Enable,omitempty" pn:"Nia0Enable"`
				Nia1Enable *string `json:"Nia1Enable,omitempty" pn:"Nia1Enable"`
				Nia2Enable *string `json:"Nia2Enable,omitempty" pn:"Nia2Enable"`
				Nia3Enable *string `json:"Nia3Enable,omitempty" pn:"Nia3Enable"`
			} `json:"IntegrityProtAlgorithm,omitempty" pn:"IntegrityProtAlgorithm"`
			SysInfoCtrlParam *struct {
				MultiBandInfoListSIB1 *string `json:"MultiBandInfoListSIB1,omitempty" pn:"MultiBandInfoListSIB1"`
				IdleModeRATMobile     map[int]*struct {
					MultiBandInfoListSIB5 *string `json:"MultiBandInfoListSIB5,omitempty" pn:"MultiBandInfoListSIB5"`
				} `json:"IdleModeRATMobile,omitempty" pn:"IdleModeRATMobile"`
			} `json:"SysInfoCtrlParam,omitempty" pn:"SysInfoCtrlParam"`
			VoNR *struct {
				IMSAdapt        *string `json:"IMSAdapt,omitempty" pn:"IMSAdapt"`
				SPSSwitchQCI1Ul *string `json:"SPSSwitchQCI1Ul,omitempty" pn:"SPSSwitchQCI1Ul"`
				VoNRSwitch      *string `json:"VoNRSwitch,omitempty" pn:"VoNRSwitch"`
				PdcpInitParam   map[int]*struct {
					RohcEn *string `json:"RohcEn,omitempty" pn:"RohcEn"`
				} `json:"PdcpInitParam,omitempty" pn:"PdcpInitParam"`
				VoNRList map[int]*struct {
					PLMNID    *string `json:"PLMNID,omitempty" pn:"PLMNID"`
					Vo5qi     *string `json:"Vo5qi,omitempty" pn:"Vo5qi"`
					VoSupport *string `json:"VoSupport,omitempty" pn:"VoSupport"`
				} `json:"VoNRList,omitempty" pn:"VoNRList"`
			} `json:"VoNR,omitempty" pn:"VoNR"`
		} `json:"RAN,omitempty" pn:"RAN"`
		NSA *struct {
			QoS map[int]*struct {
				Enable               *string `json:"Enable,omitempty" pn:"Enable"`
				DlMaxNumHarqReTx     *string `json:"DlMaxNumHarqReTx,omitempty" pn:"DlMaxNumHarqReTx"`
				DscpMarkingToUse     *string `json:"DscpMarkingToUse,omitempty" pn:"DscpMarkingToUse"`
				OutOfOrderDelivery   *string `json:"OutOfOrderDelivery,omitempty" pn:"OutOfOrderDelivery"`
				PacketDelayBudget    *string `json:"PacketDelayBudget,omitempty" pn:"PacketDelayBudget"`
				PdcpSnSizeDl         *string `json:"PdcpSnSizeDl,omitempty" pn:"PdcpSnSizeDl"`
				PdcpSnSizeUl         *string `json:"PdcpSnSizeUl,omitempty" pn:"PdcpSnSizeUl"`
				PrimaryCellGroupId   *string `json:"PrimaryCellGroupId,omitempty" pn:"PrimaryCellGroupId"`
				Priority             *string `json:"Priority,omitempty" pn:"Priority"`
				QCI                  *string `json:"QCI,omitempty" pn:"QCI"`
				QciToUse             *string `json:"QciToUse,omitempty" pn:"QciToUse"`
				RejectEnable         *string `json:"RejectEnable,omitempty" pn:"RejectEnable"`
				Treordering          *string `json:"Treordering,omitempty" pn:"Treordering"`
				Type                 *string `json:"Type,omitempty" pn:"Type"`
				UlDataSplitThreshold *string `json:"UlDataSplitThreshold,omitempty" pn:"UlDataSplitThreshold"`
			} `json:"QoS,omitempty" pn:"QoS"`
		} `json:"NSA,omitempty" pn:"NSA"`
		NGC *struct {
			QosHeavyControl       *string `json:"QosHeavyControl,omitempty" pn:"QosHeavyControl"`
			QoSNumberOfEntries    *string `json:"QoSNumberOfEntries,omitempty" pn:"QoSNumberOfEntries"`
			FiveQINumberOfEntries *string `json:"FiveQINumberOfEntries,omitempty" pn:"FiveQINumberOfEntries"`
			MaxQoSEntries         *string `json:"MaxQoSEntries,omitempty" pn:"MaxQoSEntries"`
			Max5QIEntries         *string `json:"Max5QIEntries,omitempty" pn:"Max5QIEntries"`
			QoS                   map[int]*struct {
				CommonSliceShareHeavy *string `json:"CommonSliceShareHeavy,omitempty" pn:"CommonSliceShareHeavy"`
				FiveQI                *string `json:"FiveQI,omitempty" pn:"FiveQI"`
				PLMNID                *string `json:"PLMNID,omitempty" pn:"PLMNID"`
				SNSSAI                *string `json:"SNSSAI,omitempty" pn:"SNSSAI"`
				Max5QIEntries         *string `json:"Max5QIEntries,omitempty" pn:"Max5QIEntries"`
				MaxQoSEntries         *string `json:"MaxQoSEntries,omitempty" pn:"MaxQoSEntries"`
				QosHeavy              *string `json:"QosHeavy,omitempty" pn:"QosHeavy"`
				Frequency             *string `json:"Frequency,omitempty" pn:"Frequency"`
				SsbSubSpacing         *string `json:"SsbSubSpacing,omitempty" pn:"SsbSubSpacing"`
				DLMaxAccumGFBRPerDrb  *string `json:"DLMaxAccumGFBRPerDrb,omitempty" pn:"DLMaxAccumGFBRPerDrb"`
				ULMaxAccumGFBRPerDrb  *string `json:"ULMaxAccumGFBRPerDrb,omitempty" pn:"ULMaxAccumGFBRPerDrb"`
				Dscp                  *string `json:"Dscp,omitempty" pn:"Dscp"`
				RLC                   *struct {
					Mode *string `json:"Mode,omitempty" pn:"Mode"`
					AM   *struct {
						SNSize           *string `json:"SNSize,omitempty" pn:"SNSize"`
						PollPDU          *string `json:"PollPDU,omitempty" pn:"PollPDU"`
						PollByte         *string `json:"PollByte,omitempty" pn:"PollByte"`
						MaxRetxThreshold *string `json:"MaxRetxThreshold,omitempty" pn:"MaxRetxThreshold"`
						TPollRetransmit  *string `json:"TPollRetransmit,omitempty" pn:"TPollRetransmit"`
						TReassembly      *string `json:"TReassembly,omitempty" pn:"TReassembly"`
						TStatusProhibit  *string `json:"TStatusProhibit,omitempty" pn:"TStatusProhibit"`
					} `json:"AM,omitempty" pn:"AM"`
					UM *struct {
						SNSize      *string `json:"SNSize,omitempty" pn:"SNSize"`
						TReassembly *string `json:"TReassembly,omitempty" pn:"TReassembly"`
					} `json:"UM,omitempty" pn:"UM"`
				} `json:"RLC,omitempty" pn:"RLC"`
				LCH *struct {
					BucketSizeDuration  *string `json:"BucketSizeDuration,omitempty" pn:"BucketSizeDuration"`
					LogicalChannelGroup *string `json:"LogicalChannelGroup,omitempty" pn:"LogicalChannelGroup"`
					PrioritisedBitRate  *string `json:"PrioritisedBitRate,omitempty" pn:"PrioritisedBitRate"`
					Priority            *string `json:"Priority,omitempty" pn:"Priority"`
				} `json:"LCH,omitempty" pn:"LCH"`
				PDCP *struct {
					SNSizeDL             *string `json:"SNSizeDL,omitempty" pn:"SNSizeDL"`
					SNSizeUL             *string `json:"SNSizeUL,omitempty" pn:"SNSizeUL"`
					TReorderingDl        *string `json:"TReorderingDl,omitempty" pn:"TReorderingDl"`
					TReorderingUl        *string `json:"TReorderingUl,omitempty" pn:"TReorderingUl"`
					ULDataSplitThreshold *string `json:"ULDataSplitThreshold,omitempty" pn:"ULDataSplitThreshold"`
					DiscardTimer         *string `json:"DiscardTimer,omitempty" pn:"DiscardTimer"`
					PrimaryCellGroupId   *string `json:"PrimaryCellGroupId,omitempty" pn:"PrimaryCellGroupId"`
					StatusReportRequired *string `json:"StatusReportRequired,omitempty" pn:"StatusReportRequired"`
					OutOfOrderDelivery   *string `json:"OutOfOrderDelivery,omitempty" pn:"OutOfOrderDelivery"`
					ROHC                 *struct {
						Enable       *string `json:"Enable,omitempty" pn:"Enable"`
						MaxCID       *string `json:"MaxCID,omitempty" pn:"MaxCID"`
						ContinueRohc *string `json:"ContinueRohc,omitempty" pn:"ContinueRohc"`
						UplinkOnly   *string `json:"UplinkOnly,omitempty" pn:"UplinkOnly"`
						ProfileList  *struct {
							Profile0x0001 *string `json:"Profile0x0001,omitempty" pn:"Profile0x0001"`
							Profile0x0002 *string `json:"Profile0x0002,omitempty" pn:"Profile0x0002"`
							Profile0x0003 *string `json:"Profile0x0003,omitempty" pn:"Profile0x0003"`
							Profile0x0004 *string `json:"Profile0x0004,omitempty" pn:"Profile0x0004"`
							Profile0x0006 *string `json:"Profile0x0006,omitempty" pn:"Profile0x0006"`
							Profile0x0101 *string `json:"Profile0x0101,omitempty" pn:"Profile0x0101"`
							Profile0x0102 *string `json:"Profile0x0102,omitempty" pn:"Profile0x0102"`
							Profile0x0103 *string `json:"Profile0x0103,omitempty" pn:"Profile0x0103"`
							Profile0x0104 *string `json:"Profile0x0104,omitempty" pn:"Profile0x0104"`
						} `json:"ProfileList,omitempty" pn:"ProfileList"`
					} `json:"ROHC,omitempty" pn:"ROHC"`
				} `json:"PDCP,omitempty" pn:"PDCP"`
				PREConfig *struct {
					AveragingWindow        *string `json:"AveragingWindow,omitempty" pn:"AveragingWindow"`
					MaximumDataBurstVolume *string `json:"MaximumDataBurstVolume,omitempty" pn:"MaximumDataBurstVolume"`
					PERExponent            *string `json:"PERExponent,omitempty" pn:"PERExponent"`
					PERScalar              *string `json:"PERScalar,omitempty" pn:"PERScalar"`
					PacketDelayBudget      *string `json:"PacketDelayBudget,omitempty" pn:"PacketDelayBudget"`
					Priority               *string `json:"Priority,omitempty" pn:"Priority"`
					ResourceType           *string `json:"ResourceType,omitempty" pn:"ResourceType"`
				} `json:"PREConfig,omitempty" pn:"PREConfig"`
			} `json:"QoS,omitempty" pn:"QoS"`
			NSA *struct {
				QoS map[int]*struct {
					Type               *string `json:"Type,omitempty" pn:"Type"`
					QCI                *string `json:"QCI,omitempty" pn:"QCI"`
					Enable             *string `json:"Enable,omitempty" pn:"Enable"`
					Priority           *string `json:"Priority,omitempty" pn:"Priority"`
					QciToUse           *string `json:"QciToUse,omitempty" pn:"QciToUse"`
					Treordering        *string `json:"Treordering,omitempty" pn:"Treordering"`
					PdcpSnSizeDl       *string `json:"PdcpSnSizeDl,omitempty" pn:"PdcpSnSizeDl"`
					PdcpSnSizeUl       *string `json:"PdcpSnSizeUl,omitempty" pn:"PdcpSnSizeUl"`
					RejectEnable       *string `json:"RejectEnable,omitempty" pn:"RejectEnable"`
					OutOfOrderDelivery *string `json:"OutOfOrderDelivery,omitempty" pn:"OutOfOrderDelivery"`
					DlMaxNumHarqReTx   *string `json:"DlMaxNumHarqReTx,omitempty" pn:"DlMaxNumHarqReTx"`
					PacketDelayBudget  *string `json:"PacketDelayBudget,omitempty" pn:"PacketDelayBudget"`
					PrimaryCellGroupId *string `json:"PrimaryCellGroupId,omitempty" pn:"PrimaryCellGroupId"`
				} `json:"QoS,omitempty" pn:"QoS"`
			} `json:"NSA,omitempty" pn:"NSA"`
			Default5QI *struct {
				PLMN map[int]*struct {
					PLMNID     *string `json:"PLMNID,omitempty" pn:"PLMNID"`
					Default5QI *string `json:"Default5QI,omitempty" pn:"Default5QI"`
				} `json:"PLMN,omitempty" pn:"PLMN"`
			} `json:"Default5QI,omitempty" pn:"Default5QI"`
			Slice *struct {
				SliceEnable *string `json:"SliceEnable,omitempty" pn:"SliceEnable"`
				SliceList   map[int]*struct {
					Heavy *string `json:"Heavy,omitempty" pn:"Heavy"`
				} `json:"SliceList,omitempty" pn:"SliceList"`
				FloatList map[int]*struct {
					Heavy *string `json:"Heavy,omitempty" pn:"Heavy"`
				} `json:"FloatList,omitempty" pn:"FloatList"`
			} `json:"Slice,omitempty" pn:"Slice"`
		} `json:"NGC,omitempty" pn:"NGC"`
		Capabilities *struct {
			PRBNum                  *string `json:"PRBNum,omitempty" pn:"PRBNum"`
			NoisePwdBm              *string `json:"NoisePwdBm,omitempty" pn:"NoisePwdBm"`
			MaxUEsServed            *string `json:"MaxUEsServed,omitempty" pn:"MaxUEsServed"`
			GbrDlThreshold          *string `json:"GbrDlThreshold,omitempty" pn:"GbrDlThreshold"`
			GbrUlThreshold          *string `json:"GbrUlThreshold,omitempty" pn:"GbrUlThreshold"`
			GbrThresholdSwitch      *string `json:"GbrThresholdSwitch,omitempty" pn:"GbrThresholdSwitch"`
			InactiveSwitch          *string `json:"InactiveSwitch,omitempty" pn:"InactiveSwitch"`
			SupportActiveRRCNumbers *string `json:"SupportActiveRRCNumbers,omitempty" pn:"SupportActiveRRCNumbers"`
			UeInactiveTimer         *string `json:"UeInactiveTimer,omitempty" pn:"UeInactiveTimer"`
			UserNumberReservedForHo *string `json:"UserNumberReservedForHo,omitempty" pn:"UserNumberReservedForHo"`
			UserNumberThreshold1    *string `json:"UserNumberThreshold1,omitempty" pn:"UserNumberThreshold1"`
			UserNumberThreshold2    *string `json:"UserNumberThreshold2,omitempty" pn:"UserNumberThreshold2"`
			UserNumberThreshold3    *string `json:"UserNumberThreshold3,omitempty" pn:"UserNumberThreshold3"`
			LoadBalanceConfig       *struct {
				AcceptDlPrbUsage  *string `json:"AcceptDlPrbUsage,omitempty" pn:"AcceptDlPrbUsage"`
				AcceptUlPrbUsage  *string `json:"AcceptUlPrbUsage,omitempty" pn:"AcceptUlPrbUsage"`
				DlPrbUsage        *string `json:"DlPrbUsage,omitempty" pn:"DlPrbUsage"`
				UlPrbUsage        *string `json:"UlPrbUsage,omitempty" pn:"UlPrbUsage"`
				LoadBalanceSwitch *string `json:"LoadBalanceSwitch,omitempty" pn:"LoadBalanceSwitch"`
				TimeToTrigger     *string `json:"TimeToTrigger,omitempty" pn:"TimeToTrigger"`
			} `json:"LoadBalanceConfig,omitempty" pn:"LoadBalanceConfig"`
		} `json:"Capabilities,omitempty" pn:"Capabilities"`
		VoNR *struct {
			IMSAdapt        *string `json:"IMSAdapt,omitempty" pn:"IMSAdapt"`
			VoNRSwitch      *string `json:"VoNRSwitch,omitempty" pn:"VoNRSwitch"`
			SPSSwitchQCI1Ul *string `json:"SPSSwitchQCI1Ul,omitempty" pn:"SPSSwitchQCI1Ul"`
			VoNRList        map[int]*struct {
				Vo5qi     *string `json:"Vo5qi,omitempty" pn:"Vo5qi"`
				PLMNID    *string `json:"PLMNID,omitempty" pn:"PLMNID"`
				VoSupport *string `json:"VoSupport,omitempty" pn:"VoSupport"`
			} `json:"VoNRList,omitempty" pn:"VoNRList"`
		} `json:"VoNR,omitempty" pn:"VoNR"`
	} `json:"NR,omitempty" pn:"NR"`
}
