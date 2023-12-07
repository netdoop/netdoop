package omc

import (
	"strings"
)

type Methods []string

type ParameterValues map[string]string

func (m ParameterValues) GetValue(n string) string {
	v, ok := m[n]
	if ok {
		return v
	}
	return ""
}
func (m ParameterValues) SetValue(n, v string) {
	m[n] = v
}

func (m ParameterValues) SetValues(values map[string]string) {
	for n, v := range values {
		m.SetValue(n, v)
	}
}
func (m ParameterValues) GetValues() map[string]string {
	values := map[string]string{}
	for n, v := range m {
		if strings.HasSuffix(n, ".") {
			continue
		}
		values[n] = v
	}
	return values
}

type ParameterTypes map[string]string

func (m ParameterTypes) GetValue(n string) string {
	v, ok := m[n]
	if ok {
		return v
	}
	return ""
}
func (m ParameterTypes) SetValue(n, v string) {
	m[n] = v
}

func (m ParameterTypes) SetValues(values map[string]string) {
	for n, v := range values {
		m.SetValue(n, v)
	}
}
func (m ParameterTypes) GetValues() map[string]string {
	values := map[string]string{}
	for n, v := range m {
		if strings.HasSuffix(n, ".") {
			continue
		}
		values[n] = v
	}
	return values
}

type ParameterWritables map[string]bool

func (m ParameterWritables) GetValue(n string) bool {
	v, ok := m[n]
	if ok {
		return v
	}
	return false
}
func (m ParameterWritables) SetValue(n string, v bool) {
	m[n] = v
}
func (m ParameterWritables) RemoveValue(n string) {
	delete(m, n)
}
func (m ParameterWritables) SetValues(values map[string]bool) {
	for n, v := range values {
		m.SetValue(n, v)
	}
}

type ParameterNotifications map[string]int

func (m ParameterNotifications) GetValue(n string) int {
	v, ok := m[n]
	if ok {
		return v
	}
	return 0
}
func (m ParameterNotifications) SetValue(n string, v int) {
	m[n] = v
}
func (m ParameterNotifications) SetValues(values map[string]int) {
	for n, v := range values {
		m.SetValue(n, v)
	}
}

// type Parameter struct {
// 	Name         string      `json:"Name"`
// 	Value        interface{} `json:"Value"`
// 	Type         string      `json:"Type"`
// 	Writable     bool        `json:"Writable"`
// 	Notification int         `json:"Notification"`
// }
