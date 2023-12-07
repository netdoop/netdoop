package igd

type WiFi struct {
	WLANSettings     *WLANSettings     `json:"CountryCode,omitempty" pn:"WLANSettings"`
	AccessManagement *AccessManagement `json:"AccessManagement,omitempty" pn:"AccessManagement"`
}

type WLANSettings struct {
	CountryCode  *string              `json:"CountryCode,omitempty" pn:"CountryCode"`
	Enable       *string              `json:"Enable,omitempty" pn:"Enable"`
	Band         *string              `json:"Band,omitempty" pn:"Band"`
	Mode         *string              `json:"Mode,omitempty" pn:"Mode"`
	Channel      *string              `json:"Channel,omitempty" pn:"Channel"`
	TxPower      *string              `json:"TxPower,omitempty" pn:"TxPower"`
	SSIDProfiles map[int]*SSIDProfile `json:"SSIDProfiles,omitempty" pn:"SSIDProfile"`
	WPS          *WPS                 `json:"WPS,omitempty" pn:"WPS"`
}

type SSIDProfile struct {
	Index             uint    `json:"Index,omitempty" pn:"Index"`
	Enable            *string `json:"Enable,omitempty" pn:"Enable"`
	SSID              *string `json:"SSID,omitempty" pn:"SSID"`
	MACAddress        *string `json:"MACAddress,omitempty" pn:"MACAddress"`
	Guest             *int    `json:"Guest,omitempty" pn:"Guest"`
	MaxNoOfDev        uint    `json:"MaxNoOfDev,omitempty" pn:"MaxNoOfDev"`
	HideSSIDBroadcast *int    `json:"HideSSIDBroadcast,omitempty" pn:"HideSSIDBroadcast"`
	APIsolation       *int    `json:"APIsolation,omitempty" pn:"APIsolation"`
	EncryptionMode    *string `json:"EncryptionMode,omitempty" pn:"EncryptionMode"`
	WPAKey            *string `json:"WPAKey,omitempty" pn:"WPAKey"`
	WEPKeyIndex       uint    `json:"WEPKeyIndex,omitempty" pn:"WEPKeyIndex"`
	WEPKeyLength      uint    `json:"WEPKeyLength,omitempty" pn:"WEPKeyLength"`
	WEPKey1           *string `json:"WEPKey1,omitempty" pn:"WEPKey1"`
	WEPKey2           *string `json:"WEPKey2,omitempty" pn:"WEPKey2"`
	WEPKey3           *string `json:"WEPKey3,omitempty" pn:"WEPKey3"`
	WEPKey4           *string `json:"WEPKey4,omitempty" pn:"WEPKey4"`
	AccessToIntranet  *int    `json:"AccessToIntranet,omitempty" pn:"AccessToIntranet"`
}

type WPS struct {
	Enable *string `json:"Enable,omitempty" pn:"Enable"`
}

type AccessManagement struct {
	Settings *string `json:"Settings,omitempty" pn:"Settings"`
	List     map[int]*struct {
		MACAddress *string `json:"MACAddress,omitempty" pn:"MACAddress"`
	} `json:"List,omitempty" pn:"List"`
}
