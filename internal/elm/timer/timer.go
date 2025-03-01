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
		workModel: initWorkModel(config.WorkInterval),
		keymaps: keymaps{
			quit: key.NewBinding(key.WithKeys("q", tea.KeyCtrlC.String()), key.WithHelp("q", "Quit")),
		},
	}
	return m
}

type Model struct {
	workModel WorkModel
	keymaps   keymaps
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
	}
	return m, nil
}

func (m Model) View() string {
	return ""
}

type keymaps struct {
	quit key.Binding
}
