package omc

import (
	"strings"
	"time"

	"github.com/netdoop/netdoop/models/omc/define/igd"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/store/jsontype"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (m *Device) UpdateMetaData(k string, v any) {
	if m.MetaData == nil {
		m.MetaData = &jsontype.Tags{}
	}
	m.MetaData.Set(k, v)
}
func (m *Device) UpdateMetaDatas(values map[string]any) {
	if m.MetaData == nil {
		m.MetaData = &jsontype.Tags{}
	}
	for k, v := range values {
		m.MetaData.Set(k, v)
	}
}

func (m *Device) GetProperty(k string) any {
	if m.Properties == nil {
		return nil
	}
	if v, ok := (*m.Properties)[k]; ok {
		return v
	}
	return nil
}

func (m *Device) UpdateProperty(k string, v any) {
	if m.Properties == nil {
		m.Properties = &jsontype.Tags{}
	}
	m.Properties.Set(k, v)
}

func (m *Device) UpdateProperties(values map[string]any) {
	if m.Properties == nil {
		m.Properties = &jsontype.Tags{}
	}
	for k, v := range values {
		m.Properties.Set(k, v)
	}
}

func (m *Device) UpdateParameterValue(k string, v string) {
	if m.ParameterValues == nil {
		m.ParameterValues = &ParameterValues{}
	}
	m.ParameterValues.SetValue(k, v)
}
func (m *Device) UpdateParameterValues(values map[string]string) {
	if m.ParameterValues == nil {
		m.ParameterValues = &ParameterValues{}
	}
	m.ParameterValues.SetValues(values)
}
func (m *Device) UpdateParameterWritable(k string, v bool) {
	if m.ParameterWritables == nil {
		m.ParameterWritables = &ParameterWritables{}
	}
	m.ParameterWritables.SetValue(k, v)
}
func (m *Device) UpdateParameterWritables(values map[string]bool) {
	if m.ParameterWritables == nil {
		m.ParameterWritables = &ParameterWritables{}
	}
	m.ParameterWritables.SetValues(values)
}
func (m *Device) UpdateParameterNotification(k string, v int) {
	if m.ParameterNotifications == nil {
		m.ParameterNotifications = &ParameterNotifications{}
	}
	m.ParameterNotifications.SetValue(k, v)
}
func (m *Device) UpdateParameterNotifications(values map[string]int) {
	if m.ParameterNotifications == nil {
		m.ParameterNotifications = &ParameterNotifications{}
	}
	m.ParameterNotifications.SetValues(values)
}

func (m Device) IsMethodSupported(name string) bool {
	if m.Methods == nil {
		return false
	}
	for _, v := range *m.Methods {
		if name == v {
			return true
		}
	}
	return false
}

func (m *Device) GetParameterPaths(path string, nextLevel bool) []string {
	paths := []string{}
	if m.ParameterWritables == nil {
		return paths
	}
	for n := range *m.ParameterWritables {
		parts := strings.Split(n, ".")
		if len(parts) < 2 {
			continue
		}
		p := strings.Join(parts[0:len(parts)-1], ".") + "."
		if strings.HasPrefix(p, path) {
			if nextLevel {
				if tmp := p[len(path):]; strings.Contains(tmp, ".") {
					continue
				}
			}
			paths = append(paths, p)
		}
	}
	return paths
}

func (m *Device) GetParameterNames(path string, nextLevel bool) []string {
	names := []string{}
	if m.ParameterWritables == nil {
		return names
	}
	for n := range *m.ParameterWritables {
		if strings.HasSuffix(n, ".") {
			continue
		}
		if strings.HasPrefix(n, path) {
			if nextLevel {
				if tmp := n[len(path):]; strings.Contains(tmp, ".") {
					continue
				}
			}
			names = append(names, n)
		}
	}
	return names
}

