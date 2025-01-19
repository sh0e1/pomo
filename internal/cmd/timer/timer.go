package timer

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/beeep"
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

const timeout = 25 * time.Minute

type model struct {
	timer    timer.Model
	spinner  spinner.Model
	keymap   keymap
	help     help.Model
	quitting bool
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func (k keymap) bindings() []key.Binding {
	return []key.Binding{k.start, k.stop, k.reset, k.quit}
}

func initModel() model {
	m := model{
		timer:   timer.New(timeout),
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot)),
		keymap: keymap{
			start: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "start")),
			stop:  key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "stop")),
			reset: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "reset")),
			quit:  key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"), key.WithHelp("q", "quit")),
		},
		help: help.New(),
	}

	m.keymap.start.SetEnabled(false)
	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.timer.Init(), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keymap.stop.SetEnabled(m.timer.Running())
		m.keymap.start.SetEnabled(!m.timer.Running())
		return m, cmd
	case timer.TimeoutMsg:
		m.quitting = true
		_ = beeep.Notify("pomo", "timer has expired", "")
		return m, tea.Quit
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.reset):
			m.timer.Timeout = timeout
			return m, nil
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			return m, m.timer.Toggle()
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
	return fmt.Sprintf("%s Existing in %s\n%s", m.spinner.View(), m.timer.View(), m.help.ShortHelpView(m.keymap.bindings()))
}
