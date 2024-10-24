package main

import (
	"reflect"
	"strconv"
)

type SettingItem interface {
	GetLabel() string
	GetValue() interface{}
	GetType() reflect.Type
	FilterValue() string
	Title() string
	Description() string
}

type IntSettingItem struct {
	label string
	value int
}

func (s IntSettingItem) GetLabel() string {
	return s.label
}

func (s IntSettingItem) GetValue() interface{} {
	return s.value
}

func (s IntSettingItem) GetType() reflect.Type {
	return reflect.TypeOf(s.value)
}

func (s IntSettingItem) FilterValue() string {
	return s.label
}

func (s IntSettingItem) Title() string {
	return s.label
}

func (s IntSettingItem) Description() string {
	return strconv.Itoa(s.value)
}

type FloatSettingItem struct {
	label string
	value float64
}

func (s FloatSettingItem) GetLabel() string {
	return s.label
}

func (s FloatSettingItem) GetValue() interface{} {
	return s.value
}

func (s FloatSettingItem) GetType() reflect.Type {
	return reflect.TypeOf(s.value)
}

func (s FloatSettingItem) FilterValue() string {
	return s.label
}

func (s FloatSettingItem) Title() string {
	return s.label
}

func (s FloatSettingItem) Description() string {
	return strconv.FormatFloat(s.value, 'f', -1, 64)
}

type StringSettingItem struct {
	label string
	value string
}

func (s StringSettingItem) GetLabel() string {
	return s.label
}

func (s StringSettingItem) GetValue() interface{} {
	return s.value
}

func (s StringSettingItem) GetType() reflect.Type {
	return reflect.TypeOf(s.value)
}

func (s StringSettingItem) FilterValue() string {
	return s.label
}

func (s StringSettingItem) Title() string {
	return s.label
}

func (s StringSettingItem) Description() string {
	return s.value
}

type SettingsList struct {
	title       string
	description string
	settings    []SettingItem
}

func (s SettingsList) FilterValue() string {
	return s.title
}

func (s SettingsList) Title() string {
	return s.title
}

func (s SettingsList) Description() string {
	return s.description
}

func (s SettingsList) GetSettings() []SettingItem {
	return s.settings
}
