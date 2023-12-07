package omc

import (
	"fmt"
	"strings"
	"time"

	"github.com/heypkg/store/jsontype"
	"github.com/netdoop/cwmp/acs"
	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type DeviceMethodCallState int

const (
	DeviceMethodCallStateInit            DeviceMethodCallState = 0
	DeviceMethodCallStateRequestSend     DeviceMethodCallState = 1
	DeviceMethodCallStateResponseRecv    DeviceMethodCallState = 2
	DeviceMethodCallStateResponseTimeout DeviceMethodCallState = 3
	DeviceMethodCallStateUnknow          DeviceMethodCallState = 100
)

type DeviceMethodCall struct {
	Time     jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_method_call_unique;not null"`
	Schema   string            `json:"Schema" gorm:"uniqueIndex:idx_device_method_call_unique;not null"`
	DeviceID uint              `json:"DeviceId" gorm:"uniqueIndex:idx_device_method_call_unique;not null"`
	Device   *Device           `json:"Device" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerialNumber"`
	ProductType  string `json:"ProductType" gorm:"index"`

	Updated           jsontype.JSONTime                 `json:"Updated" gorm:"autoUpdateTime"`
	State             DeviceMethodCallState             `json:"State" gorm:"index"`
	CommandKey        string                            `json:"CommandKey"`
	MethodName        string                            `json:"MethodName"`
	RequestValuesRaw  jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:request_values"`
	RequestValues     *jsontype.Tags                    `json:"RequestValues" gorm:"-"`
	ResponseValuesRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:response_values"`
	ResponseValues    *jsontype.Tags                    `json:"ResponseValues" gorm:"-"`
	FaultCode         int                               `json:"FaultCode"`
	FaultString       string                            `json:"FaultString"`
}

func (m *DeviceMethodCall) SaveData() {
	m.RequestValuesRaw = jsontype.NewJSONType(m.RequestValues)
	m.ResponseValuesRaw = jsontype.NewJSONType(m.ResponseValues)
}

func (m *DeviceMethodCall) LoadData() {
	m.RequestValues = m.RequestValuesRaw.Data
	m.ResponseValues = m.ResponseValuesRaw.Data
}

func (m *DeviceMethodCall) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *DeviceMethodCall) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func (m DeviceMethodCall) GetRequestValue(n string) string {
	if m.RequestValues == nil {
		return ""
	}
	return cast.ToString((*m.RequestValues)[n])
}

func (m DeviceMethodCall) GetMethodName() string {
	return m.MethodName
}
func (m DeviceMethodCall) GetCommandKey() string {
	return m.CommandKey
}
func (m DeviceMethodCall) GetRequestValues() map[string]string {
	values := map[string]string{}
	if m.RequestValues != nil {
		for k, v := range *m.RequestValues {
			values[k] = cast.ToString(v)
		}
	}
	return values
}

func (m *Device) GetMethodCall(commandKey string) acs.MethodCall {
	var obj DeviceMethodCall
	db := store.GetDB()
	t := time.Unix(0, cast.ToInt64(commandKey))
	if result := db.Where("device_id = ? AND time = ?", m.ID, jsontype.JSONTime(t)).First(&obj); result.Error != nil {
		return nil
	}
	return &obj
}

func (m *Device) UpdateMethodCallState(commandKey string, state DeviceMethodCallState) error {
	db := store.GetDB()
	obj, ok := m.GetMethodCall(commandKey).(*DeviceMethodCall)
	if obj != nil && ok {
		updateColumns := []string{"state"}
		updateData := &DeviceMethodCall{}
		updateData.State = state
		result := db.Model(obj).Where(
			"device_id = ? AND time = ?",
			m.ID,
			obj.Time,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update method call")
		}
	}
	return nil
}

func (m *Device) UpdateMethodCallRequestSend(commandKey string) error {
	return m.UpdateMethodCallState(commandKey, DeviceMethodCallStateRequestSend)
}

func (m *Device) UpdateMethodCallTimeout(commandKey string) error {
	return m.UpdateMethodCallState(commandKey, DeviceMethodCallStateResponseTimeout)
}

func (m *Device) UpdateMethodCallUnknow(commandKey string) error {
	return m.UpdateMethodCallState(commandKey, DeviceMethodCallStateUnknow)
}

