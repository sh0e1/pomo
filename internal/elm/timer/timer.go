package timer

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run(cfg *Config) error {
	_, err := tea.NewProgram(initModel(cfg)).Run()
	return err
}

func initModel(cfg *Config) Model {
	m := Model{
		workModel: initWorkModel(cfg),
	}
	return m
}

type Model struct {
	workModel WorkModel
}

var _ tea.Model = Model{}

func (m Model) Init() tea.Cmd {
	return m.workModel.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.workModel.Update(msg)
}

func (m Model) View() string {
	return m.workModel.View()
}