func (m *Device) FetchPmValues(db *gorm.DB, t time.Time) bool {
	lastFetchTime := time.Unix(0, 0)
	if v := m.GetProperty("LastFetchPmValuesTime"); v != nil {
		lastFetchTime = time.Unix(0, cast.ToInt64(v))
	}
	if t.Sub(lastFetchTime) > 30*time.Minute {
		product := GetProduct(m.Schema, m.Oui, m.ProductClass)
		if product != nil && product.PerformanceValueDefines != nil {
			names := []string{}
			for key := range *product.PerformanceValueDefines {
				if strings.Contains(key, ".{i}.") {
					continue
				}
				names = append(names, key)
			}
			values := map[string]any{}
			for _, n := range names {
				values[n] = ""
			}
			if len(values) > 0 {
				m.PushMethodCall(time.Now(), "GetParameterValues", values)
			}
			m.UpdateProperty("LastFetchPmValuesTime", t.UnixNano())
			return true
		}
	}
	return false
}

func (m *Device) UpdateByIgdValues(db *gorm.DB, values map[string]string) {
	// logger := utils.GetLogger()

	product := GetProduct(m.Schema, m.Oui, m.ProductClass)
	if product != nil && product.PerformanceValueDefines != nil {
		for key, def := range *product.PerformanceValueDefines {
			key = strings.ReplaceAll(key, ".{i}.", ".\\d.")
			re := utils.GetRegexp(key)
			alias := cast.ToString(def)
			for name, value := range values {
				if value != "" && re.Match([]byte(name)) {
					re2 := utils.GetRegexp("(-)?[0-9]+")
					matches := re2.FindAllString(value, -1)
					if len(matches) > 0 {
						v := cast.ToFloat64(matches[0])
						InsertDeviePerformanceValue(db, m, alias, name, v, nil)
					}
				}
			}
		}
	}

	data := igd.InternetGatewayDevice{}
	data.ReadValues(values)

	m.UpdateParameterValues(values)

	m.UpdateProperty("DeviceSummary", data.DeviceSummary)
	if deviceInfo := data.DeviceInfo; deviceInfo != nil {
		m.UpdateProperty("Manufacturer", deviceInfo.Manufacturer)
		m.UpdateProperty("ManufacturerOUI", deviceInfo.ManufacturerOUI)
		m.UpdateProperty("Description", deviceInfo.Description)
		m.UpdateProperty("ModelName", deviceInfo.ModelName)
		m.UpdateProperty("ProductClass", deviceInfo.ProductClass)
		m.UpdateProperty("SerialNumber", deviceInfo.SerialNumber)
		m.UpdateProperty("ModemFirmwareVersion", deviceInfo.ModemFirmwareVersion)
		m.UpdateProperty("HardwareVersion", deviceInfo.HardwareVersion)
		m.UpdateProperty("SoftwareVersion", deviceInfo.SoftwareVersion)
		m.UpdateProperty("AdditionalHardwareVersion", deviceInfo.AdditionalHardwareVersion)
		m.UpdateProperty("AdditionalSoftwareVersion", deviceInfo.AdditionalSoftwareVersion)
		m.UpdateProperty("SpecVersion", deviceInfo.SpecVersion)
		m.UpdateProperty("UpTime", deviceInfo.UpTime)
		m.UpdateProperty("FirstUseDate", deviceInfo.FirstUseDate)
		m.UpdateProperty("HardwarePlatform", deviceInfo.HardwarePlatform)

		m.UpdateProperty("X_VENDOR_DEVICE_TYPE", deviceInfo.X_VENDOR_DEVICE_TYPE)
		m.UpdateProperty("X_VENDOR_FEATURE_CODE", deviceInfo.X_VENDOR_FEATURE_CODE)
		if enbStatus := deviceInfo.X_VENDOR_ENODEB_STATUS; enbStatus != nil {
			if v := enbStatus.X_VENDOR_CELL_STATUS; v != nil {
				m.ActiveStatus = *v
			} else {
				m.ActiveStatus = "inactive"
			}
			m.UpdateProperty("X_VENDOR_CELL_STATUS", enbStatus.X_VENDOR_CELL_STATUS)
			m.UpdateProperty("X_VENDOR_MMEACT_NUM", enbStatus.X_VENDOR_MMEACT_NUM)
			m.UpdateProperty("X_VENDOR_UEACT_NUM", enbStatus.X_VENDOR_UEACT_NUM)
		}
		if euruInfo := deviceInfo.EURUinfo; euruInfo != nil {
			m.UpdateProperty("Eunum", euruInfo.Eunum)
			m.UpdateProperty("Runum", euruInfo.Runum)
		}
	}
	if mgmtSrv := data.ManagementServer; mgmtSrv != nil {
		m.UpdateProperty("ConnectionRequestURL", mgmtSrv.ConnectionRequestURL)
		m.UpdateProperty("ConnectionRequestUsername", mgmtSrv.ConnectionRequestUsername)
		m.UpdateProperty("ConnectionRequestPassword", mgmtSrv.ConnectionRequestPassword)
		m.UpdateProperty("UDPConnectionRequestAddress", mgmtSrv.UDPConnectionRequestAddress)
	}

	if webGui := data.WebGui; webGui != nil {
		if overview := webGui.Overview; overview != nil {
			if moduleInfo := overview.ModuleInfo; moduleInfo != nil {
				m.UpdateMetaData("IMSI", moduleInfo.IMSI)
				m.UpdateMetaData("IMEI", moduleInfo.IMEI)
				m.UpdateMetaData("ModuleModel", moduleInfo.Model)
				m.UpdateMetaData("ModuleVersion", moduleInfo.Version)
			}
			if internetStatus := overview.InternetStatus; internetStatus != nil {
				m.UpdateProperty("InternetStatus", internetStatus)
			}
			if lanStatus := overview.LANStatus; lanStatus != nil && lanStatus.MACAddress != nil {
				m.UpdateProperty("MACAddress", lanStatus.MACAddress)
			}
			if systemInfo := overview.SystemInfo; systemInfo != nil {
				m.UpdateProperty("RunTime", systemInfo.RunTime)
				m.UpdateProperty("OnlineTime", systemInfo.OnlineTime)
			}
			if gps := overview.GPS; gps != nil {
				m.UpdateProperty("Latitude", gps.Latitude)
				m.UpdateProperty("Longitude", gps.Longitude)
				m.UpdateProperty("Altitude", gps.Altitude)
			}
		}
		if network := webGui.Network; network != nil {
			if nrlte := network.NR_LTE; nrlte != nil {
				if status := nrlte.Status; status != nil {
					if nr := status.NR; nr != nil && nr.NR_CellID != nil {
						m.UpdateProperty("CellID", *nr.NR_CellID)
						m.UpdateProperty("PCI", *nr.NR_PCI)
						m.UpdateProperty("NCGI", *nr.NCGI)
						m.UpdateProperty("CGI", *nr.NCGI)
					} else if lte := status.LTE; lte != nil && lte.CellID != nil {
						m.UpdateProperty("CellID", *lte.CellID)
						m.UpdateProperty("PCI", *lte.PCI)
						m.UpdateProperty("ECGI", *lte.ECGI)
						m.UpdateProperty("CGI", *lte.ECGI)
					}
				}
			}
		}
	}
	if wanDevice, ok := data.WANDevice[1]; wanDevice != nil && ok {
		if wanConnectionDevice, ok := wanDevice.WANConnectionDevice[1]; wanConnectionDevice != nil && ok {
			if wanIPConnection, ok := wanConnectionDevice.WANIPConnection[1]; wanIPConnection != nil && ok {
				m.UpdateProperty("IPAddress", wanIPConnection.ExternalIPAddress)
			}
		}
		if wanEthernetInterfaceConfig := wanDevice.WANEthernetInterfaceConfig; wanEthernetInterfaceConfig != nil {
			m.UpdateProperty("MACAddress", wanEthernetInterfaceConfig.MACAddress)
		}
	}
	if fap := data.FAP; fap != nil {
		if gps := fap.GPS; gps != nil {
			if status := gps.ContinuousGPSStatus; status != nil {
				m.UpdateProperty("Latitude", status.X_VENDOR_Latitude)
				m.UpdateProperty("Longitude", status.X_VENDOR_Longitude)
			}
		}
		if perfMgmt := fap.PerfMgmt; perfMgmt != nil {
			if configs := perfMgmt.Config; len(configs) > 0 {
				for k, v := range configs {
					m.UpdateProperty("PerfMgmtIndex", k)
					m.UpdateProperty("PerfMgmtEnable", v.Enable)
					m.UpdateProperty("PerfMgmtPeriodicUploadInterval", v.PeriodicUploadInterval)
					break
				}
			}
		}
	}
	if data.Services != nil && len(data.Services.FAPService) > 0 {
		if service, ok := data.Services.FAPService[1]; service != nil && ok {
			if fapControl := service.FAPControl; fapControl != nil {
				if lte := fapControl.LTE; lte != nil {
					m.UpdateProperty("AdminState", lte.AdminState)
					m.UpdateProperty("OpState", lte.OpState)
					m.UpdateProperty("RFTxStatus", lte.RFTxStatus)
				}
			}
			if capabilities := service.Capabilities; capabilities != nil {
				m.UpdateProperty("MaxTxPower", capabilities.MaxTxPower)

			}
			if cellConfig := service.CellConfig; cellConfig != nil {
				if lte := cellConfig.LTE; lte != nil {
					for _, param := range lte.MmePoolConfigParam {
						if param != nil {
							m.UpdateProperty("MMEIp1", param.MMEIp1)
							m.UpdateProperty("MMEIp2", param.MMEIp2)
						}
					}

					if epc := lte.EPC; epc != nil {
						m.UpdateProperty("TAC", epc.TAC)
						for _, v := range epc.PLMNList {
							m.UpdateProperty("PLMNID", v.PLMNID)
						}
					}
					if ran := lte.RAN; ran != nil {
						if rf := ran.RF; rf != nil {
							m.UpdateProperty("EARFCNDL", rf.EARFCNDL)
							m.UpdateProperty("EARFCNUL", rf.EARFCNUL)
							m.UpdateProperty("ULBandwidth", rf.ULBandwidth)
							m.UpdateProperty("DLBandwidth", rf.DLBandwidth)
						}
					}
				}
			}
		}
	}

	if fm := data.FaultMgmt; fm != nil {
		logger := utils.GetLogger()
		if fm.CurrentAlarmNumberOfEntries != nil {
			logger.Warn("debug alarm", zap.Any("CurrentAlarmNumberOfEntries", *fm.CurrentAlarmNumberOfEntries))
			if *fm.CurrentAlarmNumberOfEntries == 0 {
				all, _ := getActiveAlarms(db, m)
				for _, activeAlarm := range all {
					if activeAlarm.EventType == "LostConnection" {
						continue
					}
					ClearCurrentAlarm(db, m, activeAlarm.AlarmIdentifier)
				}
			}
		}
		if fm.ExpeditedEventNumberOfEntries != nil {
			logger.Warn("debug alarm", zap.Any("ExpeditedEventNumberOfEntries", *fm.ExpeditedEventNumberOfEntries))
		}
		if fm.HistoryEventNumberOfEntries != nil {
			logger.Warn("debug alarm", zap.Any("HistoryEventNumberOfEntries", *fm.HistoryEventNumberOfEntries))
		}
		if fm.QueuedEventNumberOfEntries != nil {
			logger.Warn("debug alarm", zap.Any("QueuedEventNumberOfEntries", *fm.QueuedEventNumberOfEntries))
		}
		if current := fm.CurrentAlarm; current != nil {
			all, _ := getActiveAlarms(db, m)
			for _, activeAlarm := range all {
				if activeAlarm.EventType == "LostConnection" {
					continue
				}
				cleared := true
				for _, alarm := range current {
					if activeAlarm.AlarmIdentifier == *alarm.AlarmIdentifier {
						cleared = false
						break
					}
				}
				if cleared {
					ClearCurrentAlarm(db, m, activeAlarm.AlarmIdentifier)
				}
			}
			for _, alarm := range current {
				logger.Warn("debug alarm", zap.Any("CurrentAlarm", *alarm))
				UpsertCurrentAlarm(db, m, alarm)
			}
		}
		if events := fm.ExpeditedEvent; events != nil {
			for _, event := range events {
				logger.Warn("debug alarm", zap.Any("ExpeditedEvent", *event))
				UpsertExpeditedEvent(db, m, event)
			}
		}
		if events := fm.HistoryEvent; events != nil {
			for _, event := range events {
				logger.Warn("debug alarm", zap.Any("HistoryEvent", *event))
			}
		}
		if events := fm.QueuedEvent; events != nil {
			for _, event := range events {
				logger.Warn("debug alarm", zap.Any("QueuedEvent", *event))

			}
		}
	}
}
