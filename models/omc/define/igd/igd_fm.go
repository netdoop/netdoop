package igd

type FaultMgmt struct {
	SupportedAlarmNumberOfEntries *uint `json:"SupportedAlarmNumberOfEntries,omitempty" pn:"SupportedAlarmNumberOfEntries"`
	MaxCurrentAlarmEntries        *uint `json:"MaxCurrentAlarmEntries,omitempty" pn:"MaxCurrentAlarmEntries"`
	CurrentAlarmNumberOfEntries   *uint `json:"CurrentAlarmNumberOfEntries,omitempty" pn:"CurrentAlarmNumberOfEntries"`
	HistoryEventNumberOfEntries   *uint `json:"HistoryEventNumberOfEntries,omitempty" pn:"HistoryEventNumberOfEntries"`
	ExpeditedEventNumberOfEntries *uint `json:"ExpeditedEventNumberOfEntries,omitempty" pn:"ExpeditedEventNumberOfEntries"`
	QueuedEventNumberOfEntries    *uint `json:"QueuedEventNumberOfEntries,omitempty" pn:"QueuedEventNumberOfEntries"`

	SupportedAlarm map[int]*SupportedAlarm `json:"SupportedAlarm,omitempty" pn:"SupportedAlarm"`
	CurrentAlarm   map[int]*CurrentAlarm   `json:"CurrentAlarm,omitempty" pn:"CurrentAlarm"`
	HistoryEvent   map[int]*HistoryEvent   `json:"HistoryEvent,omitempty" pn:"HistoryEvent"`
	ExpeditedEvent map[int]*ExpeditedEvent `json:"ExpeditedEvent,omitempty" pn:"ExpeditedEvent"`
	QueuedEvent    map[int]*QueuedEvent    `json:"QueuedEvent,omitempty" pn:"QueuedEvent"`

	X_VENDOR_ALARM_ENABLE *string `json:"X_VENDOR_ALARM_ENABLE,omitempty" pn:"X_VENDOR_ALARM_ENABLE"`
}

type SupportedAlarm struct {
	EventType          *string `json:"EventType,omitempty" pn:"EventType"`
	ProbableCause      *string `json:"ProbableCause,omitempty" pn:"ProbableCause"`
	SpecificProblem    *string `json:"SpecificProblem,omitempty" pn:"SpecificProblem"`
	PerceivedSeverity  *string `json:"PerceivedSeverity,omitempty" pn:"PerceivedSeverity"`
	ReportingMechanism *string `json:"ReportingMechanism,omitempty" pn:"ReportingMechanism"`
}

type CurrentAlarm struct {
	AlarmIdentifier       *string `json:"AlarmIdentifier,omitempty" pn:"AlarmIdentifier"`
	AlarmRaisedTime       *string `json:"AlarmRaisedTime,omitempty" pn:"AlarmRaisedTime"`
	AlarmChangedTime      *string `json:"AlarmChangedTime,omitempty" pn:"AlarmChangedTime"`
	ManagedObjectInstance *string `json:"ManagedObjectInstance,omitempty" pn:"ManagedObjectInstance"`
	EventType             *string `json:"EventType,omitempty" pn:"EventType"`
	ProbableCause         *string `json:"ProbableCause,omitempty" pn:"ProbableCause"`
	SpecificProblem       *string `json:"SpecificProblem,omitempty" pn:"SpecificProblem"`
	PerceivedSeverity     *string `json:"PerceivedSeverity,omitempty" pn:"PerceivedSeverity"`
	AdditionalText        *string `json:"AdditionalText,omitempty" pn:"AdditionalText"`
	AdditionalInformation *string `json:"AdditionalInformation,omitempty" pn:"AdditionalInformation"`
	OUI                   *string `json:"OUI,omitempty" pn:"OUI"`
	SerialNumber          *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
}

type HistoryEvent struct {
	EventTime             *string `json:"EventTime,omitempty" pn:"EventTime"`
	AlarmIdentifier       *string `json:"AlarmIdentifier,omitempty" pn:"AlarmIdentifier"`
	NotificationType      *string `json:"NotificationType,omitempty" pn:"NotificationType"`
	ManagedObjectInstance *string `json:"ManagedObjectInstance,omitempty" pn:"ManagedObjectInstance"`
	EventType             *string `json:"EventType,omitempty" pn:"EventType"`
	ProbableCause         *string `json:"ProbableCause,omitempty" pn:"ProbableCause"`
	SpecificProblem       *string `json:"SpecificProblem,omitempty" pn:"SpecificProblem"`
	PerceivedSeverity     *string `json:"PerceivedSeverity,omitempty" pn:"PerceivedSeverity"`
	AdditionalText        *string `json:"AdditionalText,omitempty" pn:"AdditionalText"`
	AdditionalInformation *string `json:"AdditionalInformation,omitempty" pn:"AdditionalInformation"`
	OUI                   *string `json:"OUI,omitempty" pn:"OUI"`
	SerialNumber          *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
}

type ExpeditedEvent struct {
	EventTime             *string `json:"EventTime,omitempty" pn:"EventTime"`
	AlarmIdentifier       *string `json:"AlarmIdentifier,omitempty" pn:"AlarmIdentifier"`
	NotificationType      *string `json:"NotificationType,omitempty" pn:"NotificationType"`
	ManagedObjectInstance *string `json:"ManagedObjectInstance,omitempty" pn:"ManagedObjectInstance"`
	EventType             *string `json:"EventType,omitempty" pn:"EventType"`
	ProbableCause         *string `json:"ProbableCause,omitempty" pn:"ProbableCause"`
	SpecificProblem       *string `json:"SpecificProblem,omitempty" pn:"SpecificProblem"`
	PerceivedSeverity     *string `json:"PerceivedSeverity,omitempty" pn:"PerceivedSeverity"`
	AdditionalText        *string `json:"AdditionalText,omitempty" pn:"AdditionalText"`
	AdditionalInformation *string `json:"AdditionalInformation,omitempty" pn:"AdditionalInformation"`
	OUI                   *string `json:"OUI,omitempty" pn:"OUI"`
	SerialNumber          *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
}
type QueuedEvent struct {
	EventTime             *string `json:"EventTime,omitempty" pn:"EventTime"`
	AlarmIdentifier       *string `json:"AlarmIdentifier,omitempty" pn:"AlarmIdentifier"`
	NotificationType      *string `json:"NotificationType,omitempty" pn:"NotificationType"`
	ManagedObjectInstance *string `json:"ManagedObjectInstance,omitempty" pn:"ManagedObjectInstance"`
	EventType             *string `json:"EventType,omitempty" pn:"EventType"`
	ProbableCause         *string `json:"ProbableCause,omitempty" pn:"ProbableCause"`
	SpecificProblem       *string `json:"SpecificProblem,omitempty" pn:"SpecificProblem"`
	PerceivedSeverity     *string `json:"PerceivedSeverity,omitempty" pn:"PerceivedSeverity"`
	AdditionalText        *string `json:"AdditionalText,omitempty" pn:"AdditionalText"`
	AdditionalInformation *string `json:"AdditionalInformation,omitempty" pn:"AdditionalInformation"`
	OUI                   *string `json:"OUI,omitempty" pn:"OUI"`
	SerialNumber          *string `json:"SerialNumber,omitempty" pn:"SerialNumber"`
}
