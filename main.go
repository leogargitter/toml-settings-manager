package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var settingsLists []SettingsList = []SettingsList{
	SettingsList{title: "Conexão", description: "Configurações de conexão", settings: []SettingItem{
		IntSettingItem{label: "Número de conexões", value: 2},
		StringSettingItem{label: "Reponsável", value: "João"},
	}},
	SettingsList{title: "Casa", description: "Configurações de casa", settings: []SettingItem{
		IntSettingItem{label: "Número da casa", value: 1032},
		IntSettingItem{label: "Quantidade de pessoas", value: 4},
		StringSettingItem{label: "Proprietário", value: "José"},
		FloatSettingItem{label: "Renda", value: 12312.4},
	}},
}

type columnId int

const (
	mainColumn columnId = iota
	secondaryColumn
)

type Model struct {
	settingsView []list.Model
	err          error
	loaded       bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {
	settingsStart := []list.Item{}
	for _, setting := range settingsLists {
		settingsStart = append(settingsStart, setting)
	}

	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.settingsView = []list.Model{defaultList, defaultList}
	m.settingsView[mainColumn].Title = "Main Settings"
	m.settingsView[mainColumn].SetItems(settingsStart)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initLists(msg.Width, msg.Height)
		m.loaded = true
	}
	var cmd tea.Cmd
	m.settingsView[mainColumn], cmd = m.settingsView[mainColumn].Update(msg)
	m.settingsView[secondaryColumn].Title = m.settingsView[mainColumn].SelectedItem().FilterValue()
	secondaryColItems := []list.Item{}
	for _, mainSetting := range settingsLists {
		if mainSetting.Title() == m.settingsView[mainColumn].SelectedItem().FilterValue() {
			for _, secondarySetting := range mainSetting.settings {
				secondaryColItems = append(secondaryColItems, secondarySetting)
			}
		}
	}
	m.settingsView[secondaryColumn].SetItems(secondaryColItems)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(lipgloss.Left,
			m.settingsView[mainColumn].View(),
			m.settingsView[secondaryColumn].View())
	} else {
		return "loading.."
	}
}

func main() {
	m := New()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
