package timer

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
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
	timer    timer.Model
	quitting bool
	err      error
}

func initModel() model {
	return model{
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot)),
		timer:   timer.New(1 * time.Minute),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.timer.Init(), m.spinner.Tick)
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
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	return fmt.Sprintf("%s%s", m.spinner.View(), m.timer.View())
}
