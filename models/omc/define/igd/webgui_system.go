package igd

type System struct {
	ScheduleReboot *ScheduleReboot `json:"ScheduleReboot"  pn:"ScheduleReboot"`
	OnlineUpgrade  *OnlineUpgrade  `json:"OnlineUpgrade"  pn:"OnlineUpgrade"`
	DateAndTime    *DateAndTime    `json:"DateAndTime,omitempty" pn:"DateAndTime"`
	DDNS           *DDNS           `json:"DDNS,omitempty" pn:"DDNS"`
	Syslog         *Syslog         `json:"Syslog,omitempty" pn:"Syslog"`
	WebSettings    *WebSettings    `json:"WebSettings,omitempty" pn:"WebSettings"`
	Account        *Account        `json:"Account,omitempty" pn:"Account"`
}

type ScheduleReboot struct {
	Enable       *string `json:"Enable,omitempty" pn:"Enable"`
	Time         *string `json:"Time,omitempty" pn:"Time"`
	DateToReboot *string `json:"DateToReboot,omitempty" pn:"DateToReboot"`
}
type OnlineUpgrade struct {
	Enable                         *string `json:"Enable,omitempty" pn:"Enable"`
	CheckNewFWAfterConnectedEnable *string `json:"CheckNewFWAfterConnectedEnable,omitempty" pn:"CheckNewFWAfterConnectedEnable"`
	UpgradeFolder                  *string `json:"UpgradeFolder,omitempty" pn:"UpgradeFolder"`
	VersionFile                    *string `json:"VersionFile,omitempty" pn:"VersionFile"`
	Username                       *string `json:"Username,omitempty" pn:"Username"`
	Password                       *string `json:"Password,omitempty" pn:"Password"`
	CheckNewFirmwareEvery          *string `json:"CheckNewFirmwareEvery,omitempty" pn:"CheckNewFirmwareEvery"`
	StartTime                      *int    `json:"StartTime,omitempty" pn:"StartTime"`
	RandomTime                     *int    `json:"RandomTime,omitempty" pn:"RandomTime"`
}

type DateAndTime struct {
	CurrentTime       *string `json:"CurrentTime,omitempty" pn:"CurrentTime"`
	NTPServer         *string `json:"NTPServer,omitempty" pn:"NTPServer"`
	OptionalNTPServer *string `json:"OptionalNTPServer,omitempty" pn:"OptionalNTPServer"`
	TimeZone          *string `json:"TimeZone,omitempty" pn:"TimeZone"`
	DST               *DST    `json:"DST,omitempty" pn:"DST"`
}

type DST struct {
	Enable    *string `json:"Enable,omitempty" pn:"Enable"`
	StartTime *string `json:"StartTime,omitempty" pn:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" pn:"EndTime"`
	Status    *string `json:"Status,omitempty" pn:"Status"`
}

type DDNS struct {
	Enable          *string `json:"Enable,omitempty" pn:"Enable"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" pn:"ServiceProvider"`
	Domain          *string `json:"Domain,omitempty" pn:"Domain"`
	Username        *string `json:"Username,omitempty" pn:"Username"`
	Password        *string `json:"Password,omitempty" pn:"Password"`
	Refresh         *int    `json:"Refresh,omitempty" pn:"Refresh"`
	Wildcard        *int    `json:"Wildcard,omitempty" pn:"Wildcard"`
	Verification    *int    `json:"Verification,omitempty" pn:"Verification"`
	CheckEvery      *int    `json:"CheckEvery,omitempty" pn:"CheckEvery"`
	Status          *string `json:"Status,omitempty" pn:"Status"`
}

type Syslog struct {
	Level            *string `json:"Level,omitempty" pn:"Level"`
	ForwardIPAddress *string `json:"ForwardIPAddress,omitempty" pn:"ForwardIPAddress"`
	Clear            *int    `json:"Clear,omitempty" pn:"Clear"`
	Method           *string `json:"Method,omitempty" pn:"Method"`
}

type WebSettings struct {
	RefreshTime    uint    `json:"RefreshTime,omitempty" pn:"RefreshTime"`
	SessionTimeOut uint    `json:"SessionTimeOut,omitempty" pn:"SessionTimeOut"`
	HTTPEnable     *string `json:"HTTPEnable,omitempty" pn:"HTTPEnable"`
	HTTPsEnable    *string `json:"HTTPsEnable,omitempty" pn:"HTTPsEnable"`
	HTTPPort       uint    `json:"HTTPPort,omitempty" pn:"HTTPPort"`
	HTTPsPort      uint    `json:"HTTPsPort,omitempty" pn:"HTTPsPort"`
	HTTPsWANEnable *string `json:"HTTPsWANEnable,omitempty" pn:"HTTPsWANEnable"`
	PingFromWAN    *int    `json:"PingFromWAN,omitempty" pn:"PingFromWAN"`
	Language       *string `json:"Language,omitempty" pn:"Language"`
}

type Account struct {
	ChangePassword *ChangePassword `json:"ChangePassword,omitempty" pn:"ChangePassword"`
	EnableUser     *int            `json:"EnableUser,omitempty" pn:"EnableUser"`
}

type ChangePassword struct {
	User  *string `json:"User,omitempty" pn:"User"`
	Admin *string `json:"Admin,omitempty" pn:"Admin"`
}
