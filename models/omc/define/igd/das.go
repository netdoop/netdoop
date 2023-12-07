package igd

type MU struct {
	ProductClass              *string `json:"ProductClass,omitempty" pn:"ProductClass"`
	ModelName                 *string `json:"ModelName,omitempty" pn:"ModelName"`
	SerialNumber              *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
	UserLabel                 *string `json:"UserLabel,omitempty" pn:"UserLabel"`
	DnPrefix                  *string `json:"DnPrefix,omitempty" pn:"DnPrefix"`
	Manufacturer              *string `json:"Manufacturer,omitempty" pn:"Manufacturer"`
	ManufacturerData          *string `json:"ManufacturerData,omitempty" pn:"ManufacturerData"`
	ManufacturerOUI           *string `json:"ManufacturerOUI,omitempty" pn:"ManufacturerOUI"`
	VendorUnitFamilyType      *string `json:"VendorUnitFamilyType,omitempty" pn:"VendorUnitFamilyType"`
	VendorUnitTypeNumber      *string `json:"VendorUnitTypeNumber,omitempty" pn:"VendorUnitTypeNumber"`
	ProvisioningCode          *string `json:"ProvisioningCode,omitempty" pn:"ProvisioningCode"`
	HardwarePlatform          *string `json:"HardwarePlatform,omitempty" pn:"HardwarePlatform"`
	HardwareVersion           *string `json:"HardwareVersion,omitempty" pn:"HardwareVersion"`
	SoftwareVersion           *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
	AdditionalHardwareVersion *string `json:"AdditionalHardwareVersion,omitempty" pn:"AdditionalHardwareVersion"`
	AdditionalSoftwareVersion *string `json:"AdditionalSoftwareVersion,omitempty" pn:"AdditionalSoftwareVersion"`
	FirstUseDate              *string `json:"FirstUseDate,omitempty" pn:"FirstUseDate"`
	DateOfLastService         *string `json:"DateOfLastService,omitempty" pn:"DateOfLastService"`
	DateOfManufacture         *string `json:"DateOfManufacture,omitempty" pn:"DateOfManufacture"`
	UpTime                    *string `json:"UpTime,omitempty" pn:"UpTime"`
	Status                    *string `json:"Status,omitempty" pn:"Status"`
	Reboot                    *string `json:"Reboot,omitempty" pn:"Reboot"`
	ClockSource               *string `json:"ClockSource,omitempty" pn:"ClockSource"`
	ClockSynMode              *string `json:"ClockSynMode,omitempty" pn:"ClockSynMode"`
	ClockSyncMode             *string `json:"ClockSyncMode,omitempty" pn:"ClockSyncMode"`
	ClockSyncStatus           *string `json:"ClockSyncStatus,omitempty" pn:"ClockSyncStatus"`
	EuSyncSource              *string `json:"EuSyncSource,omitempty" pn:"EuSyncSource"`
	EuFrameOffset             *string `json:"EuFrameOffset,omitempty" pn:"EuFrameOffset"`
	EuSendTimingSwitch        *string `json:"EuSendTimingSwitch,omitempty" pn:"EuSendTimingSwitch"`
	SendTimingMode            *string `json:"SendTimingMode,omitempty" pn:"SendTimingMode"`

	SlotsInformation *string `json:"SlotsInformation,omitempty" pn:"SlotsInformation"`
	Slot             map[int]*struct {
		Manufacturer         *string `json:"Manufacturer,omitempty" pn:"Manufacturer"`
		ManufacturerOUI      *string `json:"ManufacturerOUI,omitempty" pn:"ManufacturerOUI"`
		VendorUnitFamilyType *string `json:"VendorUnitFamilyType,omitempty" pn:"VendorUnitFamilyType"`
		ModelName            *string `json:"ModelName,omitempty" pn:"ModelName"`
		PackPosition         *string `json:"PackPosition,omitempty" pn:"PackPosition"`
		DataModel            *string `json:"DataModel,omitempty" pn:"DataModel"`
		DataModelSpecVersion *string `json:"DataModelSpecVersion,omitempty" pn:"DataModelSpecVersion"`
		A3GPPSpecVersion     *string `json:"3GPPSpecVersion,omitempty" pn:"3GPPSpecVersion"`
		DateOfManufacture    *string `json:"DateOfManufacture,omitempty" pn:"DateOfManufacture"`
		FirstUseDate         *string `json:"FirstUseDate,omitempty" pn:"FirstUseDate"`
		UpTime               *string `json:"UpTime,omitempty" pn:"UpTime"`
		SlotsOccupied        *string `json:"SlotsOccupied,omitempty" pn:"SlotsOccupied"`
		Status               *string `json:"Status,omitempty" pn:"Status"`
		Reboot               *string `json:"Reboot,omitempty" pn:"Reboot"`
		Debug                *struct {
		} `json:"Debug,omitempty" pn:"Debug"`
		EuType *string `json:"EuType,omitempty" pn:"EuType"`
		EU     map[int]*struct {
			Manufacturer     *string `json:"Manufacturer,omitempty" pn:"Manufacturer"`
			ManufacturerOUI  *string `json:"ManufacturerOUI,omitempty" pn:"ManufacturerOUI"`
			ProvisioningCode *string `json:"ProvisioningCode,omitempty" pn:"ProvisioningCode"`
			ModelName        *string `json:"ModelName,omitempty" pn:"ModelName"`
			SerialNumber     *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
			UserLabel        *string `json:"UserLabel,omitempty" pn:"UserLabel"`
			DeviceId         *string `json:"DeviceId,omitempty" pn:"DeviceId"`
			HardwareVersion  *string `json:"HardwareVersion,omitempty" pn:"HardwareVersion"`
			SoftwareVersion  *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
			Location         *string `json:"Location,omitempty" pn:"Location"`
			Level            *string `json:"Level,omitempty" pn:"Level"`
			RouteIndex       *string `json:"RouteIndex,omitempty" pn:"RouteIndex"`
			IP               *string `json:"IP,omitempty" pn:"IP"`
			Port             *string `json:"Port,omitempty" pn:"Port"`
			OptId            *string `json:"OptId,omitempty" pn:"OptId"`
			DLCRCSum         *string `json:"DLCRCSum,omitempty" pn:"DLCRCSum"`
			ULCRCSum         *string `json:"ULCRCSum,omitempty" pn:"ULCRCSum"`
			Status           *string `json:"Status,omitempty" pn:"Status"`
			Reboot           *string `json:"Reboot,omitempty" pn:"Reboot"`
			AdminState       *string `json:"AdminState,omitempty" pn:"AdminState"`
			AdminExist       *string `json:"AdminExist,omitempty" pn:"AdminExist"`
			RU               map[int]*struct {
				Manufacturer         *string `json:"Manufacturer,omitempty" pn:"Manufacturer"`
				ManufacturerOUI      *string `json:"ManufacturerOUI,omitempty" pn:"ManufacturerOUI"`
				VendorUnitFamilyType *string `json:"VendorUnitFamilyType,omitempty" pn:"VendorUnitFamilyType"`
				VendorUnitTypeNumber *string `json:"VendorUnitTypeNumber,omitempty" pn:"VendorUnitTypeNumber"`
				ProvisioningCode     *string `json:"ProvisioningCode,omitempty" pn:"ProvisioningCode"`
				ModelName            *string `json:"ModelName,omitempty" pn:"ModelName"`
				UserLabel            *string `json:"UserLabel,omitempty" pn:"UserLabel"`
				DeviceId             *string `json:"DeviceId,omitempty" pn:"DeviceId"`
				SerialNumber         *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
				HardwareVersion      *string `json:"HardwareVersion,omitempty" pn:"HardwareVersion"`
				SoftwareVersion      *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
				DateOfManufacture    *string `json:"DateOfManufacture,omitempty" pn:"DateOfManufacture"`
				RouteIndex           *string `json:"RouteIndex,omitempty" pn:"RouteIndex"`
				Level                *string `json:"Level,omitempty" pn:"Level"`
				Location             *string `json:"Location,omitempty" pn:"Location"`
				OptId                *string `json:"OptId,omitempty" pn:"OptId"`
				Status               *string `json:"Status,omitempty" pn:"Status"`
				Reboot               *string `json:"Reboot,omitempty" pn:"Reboot"`
				AdminExist           *string `json:"AdminExist,omitempty" pn:"AdminExist"`
				FrequencyBand        *string `json:"FrequencyBand,omitempty" pn:"FrequencyBand"`
				MaxTxPower           *string `json:"MaxTxPower,omitempty" pn:"MaxTxPower"`
				RFTxStatus           *string `json:"RFTxStatus,omitempty" pn:"RFTxStatus"`
				RFChannel            map[int]*struct {
					TxGain     *string `json:"TxGain,omitempty" pn:"TxGain"`
					NoisePwdBm *string `json:"NoisePwdBm,omitempty" pn:"NoisePwdBm"`
				} `json:"RFChannel,omitempty" pn:"RFChannel"`
				SwUpgrade *struct {
					Stage        *string `json:"Stage,omitempty" pn:"Stage"`
					Status       *string `json:"Status,omitempty" pn:"Status"`
					FailureCause *string `json:"FailureCause,omitempty" pn:"FailureCause"`
				} `json:"SwUpgrade,omitempty" pn:"SwUpgrade"`
			} `json:"RU,omitempty" pn:"RU"`
			SwUpgrade *struct {
				Stage        *string `json:"Stage,omitempty" pn:"Stage"`
				Status       *string `json:"Status,omitempty" pn:"Status"`
				FailureCause *string `json:"FailureCause,omitempty" pn:"FailureCause"`
			} `json:"SwUpgrade,omitempty" pn:"SwUpgrade"`
		} `json:"EU,omitempty" pn:"EU"`
		PriEU map[int]*struct {
			ModelName         *string `json:"ModelName,omitempty" pn:"ModelName"`
			SerialNumber      *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
			HardwareVersion   *string `json:"HardwareVersion,omitempty" pn:"HardwareVersion"`
			SoftwareVersion   *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
			TopID             *string `json:"TopID,omitempty" pn:"TopID"`
			Status            *string `json:"Status,omitempty" pn:"Status"`
			Reboot            *string `json:"Reboot,omitempty" pn:"Reboot"`
			SelfCureThreshold *string `json:"SelfCureThreshold,omitempty" pn:"SelfCureThreshold"`
			SyncDelayComps    *string `json:"SyncDelayComps,omitempty" pn:"SyncDelayComps"`
			RuHeartErrReboot  *string `json:"RuHeartErrReboot,omitempty" pn:"RuHeartErrReboot"`
			NosieErrReboot    *string `json:"NosieErrReboot,omitempty" pn:"NosieErrReboot"`
			TotalConsum       *string `json:"TotalConsum,omitempty" pn:"TotalConsum"`
			LTEWorkModel      *string `json:"LTEWorkModel,omitempty" pn:"LTEWorkModel"`
			LTECell1DLFreq    *string `json:"LTECell1DLFreq,omitempty" pn:"LTECell1DLFreq"`
			LTECell1ULFreq    *string `json:"LTECell1ULFreq,omitempty" pn:"LTECell1ULFreq"`
			LTECell2DLFreq    *string `json:"LTECell2DLFreq,omitempty" pn:"LTECell2DLFreq"`
			LTECell2ULFreq    *string `json:"LTECell2ULFreq,omitempty" pn:"LTECell2ULFreq"`
			NRFreq            *string `json:"NRFreq,omitempty" pn:"NRFreq"`
			NRSync            *string `json:"NRSync,omitempty" pn:"NRSync"`
			FFTGain           *string `json:"FFTGain,omitempty" pn:"FFTGain"`
			IFFTBulk          *string `json:"IFFTBulk,omitempty" pn:"IFFTBulk"`
			EthReady          *string `json:"EthReady,omitempty" pn:"EthReady"`
			MacSetMode        *string `json:"MacSetMode,omitempty" pn:"MacSetMode"`
			AntNums           *string `json:"AntNums,omitempty" pn:"AntNums"`
			PADelay           *string `json:"PADelay,omitempty" pn:"PADelay"`
			OptId             *string `json:"OptId,omitempty" pn:"OptId"`
			RuBitErrReboot    *string `json:"RuBitErrReboot,omitempty" pn:"RuBitErrReboot"`
			BitErrCheckCycle  *string `json:"BitErrCheckCycle,omitempty" pn:"BitErrCheckCycle"`
			BitErrCheckCounts *string `json:"BitErrCheckCounts,omitempty" pn:"BitErrCheckCounts"`
			OptPower          *map[int]*struct {
				RxTxPower *string `json:"RxTxPower,omitempty" pn:"RxTxPower"`
			} `json:"OptPower,omitempty" pn:"OptPower"`
			Cell *struct {
				RRUMAC *string `json:"RRUMAC,omitempty" pn:"RRUMAC"`
				BBUMAC *string `json:"BBUMAC,omitempty" pn:"BBUMAC"`
			} `json:"Cell,omitempty" pn:"Cell"`
			RuPower *map[int]*struct {
				Power      *string `json:"Power,omitempty" pn:"Power"`
				TxPowerAvg *string `json:"TxPowerAvg,omitempty" pn:"TxPowerAvg"`
				TxPowerMax *string `json:"TxPowerMax,omitempty" pn:"TxPowerMax"`
			} `json:"RuPower,omitempty" pn:"RuPower"`
			RU map[int]*struct {
				ModelName       *string `json:"ModelName,omitempty" pn:"ModelName"`
				SerialNumber    *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
				HardwareVersion *string `json:"HardwareVersion,omitempty" pn:"HardwareVersion"`
				SoftwareVersion *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
				TopID           *string `json:"TopID,omitempty" pn:"TopID"`
				Status          *string `json:"Status,omitempty" pn:"Status"`
				Reboot          *string `json:"Reboot,omitempty" pn:"Reboot"`
				OptId           *string `json:"OptId,omitempty" pn:"OptId"`
				PowerOn         *string `json:"PowerOn,omitempty" pn:"PowerOn"`
				ULGainCtl       *string `json:"ULGainCtl,omitempty" pn:"ULGainCtl"`
				MaxTxPower      *string `json:"MaxTxPower,omitempty" pn:"MaxTxPower"`
				LTEMaxTxPower   *string `json:"LTEMaxTxPower,omitempty" pn:"LTEMaxTxPower"`
				Channel         map[int]*struct {
					RxPower *string `json:"RxPower,omitempty" pn:"RxPower"`
					TxPower *string `json:"TxPower,omitempty" pn:"TxPower"`
					DLAtt   *string `json:"DLAtt,omitempty" pn:"DLAtt"`
					ULAtt   *string `json:"ULAtt,omitempty" pn:"ULAtt"`
				} `json:"Channel,omitempty" pn:"Channel"`
				LTEChannel map[int]*struct {
					RxPower *string `json:"RxPower,omitempty" pn:"RxPower"`
					TxPower *string `json:"TxPower,omitempty" pn:"TxPower"`
					DLAtt   *string `json:"DLAtt,omitempty" pn:"DLAtt"`
					ULAtt   *string `json:"ULAtt,omitempty" pn:"ULAtt"`
				} `json:"LTEChannel,omitempty" pn:"LTEChannel"`
			} `json:"RU,omitempty" pn:"RU"`
		} `json:"PriEU,omitempty" pn:"PriEU"`
		PriEU0 map[int]*struct {
			SoftwareVersion *string `json:"SoftwareVersion,omitempty" pn:"SoftwareVersion"`
			SerialNumber    *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
			Status          *string `json:"Status,omitempty" pn:"Status"`
			Reboot          *string `json:"Reboot,omitempty" pn:"Reboot"`
			TotalOffset     *string `json:"TotalOffset,omitempty" pn:"TotalOffset"`
			FailureCause    *string `json:"FailureCause,omitempty" pn:"FailureCause"`
			Cell            *struct {
				RRUMAC *string `json:"RRUMAC,omitempty" pn:"RRUMAC"`
				BBUMAC *string `json:"BBUMAC,omitempty" pn:"BBUMAC"`
			} `json:"Cell,omitempty" pn:"Cell"`
			SwUpgrade *struct {
				Stage        *string `json:"Stage,omitempty" pn:"Stage"`
				Status       *string `json:"Status,omitempty" pn:"Status"`
				FailureCause *string `json:"FailureCause,omitempty" pn:"FailureCause"`
			} `json:"SwUpgrade,omitempty" pn:"SwUpgrade"`
		} `json:"PriEU0,omitempty" pn:"PriEU0"`
		Site *struct {
			FFTGain        *string `json:"FFTGain,omitempty" pn:"FFTGain"`
			IFFTBulk       *string `json:"IFFTBulk,omitempty" pn:"IFFTBulk"`
			PADelay        *string `json:"PADelay,omitempty" pn:"PADelay"`
			EthReady       *string `json:"EthReady,omitempty" pn:"EthReady"`
			DeliveryEnable *string `json:"DeliveryEnable,omitempty" pn:"DeliveryEnable"`
			NRStartModel   *string `json:"NRStartModel,omitempty" pn:"NRStartModel"`
			NRSync         *string `json:"NRSync,omitempty" pn:"NRSync"`
			Channel        map[int]*struct {
				DLAtt *string `json:"DLAtt,omitempty" pn:"DLAtt"`
				ULAtt *string `json:"ULAtt,omitempty" pn:"ULAtt"`
			} `json:"Channel,omitempty" pn:"Channel"`
			LTEChannel map[int]*struct {
				DLAtt *string `json:"DLAtt,omitempty" pn:"DLAtt"`
				ULAtt *string `json:"ULAtt,omitempty" pn:"ULAtt"`
			} `json:"LTEChannel,omitempty" pn:"LTEChannel"`
		} `json:"Site,omitempty" pn:"Site"`
		SwUpgrade *struct {
			Stage        *string `json:"Stage,omitempty" pn:"Stage"`
			Status       *string `json:"Status,omitempty" pn:"Status"`
			FailureCause *string `json:"FailureCause,omitempty" pn:"FailureCause"`
		} `json:"SwUpgrade,omitempty" pn:"SwUpgrade"`
	} `json:"Slot,omitempty" pn:"Slot"`
}
