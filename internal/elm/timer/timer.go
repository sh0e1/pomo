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
		isWorking: true,
		keymaps: keymaps{
			quit: key.NewBinding(key.WithKeys("q", tea.KeyCtrlC.String()), key.WithHelp("q", "Quit")),
		},
		config: config,
	}
	return m
}

type Model struct {
	workModel WorkModel
	isWorking bool

	breakModel BreakModel
	isBreaking bool

	keymaps keymaps
	config  Config
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
			model, cmd := m.workModel.Update(msg)
			m.workModel = model.(WorkModel)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) Breaking() Model {
	m.isWorking = false
	m.breakModel = initBreakModel(m.config.ShortBreakInterval)
	m.isBreaking = true
	return m
}

func (m Model) View() string {
	if m.isWorking {
		return m.workModel.View()
	}
	return ""
}

type keymaps struct {
	quit key.Binding
}
