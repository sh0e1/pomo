package timer

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func Run(config Config) error {
	_, err := tea.NewProgram(initModel(config)).Run()
	return err
}

func initModel(config Config) Model {
	m := Model{
		work:      initWorkModel(config.WorkInterval),
		isWorking: true,
		keymaps: keymaps{
			quit: key.NewBinding(key.WithKeys("q", tea.KeyCtrlC.String()), key.WithHelp("q", "Quit")),
		},
	}
	return m
}

type Model struct {
	work      WorkModel
	isWorking bool
	keymaps keymaps
}

var _ tea.Model = Model{}

func (m Model) Init() tea.Cmd {
	return m.workModel.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymaps.quit):
			return m, tea.Quit
		}
	default:
		if m.isWorking {
			model, cmd := m.work.Update(msg)
			m.work = model.(WorkModel)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.isWorking {
		return m.work.View()
	}
	return ""
}

type keymaps struct {
	quit key.Binding
}
