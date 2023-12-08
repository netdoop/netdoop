package pm

import (
	"encoding/xml"
	"time"
)

type MeasCollecFile struct {
	XMLName    xml.Name   `xml:"measCollecFile"`
	FileHeader FileHeader `xml:"fileHeader"`
	MeasData   []MeasData `xml:"measData"`
	FileFooter FileFooter `xml:"fileFooter"`
}

type FileHeader struct {
	XMLName           xml.Name   `xml:"fileHeader"`
	FileSender        FileSender `xml:"fileSender"`
	MeasCollec        MeasCollec `xml:"measCollec"`
	FileFormatVersion string     `xml:"fileFormatVersion,attr"`
	VendorName        string     `xml:"vendorName,attr,omitempty"`
	DnPrefix          string     `xml:"dnPrefix,attr,omitempty"`
}

type FileSender struct {
	XMLName     xml.Name `xml:"fileSender"`
	LocalDn     string   `xml:"localDn,attr,omitempty"`
	ElementType string   `xml:"elementType,attr,omitempty"`
}

type MeasCollec struct {
	XMLName   xml.Name  `xml:"measCollec"`
	BeginTime time.Time `xml:"beginTime,attr"`
}

type MeasData struct {
	XMLName        xml.Name       `xml:"measData"`
	ManagedElement ManagedElement `xml:"managedElement"`
	MeasInfo       []MeasInfo     `xml:"measInfo"`
}

type ManagedElement struct {
	XMLName   xml.Name `xml:"managedElement"`
	LocalDn   string   `xml:"localDn,attr,omitempty"`
	UserLabel string   `xml:"userLabel,attr,omitempty"`
	SwVersion string   `xml:"swVersion,attr,omitempty"`
}

type MeasInfo struct {
	XMLName    xml.Name   `xml:"measInfo"`
	Job        Job        `xml:"job"`
	GranPeriod GranPeriod `xml:"granPeriod"`
	RepPeriod  RepPeriod  `xml:"repPeriod"`
	MeasTypes  []MeasType `xml:"measType"`
	MeasValue  MeasValue  `xml:"measValue"`
}

type Job struct {
	XMLName xml.Name `xml:"job"`
	JobId   string   `xml:"jobId,attr"`
}

type GranPeriod struct {
	XMLName  xml.Name  `xml:"granPeriod"`
	Duration string    `xml:"duration,attr"`
	EndTime  time.Time `xml:"endTime,attr"`
}

type RepPeriod struct {
	XMLName  xml.Name `xml:"repPeriod"`
	Duration string   `xml:"duration,attr"`
}

type MeasType struct {
	XMLName xml.Name `xml:"measType"`
	P       int      `xml:"p,attr"`
	Value   string   `xml:",chardata"`
}

type MeasValue struct {
	XMLName    xml.Name `xml:"measValue"`
	MeasObjLdn string   `xml:"measObjLdn,attr"`
	R          []R      `xml:"r"`
}

type R struct {
	XMLName xml.Name `xml:"r"`
	P       int      `xml:"p,attr"`
	Value   float64  `xml:",chardata"`
}

type FileFooter struct {
	XMLName    xml.Name   `xml:"fileFooter"`
	MeasCollec MeasCollec `xml:"measCollec"`
}

type MeasCollecEndTime struct {
	XMLName xml.Name `xml:"measCollec"`
	EndTime string   `xml:"endTime,attr"`
}
