package timer

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "timer",
		Short: "Start the pomo timer",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := tea.NewProgram(initModel()).Run()
			return err
		},
	}
	return cmd
}

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func initModel() model {
	s := spinner.New(spinner.WithSpinner(spinner.Dot))
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	return m.spinner.View()
}