func (m *Device) UpdateMethodCallResponse(commandKey string,
	values map[string]any,
	faultCode int,
	faultString string,
) error {
	db := store.GetDB()
	obj, ok := m.GetMethodCall(commandKey).(*DeviceMethodCall)
	if obj != nil && ok {
		_values := jsontype.Tags{}
		for k, v := range values {
			_values[k] = v
		}
		updateColumns := []string{"state", "response_values", "fault_code", "fault_string"}
		updateData := &DeviceMethodCall{}
		updateData.State = DeviceMethodCallStateResponseRecv
		updateData.ResponseValues = &_values
		updateData.FaultCode = faultCode
		updateData.FaultString = faultString
		updateData.SaveData()
		result := db.Model(obj).Where(
			"device_id = ? AND time = ?",
			m.ID,
			obj.Time,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update method call")
		}
	}
	return nil
}

func (m *Device) PushMethodCall(t time.Time, methodName string, values map[string]any) (acs.MethodCall, error) {
	db := store.GetDB()
	commandKey := fmt.Sprintf("%v", t.UnixNano())
	tags := jsontype.Tags(values)
	call := &DeviceMethodCall{
		Time:          jsontype.JSONTime(t),
		Schema:        m.Schema,
		DeviceID:      m.ID,
		Oui:           m.Oui,
		ProductClass:  m.ProductClass,
		SerialNumber:  m.SerialNumber,
		ProductType:   m.ProductType,
		CommandKey:    commandKey,
		MethodName:    methodName,
		RequestValues: &tags,
		State:         DeviceMethodCallStateInit,
	}
	result := db.Create(call)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "create new method call")
	}
	return call, nil
}

func (m *Device) GetNextMethodCall() acs.MethodCall {
	db := store.GetDB()
	var obj DeviceMethodCall
	if result := db.Where(
		"device_id = ? AND state = ?",
		m.ID,
		DeviceMethodCallStateInit,
	).First(&obj); result.Error != nil {
		return nil
	}
	return &obj
}

func (m *Device) PushGetDeviceParameterNames(parameterPath string, nextLevel bool) (acs.MethodCall, error) {
	db := store.GetDB()
	if !m.IsMethodSupported("GetParameterNames") {
		return nil, errors.New("method not supported")
	}
	if !strings.HasSuffix(parameterPath, ".") {
		return nil, errors.New("invalid parameter path")
	}
	if parameterWritables := m.ParameterWritables; parameterWritables != nil {
		removed := []string{}
		pathSize := len(strings.Split(parameterPath, "."))
		for k, _ := range *parameterWritables {
			if strings.HasPrefix(k, parameterPath) {
				if nextLevel {
					if n := len(strings.Split(k, ".")); n == pathSize {
						removed = append(removed, k)
					}
				} else {
					removed = append(removed, k)
				}
			}
		}
		for _, k := range removed {
			parameterWritables.RemoveValue(k)
		}
		updateColumns := []string{"parameter_writables"}
		updateData := &Device{
			ParameterWritables: parameterWritables,
		}
		updateData.SaveData()
		result := db.Model(m).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return nil, errors.Wrap(result.Error, "update device")
		}
	}

	values := map[string]any{
		"NextLevel":     nextLevel,
		"ParameterPath": parameterPath,
	}
	call, err := m.PushMethodCall(time.Now(), "GetParameterNames", values)
	if err != nil {
		return nil, errors.Wrap(err, "push method call")
	}
	return call, nil
}

func (m *Device) SendConnectRequest() error {
	logger := utils.GetLogger()
	url := m.ParameterValues.GetValue("Device.ManagementServer.ConnectionRequestURL")
	username := m.ParameterValues.GetValue("Device.ManagementServer.ConnectionRequestUsername")
	password := m.ParameterValues.GetValue("Device.ManagementServer.ConnectionRequestPassword")
	udpConnectionRequestAddress := m.ParameterValues.GetValue("Device.ManagementServer.UDPConnectionRequestAddress")
	natDetected := m.ParameterValues.GetValue("Device.ManagementServer.NATDetected")

	ok, err := acs.SendHttpConnetionRequest(url, username, password)
	if err != nil {
		logger.Error("send http connection request")
	}
	if ok {
		return nil
	}

	if natDetected == "1" {
		for i := 0; i < 3; i++ {
			_, err := acs.SendUDPConnectionRequest(udpConnectionRequestAddress, username, password)
			if err != nil {
				logger.Error("send udp connection request")
			}
			time.Sleep(time.Second)
		}
	}
	return nil
}
